import { client } from "./client";

export async function ping() {
	const res = await client.ping({
		msg: "hi"
	})
	return res
}
