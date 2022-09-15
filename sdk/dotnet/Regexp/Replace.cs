// *** WARNING: this file was generated by pulumi-gen-str. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Str.Regexp
{
    public static class Replace
    {
        /// <summary>
        /// A regex based replace on a string.
        /// 
        /// The underlying regexp engine is the go "regexp" stdlib package.
        /// You can see details at for available patterns at https://pkg.go.dev/regexp/syntax.
        /// </summary>
        public static Task<ReplaceResult> InvokeAsync(ReplaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ReplaceResult>("str:regexp:replace", args ?? new ReplaceArgs(), options.WithDefaults());

        /// <summary>
        /// A regex based replace on a string.
        /// 
        /// The underlying regexp engine is the go "regexp" stdlib package.
        /// You can see details at for available patterns at https://pkg.go.dev/regexp/syntax.
        /// </summary>
        public static Output<ReplaceResult> Invoke(ReplaceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<ReplaceResult>("str:regexp:replace", args ?? new ReplaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class ReplaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The new string.
        /// 
        /// Note: Inside repl, "$" signs are interpreted as an Expand, so for instance $1 represents the text of the first submatch. 
        /// </summary>
        [Input("new", required: true)]
        public string New { get; set; } = null!;

        /// <summary>
        /// The regular expression to match against.
        /// </summary>
        [Input("old", required: true)]
        public string Old { get; set; } = null!;

        /// <summary>
        /// The string to operate over.
        /// </summary>
        [Input("string", required: true)]
        public string String { get; set; } = null!;

        public ReplaceArgs()
        {
        }
        public static new ReplaceArgs Empty => new ReplaceArgs();
    }

    public sealed class ReplaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The new string.
        /// 
        /// Note: Inside repl, "$" signs are interpreted as an Expand, so for instance $1 represents the text of the first submatch. 
        /// </summary>
        [Input("new", required: true)]
        public Input<string> New { get; set; } = null!;

        /// <summary>
        /// The regular expression to match against.
        /// </summary>
        [Input("old", required: true)]
        public Input<string> Old { get; set; } = null!;

        /// <summary>
        /// The string to operate over.
        /// </summary>
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
        /// <summary>
        /// The input "string" where each pattern matching "old" is replaced by "new".
        /// </summary>
        public readonly string Result;

        [OutputConstructor]
        private ReplaceResult(string result)
        {
            Result = result;
        }
    }
}