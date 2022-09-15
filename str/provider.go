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

package str

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumi/pulumi-str/str/regexp"
)

func Provider() p.Provider {
	return infer.NewProvider().WithFunctions(
		infer.Function[*Replace, ReplaceArgs, ReplaceResult](),
		infer.Function[*TrimPrefix, TrimPrefixArgs, TrimPrefixResult](),
		infer.Function[*TrimSuffix, TrimSuffixArgs, TrimSuffixResult](),
		infer.Function[*regexp.Replace, regexp.ReplaceArgs, regexp.ReplaceResult](),
		infer.Function[*regexp.Split, regexp.SplitArgs, regexp.SplitResult](),
	).WithModuleMap(
		map[tokens.ModuleName]tokens.ModuleName{
			"str": "index",
		},
	).WithLanguageMap(map[string]any{
		"nodejs": map[string]any{
			"respectSchemaVersion": true,
			"dependencies": map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
		},
	}).WithDescription("Basic string manipulation funcions").
		WithDisplayName("String").
		WithHomepage("https://github.com/pulumi/pulumi-str").
		WithPublisher("Pulumi").
		WithRepository("https://github.com/pulumi/pulumi-str")
}
