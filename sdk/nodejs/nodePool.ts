// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Creates node pools for Kubernetes clusters in Symbiosis.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as symbiosis from "@pulumi/symbiosis";
 *
 * const exampleCluster = new symbiosis.Cluster("exampleCluster", {region: "germany-1"});
 * const exampleNodePool = new symbiosis.NodePool("exampleNodePool", {
 *     cluster: exampleCluster.name,
 *     nodeType: "general-1",
 *     quantity: 6,
 * });
 * ```
 */
export class NodePool extends pulumi.CustomResource {
    /**
     * Get an existing NodePool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodePoolState, opts?: pulumi.CustomResourceOptions): NodePool {
        return new NodePool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'symbiosis:index/nodePool:NodePool';

    /**
     * Returns true if the given object is an instance of NodePool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NodePool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NodePool.__pulumiType;
    }

    /**
     * Name of cluster to create node pool in.
     */
    public readonly cluster!: pulumi.Output<string>;
    /**
     * Node labels to be applied to the nodes
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name of node pool
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Type of nodes for this specific pool, see docs.
     */
    public readonly nodeType!: pulumi.Output<string>;
    /**
     * Desired number of nodes for specific pool.
     */
    public readonly quantity!: pulumi.Output<number>;
    /**
     * Node taints to be applied to the nodes
     */
    public readonly taints!: pulumi.Output<outputs.NodePoolTaint[] | undefined>;

    /**
     * Create a NodePool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodePoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodePoolArgs | NodePoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NodePoolState | undefined;
            resourceInputs["cluster"] = state ? state.cluster : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeType"] = state ? state.nodeType : undefined;
            resourceInputs["quantity"] = state ? state.quantity : undefined;
            resourceInputs["taints"] = state ? state.taints : undefined;
        } else {
            const args = argsOrState as NodePoolArgs | undefined;
            if ((!args || args.cluster === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cluster'");
            }
            if ((!args || args.nodeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeType'");
            }
            if ((!args || args.quantity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'quantity'");
            }
            resourceInputs["cluster"] = args ? args.cluster : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeType"] = args ? args.nodeType : undefined;
            resourceInputs["quantity"] = args ? args.quantity : undefined;
            resourceInputs["taints"] = args ? args.taints : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NodePool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NodePool resources.
 */
export interface NodePoolState {
    /**
     * Name of cluster to create node pool in.
     */
    cluster?: pulumi.Input<string>;
    /**
     * Node labels to be applied to the nodes
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of node pool
     */
    name?: pulumi.Input<string>;
    /**
     * Type of nodes for this specific pool, see docs.
     */
    nodeType?: pulumi.Input<string>;
    /**
     * Desired number of nodes for specific pool.
     */
    quantity?: pulumi.Input<number>;
    /**
     * Node taints to be applied to the nodes
     */
    taints?: pulumi.Input<pulumi.Input<inputs.NodePoolTaint>[]>;
}

/**
 * The set of arguments for constructing a NodePool resource.
 */
export interface NodePoolArgs {
    /**
     * Name of cluster to create node pool in.
     */
    cluster: pulumi.Input<string>;
    /**
     * Node labels to be applied to the nodes
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of node pool
     */
    name?: pulumi.Input<string>;
    /**
     * Type of nodes for this specific pool, see docs.
     */
    nodeType: pulumi.Input<string>;
    /**
     * Desired number of nodes for specific pool.
     */
    quantity: pulumi.Input<number>;
    /**
     * Node taints to be applied to the nodes
     */
    taints?: pulumi.Input<pulumi.Input<inputs.NodePoolTaint>[]>;
}
