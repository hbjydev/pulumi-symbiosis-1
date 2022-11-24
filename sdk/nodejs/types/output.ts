// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface NodePoolTaint {
    /**
     * Taint effect. Can be either NoSchedule, PreferNoSchedule or NoExecute. See: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/
     */
    effect: string;
    key: string;
    value: string;
}

