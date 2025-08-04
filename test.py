from telethon.sync import TelegramClient
import asyncio

#
# # Replace with your own API ID and Hash
# api_id = '27046507'
# api_hash = 'b659a231415648f8f36a6f51e78c52be'
#
# # Create a new Telegram client session
# client = TelegramClient('xxxxxxxxx123jj333', api_id, api_hash)
#
# # Connect to Telegram
# async def send_message():
#     await client.start()
#
#     # Phone number of the user you want to send the message to
#     phone_number = '+2348156572209'  # Replace with the target user's phone number
#     message = 'Hello, this is a messagefrom oluwateniola part 4'
#
#     # Send message to the user
#     user = await client.get_entity(phone_number)
#     sent_message = await client.send_message(user, message)
#
#     # Wait for 2 minutes (120 seconds)
#     await asyncio.sleep(10)
#
#         # Delete the message after 2 minutes
#     await client.delete_messages(user, sent_message.id)
#
#
# # Run the client
# client.loop.run_until_complete(send_message())


# Function to send message
async def send_telegram_message(api_id, api_hash, phone_number, message):
    # Create a new Telegram client session
    client = TelegramClient('xxxxxxxxx123jj333', api_id, api_hash)

    # Connect to Telegram
    await client.start()

    # Send message to the user
    user = await client.get_entity(phone_number)
    sent_message = await client.send_message(user, message)

    # Wait for 2 minutes (120 seconds)
    await asyncio.sleep(10)

    # Delete the message after 2 minutes
    await client.delete_messages(user, sent_message.id)

# Function to be called from your Go program
def send_message(api_id, api_hash, phone_number, message):
    asyncio.run(send_telegram_message(api_id, api_hash, phone_number, message))

