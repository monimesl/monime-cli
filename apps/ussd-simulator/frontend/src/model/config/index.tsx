const networks = {
    m17: "Orange",
    m18: "Africell",
};

export const getNetworkName = (id: NetworkId) => networks[id]

export type NetworkId = "m17" | "m18";

export type Network = {
    id: NetworkId;
}

export interface Config {
    network: Network;
    frame?: {
        color?: "dark" | "light"
    }
}


export const newConfig= (): Config => {
    return ({
        network: {
            id: 'm17'
        }
    })
}