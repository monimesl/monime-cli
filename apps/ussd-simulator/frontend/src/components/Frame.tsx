import {PropsWithChildren, useEffect, useState} from "react";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue,} from "@/components/ui/select";
import {getNetworkName, NetworkId} from "@/model/config";
import {useConfig} from "@/model/config/provider.tsx";
import {useSession} from "@/model/session/provider.tsx";


const getCurrentTimeText = () => {
    const now = new Date();
    const hours = now.getHours().toString().padStart(2, '0');
    const minutes = now.getMinutes().toString().padStart(2, '0');
    return `${hours}:${minutes}`;
}

export default function Frame(props: PropsWithChildren & { darkBackground?: boolean }) {
    const {session} = useSession()
    const {darkBackground, children} = props;
    const {config, setConfig} = useConfig();
    const [currentTime, setCurrentTime] = useState(getCurrentTimeText())
    useEffect(() => {
        const ref = setInterval(() => {
            setCurrentTime(getCurrentTimeText());
        }, 500)
        return () => clearInterval(ref);
    }, [])
    return (
        <div
            style={{
                width: "380px",
                height: "720px",
                background: darkBackground ? "#374151" : "white",
                borderRadius: "40px",
                padding: "10px 10px 10px 10px",
                boxShadow: "0 0 20px rgba(0,0,0,0.2)",
                display: "flex",
                flexDirection: "column",
                justifyContent: "center",
                alignItems: "center",
                position: "relative",
                overflow: "hidden",
            }}
            className="mx-auto mt-10 overflow-hidden"
        >
            <div
                style={{
                    color: "#666",
                    display: "flex",
                    padding: "8px",
                    width: "100%",
                    justifyContent: "space-between",
                }}
            >
                <div style={{color: darkBackground ? "white" : "black"}}>
                    {currentTime} &nbsp;
                </div>
                <Select
                    value={config.network.id}
                    open={session.screen !== 'dial-pad' ? false : undefined}
                    onValueChange={(value) => {
                        setConfig({
                            network: {
                                id: value as NetworkId,
                            }
                        })
                    }}
                >
                    <SelectTrigger
                        style={{
                            border: "none",
                            background: "none",
                            outline: "none",
                            color: darkBackground? "white" : "black",
                        }}
                        className="max-w-fit p-0 m-0 h-fit focus:outline-none focus:ring-0 focus:border-none"
                    >
                        <SelectValue
                            placeholder="Select a network"
                            className="p-0 m-0"
                            style={{
                                border: "none",
                                background: "none",
                                outline: "none",
                                color: darkBackground ? "white" : "black",
                            }}
                        >
                            {getNetworkName(config.network.id)} 📶
                        </SelectValue>
                    </SelectTrigger>
                    <SelectContent>
                        <SelectItem value="m17">Orange</SelectItem>
                        <SelectItem value="m18">Africell</SelectItem>
                    </SelectContent>
                </Select>
            </div>
            {children}
        </div>
    );
}
