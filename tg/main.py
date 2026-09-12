from os import getenv
from pathlib import Path
from dotenv import load_dotenv

# env

dotenv_file_env = getenv("ENV_FILE", "../config/.env.tg")
dotenv_file_path = Path(dotenv_file_env).resolve()

if not load_dotenv(dotenv_path=dotenv_file_path):
	raise RuntimeError(f"Failed to load env file at: {dotenv_file_path}")

# 

print(f"test: {getenv("TEST")}")