// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * IP aliases resource allows to create and manage multiple names for a single IP address.
 * They are pretty useful to keep IPAM in sync with the DNS handling CNAME(s) from a single repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as solidserver from "@pulumi/solidserver";
 *
 * const myFirstIPAlias = new solidserver.IpAlias("myFirstIPAlias", {
 *     address: solidserver_ip_address.myFirstIPAddress.address,
 *     space: solidserver_ip_space.myFirstSpace.name,
 * });
 * ```
 */
export class IpAlias extends pulumi.CustomResource {
    /**
     * Get an existing IpAlias resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpAliasState, opts?: pulumi.CustomResourceOptions): IpAlias {
        return new IpAlias(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'solidserver:index/ipAlias:IpAlias';

    /**
     * Returns true if the given object is an instance of IpAlias.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpAlias {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpAlias.__pulumiType;
    }

    /**
     * The IP address for which the alias will be associated to.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * The FQDN of the IP address alias to create.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the space to which the address belong to.
     */
    public readonly space!: pulumi.Output<string>;
    /**
     * The type of the Alias to create (Supported: A, CNAME; Default: CNAME).
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a IpAlias resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpAliasArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpAliasArgs | IpAliasState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpAliasState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["space"] = state ? state.space : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as IpAliasArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.space === undefined) && !opts.urn) {
                throw new Error("Missing required property 'space'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["space"] = args ? args.space : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpAlias.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpAlias resources.
 */
export interface IpAliasState {
    /**
     * The IP address for which the alias will be associated to.
     */
    address?: pulumi.Input<string>;
    /**
     * The FQDN of the IP address alias to create.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the space to which the address belong to.
     */
    space?: pulumi.Input<string>;
    /**
     * The type of the Alias to create (Supported: A, CNAME; Default: CNAME).
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpAlias resource.
 */
export interface IpAliasArgs {
    /**
     * The IP address for which the alias will be associated to.
     */
    address: pulumi.Input<string>;
    /**
     * The FQDN of the IP address alias to create.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the space to which the address belong to.
     */
    space: pulumi.Input<string>;
    /**
     * The type of the Alias to create (Supported: A, CNAME; Default: CNAME).
     */
    type?: pulumi.Input<string>;
}