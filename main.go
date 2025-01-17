package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

func getAllConcernedEnvPairs(prefixFilter string, trimPrefix bool) [][]string {
	var envPairs [][]string
	for _, envPair := range os.Environ() {
		// the environment variables are fetched as plain strings such as `LANG=go`
		envTuple := strings.SplitN(envPair, "=", 2)

		if strings.HasPrefix(envTuple[0], prefixFilter) {
			if trimPrefix {
				envTuple[0] = strings.Replace(envTuple[0], prefixFilter, "", 1)
			}
			envPairs = append(envPairs, envTuple)
		}
	}

	return envPairs
}

func handler(w http.ResponseWriter, r *http.Request) {
	prefix := os.Getenv("PREFIX")
	trimPrefix := false
	if os.Getenv("TRIM") != "" {
		trimPrefix = true
	}
	displayText := ""

	// we sort them based on the first element of the tuple i.e. the key
	envPairs := getAllConcernedEnvPairs(prefix, trimPrefix)
	sort.Slice(envPairs, func(i, j int) bool {
		return envPairs[i][0] < envPairs[j][0]
	})

	if prefix != "" {
		for _, envPair := range envPairs {
			displayText += fmt.Sprintf("%s : %s\n", envPair[0], envPair[1])
		}
	} else {
		displayText += "`PREFIX` environment variable is not correctly configured"
	}

	fmt.Fprintf(w, displayText)
}

func main() {
	flag.Parse()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
