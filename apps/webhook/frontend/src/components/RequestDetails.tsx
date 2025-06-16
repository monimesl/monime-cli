import React, { useState } from "react";
import { Copy, Code2, FileText, RotateCcw } from "lucide-react";
import { Button } from "@/components/ui/button";
import JsonView from "@uiw/react-json-view";
import { lightTheme } from "@uiw/react-json-view/light";
import { darkTheme } from "@uiw/react-json-view/dark";
import type { WebhookRequest } from "../types/webhook";

interface RequestDetailsProps {
	request: WebhookRequest | null;
	onReTrigger: (request: WebhookRequest) => void;
}

export function RequestDetails({ request, onReTrigger }: RequestDetailsProps) {
	const [activeTab, setActiveTab] = useState("body");
	const [bodyFormat, setBodyFormat] = useState<"pretty" | "raw">("pretty");

	const tabs = [
		{
			id: "body",
			label: "Body",
		},
		{
			id: "summary",
			label: "Summary",
		},
	];

	if (!request) {
		return (
			<div className="flex-1 flex items-center justify-center text-muted-foreground">
				<p>Select a request to view details</p>
			</div>
		);
	}

	const copyToClipboard = async (text: string) => {
		try {
			await navigator.clipboard.writeText(text);
		} catch (err) {
			console.error("Failed to copy:", err);
		}
	};

	const getMethodColor = (method: string) => {
		switch (method) {
			case "GET":
				return "bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200";
			case "POST":
				return "bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200";
			case "PUT":
				return "bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-200";
			case "DELETE":
				return "bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200";
			default:
				return "bg-gray-100 text-gray-800 dark:bg-gray-900 dark:text-gray-200";
		}
	};

	// Check if the body is valid JSON
	const isValidJson = (str: string) => {
		try {
			JSON.parse(str);
			return true;
		} catch {
			return false;
		}
	};

	// Detect if dark mode is active
	const isDarkMode = document.documentElement.classList.contains("dark");

	const renderJsonBody = () => {
		if (!request.body) return null;

		if (bodyFormat === "raw") {
			return (
				<pre className="text-sm font-mono whitespace-pre-wrap break-words p-6 overflow-x-auto min-h-[300px]">
					<code>{request.body}</code>
				</pre>
			);
		}

		// Pretty format with JsonView
		if (isValidJson(request.body)) {
			try {
				const jsonData = JSON.parse(request.body);
				return (
					<div className="p-4">
						<JsonView
							value={jsonData}
							style={isDarkMode ? darkTheme : lightTheme}
							collapsed={2}
							displayDataTypes={false}
							enableClipboard={false}
						/>
					</div>
				);
			} catch {
				// Fallback to raw if parsing fails
				return (
					<pre className="text-sm font-mono whitespace-pre-wrap break-words p-6 overflow-x-auto min-h-[300px]">
						<code>{request.body}</code>
					</pre>
				);
			}
		}

		// Non-JSON content
		return (
			<pre className="text-sm font-mono whitespace-pre-wrap break-words p-6 overflow-x-auto min-h-[300px]">
				<code>{request.body}</code>
			</pre>
		);
	};

	return (
		<div className="flex-1 flex flex-col">
			{/* Tab Navigation */}
			<div className="border-b border-border bg-card">
				<div className="flex items-center justify-between p-4">
					<div className="flex space-x-1">
						{tabs.map((tab) => (
							<Button
								key={tab.id}
								variant={activeTab === tab.id ? "default" : "ghost"}
								size="sm"
								onClick={() => setActiveTab(tab.id)}
							>
								{tab.label}
							</Button>
						))}
					</div>
					{/* Re Trigger Button */}
					<Button
						variant="secondary"
						size="sm"
						onClick={() => onReTrigger(request)}
						className="flex items-center space-x-2"
					>
						<RotateCcw className="h-4 w-4" />
						<span>Re Trigger</span>
					</Button>
				</div>
			</div>

			{/* Tab Content */}
			<div className="flex-1 overflow-y-auto">
				<div className="p-6">
					{/* Body Tab */}
					{activeTab === "body" && (
						<div className="space-y-4">
							<div className="flex items-center justify-between">
								<h2 className="text-xl font-semibold">Request Body</h2>
								{request.body && (
									<div className="flex items-center space-x-2">
										<Button
											variant={bodyFormat === "pretty" ? "default" : "outline"}
											size="sm"
											onClick={() => setBodyFormat("pretty")}
											className="flex items-center space-x-1"
											disabled={!isValidJson(request.body)}
										>
											<Code2 className="h-3 w-3" />
											<span>Pretty</span>
										</Button>
										<Button
											variant={bodyFormat === "raw" ? "default" : "outline"}
											size="sm"
											onClick={() => setBodyFormat("raw")}
											className="flex items-center space-x-1"
										>
											<FileText className="h-3 w-3" />
											<span>Raw</span>
										</Button>
										<Button
											variant="outline"
											size="sm"
											onClick={() => copyToClipboard(request.body || "")}
											className="flex items-center space-x-1"
											title="Copy body"
										>
											<Copy className="h-3 w-3" />
											<span>Copy</span>
										</Button>
									</div>
								)}
							</div>
							<div className="bg-card rounded-lg border shadow-sm">
								{request.body ? (
									renderJsonBody()
								) : (
									<div className="p-6 text-center text-muted-foreground">
										<p>No request body</p>
									</div>
								)}
							</div>
						</div>
					)}

					{/* Summary Tab */}
					{activeTab === "summary" && (
						<div className="space-y-8">
							{/* Overview Section */}
							<div>
								<h2 className="text-xl font-semibold mb-4">Request Overview</h2>
								<div className="bg-card rounded-lg border shadow-sm p-6">
									<div className="grid grid-cols-1 md:grid-cols-2 gap-6">
										<div className="space-y-4">
											<div>
												<label className="text-sm font-medium text-muted-foreground">
													Method
												</label>
												<div className="mt-1">
													<span
														className={`px-3 py-1 rounded-md text-sm font-medium ${getMethodColor(
															request.method
														)}`}
													>
														{request.method}
													</span>
												</div>
											</div>
											<div>
												<label className="text-sm font-medium text-muted-foreground">
													URL
												</label>
												<div className="mt-1 font-mono text-sm text-blue-600 dark:text-blue-400 break-all">
													{request.url}
												</div>
											</div>
											<div>
												<label className="text-sm font-medium text-muted-foreground">
													Request ID
												</label>
												<div className="mt-1 font-mono text-sm">{request.id}</div>
											</div>
										</div>
										<div className="space-y-4">
											<div>
												<label className="text-sm font-medium text-muted-foreground">
													Timestamp
												</label>
												<div className="mt-1 text-sm">
													{request.timestamp.toLocaleString()}
												</div>
											</div>
											<div>
												<label className="text-sm font-medium text-muted-foreground">
													IP Address
												</label>
												<div className="mt-1 font-mono text-sm">{request.ip}</div>
											</div>
											<div className="grid grid-cols-2 gap-4">
												<div>
													<label className="text-sm font-medium text-muted-foreground">
														Size
													</label>
													<div className="mt-1 text-sm">{request.size} bytes</div>
												</div>
												<div>
													<label className="text-sm font-medium text-muted-foreground">
														Response Time
													</label>
													<div className="mt-1 text-sm">{request.responseTime}s</div>
												</div>
											</div>
										</div>
									</div>
								</div>
							</div>

							{/* User Agent Section */}
							<div>
								<label className="text-sm font-medium text-muted-foreground">
									User Agent
								</label>
								<div className="mt-2 bg-card rounded-lg border shadow-sm p-4">
									<div className="font-mono text-sm break-all">{request.userAgent}</div>
								</div>
							</div>

							{/* Headers Section */}
							<div>
								<div className="flex items-center justify-between mb-4">
									<h2 className="text-xl font-semibold">Request Headers</h2>
									<Button
										variant="outline"
										size="sm"
										onClick={() =>
											copyToClipboard(JSON.stringify(request.headers, null, 2))
										}
										className="flex items-center space-x-1"
									>
										<Copy className="h-3 w-3" />
										<span>Copy All</span>
									</Button>
								</div>
								<div className="bg-card rounded-lg border shadow-sm p-6">
									<div className="space-y-4">
										{Object.entries(request.headers).map(([key, value]) => (
											<div
												key={key}
												className="flex flex-col space-y-1 pb-4 border-b border-border last:border-b-0 last:pb-0"
											>
												<div className="text-sm font-medium text-muted-foreground">
													{key}
												</div>
												<div className="font-mono text-sm break-all bg-muted rounded px-2 py-1">
													{value}
												</div>
											</div>
										))}
									</div>
								</div>
							</div>
						</div>
					)}
				</div>
			</div>
		</div>
	);
}
