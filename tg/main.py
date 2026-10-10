import asyncio
from os import getenv
from pathlib import Path

import yaml
from dotenv import load_dotenv
from gen.ping import PingResponse, PingServiceStub
from grpclib.client import Channel

# config

dotenv_file_env = getenv("ENV_FILE", "../config/.env.tg")
dotenv_file_path = Path(dotenv_file_env).resolve()

if not load_dotenv(dotenv_path=dotenv_file_path):
	raise RuntimeError(f"Failed to load env file at: {dotenv_file_path}")

config_file_env = getenv("CONFIG_FILE", "../config/tg.config.yaml")
config_file_path = Path(config_file_env).resolve()

with open(config_file_path, "r") as file:
    config = yaml.safe_load(file)

with open(config_file_path, "r") as file:
    config = yaml.safe_load(file)

print(f"config: {config}")

# main

async def main():
    channel = Channel(host="127.0.0.1", port=9000)
    stub = PingServiceStub(channel)

    response: PingResponse = await stub.ping(msg="hello from tg")
    print(f"Received response from server: {response.msg}")

    channel.close()

# run

if __name__ == "__main__":
	asyncio.run(main())
