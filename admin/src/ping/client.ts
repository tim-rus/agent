import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { PingService } from "../gen/ping_pb";

const transport = createConnectTransport({
	baseUrl: "http://localhost:9000",
	useBinaryFormat: false,
});

export const client = createClient(PingService, transport);
