import os
import sys
import subprocess
import argparse
import asyncio

VENV_DIR = "venv"
PYTHON_EXECUTABLE = os.path.join(VENV_DIR, "bin", "python3") if sys.platform != "win32" else os.path.join(VENV_DIR, "Scripts", "python.exe")

def setup_virtual_environment():
    # 1. Create virtual environment if not exists
    if not os.path.exists(VENV_DIR):
        print("🔧 Creating virtual environment...")
        subprocess.check_call([sys.executable, "-m", "venv", VENV_DIR])

    # 2. Install Telethon if not already installed
    print("📦 Installing dependencies (telethon)...")
    subprocess.check_call([PYTHON_EXECUTABLE, "-m", "pip", "install", "--upgrade", "pip"])
    subprocess.check_call([PYTHON_EXECUTABLE, "-m", "pip", "install", "telethon"])

def re_run_inside_venv():
    """Re-run the current script inside the virtual environment Python interpreter."""
    print("🚀 Re-running script inside virtual environment...")
    subprocess.check_call([PYTHON_EXECUTABLE] + sys.argv)
    sys.exit(0)

def already_running_inside_venv():
    """Detect if script is already running in the venv."""
    return sys.executable.startswith(os.path.abspath(VENV_DIR))

# Async function to send a message
async def send_telegram_message(api_id, api_hash, phone_number, message):
    from telethon import TelegramClient
    from telethon.tl.functions.contacts import ImportContactsRequest
    from telethon.tl.types import InputPhoneContact

    client = TelegramClient("xxxxxxxxx123jj333", api_id, api_hash)
    await client.start()

    contact = InputPhoneContact(client_id=0, phone=phone_number, first_name="Temp", last_name="User")
    result = await client(ImportContactsRequest([contact]))

    if not result.users:
        print(f"❌ Could not find user with phone number {phone_number}.")
        return

    user = result.users[0]
    print(f"✅ Sending message to {phone_number}...")
    sent_message = await client.send_message(user, message)

    await asyncio.sleep(10)
    await client.delete_messages(user, sent_message.id)
    print("✅ Message deleted after 10 seconds.")

    await client.disconnect()

# Main execution
def main():
    parser = argparse.ArgumentParser(description="Send Telegram message")
    parser.add_argument("api_id", help="Telegram API ID")
    parser.add_argument("api_hash", help="Telegram API Hash")
    parser.add_argument("phone_number", help="Phone number of recipient (e.g., 234xxxxxxxxxx)")
    parser.add_argument("message", help="Message content")

    args = parser.parse_args()
    asyncio.run(send_telegram_message(args.api_id, args.api_hash, args.phone_number, args.message))

if __name__ == "__main__":
    if not already_running_inside_venv():
        setup_virtual_environment()
        re_run_inside_venv()
    else:
        main()
