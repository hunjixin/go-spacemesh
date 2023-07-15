import axios from 'axios';

const prefix = ""

export interface Status {
    state: string;
    machine: string;
    completedSize: string;
    commitmentSize: string;
    percent: string;
}

export async function getMachineConnection(): Promise<Map<string, string>> {
    const response = await axios.get(prefix+'/api/data');
    return new Map<string, string>(Object.entries(response.data));
}

export async function getMachineInfo(machine:string): Promise<Status> {
    const response = await axios.get(prefix+'/api/machine/info/'+machine);
    return response.data;
}
