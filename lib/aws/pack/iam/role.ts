// *** WARNING: this file was generated by the Coconut IDL Compiler (CIDLC).  ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as coconut from "@coconut/coconut";

import {ARN} from "../types";
import {InlinePolicy} from "./policy";

export class Role extends coconut.Resource implements RoleArgs {
    public readonly name: string;
    public assumeRolePolicyDocument: any;
    public readonly path?: string;
    public readonly roleName?: string;
    public managedPolicyArns?: ARN[];
    public policies?: InlinePolicy[];
    public arn: ARN;

    constructor(args: RoleArgs) {
        super();
        if (args.name === undefined) {
            throw new Error("Missing required argument 'name'");
        }
        this.name = args.name;
        if (args.assumeRolePolicyDocument === undefined) {
            throw new Error("Missing required argument 'assumeRolePolicyDocument'");
        }
        this.assumeRolePolicyDocument = args.assumeRolePolicyDocument;
        this.path = args.path;
        this.roleName = args.roleName;
        this.managedPolicyArns = args.managedPolicyArns;
        this.policies = args.policies;
    }
}

export interface RoleArgs {
    readonly name: string;
    assumeRolePolicyDocument: any;
    readonly path?: string;
    readonly roleName?: string;
    managedPolicyArns?: ARN[];
    policies?: InlinePolicy[];
}


