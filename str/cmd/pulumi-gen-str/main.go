package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	p "github.com/pulumi/pulumi-go-provider"

	"github.com/pulumi/pulumi-str/str"
	"github.com/pulumi/pulumi-str/str/version"
)

func main() {
	s, err := getSchema()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(s)
}

func getSchema() (string, error) {
	schema, err := p.GetSchema(context.Background(),
		"str", version.Version, str.Provider())
	if err != nil {
		return "", err
	}
	bytes, err := json.Marshal(schema)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
