import { writable } from "svelte/store";
import type { Server, SshConnection } from "./types";

export const serverlist = writable<Server[]>();
export const currentServer = writable<Server>();
export const currentSession = writable<SshConnection | null>();
export const showNewServerModal = writable(false);