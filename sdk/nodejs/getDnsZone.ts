// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * DNS Zone data-source allows to retrieve information about DNS zones.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as solidserver from "@pulumi/solidserver";
 *
 * const myFirstDnsZoneData = solidserver.getDnsZone({
 *     name: solidserver_dns_zone.myFirstZone.name,
 * });
 * ```
 */
export function getDnsZone(args: GetDnsZoneArgs, opts?: pulumi.InvokeOptions): Promise<GetDnsZoneResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("solidserver:index/getDnsZone:getDnsZone", {
        "dnsview": args.dnsview,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getDnsZone.
 */
export interface GetDnsZoneArgs {
    /**
     * The name of DNS view hosting the DNS zone.
     */
    dnsview?: string;
    /**
     * The Domain Name served by the DNS zone.
     */
    name: string;
}

/**
 * A collection of values returned by getDnsZone.
 */
export interface GetDnsZoneResult {
    /**
     * The class associated to the DNS zone.
     */
    readonly class: string;
    /**
     * The class parameters associated to IP space.
     */
    readonly classParameters: {[key: string]: any};
    /**
     * Automaticaly create PTR records for the DNS zone.
     */
    readonly createptr: boolean;
    /**
     * The name of DNS server or DNS SMART hosting the DNS zone.
     */
    readonly dnsserver: string;
    /**
     * The name of DNS view hosting the DNS zone.
     */
    readonly dnsview?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Domain Name served by the DNS zone.
     */
    readonly name: string;
    /**
     * The name of a space associated to the DNS zone.
     */
    readonly space: string;
    /**
     * The Type of the DNS zone.
     */
    readonly type: string;
}
/**
 * DNS Zone data-source allows to retrieve information about DNS zones.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as solidserver from "@pulumi/solidserver";
 *
 * const myFirstDnsZoneData = solidserver.getDnsZone({
 *     name: solidserver_dns_zone.myFirstZone.name,
 * });
 * ```
 */
export function getDnsZoneOutput(args: GetDnsZoneOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDnsZoneResult> {
    return pulumi.output(args).apply((a: any) => getDnsZone(a, opts))
}

/**
 * A collection of arguments for invoking getDnsZone.
 */
export interface GetDnsZoneOutputArgs {
    /**
     * The name of DNS view hosting the DNS zone.
     */
    dnsview?: pulumi.Input<string>;
    /**
     * The Domain Name served by the DNS zone.
     */
    name: pulumi.Input<string>;
}