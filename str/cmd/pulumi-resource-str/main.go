package main

import (
	"fmt"
	"os"

	p "github.com/pulumi/pulumi-go-provider"

	"github.com/pulumi/pulumi-str/str"
	"github.com/pulumi/pulumi-str/str/version"
)

func main() {
	err := p.RunProvider("str", version.Version, str.Provider())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
