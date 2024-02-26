// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * DNS View resource allows to create and configure DNS views.
 * View(s) are virutal containers mostly used to implement DNS split horizon
 * providing different answers depending on matching criterias.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as solidserver from "@pulumi/solidserver";
 *
 * const myFirstDnsView = new solidserver.DnsView("myFirstDnsView", {
 *     dnsserver: solidserver_dns_server.myFirstDnsServer.name,
 *     recursion: true,
 *     forward: "first",
 *     forwarders: [
 *         "8.8.8.8",
 *         "8.8.4.4",
 *     ],
 *     matchClients: [
 *         "172.16.0.0/12",
 *         "192.168.0.0/24",
 *     ],
 * }, {
 *     dependsOn: [solidserver_dns_server.myFirstDnsServer],
 * });
 * ```
 */
export class DnsView extends pulumi.CustomResource {
    /**
     * Get an existing DnsView resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DnsViewState, opts?: pulumi.CustomResourceOptions): DnsView {
        return new DnsView(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'solidserver:index/dnsView:DnsView';

    /**
     * Returns true if the given object is an instance of DnsView.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DnsView {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DnsView.__pulumiType;
    }

    /**
     * A list of network prefixes allowed to query the view (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    public readonly allowQueries!: pulumi.Output<string[] | undefined>;
    /**
     * A list of network prefixes allowed to query the view for recursion (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    public readonly allowRecursions!: pulumi.Output<string[] | undefined>;
    /**
     * A list of network prefixes allowed to query the view for zone transfert (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    public readonly allowTransfers!: pulumi.Output<string[] | undefined>;
    /**
     * The class associated to the DNS view.
     */
    public readonly class!: pulumi.Output<string | undefined>;
    /**
     * The class parameters associated to the view.
     */
    public readonly classParameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of DNS server or DNS SMART hosting the DNS view to create.
     */
    public readonly dnsserver!: pulumi.Output<string>;
    /**
     * The forwarding mode of the DNS SMART (Supported: none, first, only; Default: none).
     */
    public readonly forward!: pulumi.Output<string | undefined>;
    /**
     * The IP address list of the forwarder(s) configured to configure on the DNS SMART.
     */
    public readonly forwarders!: pulumi.Output<string[] | undefined>;
    /**
     * A list of network prefixes used to match the clients of the view (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    public readonly matchClients!: pulumi.Output<string[] | undefined>;
    /**
     * A list of network prefixes used to match the traffic to the view (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    public readonly matchTos!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the DNS view to create.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The level of the DNS view, where 0 represents the highest level in the views hierarchy.
     */
    public /*out*/ readonly order!: pulumi.Output<number>;
    /**
     * The recursion mode of the DNS view (Default: true).
     */
    public readonly recursion!: pulumi.Output<boolean | undefined>;

    /**
     * Create a DnsView resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DnsViewArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DnsViewArgs | DnsViewState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DnsViewState | undefined;
            resourceInputs["allowQueries"] = state ? state.allowQueries : undefined;
            resourceInputs["allowRecursions"] = state ? state.allowRecursions : undefined;
            resourceInputs["allowTransfers"] = state ? state.allowTransfers : undefined;
            resourceInputs["class"] = state ? state.class : undefined;
            resourceInputs["classParameters"] = state ? state.classParameters : undefined;
            resourceInputs["dnsserver"] = state ? state.dnsserver : undefined;
            resourceInputs["forward"] = state ? state.forward : undefined;
            resourceInputs["forwarders"] = state ? state.forwarders : undefined;
            resourceInputs["matchClients"] = state ? state.matchClients : undefined;
            resourceInputs["matchTos"] = state ? state.matchTos : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["order"] = state ? state.order : undefined;
            resourceInputs["recursion"] = state ? state.recursion : undefined;
        } else {
            const args = argsOrState as DnsViewArgs | undefined;
            if ((!args || args.dnsserver === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dnsserver'");
            }
            resourceInputs["allowQueries"] = args ? args.allowQueries : undefined;
            resourceInputs["allowRecursions"] = args ? args.allowRecursions : undefined;
            resourceInputs["allowTransfers"] = args ? args.allowTransfers : undefined;
            resourceInputs["class"] = args ? args.class : undefined;
            resourceInputs["classParameters"] = args ? args.classParameters : undefined;
            resourceInputs["dnsserver"] = args ? args.dnsserver : undefined;
            resourceInputs["forward"] = args ? args.forward : undefined;
            resourceInputs["forwarders"] = args ? args.forwarders : undefined;
            resourceInputs["matchClients"] = args ? args.matchClients : undefined;
            resourceInputs["matchTos"] = args ? args.matchTos : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["recursion"] = args ? args.recursion : undefined;
            resourceInputs["order"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DnsView.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DnsView resources.
 */
export interface DnsViewState {
    /**
     * A list of network prefixes allowed to query the view (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    allowQueries?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of network prefixes allowed to query the view for recursion (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    allowRecursions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of network prefixes allowed to query the view for zone transfert (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    allowTransfers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The class associated to the DNS view.
     */
    class?: pulumi.Input<string>;
    /**
     * The class parameters associated to the view.
     */
    classParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of DNS server or DNS SMART hosting the DNS view to create.
     */
    dnsserver?: pulumi.Input<string>;
    /**
     * The forwarding mode of the DNS SMART (Supported: none, first, only; Default: none).
     */
    forward?: pulumi.Input<string>;
    /**
     * The IP address list of the forwarder(s) configured to configure on the DNS SMART.
     */
    forwarders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of network prefixes used to match the clients of the view (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    matchClients?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of network prefixes used to match the traffic to the view (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    matchTos?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the DNS view to create.
     */
    name?: pulumi.Input<string>;
    /**
     * The level of the DNS view, where 0 represents the highest level in the views hierarchy.
     */
    order?: pulumi.Input<number>;
    /**
     * The recursion mode of the DNS view (Default: true).
     */
    recursion?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a DnsView resource.
 */
export interface DnsViewArgs {
    /**
     * A list of network prefixes allowed to query the view (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    allowQueries?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of network prefixes allowed to query the view for recursion (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    allowRecursions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of network prefixes allowed to query the view for zone transfert (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    allowTransfers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The class associated to the DNS view.
     */
    class?: pulumi.Input<string>;
    /**
     * The class parameters associated to the view.
     */
    classParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of DNS server or DNS SMART hosting the DNS view to create.
     */
    dnsserver: pulumi.Input<string>;
    /**
     * The forwarding mode of the DNS SMART (Supported: none, first, only; Default: none).
     */
    forward?: pulumi.Input<string>;
    /**
     * The IP address list of the forwarder(s) configured to configure on the DNS SMART.
     */
    forwarders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of network prefixes used to match the clients of the view (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    matchClients?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of network prefixes used to match the traffic to the view (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    matchTos?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the DNS view to create.
     */
    name?: pulumi.Input<string>;
    /**
     * The recursion mode of the DNS view (Default: true).
     */
    recursion?: pulumi.Input<boolean>;
}
