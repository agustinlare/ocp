package main

import (
	"log"
	"os"
	"strings"
)

func tryThis(e string) []byte {
	var resp []byte

	switch {
	case strings.Contains(string(e), "dial tcp"):
		panic("No estas conectado a la red.")

	case strings.Contains(string(e), "Unauthorized"):
		if os.Getenv("GO_PASSWORD") != "" {
			panic("Too many re-login...\n" + e)
		}

		var config Env = getConfig()

		log.Println("WARNING\t| Not authorized, re-login...")
		config = setPassword(config)
		os.Setenv("GO_PASSWORD", "1")

		resp = loginCluster(getContext(), "default")

	case strings.Contains(string(e), "no context exists"):
		log.Println("WARNING\nContexto inexistente.")
		// c = selectCluster()
		// resp = loginCluster(c, n)

	default:
		panic(e)
	}

	return resp
}
