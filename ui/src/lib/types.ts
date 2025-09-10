export type Server = {
    name: string;
    ip: string;
    connections: SshConnection[];
    username: string;
    password: string;
    port: number;
    tags: string[];
}

export type SshConnection = {
    id: number;
}