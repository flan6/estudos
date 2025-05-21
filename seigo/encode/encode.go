package encode

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func key(pwd string, salt []byte) ([]byte, error) {
	return pbkdf2.Key(sha1.New, pwd, salt, 100, 16)
}

// decryptAndSave decrypts an ES3 file and saves it as JSON.
func DecryptAndSave(es3Path string, password string) {
	// Read the ES3 file
	data, err := os.ReadFile(es3Path)
	if err != nil {
		fmt.Printf("Error reading ES3 file '%s': %v\n", es3Path, err)
		os.Exit(1)
	}

	// Check file size (must be at least IV size + 1 block)
	if len(data) < 16 {
		fmt.Printf("Error: ES3 file '%s' is too small\n", es3Path)
		os.Exit(1)
	}

	// Extract IV and encrypted data
	iv := data[:16]
	encrypted := data[16:]
	k, err := key(password, iv)
	if err != nil {
		fmt.Println("Error reading password:", err)
		os.Exit(1)
	}

	fmt.Printf("Derived key: %x\n", k)

	// Create AES cipher
	block, err := aes.NewCipher(k)
	if err != nil {
		fmt.Println("Error creating AES cipher:", err)
		os.Exit(1)
	}

	// Validate encrypted data length
	if len(encrypted)%aes.BlockSize != 0 {
		fmt.Printf("Error: encrypted data in '%s' is not a multiple of block size\n", es3Path)
		os.Exit(1)
	}

	// Decrypt in-place
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encrypted, encrypted)

	fmt.Println("Decrypted:", string(encrypted[:100]))

	// Remove PKCS#7 padding
	padLen := int(encrypted[len(encrypted)-1])
	if padLen == 0 || padLen > aes.BlockSize {
		fmt.Printf("Error: invalid padding in decrypted data from '%s'\n", es3Path)
		os.Exit(1)
	}
	decrypted := encrypted[:len(encrypted)-padLen]

	// Generate output JSON path
	jsonPath := strings.TrimSuffix(es3Path, filepath.Ext(es3Path)) + ".json"

	// Save the decrypted JSON
	err = os.WriteFile(jsonPath, decrypted, 0644)
	if err != nil {
		fmt.Printf("Error writing JSON file '%s': %v\n", jsonPath, err)
		os.Exit(1)
	}

	fmt.Printf("Decrypted JSON saved to '%s'\n", jsonPath)
}

// encryptAndSave encrypts a JSON file and saves it as an ES3 file.
func EncryptAndSave(jsonPath string, password string) {
	// Read the JSON file
	jsonData, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Printf("Error reading JSON file '%s': %v\n", jsonPath, err)
		os.Exit(1)
	}

	// Generate random IV
	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	if err != nil {
		fmt.Println("Error generating IV:", err)
		os.Exit(1)
	}

	k, err := key(password, iv)
	if err != nil {
		fmt.Println("Error reading password:", err)
		os.Exit(1)
	}

	// Create AES cipher
	block, err := aes.NewCipher(k)
	if err != nil {
		fmt.Println("Error creating AES cipher:", err)
		os.Exit(1)
	}

	// Add PKCS#7 padding
	padLen := aes.BlockSize - (len(jsonData) % aes.BlockSize)
	padded := make([]byte, len(jsonData)+padLen)
	copy(padded, jsonData)
	for i := len(jsonData); i < len(padded); i++ {
		padded[i] = byte(padLen)
	}

	// Encrypt the data
	mode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(padded))
	mode.CryptBlocks(encrypted, padded)

	// Combine IV and encrypted data
	es3Data := append(iv, encrypted...)

	// Generate output ES3 path
	es3Path := strings.TrimSuffix(jsonPath, filepath.Ext(jsonPath)) + ".es3"

	// Save the ES3 file
	err = os.WriteFile(es3Path, es3Data, 0644)
	if err != nil {
		fmt.Printf("Error writing ES3 file '%s': %v\n", es3Path, err)
		os.Exit(1)
	}

	fmt.Printf("Encrypted ES3 saved to '%s'\n", es3Path)
}
