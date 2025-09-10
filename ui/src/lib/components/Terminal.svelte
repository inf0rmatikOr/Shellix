<script lang="ts">
    import { onDestroy, onMount } from "svelte";
    import { Terminal } from "xterm";
    import { FitAddon } from "xterm-addon-fit";
    import "xterm/css/xterm.css";

    export let username: string;
    export let password: string;
    export let host: string;
    export let port: number;

    let terminalDiv: HTMLDivElement;
    const fitAddon = new FitAddon();
    let term: Terminal;
    let ws: WebSocket;

    function resizeTerminal() {
        fitAddon.fit();
    }

    onMount(() => {
        term = new Terminal({
            fontSize: 14,
            theme: { background: "black", foreground: "#ffffff" },
            cursorBlink: true
        });

        term.loadAddon(fitAddon);
        term.open(terminalDiv);
        fitAddon.fit();

        // WebSocket-Verbindung aufbauen
        ws = new WebSocket("ws://localhost:8080/ssh");

        ws.onopen = () => {
            // Credentials an Server senden
            ws.send(
                JSON.stringify({
                    user: username,
                    password: password,
                    host: host + ":" + port.toString(),
                })
            );

            term.writeln("Verbinde mit SSH...");
        };

        // SSH-Ausgabe ins Terminal
        ws.onmessage = (event) => {
            term.write(event.data);
        };

        ws.onerror = (err) => {
            term.writeln("\r\n[Fehler bei Verbindung zum Backend]");
            console.error(err);
        };

        ws.onclose = () => {
            term.writeln("\r\n[Verbindung beendet]");
        };

        // Eingaben vom User -> an Server
        term.onData((data) => {
            ws.send(data);
        });

        window.addEventListener("resize", resizeTerminal);
    });

    onDestroy(() => {
        window.removeEventListener("resize", resizeTerminal);
        if (ws) ws.close();
    });
</script>

<div bind:this={terminalDiv} class="w-full h-full p-1"></div>
