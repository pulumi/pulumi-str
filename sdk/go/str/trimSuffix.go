// Code generated by pulumi-gen-str DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package str

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Trim a suffix from a string.
func TrimSuffix(ctx *pulumi.Context, args *TrimSuffixArgs, opts ...pulumi.InvokeOption) (*TrimSuffixResult, error) {
	var rv TrimSuffixResult
	err := ctx.Invoke("str:index:trimSuffix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type TrimSuffixArgs struct {
	String string `pulumi:"string"`
	Suffix string `pulumi:"suffix"`
}

type TrimSuffixResult struct {
	Result string `pulumi:"result"`
}

func TrimSuffixOutput(ctx *pulumi.Context, args TrimSuffixOutputArgs, opts ...pulumi.InvokeOption) TrimSuffixResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (TrimSuffixResult, error) {
			args := v.(TrimSuffixArgs)
			r, err := TrimSuffix(ctx, &args, opts...)
			var s TrimSuffixResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(TrimSuffixResultOutput)
}

type TrimSuffixOutputArgs struct {
	String pulumi.StringInput `pulumi:"string"`
	Suffix pulumi.StringInput `pulumi:"suffix"`
}

func (TrimSuffixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrimSuffixArgs)(nil)).Elem()
}

type TrimSuffixResultOutput struct{ *pulumi.OutputState }

func (TrimSuffixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrimSuffixResult)(nil)).Elem()
}

func (o TrimSuffixResultOutput) ToTrimSuffixResultOutput() TrimSuffixResultOutput {
	return o
}

func (o TrimSuffixResultOutput) ToTrimSuffixResultOutputWithContext(ctx context.Context) TrimSuffixResultOutput {
	return o
}

func (o TrimSuffixResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v TrimSuffixResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TrimSuffixResultOutput{})
}
