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
	"context"
	"strings"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Replace struct{}
type ReplaceArgs struct {
	S   string `pulumi:"string"`
	Old string `pulumi:"old"`
	New string `pulumi:"new"`
}

func (r *Replace) Annotate(a infer.Annotator) {
	a.Describe(r, `Replace returns a copy of the string s with all
non-overlapping instances of old replaced by new.
If old is empty, it matches at the beginning of the string
and after each UTF-8 sequence, yielding up to k+1 replacements
for a k-rune string.`)
}

type ReplaceResult struct {
	Result string `pulumi:"result"`
}

func (*Replace) Invoke(ctx context.Context, req infer.FunctionRequest[ReplaceArgs]) (infer.FunctionResponse[ReplaceResult], error) {
	input := req.Input
	return infer.FunctionResponse[ReplaceResult]{
		Output: ReplaceResult{
			Result: strings.ReplaceAll(input.S, input.Old, input.New),
		},
	}, nil
}

type TrimPrefix struct{}
type TrimPrefixArgs struct {
	S      string `pulumi:"string"`
	Prefix string `pulumi:"prefix"`
}
type TrimPrefixResult struct {
	Result string `pulumi:"result"`
}

func (*TrimPrefix) Invoke(ctx context.Context, req infer.FunctionRequest[TrimPrefixArgs]) (infer.FunctionResponse[TrimPrefixResult], error) {
	return infer.FunctionResponse[TrimPrefixResult]{
		Output: TrimPrefixResult{
			Result: strings.TrimPrefix(req.Input.S, req.Input.Prefix),
		},
	}, nil
}

func (t *TrimPrefix) Annotate(a infer.Annotator) {
	a.Describe(t, "Trim a prefix from a string.")
}

type TrimSuffix struct{}
type TrimSuffixArgs struct {
	S      string `pulumi:"string"`
	Suffix string `pulumi:"suffix"`
}
type TrimSuffixResult struct {
	Result string `pulumi:"result"`
}

func (*TrimSuffix) Invoke(ctx context.Context, req infer.FunctionRequest[TrimSuffixArgs]) (infer.FunctionResponse[TrimSuffixResult], error) {
	return infer.FunctionResponse[TrimSuffixResult]{
		Output: TrimSuffixResult{
			Result: strings.TrimSuffix(req.Input.S, req.Input.Suffix),
		},
	}, nil
}

func (t *TrimSuffix) Annotate(a infer.Annotator) {
	a.Describe(t, "Trim a suffix from a string.")
}
