<script lang="ts">
    import Modal from "./lib/components/Modal.svelte";
    import SideBar from "./lib/components/SideBar.svelte";
    import Terminal from "./lib/components/Terminal.svelte";
    import { currentServer, currentSession, showNewServerModal } from "./lib/stores";

    currentServer.set({
        name: "Localhost",
        ip: "202.61.206.136",
        username: "leon",
        password: "Paco0705!",
        port: 22,
        connections: [],
        tags: []
    })

    function closeSession(index: number) {
        currentServer.update(server => {
            if (server) {
                server.connections.splice(index, 1);
            }
            return server;
        });

        if ($currentServer?.connections.length > 0) {
            currentSession.set($currentServer.connections[0]);
            console.log($currentSession);
        } else {
            currentSession.set({ id: 0 });
        }
    }

    function addSession() {
        // get new id
        let newId = 1;
        if ($currentServer?.connections.length > 0) {
            newId = $currentServer.connections[$currentServer.connections.length - 1].id + 1;
        }   

        // check if newId already exists
        while ($currentServer?.connections.find(con => con.id === newId)) {
            newId++;
        }

        currentServer.update(server => {
            if (server) {
                server.connections.push({
                    id: newId
                });
            }
            return server;
        });

        currentSession.set($currentServer?.connections[$currentServer.connections.length - 1]);
    }

    function openSession(index: number) {
        currentSession.set($currentServer?.connections[index]);
    }
</script>

{#if $showNewServerModal}
    <Modal on:close={() => showNewServerModal.set(false)}>
        <span slot="header" class="text-white text-lg font-semibold">Neuen Server hinzufügen</span>
        <form on:submit|preventDefault class="flex flex-col gap-y-4 py-4">
            <div class="flex flex-col">
                <label class="text-white" for="name">Name</label>
                <input class="bg-neutral-800 text-white border-2 outline-none border-neutral-700 rounded-lg p-2" type="text" id="name" required />
            </div>
            <div class="flex flex-col">
                <label class="text-white" for="ip">IP-Adresse</label>
                <input class="bg-neutral-800 text-white border-2 outline-none border-neutral-700 rounded-lg p-2" type="text" id="ip" required />
            </div>
            <div class="flex flex-col">
                <label class="text-white" for="username">Benutzername</label>
                <input class="bg-neutral-800 text-white border-2 outline-none border-neutral-700 rounded-lg p-2" type="text" id="username" required />
            </div>
            <div class="flex flex-col">
                <label class="text-white" for="password">Passwort</label>
                <input class="bg-neutral-800 text-white border-2 outline-none border-neutral-700 rounded-lg p-2" type="password" id="password" required />
            </div>
            <div class="flex flex-col">
                <label class="text-white" for="port">Port</label>
                <input class="bg-neutral-800 text-white border-2 outline-none border-neutral-700 rounded-lg p-2" type="number" id="port" required />
            </div>
            <button class="mt-4 bg-blue-500 transition-all hover:bg-blue-600 cursor-pointer text-white rounded-lg p-2" type="submit">Hinzufügen</button>
        </form>
    </Modal>
{/if}

<div class="h-screen flex bg-neutral-950 p-2">
    <SideBar />
    <div class="w-full h-full ml-4 mr-2 flex flex-col">
        <div class="rounded-xl p-1 border-2 flex gap-2 border-neutral-800 bg-neutral-900">
            {#if $currentServer && $currentServer.connections.length > 0}
                <div class="flex h-auto gap-2 flex-wrap">
                    {#each $currentServer.connections as con, i}
                        <div class="rounded-xl text-white h-12 border-2 border-neutral-800 transition-all bg-[#1a1a1a] flex justify-center items-center text-center">
                            <button on:click={() => openSession(i)} class="flex hover:cursor-pointer items-center hover:rounded-l-xl px-2 hover:bg-[#1f1f1f] h-full">
                                <svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="gray"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-terminal"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M5 7l5 5l-5 5" /><path d="M12 19l7 0" /></svg>
                                <h1 class="ml-2 text-sm">Session {con.id}</h1>
                            </button>
                            <button on:click={() => closeSession(i)} class="border-l-2 px-2 hover:cursor-pointer hover:rounded-r-xl h-full hover:bg-[#1f1f1f] flex justify-center items-center text-center border-neutral-800" aria-label="close_session">
                                <svg  xmlns="http://www.w3.org/2000/svg"  width="20"  height="20"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-x"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M18 6l-12 12" /><path d="M6 6l12 12" /></svg>
                            </button>
                        </div>
                    {/each}
                    <button on:click={() => addSession()} class="rounded-xl hover:cursor-pointer w-10 border-2 border-neutral-800 transition-all hover:bg-[#1f1f1f] bg-[#1a1a1a] flex justify-center items-center text-center" aria-label="add_session">
                        <svg  xmlns="http://www.w3.org/2000/svg"  width="20"  height="20"  viewBox="0 0 24 24"  fill="none"  stroke="white"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-plus"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M12 5l0 14" /><path d="M5 12l14 0" /></svg>
                    </button>
                </div>
            {:else}
                <button on:click={() => addSession()} class="rounded-xl h-12 hover:cursor-pointer w-10 border-2 border-neutral-800 transition-all hover:bg-[#1f1f1f] bg-[#1a1a1a] flex justify-center items-center text-center" aria-label="add_session">
                    <svg  xmlns="http://www.w3.org/2000/svg"  width="20"  height="20"  viewBox="0 0 24 24"  fill="none"  stroke="white"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-plus"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M12 5l0 14" /><path d="M5 12l14 0" /></svg>
                </button>
            {/if}
        </div>
        <div class="w-full h-full mt-4 rounded-xl border-2 border-neutral-800 bg-neutral-900 p-1.5">
            <div class="h-full w-full bg-black overflow-hidden">
                {#each $currentServer.connections as con (con.id)}
                    {#if $currentSession }
                        <div class={con.id === $currentSession.id ? "block h-full w-full" : "hidden"}>
                            <Terminal
                                username={$currentServer.username}
                                password={$currentServer.password}
                                host={$currentServer.ip}
                                port={$currentServer.port}
                            />
                        </div>
                    {/if}
                {:else}
                    <div class="h-full w-full flex justify-center items-center text-center">
                        <h1 class="text-white text-2xl">Keine aktiven Sessions. Klicke auf das <span class="font-bold">+</span> Symbol um eine neue Session zu erstellen.</h1>
                    </div>
                {/each}
            </div>
        </div>
    </div>
</div>