// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * DNS Server resource allows to register and configure DNS servers.
 * Most of the time, they are just added to a SMART, but they can remain standalone.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as solidserver from "@pulumi/solidserver";
 *
 * const myFirstDnsServer = new solidserver.DnsServer("myFirstDnsServer", {
 *     address: "127.0.0.1",
 *     allowQueries: [
 *         "172.16.0.0/12",
 *         "10.0.0.0/8",
 *         "192.168.0.0/24",
 *     ],
 *     allowRecursions: [
 *         "172.16.0.0/12",
 *         "10.0.0.0/8",
 *         "192.168.0.0/24",
 *     ],
 *     comment: "My First DNS Server Autmatically created",
 *     forward: "first",
 *     forwarders: [
 *         "10.0.0.42",
 *         "10.0.0.43",
 *     ],
 *     login: "admin",
 *     password: "admin",
 *     smart: solidserver_dns_smart.myFirstDnsSMART.name,
 *     smartRole: "master",
 * });
 * ```
 */
export class DnsServer extends pulumi.CustomResource {
    /**
     * Get an existing DnsServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DnsServerState, opts?: pulumi.CustomResourceOptions): DnsServer {
        return new DnsServer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'solidserver:index/dnsServer:DnsServer';

    /**
     * Returns true if the given object is an instance of DnsServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DnsServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DnsServer.__pulumiType;
    }

    /**
     * The IPv4 address of the DNS server to create.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * A list of network prefixes allowed to query the DNS server (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    public readonly allowQueries!: pulumi.Output<string[] | undefined>;
    /**
     * A list of network prefixes allowed to query the DNS server for recursion (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    public readonly allowRecursions!: pulumi.Output<string[] | undefined>;
    /**
     * A list of network prefixes allowed to query the DNS server for zone transfert (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    public readonly allowTransfers!: pulumi.Output<string[] | undefined>;
    /**
     * The class associated to the DNS server.
     */
    public readonly class!: pulumi.Output<string | undefined>;
    /**
     * The class parameters associated to the DNS server.
     */
    public readonly classParameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Custom information about the DNS server.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * The forwarding mode of the DNS server (Supported: none, first, only; Default: none).
     */
    public readonly forward!: pulumi.Output<string | undefined>;
    /**
     * The list of forwarders' IP address to be used by the DNS server.
     */
    public readonly forwarders!: pulumi.Output<string[] | undefined>;
    /**
     * The login to use for enrolling of the DNS server.
     */
    public readonly login!: pulumi.Output<string>;
    /**
     * The name of the DNS server to create.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The password to use the enrolling of the DNS server.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The recursion mode of the DNS server (Default: true).
     */
    public readonly recursion!: pulumi.Output<boolean | undefined>;
    /**
     * The DNS SMART the DNS server must join.
     */
    public readonly smart!: pulumi.Output<string | undefined>;
    /**
     * The role the DNS server will play within the SMART (Supported: master, slave; Default: slave).
     */
    public readonly smartRole!: pulumi.Output<string | undefined>;
    /**
     * The type of DNS server (Supported: ipm (SOLIDserver or Linux Package); Default: ipm).
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DnsServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DnsServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DnsServerArgs | DnsServerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DnsServerState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["allowQueries"] = state ? state.allowQueries : undefined;
            resourceInputs["allowRecursions"] = state ? state.allowRecursions : undefined;
            resourceInputs["allowTransfers"] = state ? state.allowTransfers : undefined;
            resourceInputs["class"] = state ? state.class : undefined;
            resourceInputs["classParameters"] = state ? state.classParameters : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["forward"] = state ? state.forward : undefined;
            resourceInputs["forwarders"] = state ? state.forwarders : undefined;
            resourceInputs["login"] = state ? state.login : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["recursion"] = state ? state.recursion : undefined;
            resourceInputs["smart"] = state ? state.smart : undefined;
            resourceInputs["smartRole"] = state ? state.smartRole : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as DnsServerArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.login === undefined) && !opts.urn) {
                throw new Error("Missing required property 'login'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["allowQueries"] = args ? args.allowQueries : undefined;
            resourceInputs["allowRecursions"] = args ? args.allowRecursions : undefined;
            resourceInputs["allowTransfers"] = args ? args.allowTransfers : undefined;
            resourceInputs["class"] = args ? args.class : undefined;
            resourceInputs["classParameters"] = args ? args.classParameters : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["forward"] = args ? args.forward : undefined;
            resourceInputs["forwarders"] = args ? args.forwarders : undefined;
            resourceInputs["login"] = args ? args.login : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["recursion"] = args ? args.recursion : undefined;
            resourceInputs["smart"] = args ? args.smart : undefined;
            resourceInputs["smartRole"] = args ? args.smartRole : undefined;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DnsServer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DnsServer resources.
 */
export interface DnsServerState {
    /**
     * The IPv4 address of the DNS server to create.
     */
    address?: pulumi.Input<string>;
    /**
     * A list of network prefixes allowed to query the DNS server (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    allowQueries?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of network prefixes allowed to query the DNS server for recursion (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    allowRecursions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of network prefixes allowed to query the DNS server for zone transfert (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    allowTransfers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The class associated to the DNS server.
     */
    class?: pulumi.Input<string>;
    /**
     * The class parameters associated to the DNS server.
     */
    classParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Custom information about the DNS server.
     */
    comment?: pulumi.Input<string>;
    /**
     * The forwarding mode of the DNS server (Supported: none, first, only; Default: none).
     */
    forward?: pulumi.Input<string>;
    /**
     * The list of forwarders' IP address to be used by the DNS server.
     */
    forwarders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The login to use for enrolling of the DNS server.
     */
    login?: pulumi.Input<string>;
    /**
     * The name of the DNS server to create.
     */
    name?: pulumi.Input<string>;
    /**
     * The password to use the enrolling of the DNS server.
     */
    password?: pulumi.Input<string>;
    /**
     * The recursion mode of the DNS server (Default: true).
     */
    recursion?: pulumi.Input<boolean>;
    /**
     * The DNS SMART the DNS server must join.
     */
    smart?: pulumi.Input<string>;
    /**
     * The role the DNS server will play within the SMART (Supported: master, slave; Default: slave).
     */
    smartRole?: pulumi.Input<string>;
    /**
     * The type of DNS server (Supported: ipm (SOLIDserver or Linux Package); Default: ipm).
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DnsServer resource.
 */
export interface DnsServerArgs {
    /**
     * The IPv4 address of the DNS server to create.
     */
    address: pulumi.Input<string>;
    /**
     * A list of network prefixes allowed to query the DNS server (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    allowQueries?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of network prefixes allowed to query the DNS server for recursion (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    allowRecursions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of network prefixes allowed to query the DNS server for zone transfert (named ACL(s) are not supported using this provider).  Use '!' to negate an entry.
     */
    allowTransfers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The class associated to the DNS server.
     */
    class?: pulumi.Input<string>;
    /**
     * The class parameters associated to the DNS server.
     */
    classParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Custom information about the DNS server.
     */
    comment?: pulumi.Input<string>;
    /**
     * The forwarding mode of the DNS server (Supported: none, first, only; Default: none).
     */
    forward?: pulumi.Input<string>;
    /**
     * The list of forwarders' IP address to be used by the DNS server.
     */
    forwarders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The login to use for enrolling of the DNS server.
     */
    login: pulumi.Input<string>;
    /**
     * The name of the DNS server to create.
     */
    name?: pulumi.Input<string>;
    /**
     * The password to use the enrolling of the DNS server.
     */
    password: pulumi.Input<string>;
    /**
     * The recursion mode of the DNS server (Default: true).
     */
    recursion?: pulumi.Input<boolean>;
    /**
     * The DNS SMART the DNS server must join.
     */
    smart?: pulumi.Input<string>;
    /**
     * The role the DNS server will play within the SMART (Supported: master, slave; Default: slave).
     */
    smartRole?: pulumi.Input<string>;
}
