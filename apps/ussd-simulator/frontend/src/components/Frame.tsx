import { useState } from "react";
import {
	Select,
	SelectContent,
	SelectItem,
	SelectTrigger,
	SelectValue,
} from "@/components/ui/select";

const networks = {
	m17: "Orange",
	m18: "Africell",
};

type Network = "m17" | "m18";

const getNetwork = (network: Network) => {
	return networks[network];
};

export default function Frame({
	children,
	frameColor = "light",
}: {
	children: React.ReactNode;
	frameColor?: "dark" | "light";
}) {
	const [network, setNetwork] = useState<Network>("m17");
	return (
		<div
			style={{
				width: "380px",
				height: "720px",
				background: frameColor === "dark" ? "#374151" : "white",
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
				<div style={{ color: frameColor === "dark" ? "white" : "black" }}>
					9:41 &nbsp;
				</div>
				<Select
					value={network}
					onValueChange={(value) => setNetwork(value as Network)}
				>
					<SelectTrigger
						style={{
							border: "none",
							background: "none",
							outline: "none",
							color: frameColor === "dark" ? "white" : "black",
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
								color: frameColor === "dark" ? "white" : "black",
							}}
						>
							{getNetwork(network)} 📶
						</SelectValue>
					</SelectTrigger>
					<SelectContent>
						<SelectItem value="m17">Orange</SelectItem>
						<SelectItem value="m18">Africell</SelectItem>
					</SelectContent>
				</Select>
				{/* <div
					style={{
						border: "none",
						background: "none",
						outline: "none",
						color: frameColor === "dark" ? "white" : "black",
					}}
				></div> */}
			</div>
			{children}
		</div>
	);
}
