"use client";

import {Button} from "@/components/ui/button.tsx";

export default function PromptScreen() {
    return (
        <div className="flex items-center flex-col justify-center h-full bg-gray-700 p-0 w-full">
            <>
                {/* Main Content */}
                <div className="flex-1 flex flex-col justify-center items-center text-center px-6">
                    <div className="text-white mb-8 w-full">
                        <p className="text-base">Welcome to Monime<br/>Please enter code to pay <br/>- - -</p>
                    </div>
                </div>

                {/* Bottom Buttons */}
                <div className="flex gap-4 px-6 pb-8 w-full">
                    <Button
                        variant="ghost"
                        className="flex-1 w-full h-14 bg-white hover:bg-gray-100 text-black text-lg font-medium rounded-xl"
                    >
                        Dismiss
                    </Button>
                    <Button
                        className="flex-1 w-full h-14 bg-black hover:bg-gray-800 text-white text-lg font-medium rounded-xl"
                        // onClick={handleReply}
                    >
                        Reply
                    </Button>
                </div>
            </>
        </div>
    );
}
