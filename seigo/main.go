package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/flan6/seigo/encode"
	"github.com/flan6/seigo/internal/secret"
	"github.com/flan6/seigo/options"
)

// main parses command-line flags and directs the program flow.
func main() {
	op := (&options.ArgOptions{}).Parse()

	password := secret.Password()
	fmt.Println(password)

	// dst := base64.StdEncoding.EncodeToString([]byte(k))
	// fmt.Println(string(dst))

	// key, err := hex.DecodeString(k)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	fmt.Println(len(password))

	filesArgs := flag.Args()

	if len(filesArgs) < 1 {
		fmt.Println("Invalid file path.")
		os.Exit(1)
	}

	// Execute the appropriate action
	switch {
	case op.R:
		encode.DecryptAndSave(filesArgs[0], password)
	case op.S:
		encode.EncryptAndSave(filesArgs[0], password)
	}
}
