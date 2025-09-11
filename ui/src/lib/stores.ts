import { writable } from "svelte/store";
import type { User, Server, SshConnection } from "./types";

export const serverlist = writable<Server[]>([]);
export const currentServer = writable<Server>();
export const currentSession = writable<SshConnection | null>();
export const showNewServerModal = writable(false);
export const currentUser = writable<User | null>(null);
export const optionsVisible = writable(false);
export const currentOptionsServer = writable<Server | null>(null);