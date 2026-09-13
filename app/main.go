package main

import (
	"fmt"
	"os"
	"io"
	"zlib"
)

// Usage: your_program.sh <command> <arg1> <arg2> ...
func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintf(os.Stderr, "Logs from your program will appear here!\n")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
		os.Exit(1)
	}

	switch command := os.Args[1]; command {
	case "init":
		for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
			}
		}
		
		headFileContents := []byte("ref: refs/heads/main\n")
		if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
		}
		
		fmt.Println("Initialized git directory")
		
	case "cat-file":
		if flag, hash := os.Args[2], os.Args[3]; flag == "-p" {
			file, err := os.Open(fmt.Sprintf(".git/objects/%s/%s", hash[:2], hash[2:]))
			if err != nil {
				fmt.Fprintf(os.Stderr, "error when reading from file")
				return
			} else {
				r, err := zlib.NewReader(file)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error when reading from file")
				}
				data, err := io.ReadAll(r)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error when reading from file")
				}

				defer r.Close()
				fmt.Println(string(data))				
			}
			defer file.Close()
		} else {
			fmt.Fprintf(os.Stderr, "unknown options to cat-file")
		} 	

	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
		os.Exit(1)
	}
}
