package main

import (
	"encoding/json"
	"fmt"
	"os"

	"infra-policy-scanner/internal/scanner"
)

func main() {
	input := []scanner.Resource{
		{Name: "bucket-a", Public: true},
		{Name: "db-a", Public: false},
	}
	findings := scanner.Scan(input)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(findings); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
