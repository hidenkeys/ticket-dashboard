package utils

import (
	"crypto/rand"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"log"
	"math"
	"math/big"
	"net/smtp"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"
)

// GenerateOTP generates a 6-digit OTP
func GenerateOTP() string {
	// Generate a cryptographically secure random number between 100000 and 999999
	newInt := big.NewInt(900000)
	n, err := rand.Int(rand.Reader, newInt)
	if err != nil {
		// Fallback - this should be handled properly in production
		panic("failed to generate secure random number")
	}
	otp := n.Int64() + 100000
	return fmt.Sprintf("%d", otp)
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWTToken(userID, email, phone string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"email":  email,
		"phone":  phone,                                 // Include role in JWT
		"exp":    time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func SendEmail00(to, subject, body string) error {
	from := os.Getenv("SMTP_FROM")
	password := os.Getenv("SMTP_PASSWORD")

	// Use Zoho SMTP host in PlainAuth
	auth := smtp.PlainAuth("", from, password, "smtp.zoho.com")

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\n\n" +
		body

	err := smtp.SendMail("smtp.zoho.com:587", auth, from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}
	return nil
}

// Helper function to compare names ignoring case, order, and allowing similarity
func AreNamesSimilar(name1, name2 string) bool {
	// Convert both names to lowercase for case-insensitive comparison
	name1 = strings.ToLower(name1)
	name2 = strings.ToLower(name2)

	// Split names into parts (assuming space separates first and last names)
	name1Parts := strings.Fields(name1)
	name2Parts := strings.Fields(name2)

	// Sort the parts of both names to handle order mismatch (e.g., "John Doe" vs "Doe John")
	sort.Strings(name1Parts)
	sort.Strings(name2Parts)

	// Join the sorted parts back into a single string for comparison
	name1 = strings.Join(name1Parts, " ")
	name2 = strings.Join(name2Parts, " ")

	// Calculate the Levenshtein distance between the names
	distance := levenshteinDistance(name1, name2)

	// You can adjust this threshold based on how lenient you want the match to be
	const threshold = 3 // Example threshold for similarity (lower = stricter)
	return distance <= threshold
}

// Levenshtein distance calculation
func levenshteinDistance(a, b string) int {
	lenA := len(a)
	lenB := len(b)

	matrix := make([][]int, lenA+1)
	for i := range matrix {
		matrix[i] = make([]int, lenB+1)
	}

	// Initialize the matrix
	for i := 0; i <= lenA; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= lenB; j++ {
		matrix[0][j] = j
	}

	// Fill the matrix with the Levenshtein distance
	for i := 1; i <= lenA; i++ {
		for j := 1; j <= lenB; j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}

			// Corrected the type mismatch by converting int to float64
			matrix[i][j] = int(math.Min(
				float64(matrix[i-1][j])+1, // Deletion
				math.Min(
					float64(matrix[i][j-1])+1,               // Insertion
					float64(matrix[i-1][j-1])+float64(cost), // Substitution (converted to float64)
				),
			))
		}
	}

	return matrix[lenA][lenB]
}

func SendTelegramMessage(apiID, apiHash, phoneNumber, message string) {
	// Constructing the command to execute the Python script with the required arguments
	cmd := exec.Command("python3", "/Users/oluwateniolasobande/GolandProjects/awesomeProject/ticket-dashboard/test2.py", apiID, apiHash, phoneNumber, message)

	// Running the command and getting the output or error
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing Python script: %s\n", err)
	}

	// Output the result from the Python script
	fmt.Printf("Output: %s\n", output)
}
