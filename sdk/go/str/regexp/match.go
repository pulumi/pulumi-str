// Code generated by pulumi-gen-str DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package regexp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Match reports whether the string s contains any match of the regular expression pattern.
func Match(ctx *pulumi.Context, args *MatchArgs, opts ...pulumi.InvokeOption) (*MatchResult, error) {
	var rv MatchResult
	err := ctx.Invoke("str:regexp:match", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type MatchArgs struct {
	Pattern string `pulumi:"pattern"`
	String  string `pulumi:"string"`
}

type MatchResult struct {
	Matches bool `pulumi:"matches"`
}

func MatchOutput(ctx *pulumi.Context, args MatchOutputArgs, opts ...pulumi.InvokeOption) MatchResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (MatchResult, error) {
			args := v.(MatchArgs)
			r, err := Match(ctx, &args, opts...)
			var s MatchResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(MatchResultOutput)
}

type MatchOutputArgs struct {
	Pattern pulumi.StringInput `pulumi:"pattern"`
	String  pulumi.StringInput `pulumi:"string"`
}

func (MatchOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MatchArgs)(nil)).Elem()
}

type MatchResultOutput struct{ *pulumi.OutputState }

func (MatchResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MatchResult)(nil)).Elem()
}

func (o MatchResultOutput) ToMatchResultOutput() MatchResultOutput {
	return o
}

func (o MatchResultOutput) ToMatchResultOutputWithContext(ctx context.Context) MatchResultOutput {
	return o
}

func (o MatchResultOutput) Matches() pulumi.BoolOutput {
	return o.ApplyT(func(v MatchResult) bool { return v.Matches }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(MatchResultOutput{})
}
