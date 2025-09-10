package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type SSHCredentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

func sshHandler(w http.ResponseWriter, r *http.Request) {
	// WebSocket upgraden
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer ws.Close()

	// Erste Nachricht = Verbindungsdaten
	_, msg, err := ws.ReadMessage()
	if err != nil {
		log.Println("Failed to read credentials:", err)
		return
	}

	var creds SSHCredentials
	if err := json.Unmarshal(msg, &creds); err != nil {
		ws.WriteMessage(websocket.TextMessage, []byte("Invalid credentials JSON"))
		return
	}

	// SSH Config mit Daten vom Frontend
	config := &ssh.ClientConfig{
		User: creds.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(creds.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Verbindung herstellen
	client, err := ssh.Dial("tcp", creds.Host, config)
	if err != nil {
		ws.WriteMessage(websocket.TextMessage, []byte("SSH connect failed: "+err.Error()))
		return
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		ws.WriteMessage(websocket.TextMessage, []byte("SSH session failed: "+err.Error()))
		return
	}
	defer session.Close()

	// Pseudo-Terminal anfordern
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
		ws.WriteMessage(websocket.TextMessage, []byte("RequestPty failed: "+err.Error()))
		return
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		ws.WriteMessage(websocket.TextMessage, []byte("stdin failed: "+err.Error()))
		return
	}
	stdout, err := session.StdoutPipe()
	if err != nil {
		ws.WriteMessage(websocket.TextMessage, []byte("stdout failed: "+err.Error()))
		return
	}

	// SSH Shell starten
	if err := session.Shell(); err != nil {
		ws.WriteMessage(websocket.TextMessage, []byte("Shell failed: "+err.Error()))
		return
	}

	// SSH -> WS
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if err != nil {
				break
			}
			ws.WriteMessage(websocket.TextMessage, buf[:n])
		}
	}()

	// WS -> SSH
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		stdin.Write(msg)
	}
}

//go:embed ui/dist/*
var staticFiles embed.FS

func main() {
	http.HandleFunc("/ssh", sshHandler)
	subFS, _ := fs.Sub(staticFiles, "ui/dist")
	http.Handle("/", http.FileServer(http.FS(subFS)))

	fmt.Println("Server läuft auf :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
