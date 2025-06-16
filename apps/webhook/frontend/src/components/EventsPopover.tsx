import { useState } from "react";
import { Settings, ChevronDown } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
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

interface WebhookEvent {
	id: string;
	label: string;
	description: string;
}

const availableEvents: WebhookEvent[] = [
	{
		id: "user.created",
		label: "User Created",
		description: "Triggered when a new user registers",
	},
	{
		id: "user.updated",
		label: "User Updated",
		description: "Triggered when user profile is updated",
	},
	{
		id: "user.deleted",
		label: "User Deleted",
		description: "Triggered when a user account is deleted",
	},
	{
		id: "payment.completed",
		label: "Payment Completed",
		description: "Triggered when a payment is successfully processed",
	},
	{
		id: "payment.failed",
		label: "Payment Failed",
		description: "Triggered when a payment fails",
	},
	{
		id: "payment.refunded",
		label: "Payment Refunded",
		description: "Triggered when a payment is refunded",
	},
	{
		id: "order.created",
		label: "Order Created",
		description: "Triggered when a new order is placed",
	},
	{
		id: "order.shipped",
		label: "Order Shipped",
		description: "Triggered when an order is shipped",
	},
	{
		id: "order.delivered",
		label: "Order Delivered",
		description: "Triggered when an order is delivered",
	},
	{
		id: "order.cancelled",
		label: "Order Cancelled",
		description: "Triggered when an order is cancelled",
	},
	{
		id: "subscription.created",
		label: "Subscription Created",
		description: "Triggered when a new subscription is created",
	},
	{
		id: "subscription.updated",
		label: "Subscription Updated",
		description: "Triggered when subscription details are updated",
	},
	{
		id: "subscription.cancelled",
		label: "Subscription Cancelled",
		description: "Triggered when a subscription is cancelled",
	},
	{
		id: "file.uploaded",
		label: "File Uploaded",
		description: "Triggered when a file is uploaded",
	},
	{
		id: "file.deleted",
		label: "File Deleted",
		description: "Triggered when a file is deleted",
	},
	{
		id: "email.sent",
		label: "Email Sent",
		description: "Triggered when an email is sent",
	},
	{
		id: "email.bounced",
		label: "Email Bounced",
		description: "Triggered when an email bounces",
	},
	{
		id: "notification.sent",
		label: "Notification Sent",
		description: "Triggered when a notification is sent",
	},
];

interface EventsPopoverProps {
	selectedEvents?: string[];
	onEventsChange?: (events: string[]) => void;
}

export function EventsPopover({
	selectedEvents = [],
	onEventsChange,
}: EventsPopoverProps) {
	const [showEventsPopover, setShowEventsPopover] = useState(false);
	const [tempSelectedEvents, setTempSelectedEvents] =
		useState<string[]>(selectedEvents);

	const handleEventToggle = (eventId: string) => {
		setTempSelectedEvents((prev) =>
			prev.includes(eventId)
				? prev.filter((id) => id !== eventId)
				: [...prev, eventId]
		);
	};

	const handleDone = () => {
		onEventsChange?.(tempSelectedEvents);
		setShowEventsPopover(false);
	};

	const handleCancel = () => {
		setTempSelectedEvents(selectedEvents);
		setShowEventsPopover(false);
	};

	const selectedCount = tempSelectedEvents.length;

	return (
		<Popover open={showEventsPopover} onOpenChange={setShowEventsPopover}>
			<PopoverTrigger asChild>
				<Button
					variant="secondary"
					size="sm"
					className="flex items-center space-x-2"
				>
					<Settings className="h-4 w-4" />
					<span>Events</span>
					{selectedCount > 0 && (
						<span className="bg-primary text-primary-foreground text-xs px-1.5 py-0.5 rounded-full min-w-[1.25rem] text-center">
							{selectedCount}
						</span>
					)}
					<ChevronDown className="h-4 w-4" />
				</Button>
			</PopoverTrigger>
			<PopoverContent className="w-96 p-0" align="end">
				<div className="space-y-4 p-4">
					<div>
						<h3 className="font-semibold text-sm mb-2">Configure Webhook Events</h3>
						<p className="text-xs text-muted-foreground mb-4">
							Select which events should trigger this webhook
						</p>
					</div>

					{/* Command Component for Event Selection */}
					<Command className="border rounded-lg">
						<CommandInput placeholder="Search events..." />
						<CommandList className="max-h-64">
							<CommandEmpty>No events found.</CommandEmpty>
							<CommandGroup>
								{availableEvents.map((event) => {
									const isSelected = tempSelectedEvents.includes(event.id);
									return (
										<CommandItem
											key={event.id}
											value={`${event.label} ${event.description} ${event.id}`}
											onSelect={() => handleEventToggle(event.id)}
											className="cursor-pointer"
										>
											<div className="flex items-center space-x-3 w-full">
												<Checkbox
													checked={isSelected}
													onCheckedChange={() => handleEventToggle(event.id)}
													className="flex-shrink-0"
												/>
												<div className="flex-1 text-left">
													<div className="font-medium text-sm">{event.label}</div>
													<div className="text-xs text-muted-foreground">
														{event.description}
													</div>
												</div>
											</div>
										</CommandItem>
									);
								})}
							</CommandGroup>
						</CommandList>
					</Command>

					{/* Selected Events Summary */}
					{selectedCount > 0 && (
						<div className="p-3 bg-muted rounded-lg">
							<div className="text-sm font-medium mb-1">
								{selectedCount} event{selectedCount !== 1 ? "s" : ""} selected
							</div>
							<div className="text-xs text-muted-foreground">
								{tempSelectedEvents
									.slice(0, 3)
									.map((eventId) => {
										const event = availableEvents.find((e) => e.id === eventId);
										return event?.label;
									})
									.join(", ")}
								{selectedCount > 3 && ` and ${selectedCount - 3} more`}
							</div>
						</div>
					)}

					{/* Action Buttons */}
					<div className="flex justify-end space-x-2 pt-2 border-t border-border">
						<Button variant="ghost" size="sm" onClick={handleCancel}>
							Cancel
						</Button>
						<Button size="sm" onClick={handleDone}>
							Done
						</Button>
					</div>
				</div>
			</PopoverContent>
		</Popover>
	);
}
