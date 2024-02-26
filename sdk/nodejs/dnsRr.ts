// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * DNS RR resource allows to create and manage DNS resource records of type A, AAAA, PTR, CNAME, DNAME, NS.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as solidserver from "@pulumi/solidserver";
 *
 * const myFirstIPPTR = solidserver.getIpPtr({
 *     address: solidserver_ip_address.myFirstIPAddress.address,
 * });
 * const aaRecord = new solidserver.DnsRr("aaRecord", {
 *     dnsserver: "ns.mycompany.priv",
 *     dnsview: "Internal",
 *     dnszone: "mycompany.priv",
 *     type: "PTR",
 *     value: "myapp.mycompany.priv",
 * });
 * ```
 */
export class DnsRr extends pulumi.CustomResource {
    /**
     * Get an existing DnsRr resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DnsRrState, opts?: pulumi.CustomResourceOptions): DnsRr {
        return new DnsRr(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'solidserver:index/dnsRr:DnsRr';

    /**
     * Returns true if the given object is an instance of DnsRr.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DnsRr {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DnsRr.__pulumiType;
    }

    /**
     * The class associated to the DNS view.
     */
    public readonly class!: pulumi.Output<string | undefined>;
    /**
     * The class parameters associated to the view.
     */
    public readonly classParameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The managed SMART DNS server name, or DNS server name hosting the RR's zone.
     */
    public readonly dnsserver!: pulumi.Output<string>;
    /**
     * The View name of the RR to create.
     */
    public readonly dnsview!: pulumi.Output<string | undefined>;
    /**
     * The Zone name of the RR to create.
     */
    public readonly dnszone!: pulumi.Output<string | undefined>;
    /**
     * The Fully Qualified Domain Name of the RR to create.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The DNS Time To Live of the RR to create.
     */
    public readonly ttl!: pulumi.Output<number | undefined>;
    /**
     * The type of the RR to create (Supported: A, AAAA, PTR, CNAME, DNAME and NS).
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The value od the RR to create.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a DnsRr resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DnsRrArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DnsRrArgs | DnsRrState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DnsRrState | undefined;
            resourceInputs["class"] = state ? state.class : undefined;
            resourceInputs["classParameters"] = state ? state.classParameters : undefined;
            resourceInputs["dnsserver"] = state ? state.dnsserver : undefined;
            resourceInputs["dnsview"] = state ? state.dnsview : undefined;
            resourceInputs["dnszone"] = state ? state.dnszone : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as DnsRrArgs | undefined;
            if ((!args || args.dnsserver === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dnsserver'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["class"] = args ? args.class : undefined;
            resourceInputs["classParameters"] = args ? args.classParameters : undefined;
            resourceInputs["dnsserver"] = args ? args.dnsserver : undefined;
            resourceInputs["dnsview"] = args ? args.dnsview : undefined;
            resourceInputs["dnszone"] = args ? args.dnszone : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DnsRr.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DnsRr resources.
 */
export interface DnsRrState {
    /**
     * The class associated to the DNS view.
     */
    class?: pulumi.Input<string>;
    /**
     * The class parameters associated to the view.
     */
    classParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The managed SMART DNS server name, or DNS server name hosting the RR's zone.
     */
    dnsserver?: pulumi.Input<string>;
    /**
     * The View name of the RR to create.
     */
    dnsview?: pulumi.Input<string>;
    /**
     * The Zone name of the RR to create.
     */
    dnszone?: pulumi.Input<string>;
    /**
     * The Fully Qualified Domain Name of the RR to create.
     */
    name?: pulumi.Input<string>;
    /**
     * The DNS Time To Live of the RR to create.
     */
    ttl?: pulumi.Input<number>;
    /**
     * The type of the RR to create (Supported: A, AAAA, PTR, CNAME, DNAME and NS).
     */
    type?: pulumi.Input<string>;
    /**
     * The value od the RR to create.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DnsRr resource.
 */
export interface DnsRrArgs {
    /**
     * The class associated to the DNS view.
     */
    class?: pulumi.Input<string>;
    /**
     * The class parameters associated to the view.
     */
    classParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The managed SMART DNS server name, or DNS server name hosting the RR's zone.
     */
    dnsserver: pulumi.Input<string>;
    /**
     * The View name of the RR to create.
     */
    dnsview?: pulumi.Input<string>;
    /**
     * The Zone name of the RR to create.
     */
    dnszone?: pulumi.Input<string>;
    /**
     * The Fully Qualified Domain Name of the RR to create.
     */
    name?: pulumi.Input<string>;
    /**
     * The DNS Time To Live of the RR to create.
     */
    ttl?: pulumi.Input<number>;
    /**
     * The type of the RR to create (Supported: A, AAAA, PTR, CNAME, DNAME and NS).
     */
    type: pulumi.Input<string>;
    /**
     * The value od the RR to create.
     */
    value: pulumi.Input<string>;
}