package options

import (
	"flag"
	"fmt"
	"os"
)

// represents the available [OPTIONS]
type ArgOptions struct {
	R bool // represents the -r arg to read and decrypt es3
	S bool // represents the -s arg to encrypt and save es3
}

// Parse wraps flags.Parse function
// read args and sets defaults to [ArgOptions]
func (a *ArgOptions) Parse() *ArgOptions {
	flag.BoolVar(&a.R, "r", false, "reads and decrypt es3")
	flag.BoolVar(&a.S, "s", false, "encrypt and saves es3")

	flag.Parse()

	// Ensure only one of -r or -s is specified
	if a.R == a.S {
		fmt.Println("Usage: specify either -r <path> or -s <path>, but not both.")
		fmt.Println("  -r <path>: Decrypt ES3 file to JSON")
		fmt.Println("  -s <path>: Encrypt JSON file to ES3")
		fmt.Println("  -k <hex_key>: 32-character hex-encoded AES key")
		os.Exit(1)
	}

	return a
}
