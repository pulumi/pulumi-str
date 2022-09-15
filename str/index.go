package str

import (
	"strings"

	p "github.com/pulumi/pulumi-go-provider"
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

func (*Replace) Call(ctx p.Context, input ReplaceArgs) (ReplaceResult, error) {
	return ReplaceResult{
		Result: strings.ReplaceAll(input.S, input.Old, input.New),
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

func (*TrimPrefix) Call(ctx p.Context, input TrimPrefixArgs) (TrimPrefixResult, error) {
	return TrimPrefixResult{
		Result: strings.TrimPrefix(input.S, input.Prefix),
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

func (*TrimSuffix) Call(ctx p.Context, input TrimSuffixArgs) (TrimSuffixResult, error) {
	return TrimSuffixResult{
		Result: strings.TrimSuffix(input.S, input.Suffix),
	}, nil
}

func (t *TrimSuffix) Annotate(a infer.Annotator) {
	a.Describe(t, "Trim a suffix from a string.")
}
