// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Trim a suffix from a string.
 */
export function trimSuffix(args: TrimSuffixArgs, opts?: pulumi.InvokeOptions): Promise<TrimSuffixResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("str:index:trimSuffix", {
        "string": args.string,
        "suffix": args.suffix,
    }, opts);
}

export interface TrimSuffixArgs {
    string: string;
    suffix: string;
}

export interface TrimSuffixResult {
    readonly result: string;
}

export function trimSuffixOutput(args: TrimSuffixOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<TrimSuffixResult> {
    return pulumi.output(args).apply(a => trimSuffix(a, opts))
}

export interface TrimSuffixOutputArgs {
    string: pulumi.Input<string>;
    suffix: pulumi.Input<string>;
}
