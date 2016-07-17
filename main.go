package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	ci "github.com/ipfs/go-libp2p-crypto"
	peer "github.com/ipfs/go-libp2p-peer"
)

func main() {
	vanity := flag.String("name", "", "write the name to appear in the begining of the key")
	size := flag.Int("bitsize", 2048, "select the bitsize of the key to generate")
	typ := flag.String("type", "RSA", "select type of key to generate")

	flag.Parse()

	var atyp int
	switch *typ {
	case "RSA":
		atyp = ci.RSA
	default:
		fmt.Fprintln(os.Stderr, "unrecognized key type: ", *typ)
		os.Exit(1)
	}

	r, err := regexp.Compile("^Qm" + *vanity + ".+$")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Generating some %d bit %s keys...\n", *size, *typ)

	var priv ci.PrivKey
	var pub ci.PubKey
	var id = ""
	var i = 0

	for !r.MatchString(id) {

		priv, pub, err = ci.GenerateKeyPair(atyp, *size)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		pid, err := peer.IDFromPublicKey(pub)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		id = pid.Pretty()

		i++

		if i % 10 == 0 {
			fmt.Fprintf(os.Stderr, "%s", ".")
		}

		if i % 500 == 0 {
			fmt.Fprintln(os.Stderr, "", i)
		}
	}

	fmt.Fprintln(os.Stderr, "\nSuccess!")

	fmt.Fprintf(os.Stderr, "ID for generated key: %s\n", id)

	data, err := priv.Bytes()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Stdout.Write(data)
}
