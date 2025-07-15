"use client";
import {useEffect} from "react";

import {useExchange} from "@/lib/apis.ts";
import Spinner from "@/components/ui/spinner.tsx";
import {useSession} from "@/model/session/provider.tsx";
import {useConfig} from "@/model/config/provider.tsx";
import {sanitizeMessage} from "@/lib/utils.ts";

export default function LoadingScreen() {
    const {config} = useConfig()
    const {session, setSession} = useSession();
    const {response, error} = useExchange(config, session);
    useEffect(() => {
        if (!error) return
        console.error("USSD exception error:", error);
        setSession({
            screen: 'terminal',
            outputs: {
                message: 'Error performing request<br/>Unknown Error',
            }
        })
    }, [error, setSession])
    useEffect(() => {
        if (!response) return
        setSession({
            id: response.sessionId,
            screen: response.terminate ? 'terminal' : 'feedback',
            outputs: {
                message: sanitizeMessage(response.responseMessage),
            },
        })
    }, [response, setSession]);
    return (
        <div className="flex items-center flex-col justify-center h-full bg-gray-700 p-0 w-full">
            <>
                {/* Main Content */}
                <div className="flex-1 flex flex-col justify-center items-center text-center px-6">
                    <div className="text-white mb-8 w-full">
                        <Spinner/>
                        <p className="text-base mt-3">Please wait...</p>
                    </div>
                </div>
            </>
        </div>
    );
}
