package main

import (
	str "github.com/pulumi/pulumi-str/sdk/go/str"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		out := str.ReplaceOutput(ctx, str.ReplaceOutputArgs{
			String: pulumi.String("Foo"),
			Old:    pulumi.String("o"),
			New:    pulumi.String("O"),
		})
		s := out.Result().ApplyT(func(s string) string {
			if s != "FOO" {
				panic(s)
			}
			return s
		})
		ctx.Export("s", s)
		return nil
	})
}
