package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/transport"
)

// Function to send a message to a user and delete it after 2 minutes
func sendTelegramMessage(apiID int, apiHash, phoneNumber, message string) error {
	// Create the Telegram
	client := telegram.NewClient(apiID, apiHash)

	// Create the context for interacting with Telegram API
	ctx := context.Background()

	// Run client operations in the provided context
	return client.Run(ctx, func(ctx context.Context) error {
		// Sign in using the phone number (this assumes 2FA is not required)
		if _, err := client.Auth().SignIn(ctx, phoneNumber, ""); err != nil {
			return fmt.Errorf("failed to sign in: %w", err)
		}

		// Get the user by phone number
		user, err := client.Users().GetUserByUsername(ctx, phoneNumber)
		if err != nil {
			return fmt.Errorf("failed to get user: %w", err)
		}

		// Send the message to the user
		messageSent, err := client.Messages().SendMessage(ctx, &tg.SendMessageRequest{
			Peer:    user.ID,
			Message: message,
		})
		if err != nil {
			return fmt.Errorf("failed to send message: %w", err)
		}

		// Wait for 2 minutes
		time.Sleep(2 * time.Minute)

		// Delete the message after 2 minutes
		_, err = client.Messages().DeleteMessages(ctx, &tg.DeleteMessagesRequest{
			Peer:   user.ID,
			ID:     []int{messageSent.ID},
			Revoke: true, // Revoke the message
		})
		if err != nil {
			return fmt.Errorf("failed to delete message: %w", err)
		}

		return nil
	})
}

func main() {
	// Telegram API credentials
	apiID := 27046507                             // Replace with your own API ID
	apiHash := "b659a231415648f8f36a6f51e78c52be" // Replace with your own API Hash
	phoneNumber := "+2348156572209"               // Replace with the target user's phone number
	message := "Hello, this is a message from Oluwateniola part 4"

	// Send the message
	err := sendTelegramMessage(apiID, apiHash, phoneNumber, message)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	} else {
		fmt.Println("Message sent and deleted successfully")
	}
}
