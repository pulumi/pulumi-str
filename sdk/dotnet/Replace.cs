// *** WARNING: this file was generated by pulumi-gen-str. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Str
{
    public static class Replace
    {
        /// <summary>
        /// Replace returns a copy of the string s with all
        /// non-overlapping instances of old replaced by new.
        /// If old is empty, it matches at the beginning of the string
        /// and after each UTF-8 sequence, yielding up to k+1 replacements
        /// for a k-rune string.
        /// </summary>
        public static Task<ReplaceResult> InvokeAsync(ReplaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ReplaceResult>("str:index:replace", args ?? new ReplaceArgs(), options.WithDefaults());

        /// <summary>
        /// Replace returns a copy of the string s with all
        /// non-overlapping instances of old replaced by new.
        /// If old is empty, it matches at the beginning of the string
        /// and after each UTF-8 sequence, yielding up to k+1 replacements
        /// for a k-rune string.
        /// </summary>
        public static Output<ReplaceResult> Invoke(ReplaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ReplaceResult>("str:index:replace", args ?? new ReplaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class ReplaceArgs : global::Pulumi.InvokeArgs
    {
        [Input("new", required: true)]
        public string New { get; set; } = null!;

        [Input("old", required: true)]
        public string Old { get; set; } = null!;

        [Input("string", required: true)]
        public string String { get; set; } = null!;

        public ReplaceArgs()
        {
        }
        public static new ReplaceArgs Empty => new ReplaceArgs();
    }

    public sealed class ReplaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("new", required: true)]
        public Input<string> New { get; set; } = null!;

        [Input("old", required: true)]
        public Input<string> Old { get; set; } = null!;

        [Input("string", required: true)]
        public Input<string> String { get; set; } = null!;

        public ReplaceInvokeArgs()
        {
        }
        public static new ReplaceInvokeArgs Empty => new ReplaceInvokeArgs();
    }


    [OutputType]
    public sealed class ReplaceResult
    {
        public readonly string Result;

        [OutputConstructor]
        private ReplaceResult(string result)
        {
            Result = result;
        }
    }
}
