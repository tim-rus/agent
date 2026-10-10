import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { PingService } from "../gen/ping_pb";

const transport = createConnectTransport({
	baseUrl: import.meta.env["VITE_RPC_BASE_URL"] || "http://localhost:8080",
	useBinaryFormat: false,
});

export const client = createClient(PingService, transport);
