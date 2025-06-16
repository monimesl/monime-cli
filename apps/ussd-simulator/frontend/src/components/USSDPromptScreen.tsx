"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import Frame from "./Frame";

export default function USSDPromptScreen() {
	const [showReplyScreen, setShowReplyScreen] = useState(false);
	const [replyMessage, setReplyMessage] = useState("");
	const maxCharacters = 182;

	const handleReply = () => {
		setShowReplyScreen(true);
	};

	const handleCancel = () => {
		setShowReplyScreen(false);
		setReplyMessage("");
	};

	const handleSendReply = () => {
		if (replyMessage.trim()) {
			alert(`Reply sent: ${replyMessage}`);
			setShowReplyScreen(false);
			setReplyMessage("");
		}
	};

	const remainingCharacters = maxCharacters - replyMessage.length;

	return (
		<Frame frameColor="dark">
			<div className="flex items-center flex-col justify-center h-full bg-gray-700 p-0 w-full">
				{showReplyScreen ? (
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
									disabled={!replyMessage.trim()}
								>
									Reply
								</Button>
							</div>
						</div>

						{/* Message Content */}
						<div className="flex-1 px-4 py-6 text-center">
							<div className="text-white mt-10">
								<div className="text-lg font-medium mb-4">Monime</div>
								<div className="text-sm leading-relaxed">
									You are about to pay Sulailas Multi Cuisine Restaurant
									<br />
									1:Confirm
									<br />
									2:Cancel
									<br />
									<br />- - -
									<br />
									<br />
									00:menu
									<br />
									0:back
								</div>
							</div>
						</div>

						{/* Input Field */}
						<div className="px-4 mb-4 w-full">
							<textarea
								value={replyMessage}
								onChange={(e) => setReplyMessage(e.target.value)}
								placeholder="Type your reply..."
								className="w-full h-16 bg-gray-600 border border-gray-500 rounded-lg px-3 py-2 text-white placeholder-gray-400 resize-none focus:outline-none focus:border-blue-400"
								maxLength={maxCharacters}
							/>
							<div className="text-gray-400 text-sm mt-1 text-center">
								{remainingCharacters} characters remaining
							</div>
						</div>
						<div className="h-40" />
					</>
				) : (
					<>
						{/* Main Content */}
						<div className="flex-1 flex flex-col justify-center items-center text-center px-6">
							<div className="text-white mb-8 w-full">
								<h1 className="text-2xl font-light mb-4">Welcome to Monime</h1>
								<p className="text-lg">Please enter code to pay</p>
								<br />- - -<br />
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
								onClick={handleReply}
							>
								Reply
							</Button>
						</div>
					</>
				)}
			</div>
		</Frame>
	);
}
