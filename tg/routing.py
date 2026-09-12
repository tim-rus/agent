from aiogram import Dispatcher
from aiogram.types import Message

def setupRoutes(dp: Dispatcher):

	@dp.message()
	async def msg_handler(message: Message) -> None:
		try:
			await message.send_copy(chat_id=message.chat.id)

		except TypeError:
			await message.answer("Unsupported message")