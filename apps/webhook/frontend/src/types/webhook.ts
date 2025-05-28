export interface WebhookRequest {
  id: string;
  method: string;
  url: string;
  timestamp: Date;
  headers: Record<string, string>;
  queryParams: Record<string, string>;
  body: string | null;
  cookies: Record<string, string>;
  userAgent: string;
  ip: string;
  size: number;
  responseTime: number;
  status: number;
}