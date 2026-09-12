import asyncio
from os import getenv
from pathlib import Path
from dotenv import load_dotenv
import yaml

from pg.chat import ChatServiceStub
from grpclib.client import Channel

from tgmanager import TG
import routing

# env

dotenv_file_env = getenv("ENV_FILE", "./.env")
dotenv_file_path = Path(dotenv_file_env).resolve()

if not load_dotenv(dotenv_path=dotenv_file_path):
	raise RuntimeError(f"Failed to load env file at: {dotenv_file_path}")

# config

config_file_env = getenv("CONFIG_FILE", "./.config.yaml")
config_file_path = Path(config_file_env).resolve()

with open(config_file_path, "r") as file:
    config = yaml.safe_load(file)

# params

tgToken = getenv("TG_TOKEN")
if not tgToken:
	raise RuntimeError(f"TG_TOKEN required")

proxyString = getenv("PROXY_STRING")
if not proxyString:
	raise RuntimeError(f"PROXY_STRING required")


# main

async def main():
	channel = Channel(host=config['rpc']['host'], port=config['rpc']['port'])
	chatRPC = ChatServiceStub(channel)

	bot = TG(token=tgToken, proxy_url=proxyString)
	routing.setupRoutes(bot.dp, chatRPC)

	print("Starting bot")
	await bot.connect()

	# cleanup
	channel.close()

# run

if __name__ == "__main__":
	asyncio.run(main())