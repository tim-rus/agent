import { Button } from "../components/ui/button";
import { useCounterStore } from "../stores/counter";

export function Home(){
	const count = useCounterStore((state:any) => state.count)
	const increment = useCounterStore((state:any) => state.increment)

	return (
		<>
			<div className="bg-amber-300">Hello from ui with router, zustand, tailwindcss and shadcn</div>
			<Button onClick={increment}>Click me ({count})</Button>
		</>
	)
}