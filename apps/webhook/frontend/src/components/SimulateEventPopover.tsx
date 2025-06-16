import { useState } from "react";
import { Play, ChevronDown } from "lucide-react";
import { Button } from "@/components/ui/button";
import {
	Popover,
	PopoverContent,
	PopoverTrigger,
} from "@/components/ui/popover";
import {
	Command,
	CommandEmpty,
	CommandGroup,
	CommandInput,
	CommandItem,
	CommandList,
} from "@/components/ui/command";

interface SimulationEvent {
	id: string;
	label: string;
	description: string;
}

const simulationEvents: SimulationEvent[] = [
	{
		id: "user.created",
		label: "User Created",
		description: "Simulate a new user registration",
	},
	{
		id: "payment.completed",
		label: "Payment Completed",
		description: "Simulate a successful payment",
	},
	{
		id: "order.shipped",
		label: "Order Shipped",
		description: "Simulate an order shipment notification",
	},
	{
		id: "subscription.cancelled",
		label: "Subscription Cancelled",
		description: "Simulate a subscription cancellation",
	},
	{
		id: "file.uploaded",
		label: "File Uploaded",
		description: "Simulate a file upload event",
	},
	{
		id: "email.bounced",
		label: "Email Bounced",
		description: "Simulate an email bounce notification",
	},
];

interface SimulateEventPopoverProps {
	onSimulate?: (eventId: string) => void;
}

export function SimulateEventPopover({
	onSimulate,
}: SimulateEventPopoverProps) {
	const [showSimulatePopover, setShowSimulatePopover] = useState(false);
	const [selectedEvent, setSelectedEvent] = useState<string>("");

	const handleSimulate = () => {
		if (selectedEvent) {
			onSimulate?.(selectedEvent);
			setShowSimulatePopover(false);
			setSelectedEvent("");
		}
	};

	const handleCancel = () => {
		setShowSimulatePopover(false);
		setSelectedEvent("");
	};

	const handleSelectEvent = (eventId: string) => {
		setSelectedEvent(eventId === selectedEvent ? "" : eventId);
	};

	const selectedEventData = simulationEvents.find(
		(event) => event.id === selectedEvent
	);

	return (
		<Popover open={showSimulatePopover} onOpenChange={setShowSimulatePopover}>
			<PopoverTrigger asChild>
				<Button
					variant="secondary"
					size="sm"
					className="flex items-center space-x-2"
				>
					<Play className="h-4 w-4" />
					<span>Simulate Event</span>
					<ChevronDown className="h-4 w-4" />
				</Button>
			</PopoverTrigger>
			<PopoverContent className="w-96 p-0" align="end">
				<div className="space-y-4 p-4">
					<div>
						<h3 className="font-semibold text-sm mb-2">Simulate Webhook Event</h3>
						<p className="text-xs text-muted-foreground mb-4">
							Send a test webhook event to your endpoint
						</p>
					</div>

					{/* Command Component for Event Selection */}
					<Command className="border rounded-lg">
						<CommandInput placeholder="Search events..." />
						<CommandList className="max-h-48">
							<CommandEmpty>No events found.</CommandEmpty>
							<CommandGroup>
								{simulationEvents.map((event) => (
									<CommandItem
										key={event.id}
										value={`${event.label} ${event.description} ${event.id}`}
										onSelect={() => handleSelectEvent(event.id)}
										className={`cursor-pointer ${
											selectedEvent === event.id ? "bg-accent" : ""
										}`}
									>
										<div className="flex items-center space-x-2 w-full">
											<div
												className={`w-2 h-2 rounded-full ${
													selectedEvent === event.id
														? "bg-primary"
														: "bg-transparent border border-border"
												}`}
											/>
											<div className="flex-1 text-left">
												<div className="font-medium text-sm">{event.label}</div>
												<div className="text-xs text-muted-foreground">
													{event.description}
												</div>
											</div>
										</div>
									</CommandItem>
								))}
							</CommandGroup>
						</CommandList>
					</Command>

					{/* Selected Event Preview */}
					{selectedEventData && (
						<div className="p-3 bg-muted rounded-lg">
							<div className="text-sm font-medium">{selectedEventData.label}</div>
							<div className="text-xs text-muted-foreground mt-1">
								Event ID:{" "}
								<code className="bg-background px-1 rounded">
									{selectedEventData.id}
								</code>
							</div>
						</div>
					)}

					{/* Action Buttons */}
					<div className="flex justify-end space-x-2 pt-2 border-t border-border">
						<Button variant="ghost" size="sm" onClick={handleCancel}>
							Cancel
						</Button>
						<Button size="sm" onClick={handleSimulate} disabled={!selectedEvent}>
							Simulate
						</Button>
					</div>
				</div>
			</PopoverContent>
		</Popover>
	);
}
