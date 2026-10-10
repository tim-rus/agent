import { useState } from "react";
import { useCounterStore } from "../stores/counter";
import { errAsync } from "../lib/errors";
import { ping } from "../ping/ping";
import { Button } from "../components/ui/button";

export function Home(){
	const count = useCounterStore((state:any) => state.count)
	const increment = useCounterStore((state:any) => state.increment)

	const [data, setData] = useState<string>("waiting")

	const click = async () => {
		const [res, err] = await errAsync(ping())
		if (err) {
			setData(`error: ${err.message}`)
			return
		}
		console.log(res);
		if (res && res.msg) {
			setData(res.msg)
		}
	}

	return (
		<>
			<div className="bg-amber-300">Hello from ui with router, zustand, tailwindcss and shadcn</div>
			<Button onClick={increment}>Click me ({count})</Button>
			<div className="bg-green-300">{data}</div>
			<Button onClick={click}>Click me</Button>
		</>
	)
}
