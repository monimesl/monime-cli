import { useEffect, useState } from "react";
import { WebhookHeader } from "./components/WebhookHeader";
import { RequestList } from "./components/RequestList";
import { RequestDetails } from "./components/RequestDetails";
import { generateWebhookId, mockRequests } from "./utils/webhook";
import type { WebhookRequest } from "./types/webhook";

export function App() {
	const [webhookId] = useState(() => generateWebhookId());
	const [requests, setRequests] = useState<WebhookRequest[]>(mockRequests);
	const [selectedRequest, setSelectedRequest] = useState<WebhookRequest | null>(
		requests[0] || null
	);
	const [tunnelUrl, setTunnelUrl] = useState("");
	const [tunnelLive, setTunnelLive] = useState(false);
	// In a real app, this would come from your workspace/project context
	const [spaceName] = useState("Lajor");

	useEffect(() => {
		const fetchTunnelUrl = async () => {
			const response = await fetch(`${import.meta.env.VITE_API_URL}/tunnel/url`);
			const data = await response.json();
			setTunnelUrl(data.url);
		};
		fetchTunnelUrl();
	}, []);

	const webhookUrl = `https://webhook-tester.com/${webhookId}`;

	const handleReTrigger = (request: WebhookRequest) => {
		console.log("Re-triggering request:", request.id);
		// Here you would implement the actual re-trigger logic
	};

	return (
		<div className="w-full min-h-screen bg-background flex flex-col">
			<WebhookHeader
				webhookUrl={webhookUrl}
				tunnelUrl={tunnelUrl}
				onTunnelUrlChange={setTunnelUrl}
				tunnelLive={tunnelLive}
				onTunnelLiveChange={setTunnelLive}
			/>
			<div className="flex flex-1 overflow-hidden">
				<RequestList
					requests={requests}
					selectedRequest={selectedRequest}
					onSelectRequest={setSelectedRequest}
					spaceName={spaceName}
				/>
				<RequestDetails request={selectedRequest} onReTrigger={handleReTrigger} />
			</div>
		</div>
	);
}
