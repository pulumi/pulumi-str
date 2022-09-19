// *** WARNING: this file was generated by pulumi-gen-str. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Str.Regexp
{
    public static class Match
    {
        /// <summary>
        /// Match reports whether the string s contains any match of the regular expression pattern.
        /// </summary>
        public static Task<MatchResult> InvokeAsync(MatchArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<MatchResult>("str:regexp:match", args ?? new MatchArgs(), options.WithDefaults());

        /// <summary>
        /// Match reports whether the string s contains any match of the regular expression pattern.
        /// </summary>
        public static Output<MatchResult> Invoke(MatchInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<MatchResult>("str:regexp:match", args ?? new MatchInvokeArgs(), options.WithDefaults());
    }


    public sealed class MatchArgs : global::Pulumi.InvokeArgs
    {
        [Input("pattern", required: true)]
        public string Pattern { get; set; } = null!;

        [Input("string", required: true)]
        public string String { get; set; } = null!;

        public MatchArgs()
        {
        }
        public static new MatchArgs Empty => new MatchArgs();
    }

    public sealed class MatchInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("pattern", required: true)]
        public Input<string> Pattern { get; set; } = null!;

        [Input("string", required: true)]
        public Input<string> String { get; set; } = null!;

        public MatchInvokeArgs()
        {
        }
        public static new MatchInvokeArgs Empty => new MatchInvokeArgs();
    }


    [OutputType]
    public sealed class MatchResult
    {
        public readonly bool Matches;

        [OutputConstructor]
        private MatchResult(bool matches)
        {
            Matches = matches;
        }
    }
}