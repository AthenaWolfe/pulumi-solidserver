// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Space data-source allows to retrieve information about reserved IP spaces, including meta-data.
 * Spaces are the highest level objets in the SOLIDserver's IPAM module organization,
 * the entry point of any IPv4 or IPv6 addressing plan. It allows to manage unique ranges
 * of IP addresses.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as solidserver from "@pulumi/solidserver";
 *
 * const myFirstSpaceData = solidserver.getIpSpace({
 *     name: solidserver_ip_space.myFirstSpace.name,
 * });
 * ```
 */
export function getIpSpace(args: GetIpSpaceArgs, opts?: pulumi.InvokeOptions): Promise<GetIpSpaceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("solidserver:index/getIpSpace:getIpSpace", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpSpace.
 */
export interface GetIpSpaceArgs {
    /**
     * The name of the IP space.
     */
    name: string;
}

/**
 * A collection of values returned by getIpSpace.
 */
export interface GetIpSpaceResult {
    /**
     * The class associated to the IP space.
     */
    readonly class: string;
    /**
     * The class parameters associated to IP space.
     */
    readonly classParameters: {[key: string]: any};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the IP space.
     */
    readonly name: string;
}
/**
 * Space data-source allows to retrieve information about reserved IP spaces, including meta-data.
 * Spaces are the highest level objets in the SOLIDserver's IPAM module organization,
 * the entry point of any IPv4 or IPv6 addressing plan. It allows to manage unique ranges
 * of IP addresses.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as solidserver from "@pulumi/solidserver";
 *
 * const myFirstSpaceData = solidserver.getIpSpace({
 *     name: solidserver_ip_space.myFirstSpace.name,
 * });
 * ```
 */
export function getIpSpaceOutput(args: GetIpSpaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIpSpaceResult> {
    return pulumi.output(args).apply((a: any) => getIpSpace(a, opts))
}

/**
 * A collection of arguments for invoking getIpSpace.
 */
export interface GetIpSpaceOutputArgs {
    /**
     * The name of the IP space.
     */
    name: pulumi.Input<string>;
}