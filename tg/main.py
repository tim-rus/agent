import asyncio

from gen.ping import PingResponse, PingServiceStub
from grpclib.client import Channel


async def main():
    channel = Channel(host="127.0.0.1", port=9000)
    stub = PingServiceStub(channel)

    response: PingResponse = await stub.ping(msg="hello from tg")
    print(f"Received response from server: {response.msg}")

    channel.close()

if __name__ == "__main__":
	asyncio.run(main())
