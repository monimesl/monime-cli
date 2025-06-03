"use client";

import { useState } from "react";
import { PhoneCall, Delete } from "lucide-react";
import { Button } from "@/components/ui/button";
import Frame from "./Frame";

export default function PhoneKeyPadScreen() {
	const [dialedNumber, setDialedNumber] = useState("");

	const dialPadNumbers = [
		{ number: "1", letters: "" },
		{ number: "2", letters: "ABC" },
		{ number: "3", letters: "DEF" },
		{ number: "4", letters: "GHI" },
		{ number: "5", letters: "JKL" },
		{ number: "6", letters: "MNO" },
		{ number: "7", letters: "PQRS" },
		{ number: "8", letters: "TUV" },
		{ number: "9", letters: "WXYZ" },
		{ number: "*", letters: "" },
		{ number: "0", letters: "+" },
		{ number: "#", letters: "" },
	];

	const handleNumberPress = (number: string) => {
		setDialedNumber((prev) => prev + number);
	};

	const handleDelete = () => {
		setDialedNumber((prev) => prev.slice(0, -1));
	};

	const handleCall = () => {
		if (dialedNumber) {
			alert(`Calling ${dialedNumber}...`);
		}
	};

	return (
		<Frame>
			<div className="flex-1 flex flex-col px-4 py-6">
				{/* Dialed Number Display */}
				<div className="text-center mb-8 mt-4">
					<input
						type="tel"
						value={dialedNumber}
						autoFocus
						onChange={(e) => {
							// Filter to only allow valid phone number characters
							const filtered = e.target.value.replace(/[^0-9*#+\-() ]/g, "");
							setDialedNumber(filtered);
						}}
						onPaste={(e) => {
							e.preventDefault();
							const pastedText = e.clipboardData.getData("text");
							// Filter pasted content to only include valid phone number characters
							const filtered = pastedText.replace(/[^0-9*#+\-() ]/g, "");
							setDialedNumber((prev) => prev + filtered);
						}}
						className="text-4xl font-bold text-black min-h-[40px] bg-transparent border-none outline-none text-center w-full placeholder-gray-400 rounded-lg px-2 transition-colors"
						autoComplete="tel"
					/>
				</div>

				{/* Dial Pad */}
				<div className="flex-1 flex flex-col justify-center">
					<div className="grid grid-cols-3 place-items-center justify-center gap-4 mx-auto w-full">
						{dialPadNumbers.map((item, index) => (
							<Button
								key={index}
								variant="ghost"
								className="h-20 w-20 rounded-full bg-gray-100 hover:bg-gray-200 flex flex-col items-center justify-center text-black border-0 shadow-sm"
								onClick={() => handleNumberPress(item.number)}
							>
								<span className="text-2xl font-normal">{item.number}</span>
								{/* {item.letters && (
									<span className="text-xs text-gray-600 mt-1">{item.letters}</span>
								)} */}
							</Button>
						))}
					</div>

					{/* Action Buttons */}
					<div className="flex justify-center items-center gap-6 mt-6">
						<div className="w-12 h-12"></div> {/* Spacer for symmetry */}
						<Button
							className="w-20 h-20 rounded-full bg-green-500 hover:bg-green-600"
							onClick={handleCall}
							disabled={!dialedNumber}
						>
							<PhoneCall className="w-6 h-6 text-white" />
						</Button>
						<Button
							variant="ghost"
							size="icon"
							className="w-16 h-16 rounded-full bg-gray-100 hover:bg-gray-200"
							onClick={handleDelete}
							disabled={!dialedNumber}
						>
							<Delete className="w-5 h-5 text-gray-600" />
						</Button>
					</div>
				</div>
			</div>
		</Frame>
	);
}
