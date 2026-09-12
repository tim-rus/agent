from aiogram.client.session.aiohttp import AiohttpSession
from aiogram import Bot, Dispatcher

class TG:
	def __init__(self, token: str, proxy_url: str):
		# TODO: check params
		session = AiohttpSession(proxy=proxy_url)
		self.bot = Bot(token=token, session=session)
		self.dp = Dispatcher()


	async def connect(self):
		print("Start long polling...")
		await self.dp.start_polling(self.bot)
