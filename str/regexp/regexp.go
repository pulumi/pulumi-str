package regexp

import (
	"fmt"
	"regexp"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Replace struct{}

type ReplaceArgs struct {
	S   string `pulumi:"string"`
	Old string `pulumi:"old"`
	New string `pulumi:"new"`
}
type ReplaceResult struct {
	Result string `pulumi:"result"`
}

func (*Replace) Call(ctx p.Context, input ReplaceArgs) (ReplaceResult, error) {
	rgx, err := regexp.Compile(input.Old)
	if err != nil {
		return ReplaceResult{}, err
	}
	return ReplaceResult{
		Result: rgx.ReplaceAllString(input.S, input.New),
	}, nil
}

func (r *Replace) Annotate(a infer.Annotator) {
	a.Describe(r, `A regex based replace on a string.

The underlying regexp engine is the go "regexp" stdlib package.
You can see details at for available patterns at https://pkg.go.dev/regexp/syntax.`)
}

func (r *ReplaceArgs) Annotate(a infer.Annotator) {
	a.Describe(&r.New,
		`The new string.

Note: Inside repl, "$" signs are interpreted as an Expand, so for instance $1 represents the text of the first submatch. `)
	a.Describe(&r.Old, "The regular expression to match against.")
	a.Describe(&r.S, "The string to operate over.")
}

func (r *ReplaceResult) Annotate(a infer.Annotator) {
	a.Describe(&r.Result, `The input "string" where each pattern matching "old" is replaced by "new".`)
}

type Split struct{}
type SplitArgs struct {
	S  string `pulumi:"string"`
	On string `pulumi:"on"`
	N  *int   `pulumi:"count,optional"`
}

type SplitResult struct {
	Result []string `pulumi:"result"`
}

func (s *Split) Annotate(a infer.Annotator) {
	a.Describe(s, "Split a string on a regex.")
}

func (s *SplitArgs) Annotate(a infer.Annotator) {
	a.Describe(&s.N, "`count` determines the number of substrings to return. \n"+
		"If `count` is not provided, it defaults to substrings.\n"+
		"If `count` is provided then the last substring will be the unsplit remainder.\n"+
		"It is an error to pass `count < 1`.")
	a.Describe(&s.On, "The regex to split on.")
	a.Describe(&s.S, "The string on which to split.")
}

func (s *SplitResult) Annotate(a infer.Annotator) {
	a.Describe(&s.Result, "The result of the string split.")
}

func (*Split) Call(ctx p.Context, input SplitArgs) (SplitResult, error) {
	rgx, err := regexp.Compile(input.On)
	if err != nil {
		return SplitResult{}, err
	}
	n := -1
	if input.N != nil {
		if *input.N <= 0 {
			return SplitResult{}, fmt.Errorf("count <= 0 is not allowed")
		}
		n = *input.N
	}
	return SplitResult{
		Result: rgx.Split(input.S, n),
	}, nil
}
