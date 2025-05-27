import React from 'react';
import { formatDistanceToNow } from 'date-fns';
import type { WebhookRequest } from '../types/webhook';
interface RequestListProps {
  requests: WebhookRequest[];
  selectedRequest: WebhookRequest | null;
  onSelectRequest: (request: WebhookRequest) => void;
}
export function RequestList({
  requests,
  selectedRequest,
  onSelectRequest
}: RequestListProps) {
  const getMethodColor = (method: string) => {
    switch (method) {
      case 'GET':
        return 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200';
      case 'POST':
        return 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200';
      case 'PUT':
        return 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-200';
      case 'DELETE':
        return 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200';
      default:
        return 'bg-gray-100 text-gray-800 dark:bg-gray-900 dark:text-gray-200';
    }
  };
  return <div className="w-80 border-r border-border bg-card">
      <div className="p-4 border-b border-border">
        <h2 className="font-semibold text-sm text-foreground">
          REQUESTS ({requests.length})
        </h2>
        <p className="text-xs text-muted-foreground mt-1">Newest First</p>
      </div>
      <div className="overflow-y-auto">
        {requests.map(request => <div key={request.id} onClick={() => onSelectRequest(request)} className={`p-4 border-b border-border cursor-pointer transition-colors ${selectedRequest?.id === request.id ? 'bg-accent' : 'hover:bg-muted/50'}`}>
            <div className="flex items-center space-x-2 mb-2">
              <span className={`px-2 py-1 rounded text-xs font-medium ${getMethodColor(request.method)}`}>
                {request.method}
              </span>
              <span className="text-xs text-muted-foreground truncate">
                {request.id}
              </span>
            </div>
            <div className="text-xs text-muted-foreground">
              {formatDistanceToNow(request.timestamp, {
            addSuffix: true
          })}
            </div>
            <div className="text-xs text-muted-foreground mt-1">
              {request.size} bytes • {request.responseTime}s
            </div>
          </div>)}
      </div>
    </div>;
}