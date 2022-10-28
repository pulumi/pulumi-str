// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Trim a prefix from a string.
 */
export function trimPrefix(args: TrimPrefixArgs, opts?: pulumi.InvokeOptions): Promise<TrimPrefixResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("str:index:trimPrefix", {
        "prefix": args.prefix,
        "string": args.string,
    }, opts);
}

export interface TrimPrefixArgs {
    prefix: string;
    string: string;
}

export interface TrimPrefixResult {
    readonly result: string;
}

export function trimPrefixOutput(args: TrimPrefixOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<TrimPrefixResult> {
    return pulumi.output(args).apply(a => trimPrefix(a, opts))
}

export interface TrimPrefixOutputArgs {
    prefix: pulumi.Input<string>;
    string: pulumi.Input<string>;
}
