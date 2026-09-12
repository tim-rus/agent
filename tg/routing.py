from aiogram import Dispatcher
from aiogram.types import Message
from pg.ping import PINGServiceStub

def setupRoutes(dp: Dispatcher, rpc: PINGServiceStub):

	@dp.message()
	async def msg_handler(message: Message) -> None:
		try:
			await message.send_copy(chat_id=message.chat.id)
			res = await rpc.ping(msg=message.text)
			print(f"msg: {res.msg}")

		except TypeError:
			await message.answer("Unsupported message")