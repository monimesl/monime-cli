import React, { useState } from 'react';
import { Copy, Settings, Check, Play, ChevronDown } from 'lucide-react';
interface WebhookHeaderProps {
  webhookUrl: string;
  tunnelUrl: string;
  onTunnelUrlChange: (url: string) => void;
  tunnelLive: boolean;
  onTunnelLiveChange: (live: boolean) => void;
}
const simulationEvents = [{
  id: 'user.created',
  label: 'User Created',
  description: 'Simulate a new user registration'
}, {
  id: 'payment.completed',
  label: 'Payment Completed',
  description: 'Simulate a successful payment'
}, {
  id: 'order.shipped',
  label: 'Order Shipped',
  description: 'Simulate an order shipment notification'
}, {
  id: 'subscription.cancelled',
  label: 'Subscription Cancelled',
  description: 'Simulate a subscription cancellation'
}, {
  id: 'file.uploaded',
  label: 'File Uploaded',
  description: 'Simulate a file upload event'
}, {
  id: 'email.bounced',
  label: 'Email Bounced',
  description: 'Simulate an email bounce notification'
}];
export function WebhookHeader({
  webhookUrl,
  tunnelUrl,
  onTunnelUrlChange,
  tunnelLive,
  onTunnelLiveChange
}: WebhookHeaderProps) {
  const [copied, setCopied] = useState(false);
  const [showTunnelPopover, setShowTunnelPopover] = useState(false);
  const [showSimulatePopover, setShowSimulatePopover] = useState(false);
  const [selectedEvent, setSelectedEvent] = useState<string>('');
  const [searchTerm, setSearchTerm] = useState('');
  const copyToClipboard = async () => {
    try {
      await navigator.clipboard.writeText(webhookUrl);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    } catch (err) {
      console.error('Failed to copy URL:', err);
    }
  };
  const filteredEvents = simulationEvents.filter(event => event.label.toLowerCase().includes(searchTerm.toLowerCase()) || event.id.toLowerCase().includes(searchTerm.toLowerCase()));
  const handleSimulate = () => {
    if (selectedEvent) {
      console.log('Simulating event:', selectedEvent);
      // Here you would trigger the actual simulation
      setShowSimulatePopover(false);
      setSelectedEvent('');
      setSearchTerm('');
    }
  };
  return <header className="bg-card border-b border-border p-4">
      <div className="flex items-center justify-between">
        <div className="flex items-center space-x-4">
          <h1 className="text-xl font-semibold text-foreground">
            Webhook Tester
          </h1>
          <div className="flex items-center space-x-2 bg-muted rounded-lg p-2">
            <code className="text-sm font-mono text-muted-foreground">
              {webhookUrl.split('/').pop()}
            </code>
            <button onClick={copyToClipboard} className="p-1 hover:bg-background rounded transition-colors" title="Copy webhook URL">
              {copied ? <Check className="h-4 w-4 text-green-600" /> : <Copy className="h-4 w-4 text-muted-foreground" />}
            </button>
          </div>
        </div>
        <div className="flex items-center space-x-2">
          {/* Simulate Event Button */}
          <div className="relative">
            <button onClick={() => setShowSimulatePopover(!showSimulatePopover)} className="flex items-center space-x-2 bg-secondary text-secondary-foreground px-4 py-2 rounded-lg hover:bg-secondary/90 transition-colors">
              <Play className="h-4 w-4" />
              <span>Simulate Event</span>
              <ChevronDown className="h-4 w-4" />
            </button>
            {showSimulatePopover && <div className="absolute right-0 top-full mt-2 w-96 bg-popover border border-border rounded-lg shadow-lg p-4 z-50">
                <h3 className="font-semibold text-sm mb-2">
                  Simulate Webhook Event
                </h3>
                <p className="text-xs text-muted-foreground mb-4">
                  Send a test webhook event to your endpoint
                </p>
                <div className="space-y-4">
                  {/* Search Input */}
                  <input type="text" placeholder="Search events..." value={searchTerm} onChange={e => setSearchTerm(e.target.value)} className="w-full px-3 py-2 border border-input rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-ring" />
                  {/* Event List */}
                  <div className="max-h-48 overflow-y-auto space-y-1">
                    {filteredEvents.map(event => <button key={event.id} onClick={() => setSelectedEvent(event.id)} className={`w-full text-left p-3 rounded-md transition-colors ${selectedEvent === event.id ? 'bg-accent text-accent-foreground' : 'hover:bg-muted'}`}>
                        <div className="font-medium text-sm">{event.label}</div>
                        <div className="text-xs text-muted-foreground">
                          {event.description}
                        </div>
                      </button>)}
                  </div>
                  {/* Action Buttons */}
                  <div className="flex justify-end space-x-2 pt-2 border-t border-border">
                    <button onClick={() => {
                  setShowSimulatePopover(false);
                  setSelectedEvent('');
                  setSearchTerm('');
                }} className="px-3 py-1 text-sm text-muted-foreground hover:text-foreground">
                      Cancel
                    </button>
                    <button onClick={handleSimulate} disabled={!selectedEvent} className="px-3 py-1 text-sm bg-primary text-primary-foreground rounded hover:bg-primary/90 disabled:opacity-50 disabled:cursor-not-allowed">
                      Simulate
                    </button>
                  </div>
                </div>
              </div>}
          </div>
          {/* Tunnel Request Button */}
          <div className="relative">
            <button onClick={() => setShowTunnelPopover(!showTunnelPopover)} className="flex items-center space-x-2 bg-primary text-primary-foreground px-4 py-2 rounded-lg hover:bg-primary/90 transition-colors">
              <div className="flex items-center space-x-2">
                <div className={`w-2 h-2 rounded-full ${tunnelLive ? 'bg-green-400' : 'bg-gray-400'}`} />
                <Settings className="h-4 w-4" />
                <span>Tunnel Request</span>
              </div>
            </button>
            {showTunnelPopover && <div className="absolute right-0 top-full mt-2 w-80 bg-popover border border-border rounded-lg shadow-lg p-4 z-50">
                <h3 className="font-semibold text-sm mb-2">
                  Tunnel Requests To
                </h3>
                <p className="text-xs text-muted-foreground mb-3">
                  Forward all webhook requests to your own endpoint
                </p>
                <div className="space-y-4">
                  {/* Live Switch */}
                  <div className="flex items-center justify-between p-3 bg-muted rounded-lg">
                    <div className="flex items-center space-x-3">
                      <div className={`w-3 h-3 rounded-full ${tunnelLive ? 'bg-green-500' : 'bg-gray-400'}`} />
                      <div>
                        <div className="text-sm font-medium">
                          {tunnelLive ? 'Live' : 'Inactive'}
                        </div>
                        <div className="text-xs text-muted-foreground">
                          {tunnelLive ? 'Tunneling is active' : 'Tunneling is disabled'}
                        </div>
                      </div>
                    </div>
                    <button onClick={() => onTunnelLiveChange(!tunnelLive)} className={`relative inline-flex h-6 w-11 items-center rounded-full transition-colors ${tunnelLive ? 'bg-green-600' : 'bg-gray-200'}`}>
                      <span className={`inline-block h-4 w-4 transform rounded-full bg-white transition-transform ${tunnelLive ? 'translate-x-6' : 'translate-x-1'}`} />
                    </button>
                  </div>
                  <input type="url" placeholder="https://your-api.com/webhook" value={tunnelUrl} onChange={e => onTunnelUrlChange(e.target.value)} className="w-full px-3 py-2 border border-input rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-ring" />
                  <div className="flex justify-end space-x-2">
                    <button onClick={() => setShowTunnelPopover(false)} className="px-3 py-1 text-sm text-muted-foreground hover:text-foreground">
                      Cancel
                    </button>
                    <button onClick={() => setShowTunnelPopover(false)} className="px-3 py-1 text-sm bg-primary text-primary-foreground rounded hover:bg-primary/90">
                      Save
                    </button>
                  </div>
                </div>
              </div>}
          </div>
        </div>
      </div>
    </header>;
}