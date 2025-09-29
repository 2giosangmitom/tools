package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Struct to hold configuration options
type opts struct {
	length           int
	includeNumbers   bool
	includeSymbols   bool
	includeUppercase bool
	includeLowercase bool
	includeAlpha     bool
}

// Function to build and print the password based on the configuration
func build(config *opts) (string, error) {
	// Character pools
	var charPool strings.Builder

	// Add character sets based on configuration
	if config.includeAlpha {
		if config.includeUppercase {
			charPool.WriteString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		}
		if config.includeLowercase {
			charPool.WriteString("abcdefghijklmnopqrstuvwxyz")
		}
	}
	if config.includeNumbers {
		charPool.WriteString("0123456789")
	}
	if config.includeSymbols {
		charPool.WriteString("!@#$%^&*()-_=+[]{}|;:,.<>?/")
	}

	pool := charPool.String()
	if len(pool) == 0 {
		return "", fmt.Errorf("character pool is empty, please check your configuration")
	}

	// Generate password
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	password := make([]byte, config.length)
	for i := range password {
		password[i] = pool[r.Intn(len(pool))]
	}

	return string(password), nil
}

func main() {
	// Flags
	var config opts

	flag.IntVar(&config.length, "len", 12, "Length of the password to generate")
	flag.BoolVar(&config.includeNumbers, "num", false, "Include numbers in the password")
	flag.BoolVar(&config.includeSymbols, "sym", false, "Include symbols in the password")
	flag.BoolVar(&config.includeUppercase, "up", true, "Include uppercase letters in the password (ignored if -alpha is false)")
	flag.BoolVar(&config.includeLowercase, "low", true, "Include lowercase letters in the password (ignored if -alpha is false)")
	// Include alphabetic characters by default
	flag.BoolVar(&config.includeAlpha, "alpha", true, "Include alphabetic characters in the password")

	// Parse command-line flags
	flag.Parse()

	// Generate and print the password based on the configuration
	password, err := build(&config)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(password)
}
