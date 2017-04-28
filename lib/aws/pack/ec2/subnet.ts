// *** WARNING: this file was generated by the Coconut IDL Compiler (CIDLC).  ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as coconut from "@coconut/coconut";

import {VPC} from "./vpc";

export class Subnet extends coconut.Resource implements SubnetArgs {
    public readonly name: string;
    public readonly cidrBlock: string;
    public readonly vpc: VPC;
    public readonly availabilityZone?: string;
    public mapPublicIpOnLaunch?: boolean;

    constructor(args: SubnetArgs) {
        super();
        if (args.name === undefined) {
            throw new Error("Missing required argument 'name'");
        }
        this.name = args.name;
        if (args.cidrBlock === undefined) {
            throw new Error("Missing required argument 'cidrBlock'");
        }
        this.cidrBlock = args.cidrBlock;
        if (args.vpc === undefined) {
            throw new Error("Missing required argument 'vpc'");
        }
        this.vpc = args.vpc;
        this.availabilityZone = args.availabilityZone;
        this.mapPublicIpOnLaunch = args.mapPublicIpOnLaunch;
    }
}

export interface SubnetArgs {
    readonly name: string;
    readonly cidrBlock: string;
    readonly vpc: VPC;
    readonly availabilityZone?: string;
    mapPublicIpOnLaunch?: boolean;
}


