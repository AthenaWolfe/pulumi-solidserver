// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * User resource allows to associate users with groups managing
 * SOLIDserver permissions.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as solidserver from "@pulumi/solidserver";
 *
 * const tGroup01 = new solidserver.Usergroup("tGroup01", {description: "descr01"});
 * ```
 */
export class Usergroup extends pulumi.CustomResource {
    /**
     * Get an existing Usergroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UsergroupState, opts?: pulumi.CustomResourceOptions): Usergroup {
        return new Usergroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'solidserver:index/usergroup:Usergroup';

    /**
     * Returns true if the given object is an instance of Usergroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Usergroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Usergroup.__pulumiType;
    }

    /**
     * The description of the group
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the group
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Usergroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UsergroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UsergroupArgs | UsergroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UsergroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as UsergroupArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Usergroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Usergroup resources.
 */
export interface UsergroupState {
    /**
     * The description of the group
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the group
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Usergroup resource.
 */
export interface UsergroupArgs {
    /**
     * The description of the group
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the group
     */
    name?: pulumi.Input<string>;
}
