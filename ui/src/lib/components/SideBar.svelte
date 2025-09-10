<script>
    import { serverlist, showNewServerModal } from "../stores";

    serverlist.set([{name: "MTL Web", ip: "202.61.206.136", username: "leon", password: "Paco0705!", port: 22, connections: [], tags: ["MTLWeb", "Debian"]}]);
</script>

<div class="h-full rounded-xl border-2 w-58 border-neutral-800 bg-neutral-900 flex flex-col justify-between">
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
            {#each $serverlist as server}
                <div class="flex flex-col w-full border-2 rounded-xl border-neutral-800 bg-neutral-950">
                    <div class="px-2 flex w-full justify-between items-center h-12 text-white">
                        <div class="flex items-center">
                            <span class="relative flex size-2">
                            <span class="absolute inline-flex h-full w-full animate-ping rounded-full bg-green-500 opacity-75"></span>
                            <span class="relative inline-flex size-2 rounded-full bg-green-600"></span>
                            </span>
                            <h1 class="ml-3 text-sm font-semibold">{server.name}</h1>
                        </div>
                        <button class="border-2 hover:cursor-pointer rounded-lg flex items-center justify-center w-7 h-7 bg-neutral-900 border-neutral-800" aria-label="open_terminal">
                            <svg  xmlns="http://www.w3.org/2000/svg"  width="18"  height="18"  viewBox="0 0 24 24"  fill="none"  stroke="white"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-terminal"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M5 7l5 5l-5 5" /><path d="M12 19l7 0" /></svg>
                        </button>
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
            <h1 class="ml-2 text-sm text-white">Leon</h1>
        </div>
        <button class="text-white hover:cursor-pointer hover:text-neutral-400 transition-all" aria-label="Logout">
            <svg  xmlns="http://www.w3.org/2000/svg"  width="20"  height="20"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-logout-2"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M10 8v-2a2 2 0 0 1 2 -2h7a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-7a2 2 0 0 1 -2 -2v-2" /><path d="M15 12h-12l3 -3" /><path d="M6 15l-3 -3" /></svg>
        </button>
    </div>
</div>