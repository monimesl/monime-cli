"use client";

import {useState} from "react";
import {newSession} from "@/model/session";
import {Button} from "@/components/ui/button.tsx";
import {useSession} from "@/model/session/provider.tsx";

const maxCharacters = 182;


export default function ReplyScreen() {
    const { setSession } = useSession();
    const [reply, setReply] = useState("");

    const handleCancel = () => {
        setSession({
            ...newSession('dial-pad'),
            replace: true,
        })
    };
    const handleSendReply = () => {
       const sanitizedReply = reply.trim()
        if (!sanitizedReply)return
        setSession({
            screen: 'loading',
            inputs: {
                reply: sanitizedReply,
            }
        })
    };
    const remainingCharacters = maxCharacters - reply.length;
    return (
        <div className="flex items-center flex-col justify-center h-full bg-gray-700 p-0 w-full">
            <>
                {/* Reply Header */}
                <div className="flex justify-between items-center px-4 py-3 border-b border-gray-600 w-full">
                    <div className="flex items-center gap-2">
                        <Button
                            variant="ghost"
                            size={"lg"}
                            className="text-blue-400 hover:text-blue-300 hover:bg-transparent p-0 h-auto font-normal"
                            onClick={handleCancel}
                        >
                            Cancel
                        </Button>
                    </div>
                    <div className="text-white font-medium">Reply</div>
                    <div className="flex items-center gap-2">
                        <Button
                            variant="ghost"
                            size={"lg"}
                            className="text-blue-400 hover:text-blue-300 hover:bg-transparent p-0 h-auto font-normal"
                            onClick={handleSendReply}
                            disabled={!reply.trim()}
                        >
                            Reply
                        </Button>
                    </div>
                </div>

                {/* Message Content */}
                <div className="flex-1 px-4 py-6 text-center">
                    <div className="text-white mt-10">
                        <div className="text-sm leading-relaxed">
                            Monime<br/>
                            You are about to pay Sulailas Multi Cuisine Restaurant
                            <br/>
                            1: Confirm
                            <br/>
                            2: Cancel
                            <br/>
                            <br/>- - -
                            <br/>
                            <br/>
                            00:menu
                            <br/>
                            0:back
                        </div>
                    </div>
                </div>

                {/* Input Field */}
                <div className="px-4 mb-4 w-full">
                    <textarea
                        value={reply}
                        onChange={(e) => setReply(e.target.value)}
                        placeholder="Type your reply..."
                        className="w-full h-16 bg-gray-600 border border-gray-500 rounded-lg px-3 py-2 text-white placeholder-gray-400 resize-none focus:outline-none focus:border-blue-400"
                        maxLength={maxCharacters}
                    />
                    <div className="text-gray-400 text-sm mt-1 text-center">
                        {remainingCharacters} characters remaining
                    </div>
                </div>
                <div className="h-40"/>
            </>
        </div>
    );
}
