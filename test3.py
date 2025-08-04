import subprocess
import sys
import os
import asyncio
import argparse

# Function to create and activate a virtual environment, and install required dependencies
def setup_virtual_environment():
    # Check if virtual environment already exists
    if not os.path.exists("venv"):
        # Create a virtual environment
        subprocess.check_call([sys.executable, "-m", "venv", "venv"])

    # Install dependencies (telethon)
    subprocess.check_call([os.path.join("venv", "bin", "pip"), "install", "telethon"])

# Async function to send a message to a Telegram user
async def send_telegram_message(api_id, api_hash, phone_number, message):
    from telethon import TelegramClient
    from telethon.tl.functions.contacts import ImportContactsRequest
    from telethon.tl.types import InputPhoneContact

    # Create a Telegram client session
    client = TelegramClient('xxxxxxxxx123jj333', api_id, api_hash)

    # Connect to Telegram
    await client.start()

    # Add the user as a temporary contact
    contact = InputPhoneContact(client_id=0, phone=phone_number, first_name="Temp", last_name="User")
    result = await client(ImportContactsRequest([contact]))

    if not result.users:
        print(f"❌ Could not find user with phone number {phone_number}.")
        return

    user = result.users[0]

    # Send the message
    print(f"✅ Sending message to {phone_number}...")
    sent_message = await client.send_message(user, message)

    # Wait 10 seconds (change this as needed)
    await asyncio.sleep(10)

    # Delete the message
    await client.delete_messages(user, sent_message.id)
    print("✅ Message deleted after 10 seconds.")

    # Optional: Remove the contact (if you want)
    # await client.delete_contacts([user.id])

    await client.disconnect()

# Function to parse arguments and call the send function
def run(api_id, api_hash, phone_number, message):
      # Run the async function to send the message
    setup_virtual_environment()
    asyncio.run(send_telegram_message(api_id, api_hash, phone_number, message))

if __name__ == "__main__":
    # Parse command-line arguments
    parser = argparse.ArgumentParser(description="Send message via Telegram.")
    parser.add_argument("api_id", help="Telegram API ID")
    parser.add_argument("api_hash", help="Telegram API Hash")
    parser.add_argument("phone_number", help="Recipient phone number")
    parser.add_argument("message", help="Message to send")

    args = parser.parse_args()
    run(args.api_id, args.api_hash, args.phone_number, args.message)
