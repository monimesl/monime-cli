import React, { useState, ReactNode } from "react";
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
import { Check, ChevronsUpDown } from "lucide-react";
import { cn } from "@/lib/utils";

export interface ComboboxItem {
	value: string;
	label: string;
	[key: string]: any; // Allow additional properties
}

export interface ComboboxProps {
	// Core functionality
	items: ComboboxItem[];
	value?: string;
	onValueChange?: (value: string) => void;
	placeholder?: string;
	searchPlaceholder?: string;
	emptyMessage?: string;

	// Trigger customization
	trigger?: ReactNode;
	triggerClassName?: string;
	triggerVariant?:
		| "default"
		| "destructive"
		| "outline"
		| "secondary"
		| "ghost"
		| "link";

	// Popover customization
	popoverContentProps?: React.ComponentProps<typeof PopoverContent>;
	popoverAlign?: "start" | "center" | "end";
	popoverSide?: "top" | "right" | "bottom" | "left";

	// Command customization
	commandProps?: React.ComponentProps<typeof Command>;
	commandInputProps?: React.ComponentProps<typeof CommandInput>;
	commandListProps?: React.ComponentProps<typeof CommandList>;
	commandGroupProps?: React.ComponentProps<typeof CommandGroup>;

	// Render props for custom item rendering
	renderItem?: (item: ComboboxItem, isSelected: boolean) => ReactNode;

	// Additional props
	disabled?: boolean;
	className?: string;
}

export function Combobox({
	items,
	value,
	onValueChange,
	placeholder = "Select item...",
	searchPlaceholder = "Search...",
	emptyMessage = "No items found.",
	trigger,
	triggerClassName,
	triggerVariant = "outline",
	popoverContentProps,
	popoverAlign = "start",
	popoverSide = "bottom",
	commandProps,
	commandInputProps,
	commandListProps,
	commandGroupProps,
	renderItem,
	disabled = false,
	className,
}: ComboboxProps) {
	const [open, setOpen] = useState(false);

	const selectedItem = items.find((item) => item.value === value);

	const handleSelect = (selectedValue: string) => {
		const newValue = selectedValue === value ? "" : selectedValue;
		onValueChange?.(newValue);
		setOpen(false);
	};

	const defaultTrigger = (
		<Button
			variant={triggerVariant}
			role="combobox"
			aria-expanded={open}
			disabled={disabled}
			className={cn("justify-between", triggerClassName)}
		>
			{selectedItem ? selectedItem.label : placeholder}
			<ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
		</Button>
	);

	const defaultRenderItem = (item: ComboboxItem, isSelected: boolean) => (
		<>
			<Check
				className={cn("mr-2 h-4 w-4", isSelected ? "opacity-100" : "opacity-0")}
			/>
			{item.label}
		</>
	);

	return (
		<div className={className}>
			<Popover open={open} onOpenChange={setOpen}>
				<PopoverTrigger asChild>{trigger || defaultTrigger}</PopoverTrigger>
				<PopoverContent
					align={popoverAlign}
					side={popoverSide}
					className={cn("w-[200px] p-0", popoverContentProps?.className)}
					{...popoverContentProps}
				>
					<Command {...commandProps}>
						<CommandInput placeholder={searchPlaceholder} {...commandInputProps} />
						<CommandList {...commandListProps}>
							<CommandEmpty>{emptyMessage}</CommandEmpty>
							<CommandGroup {...commandGroupProps}>
								{items.map((item) => {
									const isSelected = value === item.value;
									return (
										<CommandItem
											key={item.value}
											value={item.value}
											onSelect={handleSelect}
										>
											{renderItem
												? renderItem(item, isSelected)
												: defaultRenderItem(item, isSelected)}
										</CommandItem>
									);
								})}
							</CommandGroup>
						</CommandList>
					</Command>
				</PopoverContent>
			</Popover>
		</div>
	);
}
