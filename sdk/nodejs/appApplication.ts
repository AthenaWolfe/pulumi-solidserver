// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Application resource allows to create and manage applications that can be used to implement traffic policies in order
 * to optimize the routing of the associated traffic according to the selected loadbalancing strategy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as solidserver from "@pulumi/solidserver";
 *
 * const myFirstApplicaton = new solidserver.AppApplication("myFirstApplicaton", {
 *     "class": "INTERNAL_APP",
 *     classParameters: {
 *         contact: "a.smith@mycompany.priv",
 *         owner: "MR. Smith",
 *     },
 *     fqdn: "myfirstapp.priv",
 *     gslbMembers: [
 *         "ns0.priv",
 *         "ns1.priv",
 *     ],
 * });
 * ```
 */
export class AppApplication extends pulumi.CustomResource {
    /**
     * Get an existing AppApplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppApplicationState, opts?: pulumi.CustomResourceOptions): AppApplication {
        return new AppApplication(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'solidserver:index/appApplication:AppApplication';

    /**
     * Returns true if the given object is an instance of AppApplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppApplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppApplication.__pulumiType;
    }

    /**
     * The class associated to the application.
     */
    public readonly class!: pulumi.Output<string | undefined>;
    /**
     * The class parameters associated to application.
     */
    public readonly classParameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Fully Qualified Domain Name of the application to create.
     */
    public readonly fqdn!: pulumi.Output<string>;
    /**
     * The names of the GSLB servers applying the application traffic policy.
     */
    public readonly gslbMembers!: pulumi.Output<string[]>;
    /**
     * The name of the application to create.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a AppApplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppApplicationArgs | AppApplicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppApplicationState | undefined;
            resourceInputs["class"] = state ? state.class : undefined;
            resourceInputs["classParameters"] = state ? state.classParameters : undefined;
            resourceInputs["fqdn"] = state ? state.fqdn : undefined;
            resourceInputs["gslbMembers"] = state ? state.gslbMembers : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as AppApplicationArgs | undefined;
            if ((!args || args.fqdn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fqdn'");
            }
            if ((!args || args.gslbMembers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gslbMembers'");
            }
            resourceInputs["class"] = args ? args.class : undefined;
            resourceInputs["classParameters"] = args ? args.classParameters : undefined;
            resourceInputs["fqdn"] = args ? args.fqdn : undefined;
            resourceInputs["gslbMembers"] = args ? args.gslbMembers : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AppApplication.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppApplication resources.
 */
export interface AppApplicationState {
    /**
     * The class associated to the application.
     */
    class?: pulumi.Input<string>;
    /**
     * The class parameters associated to application.
     */
    classParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Fully Qualified Domain Name of the application to create.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * The names of the GSLB servers applying the application traffic policy.
     */
    gslbMembers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the application to create.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppApplication resource.
 */
export interface AppApplicationArgs {
    /**
     * The class associated to the application.
     */
    class?: pulumi.Input<string>;
    /**
     * The class parameters associated to application.
     */
    classParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Fully Qualified Domain Name of the application to create.
     */
    fqdn: pulumi.Input<string>;
    /**
     * The names of the GSLB servers applying the application traffic policy.
     */
    gslbMembers: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the application to create.
     */
    name?: pulumi.Input<string>;
}