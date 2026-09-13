import { useState, type KeyboardEvent } from "react";
import { Button } from "../components/ui/button";
import { useCounterStore } from "../stores/counter";
import { ChatService } from "../pg/chat_pb";
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { errAsync } from "../lib/errors";

const transport = createConnectTransport({
	baseUrl: import.meta.env["RPC_BASE_URL"] || "http://localhost:8080",
	useBinaryFormat: false,
});

const chat = createClient(ChatService, transport)

export function Home(){
	const count = useCounterStore((state:any) => state.count)
	const increment = useCounterStore((state:any) => state.increment)

	const [history, setHistory] = useState<string[]>([])

	const update = async (e:KeyboardEvent) => {
		if (e.key != "Enter") return 

		const [res, err] = await errAsync(chat.ask({
			prompt: (e.target as HTMLInputElement).value
		}))
		if (err) {
			(e.target as HTMLInputElement).value = ""
			console.error("ERR", err);
			return
		}

		setHistory([...history, (e.target as HTMLInputElement).value, res.completion]);

		(e.target as HTMLInputElement).value = ""
	}

	return (
		<div className="w-full flex flex-col gap-4">
			<div>
				<div className="bg-amber-300">Hello from ui with router, zustand, tailwindcss and shadcn</div>
				<Button onClick={increment}>Click me ({count})</Button>
			</div>
			<div className="flex-1 bg-amber-300"><input type="text" onKeyDown={update} /></div>
			<div>
				{history.map(m=>(<div>{m}</div>))}
			</div>
		</div>
	)
}