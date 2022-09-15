package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	p "github.com/pulumi/pulumi-go-provider"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pythongen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/spf13/cobra"

	"github.com/pulumi/pulumi-str/str"
	"github.com/pulumi/pulumi-str/str/version"
)

const NAME string = "str"

func main() {
	c := &cobra.Command{
		Use:     fmt.Sprintf("pulumi-gen-%s", NAME),
		Version: version.Version,
	}
	c.AddCommand(schemaCmd(), sdkCmd())
	endIf(c.Execute())

}
func endIf(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
func schemaCmd() *cobra.Command {
	var file string
	c := &cobra.Command{
		Use:   "schema",
		Short: fmt.Sprintf("Emit the schema for the `%s` package.", NAME),
		Run: func(cmd *cobra.Command, args []string) {

			s, err := getSchema()
			endIf(err)
			if file != "" {
				f, err := os.Create(file)
				endIf(err)
				defer f.Close()
				n, err := f.WriteString(s)
				endIf(err)
				if n != len(s) {
					endIf(fmt.Errorf("only wrote %d/%d bytes", n, len(s)))
				}
			} else {
				fmt.Print(s)
			}
		},
	}
	c.PersistentFlags().StringVarP(&file, "file", "f", "",
		"Output to a file instead of stdout.")
	return c
}

func sdkCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "language [language]",
		Short: fmt.Sprintf("Emit the SDK(s) for the `%s` package.", NAME),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				endIf(fmt.Errorf("missing required positional arg: language"))
			}
			language := args[0]
			var f func(string, *pschema.Package, map[string][]byte) (map[string][]byte, error)
			switch language {
			case "typescript", "nodejs":
				f = nodejsgen.GeneratePackage
			case "python":
				f = pythongen.GeneratePackage
			case "go":
				f = func(s string, p *pschema.Package, m map[string][]byte) (map[string][]byte, error) {
					return gogen.GeneratePackage(s, p)
				}
			case "csharp", "c#", "dotnet":
				f = dotnetgen.GeneratePackage
			default:
				endIf(fmt.Errorf("unrecognized language: %q", language))
			}
			spec, err := p.GetSchema(context.Background(), NAME, version.Version, str.Provider())
			endIf(err)
			pkg, err := pschema.ImportSpec(spec, nil)
			endIf(err)
			sdk, err := f(fmt.Sprintf("pulumi-gen-%s", NAME), pkg, nil)
			endIf(err)
			if err := os.MkdirAll("sdk", 0700); !os.IsNotExist(err) {
				endIf(err)
			}
			root := filepath.Join("sdk", language)
			if err := os.RemoveAll(root); err != nil && !os.IsNotExist(err) {
				endIf(fmt.Errorf("removing root: %w", err))
			}
			for path, file := range sdk {
				path := filepath.Join(root, path)
				err = os.MkdirAll(filepath.Dir(path), 0700)
				endIf(err)
				err = os.WriteFile(path, file, 0600)
				endIf(err)
			}
		},
	}
	return c
}

func getSchema() (string, error) {
	schema, err := p.GetSchema(context.Background(),
		NAME, version.Version, str.Provider())
	if err != nil {
		return "", err
	}
	bytes, err := json.Marshal(schema)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
