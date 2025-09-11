<script lang="ts">
    import { onMount } from "svelte";
    import { currentOptionsServer, currentServer, currentSession, optionsVisible, serverlist, showNewServerModal } from "../stores";
    import Modal from "./Modal.svelte";

    // serverlist.set([{name: "MTL Web", ip: "202.61.206.136", username: "leon", password: "Paco0705!", port: 22, connections: [], tags: ["MTLWeb", "Debian"]}]);
    let username = "";
    let optionsServer = {
        name: "",
        ip: "",
        username: "",
        password: "",
        port: 22,
        tags: [] as string[],
    };

    onMount(() => {
        fetch("http://localhost:8080/api/username")
            .then(res => res.json())
            .then(data => {
                username = data.username;
            })
            .catch(err => {
                console.error("Error fetching username:", err);
                username = "User";
            });
    });
    
    function openServer(index: number) {
        currentServer.set($serverlist[index]);
        currentSession.set({ id: 0 });
    }

    function openOptions(index: number) {
        optionsVisible.set(true);
        optionsServer = { 
            name: $serverlist[index].name,
            ip: $serverlist[index].ip,
            username: $serverlist[index].username,
            password: $serverlist[index].password,
            port: $serverlist[index].port,
            tags: [...$serverlist[index].tags],
        };
    }

    function removeServer() {
        if (optionsServer.name === "" || optionsServer.ip === "" || optionsServer.username === "" || optionsServer.password === "" || optionsServer.port === 0) {
            alert("Bitte füllen Sie alle Pflichtfelder aus.");
            return;
        }

        serverlist.update(servers => {
            return servers.filter(s => !(s.name === optionsServer.name && s.ip === optionsServer.ip && s.username === optionsServer.username && s.port === optionsServer.port));
        });

        optionsVisible.set(false);
        optionsServer = {
            name: "",
            ip: "",
            username: "",
            password: "",
            port: 22,
            tags: [],
        };
    }

    // FIXME
    function saveOptions() {
        if (optionsServer.name === "" || optionsServer.ip === "" || optionsServer.username === "" || optionsServer.password === "" || optionsServer.port === 0) {
            alert("Bitte füllen Sie alle Pflichtfelder aus.");
            return;
        }

        serverlist.update(servers => {
            const index = servers.findIndex(s => s.name === $currentOptionsServer?.name && s.ip === $currentOptionsServer?.ip && s.username === $currentOptionsServer?.username && s.port === $currentOptionsServer?.port);
            if (index !== -1) {
                servers[index] = { 
                    ...optionsServer, 
                    tags: optionsServer.tags.map(tag => tag.trim()).filter(tag => tag !== ""), 
                    connections: servers[index].connections 
                };
            }
            return servers;
        });

        optionsVisible.set(false);
        optionsServer = {
            name: "",
            ip: "",
            username: "",
            password: "",
            port: 22,
            tags: [],
        };
    }
</script>

{#if $optionsVisible}
    <Modal on:close={() => optionsVisible.set(false)}>
        <span slot="header" class="text-white text-lg font-semibold">Server bearbeiten</span>
        <form on:submit|preventDefault={saveOptions} class="flex flex-col gap-y-4 py-4">
            <div class="flex flex-col">
                <label class="text-white" for="name">* Name</label>
                <input bind:value={optionsServer.name} class="bg-neutral-800 text-white border-2 outline-none border-neutral-700 rounded-lg p-2" type="text" id="name" required />
            </div>
            <div class="flex flex-col">
                <label class="text-white" for="ip">* IP-Adresse</label>
                <input bind:value={optionsServer.ip} class="bg-neutral-800 text-white border-2 outline-none border-neutral-700 rounded-lg p-2" type="text" id="ip" required />
            </div>
            <div class="flex flex-col">
                <label class="text-white" for="username">* Benutzername</label>
                <input bind:value={optionsServer.username} class="bg-neutral-800 text-white border-2 outline-none border-neutral-700 rounded-lg p-2" type="text" id="username" required />
            </div>
            <div class="flex flex-col">
                <label class="text-white" for="password">* Passwort</label>
                <input bind:value={optionsServer.password} class="bg-neutral-800 text-white border-2 outline-none border-neutral-700 rounded-lg p-2" type="password" id="password" required />
            </div>
            <div class="flex flex-col">
                <label class="text-white" for="port">* Port <span class="text-neutral-400 italic">(Standard is 22)</span></label>
                <input bind:value={optionsServer.port} class="bg-neutral-800 text-white border-2 outline-none border-neutral-700 rounded-lg p-2" type="number" id="port" required />
            </div>
            <div class="flex flex-col">
                <label class="text-white" for="tags">Tags <span class="text-neutral-400 italic">(Comma separated)</span></label>
                <input bind:value={optionsServer.tags} class="bg-neutral-800 text-white border-2 outline-none border-neutral-700 rounded-lg p-2" type="text" id="tags" />
            </div>
            <!-- delete button -->
            <div class="flex gap-2 w-full justify-between items-center">
                <button on:click={() => removeServer()} class="bg-red-600 w-full transition-all hover:bg-red-700 cursor-pointer text-white rounded-lg p-2" type="button">Löschen</button>
                <button type="submit" class="bg-lime-600 w-full transition-all hover:bg-lime-700 cursor-pointer text-white rounded-lg p-2">Speichern</button>
            </div>
        </form>
    </Modal>
{/if}

<div class="h-full rounded-xl border-2 w-64 border-neutral-800 bg-neutral-900 flex flex-col justify-between">
    <div>
        <div class="h-12 border-b-2 border-neutral-800 flex justify-center items-center">
            <h1 class="text-white text-xl font-semibold text-center w-full">Shellix</h1>
        </div>
        <div class="h-14 border-b-2 border-neutral-800 flex justify-center items-center">
            <button on:click={() => showNewServerModal.set(true)} class="w-11/12 h-8 cursor-pointer bg-neutral-800 hover:bg-neutral-700 rounded-lg text-white font-semibold transition-all" aria-label="Add_Server">
                Server hinzufügen
            </button>
        </div>
        <div class="py-4 px-2 flex flex-col gap-y-2">
            {#each $serverlist as server, i}
                <div class="flex flex-col w-full border-2 rounded-xl border-neutral-800 bg-neutral-950">
                    <div class="px-2 flex w-full justify-between items-center h-12 text-white">
                        <div class="flex items-center">
                            <span class="relative flex size-2">
                            <span class="absolute inline-flex h-full w-full animate-ping rounded-full bg-green-500 opacity-75"></span>
                            <span class="relative inline-flex size-2 rounded-full bg-green-600"></span>
                            </span>
                            <h1 class="ml-3 text-sm font-semibold">{server.name}</h1>
                        </div>
                        <div class="flex gap-x-2">
                            <button on:click={() => openOptions(i)} class="border-2 hover:cursor-pointer rounded-lg flex items-center justify-center w-7 h-7 bg-neutral-900 border-neutral-800" aria-label="settings">
                                <svg  xmlns="http://www.w3.org/2000/svg"  width="18"  height="18"  viewBox="0 0 24 24"  fill="none"  stroke="white"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-settings"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M10.325 4.317c.426 -1.756 2.924 -1.756 3.35 0a1.724 1.724 0 0 0 2.573 1.066c1.543 -.94 3.31 .826 2.37 2.37a1.724 1.724 0 0 0 1.065 2.572c1.756 .426 1.756 2.924 0 3.35a1.724 1.724 0 0 0 -1.066 2.573c.94 1.543 -.826 3.31 -2.37 2.37a1.724 1.724 0 0 0 -2.572 1.065c-.426 1.756 -2.924 1.756 -3.35 0a1.724 1.724 0 0 0 -2.573 -1.066c-1.543 .94 -3.31 -.826 -2.37 -2.37a1.724 1.724 0 0 0 -1.065 -2.572c-1.756 -.426 -1.756 -2.924 0 -3.35a1.724 1.724 0 0 0 1.066 -2.573c-.94 -1.543 .826 -3.31 2.37 -2.37a1.724 1.724 0 0 0 .572 -.308z" /><circle cx="12" cy="12" r="3" /></svg>
                            </button>
                            <button on:click={() => openServer(i)} class="border-2 hover:cursor-pointer rounded-lg flex items-center justify-center w-7 h-7 bg-neutral-900 border-neutral-800" aria-label="open_terminal">
                                <svg  xmlns="http://www.w3.org/2000/svg"  width="18"  height="18"  viewBox="0 0 24 24"  fill="none"  stroke="white"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-terminal"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M5 7l5 5l-5 5" /><path d="M12 19l7 0" /></svg>
                            </button>
                        </div>
                    </div>
                    {#if server.tags.length > 0}
                        <div class="flex p-2 flex-wrap gap-2">
                            {#each server.tags as tag}
                                <div class="bg-neutral-900 px-1.5 border-2 rounded-xl border-neutral-800">
                                    <h1 class="text-white text-xs">{tag}</h1>
                                </div>
                            {/each}
                        </div>
                    {/if}
                </div>  
            {/each}
        </div>
    </div>
    <div class="h-12 border-t-2 border-neutral-800 flex justify-between items-center px-4">
        <div class="flex">
            <svg  xmlns="http://www.w3.org/2000/svg"  width="20"  height="20"  viewBox="0 0 24 24"  fill="none"  stroke="white"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-user"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M8 7a4 4 0 1 0 8 0a4 4 0 0 0 -8 0" /><path d="M6 21v-2a4 4 0 0 1 4 -4h4a4 4 0 0 1 4 4v2" /></svg>
            <h1 class="ml-2 text-sm text-white">{username}</h1>
        </div>
        <div></div>
    </div>
</div>