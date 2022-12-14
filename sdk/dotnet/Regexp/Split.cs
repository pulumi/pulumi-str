// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Str.Regexp
{
    public static class Split
    {
        /// <summary>
        /// Split a string on a regex.
        /// </summary>
        public static Task<SplitResult> InvokeAsync(SplitArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<SplitResult>("str:regexp:split", args ?? new SplitArgs(), options.WithDefaults());

        /// <summary>
        /// Split a string on a regex.
        /// </summary>
        public static Output<SplitResult> Invoke(SplitInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<SplitResult>("str:regexp:split", args ?? new SplitInvokeArgs(), options.WithDefaults());
    }


    public sealed class SplitArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// `count` determines the number of substrings to return. 
        /// If `count` is not provided, it defaults to substrings.
        /// If `count` is provided then the last substring will be the unsplit remainder.
        /// It is an error to pass `count &lt; 1`.
        /// </summary>
        [Input("count")]
        public int? Count { get; set; }

        /// <summary>
        /// The regex to split on.
        /// </summary>
        [Input("on", required: true)]
        public string On { get; set; } = null!;

        /// <summary>
        /// The string on which to split.
        /// </summary>
        [Input("string", required: true)]
        public string String { get; set; } = null!;

        public SplitArgs()
        {
        }
        public static new SplitArgs Empty => new SplitArgs();
    }

    public sealed class SplitInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// `count` determines the number of substrings to return. 
        /// If `count` is not provided, it defaults to substrings.
        /// If `count` is provided then the last substring will be the unsplit remainder.
        /// It is an error to pass `count &lt; 1`.
        /// </summary>
        [Input("count")]
        public Input<int>? Count { get; set; }

        /// <summary>
        /// The regex to split on.
        /// </summary>
        [Input("on", required: true)]
        public Input<string> On { get; set; } = null!;

        /// <summary>
        /// The string on which to split.
        /// </summary>
        [Input("string", required: true)]
        public Input<string> String { get; set; } = null!;

        public SplitInvokeArgs()
        {
        }
        public static new SplitInvokeArgs Empty => new SplitInvokeArgs();
    }


    [OutputType]
    public sealed class SplitResult
    {
        /// <summary>
        /// The result of the string split.
        /// </summary>
        public readonly ImmutableArray<string> Result;

        [OutputConstructor]
        private SplitResult(ImmutableArray<string> result)
        {
            Result = result;
        }
    }
}
