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
    from telethon.sync import TelegramClient

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

# Function to parse arguments and call the send function
def run(api_id, api_hash, phone_number, message):
    # Set up the virtual environment and install dependencies
    setup_virtual_environment()

    # Run the async function to send the message
    asyncio.run(send_telegram_message(api_id, api_hash, phone_number, message))

if __name__ == "__main__":
    # Parse command-line arguments
    parser = argparse.ArgumentParser(description="Send message via Telegram.")
    parser.add_argument("api_id", help="Telegram API ID")
    parser.add_argument("api_hash", help="Telegram API Hash")
    parser.add_argument("phone_number", help="Recipient phone number")
    parser.add_argument("message", help="Message to send")

    args = parser.parse_args()

    # Run the message sending process
    run(args.api_id, args.api_hash, args.phone_number, args.message)
