import type { WebhookRequest } from "../types/webhook";
export function generateWebhookId(): string {
  return Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15);
}
export const mockRequests: WebhookRequest[] = [{
  id: "req_001",
  method: "GET",
  url: "https://webhook-tester.com/abc123",
  timestamp: new Date(Date.now() - 1000 * 60 * 5),
  headers: {
    "accept": "*/*",
    "user-agent": "curl/8.7.1",
    "host": "webhook-tester.com",
    "x-forwarded-for": "192.168.1.1",
    "content-type": "application/json"
  },
  queryParams: {
    "source": "test",
    "version": "1.0"
  },
  body: JSON.stringify({
    message: "Hello webhook!",
    userId: 123
  }),
  cookies: {
    "session_id": "abc123def456",
    "user_pref": "dark_mode"
  },
  userAgent: "curl/8.7.1",
  ip: "192.168.1.1",
  size: 256,
  responseTime: 0.045,
  status: 200
}, {
  id: "req_002",
  method: "POST",
  url: "https://webhook-tester.com/abc123",
  timestamp: new Date(Date.now() - 1000 * 60 * 10),
  headers: {
    "accept": "application/json",
    "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
    "host": "webhook-tester.com",
    "content-type": "application/json",
    "authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
  },
  queryParams: {},
  body: JSON.stringify({
    event: "user.created",
    data: {
      id: 456,
      email: "user@example.com",
      name: "John Doe"
    },
    timestamp: "2025-01-26T18:30:00Z"
  }),
  cookies: {},
  userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
  ip: "203.0.113.42",
  size: 512,
  responseTime: 0.123,
  status: 200
}];