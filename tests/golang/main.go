// Copyright 2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
