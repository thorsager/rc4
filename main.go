package main

import (
	"crypto/rc4"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	flag "github.com/spf13/pflag"
)

var (
	flagInput   string
	flagOutput  string
	flagKeyFile string
)

func init() {
	flag.StringVarP(&flagInput, "input-file", "i", "", "input-file name")
	flag.StringVarP(&flagOutput, "output-file", "o", "", "output-file name")
	flag.StringVarP(&flagKeyFile, "key-file", "k", "", "key-file name")
}

func main() {
	flag.Parse()

	if flagInput == "" {
		fmt.Fprintln(os.Stderr, "input-file is required")
		flag.Usage()
		os.Exit(1)
	}

	if flagOutput == "" {
		fmt.Fprintln(os.Stderr, "output-file is required")
		flag.Usage()
		os.Exit(1)
	}

	if flagKeyFile == "" {
		fmt.Fprintln(os.Stderr, "key-file is required")
		flag.Usage()
		os.Exit(1)
	}

	key, err := ioutil.ReadFile(flagKeyFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open file %s: %s\n", flagKeyFile, err)
		os.Exit(2)
	}

	input, err := ioutil.ReadFile(flagInput)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open file %s: %s\n", flagInput, err)
		os.Exit(2)
	}

	cipher, err := rc4.NewCipher(key)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to create ciper: %s\n", err)
		os.Exit(3)
	}

	output := make([]byte, len(input))

	cipher.XORKeyStream(output, input)

	err = ioutil.WriteFile(flagOutput, output, fs.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to write file %s: %s\n", flagOutput, err)
		os.Exit(2)
	}
}
