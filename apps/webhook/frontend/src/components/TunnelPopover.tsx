import { useState } from "react";
import { Settings } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Switch } from "@/components/ui/switch";
import {
	Popover,
	PopoverContent,
	PopoverTrigger,
} from "@/components/ui/popover";

interface TunnelPopoverProps {
	tunnelUrl: string;
	onTunnelUrlChange: (url: string) => void;
	tunnelLive: boolean;
	onTunnelLiveChange: (live: boolean) => void;
}

export function TunnelPopover({
	tunnelUrl,
	onTunnelUrlChange,
	tunnelLive,
	onTunnelLiveChange,
}: TunnelPopoverProps) {
	const [showTunnelPopover, setShowTunnelPopover] = useState(false);

	return (
		<Popover open={showTunnelPopover} onOpenChange={setShowTunnelPopover}>
			<PopoverTrigger asChild>
				<Button size="sm" className="flex items-center space-x-2">
					<div className="flex items-center space-x-2">
						<div
							className={`w-2 h-2 rounded-full ${
								tunnelLive ? "bg-green-400" : "bg-gray-400"
							}`}
						/>
						<Settings className="h-4 w-4" />
						<span>Tunnel Request</span>
					</div>
				</Button>
			</PopoverTrigger>
			<PopoverContent className="w-80" align="end">
				<div className="space-y-4">
					<div>
						<h3 className="font-semibold text-sm mb-2">Tunnel Requests To</h3>
						<p className="text-xs text-muted-foreground mb-3">
							Forward all webhook requests to your own endpoint
						</p>
					</div>
					{/* Live Switch */}
					<div className="flex items-center justify-between p-3 bg-muted rounded-lg">
						<div className="flex items-center space-x-3">
							<div
								className={`w-3 h-3 rounded-full ${
									tunnelLive ? "bg-green-500" : "bg-gray-400"
								}`}
							/>
							<div>
								<div className="text-sm font-medium">
									{tunnelLive ? "Live" : "Inactive"}
								</div>
								<div className="text-xs text-muted-foreground">
									{tunnelLive ? "Tunneling is active" : "Tunneling is disabled"}
								</div>
							</div>
						</div>
						<Switch checked={tunnelLive} onCheckedChange={onTunnelLiveChange} />
					</div>
					<Input
						type="url"
						placeholder="https://your-api.com/webhook"
						value={tunnelUrl}
						onChange={(e) => onTunnelUrlChange(e.target.value)}
					/>
					<div className="flex justify-end space-x-2">
						<Button
							variant="ghost"
							size="sm"
							onClick={() => setShowTunnelPopover(false)}
						>
							Cancel
						</Button>
						<Button size="sm" onClick={() => setShowTunnelPopover(false)}>
							Save
						</Button>
					</div>
				</div>
			</PopoverContent>
		</Popover>
	);
}
