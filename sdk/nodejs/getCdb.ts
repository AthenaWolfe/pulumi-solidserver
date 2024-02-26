// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Custom DB data-source allows to retrieve information about custom database stored within SOLIDserver.
 * This custom database entries can be leveraged within object classes and wizards in order to store custom meta-data.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as solidserver from "@pulumi/solidserver";
 *
 * const myFirstCustomDB = solidserver.getCdb({
 *     name: "myFirstCustomDB",
 * });
 * ```
 */
export function getCdb(args: GetCdbArgs, opts?: pulumi.InvokeOptions): Promise<GetCdbResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("solidserver:index/getCdb:getCdb", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCdb.
 */
export interface GetCdbArgs {
    /**
     * The name of the custom DB.
     */
    name: string;
}

/**
 * A collection of values returned by getCdb.
 */
export interface GetCdbResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the label 1
     */
    readonly label1: string;
    /**
     * The name of the label 10
     */
    readonly label10: string;
    /**
     * The name of the label 2
     */
    readonly label2: string;
    /**
     * The name of the label 3
     */
    readonly label3: string;
    /**
     * The name of the label 4
     */
    readonly label4: string;
    /**
     * The name of the label 5
     */
    readonly label5: string;
    /**
     * The name of the label 6
     */
    readonly label6: string;
    /**
     * The name of the label 7
     */
    readonly label7: string;
    /**
     * The name of the label 8
     */
    readonly label8: string;
    /**
     * The name of the label 9
     */
    readonly label9: string;
    /**
     * The name of the custom DB.
     */
    readonly name: string;
}
/**
 * Custom DB data-source allows to retrieve information about custom database stored within SOLIDserver.
 * This custom database entries can be leveraged within object classes and wizards in order to store custom meta-data.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as solidserver from "@pulumi/solidserver";
 *
 * const myFirstCustomDB = solidserver.getCdb({
 *     name: "myFirstCustomDB",
 * });
 * ```
 */
export function getCdbOutput(args: GetCdbOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCdbResult> {
    return pulumi.output(args).apply((a: any) => getCdb(a, opts))
}

/**
 * A collection of arguments for invoking getCdb.
 */
export interface GetCdbOutputArgs {
    /**
     * The name of the custom DB.
     */
    name: pulumi.Input<string>;
}
