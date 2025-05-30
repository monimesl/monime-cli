import { useState } from "react";
import { Copy, Check } from "lucide-react";
import { Button } from "@/components/ui/button";
import { TunnelPopover } from "./TunnelPopover";
import { SimulateEventPopover } from "./SimulateEventPopover";
import { EventsPopover } from "./EventsPopover";

interface WebhookHeaderProps {
	webhookUrl: string;
	tunnelUrl: string;
	onTunnelUrlChange: (url: string) => void;
	tunnelLive: boolean;
	onTunnelLiveChange: (live: boolean) => void;
}

export function WebhookHeader({
	webhookUrl,
	tunnelUrl,
	onTunnelUrlChange,
	tunnelLive,
	onTunnelLiveChange,
}: WebhookHeaderProps) {
	const [copied, setCopied] = useState(false);
	const [selectedEvents, setSelectedEvents] = useState<string[]>([]);

	const copyToClipboard = async () => {
		try {
			await navigator.clipboard.writeText(webhookUrl);
			setCopied(true);
			setTimeout(() => setCopied(false), 2000);
		} catch (err) {
			console.error("Failed to copy URL:", err);
		}
	};

	const handleSimulate = (eventId: string) => {
		console.log("Simulating event:", eventId);
		// Here you would trigger the actual simulation
	};

	const handleEventsChange = (events: string[]) => {
		setSelectedEvents(events);
		console.log("Selected events updated:", events);
		// Here you would update the webhook configuration
	};

	return (
		<header className="bg-card border-b border-border p-2">
			<div className="flex items-center justify-between">
				<div className="flex items-center space-x-4">
					<div className="flex items-center space-x-2">
						<img src="/monime-logo.png" alt="Monime Logo" className="w-8 h-8" />
						<h1 className="text-lg font-semibold text-foreground">Monime Webhook</h1>
					</div>
				</div>
				<div className="flex items-center space-x-2">
					<div className="flex items-center space-x-2 bg-muted rounded-lg p-2">
						<code className="text-sm font-mono text-muted-foreground truncate max-w-[200px]">
							{webhookUrl}
						</code>
						<Button
							variant="ghost"
							size="sm"
							onClick={copyToClipboard}
							className="p-1 h-auto"
							title="Copy webhook URL"
						>
							{copied ? (
								<Check className="h-4 w-4 text-green-600" />
							) : (
								<Copy className="h-4 w-4 text-muted-foreground" />
							)}
						</Button>
					</div>
					<EventsPopover
						selectedEvents={selectedEvents}
						onEventsChange={handleEventsChange}
					/>
					<SimulateEventPopover onSimulate={handleSimulate} />
					<TunnelPopover
						tunnelUrl={tunnelUrl}
						onTunnelUrlChange={onTunnelUrlChange}
						tunnelLive={tunnelLive}
						onTunnelLiveChange={onTunnelLiveChange}
					/>
				</div>
			</div>
		</header>
	);
}
