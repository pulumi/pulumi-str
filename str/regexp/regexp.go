package regexp

import (
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
