// *** WARNING: this file was generated by pulumi-gen-str. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Str
{
    public static class TrimSuffix
    {
        /// <summary>
        /// Trim a suffix from a string.
        /// </summary>
        public static Task<TrimSuffixResult> InvokeAsync(TrimSuffixArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<TrimSuffixResult>("str:index:trimSuffix", args ?? new TrimSuffixArgs(), options.WithDefaults());

        /// <summary>
        /// Trim a suffix from a string.
        /// </summary>
        public static Output<TrimSuffixResult> Invoke(TrimSuffixInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<TrimSuffixResult>("str:index:trimSuffix", args ?? new TrimSuffixInvokeArgs(), options.WithDefaults());
    }


    public sealed class TrimSuffixArgs : global::Pulumi.InvokeArgs
    {
        [Input("string", required: true)]
        public string String { get; set; } = null!;

        [Input("suffix", required: true)]
        public string Suffix { get; set; } = null!;

        public TrimSuffixArgs()
        {
        }
        public static new TrimSuffixArgs Empty => new TrimSuffixArgs();
    }

    public sealed class TrimSuffixInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("string", required: true)]
        public Input<string> String { get; set; } = null!;

        [Input("suffix", required: true)]
        public Input<string> Suffix { get; set; } = null!;

        public TrimSuffixInvokeArgs()
        {
        }
        public static new TrimSuffixInvokeArgs Empty => new TrimSuffixInvokeArgs();
    }


    [OutputType]
    public sealed class TrimSuffixResult
    {
        public readonly string Result;

        [OutputConstructor]
        private TrimSuffixResult(string result)
        {
            Result = result;
        }
    }
}
