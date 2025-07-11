"use client";
import {useEffect} from "react";

import {useSendReply} from "@/lib/apis.ts";
import Spinner from "@/components/ui/spinner.tsx";
import {useSession} from "@/model/session/provider.tsx";

export default function LoadingScreen() {
    const {session, setSession} = useSession();
    const {reply, error} = useSendReply(session);
    useEffect(() => {
        if (!error) return
        setSession({
            screen: 'terminal',
            outputs: {
                message: 'Error performing request<br/>Unknown Error',
            }
        })
    }, [error, setSession])
    useEffect(() => {
        if (!reply) return
        setSession({
            screen: reply.terminate ? 'terminal' : 'feedback',
            outputs: {
                message: reply.responseMessage,
            }
        })
    }, [reply, setSession]);
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
