// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package regexp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A regex based replace on a string.
//
// The underlying regexp engine is the go "regexp" stdlib package.
// You can see details at for available patterns at https://pkg.go.dev/regexp/syntax.
func Replace(ctx *pulumi.Context, args *ReplaceArgs, opts ...pulumi.InvokeOption) (*ReplaceResult, error) {
	var rv ReplaceResult
	err := ctx.Invoke("str:regexp:replace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ReplaceArgs struct {
	// The new string.
	//
	// Note: Inside repl, "$" signs are interpreted as an Expand, so for instance
	// $1 represents the text of the first submatch.
	New string `pulumi:"new"`
	// The regular expression to match against.
	Old string `pulumi:"old"`
	// The string to operate over.
	String string `pulumi:"string"`
}

type ReplaceResult struct {
	// The input "string" where each pattern matching "old" is replaced by "new".
	Result string `pulumi:"result"`
}

func ReplaceOutput(ctx *pulumi.Context, args ReplaceOutputArgs, opts ...pulumi.InvokeOption) ReplaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ReplaceResult, error) {
			args := v.(ReplaceArgs)
			r, err := Replace(ctx, &args, opts...)
			var s ReplaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ReplaceResultOutput)
}

type ReplaceOutputArgs struct {
	// The new string.
	//
	// Note: Inside repl, "$" signs are interpreted as an Expand, so for instance
	// $1 represents the text of the first submatch.
	New pulumi.StringInput `pulumi:"new"`
	// The regular expression to match against.
	Old pulumi.StringInput `pulumi:"old"`
	// The string to operate over.
	String pulumi.StringInput `pulumi:"string"`
}

func (ReplaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplaceArgs)(nil)).Elem()
}

type ReplaceResultOutput struct{ *pulumi.OutputState }

func (ReplaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplaceResult)(nil)).Elem()
}

func (o ReplaceResultOutput) ToReplaceResultOutput() ReplaceResultOutput {
	return o
}

func (o ReplaceResultOutput) ToReplaceResultOutputWithContext(ctx context.Context) ReplaceResultOutput {
	return o
}

// The input "string" where each pattern matching "old" is replaced by "new".
func (o ReplaceResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v ReplaceResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplaceResultOutput{})
}
