package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
)

func main() {
	encode := flag.Bool("e", false, "Encode")
	decode := flag.Bool("d", false, "Decode")
	flag.Parse()

	if !*encode && !*decode {
		fmt.Fprintln(os.Stderr, "No no no. Please specify -e (encode) or -d (decode).")
		os.Exit(1)
	}

	if *encode && *decode {
		fmt.Fprintln(os.Stderr, "No no no. Please specify -e (encode) xor -d (decode).")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch {
		case *encode:
			fmt.Println(url.QueryEscape(scanner.Text()))
		case *decode:
			s, err := url.QueryUnescape(scanner.Text())
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println(s)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
