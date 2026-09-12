import re
from aiogram import Dispatcher
from aiogram.types import Message
from aiogram.enums import ParseMode

from md2tgmd import escape

from pg.chat import ChatServiceStub

def setupRoutes(dp: Dispatcher, rpc: ChatServiceStub):

	@dp.message()
	async def msg_handler(message: Message) -> None:
		try:
			print(f"prompt: {message.text}")
			res = await rpc.ask(prompt=message.text)
			print(f"response: {res.completion}")
			safe_text = escape(res.completion)
			await message.answer(text=safe_text, parse_mode=ParseMode.MARKDOWN_V2)

		except Exception as err:
			print(f"error: {err}")
			await message.answer("failed to reply")