import asyncio
from os import getenv
from pathlib import Path
from dotenv import load_dotenv

from tgmanager import TG
import routing

# env

dotenv_file_env = getenv("ENV_FILE", "../config/.env.tg")
dotenv_file_path = Path(dotenv_file_env).resolve()

if not load_dotenv(dotenv_path=dotenv_file_path):
	raise RuntimeError(f"Failed to load env file at: {dotenv_file_path}")

# params

tgToken = getenv("TG_TOKEN")
if not tgToken:
	raise RuntimeError(f"TG_TOKEN required")

proxyString = getenv("PROXY_STRING")
if not proxyString:
	raise RuntimeError(f"PROXY_STRING required")

# services

bot = TG(token=tgToken, proxy_url=proxyString)
routing.setupRoutes(bot.dp)

# main

async def main():
	print("Starting bot")
	await bot.connect()

# run

if __name__ == "__main__":
	asyncio.run(main())