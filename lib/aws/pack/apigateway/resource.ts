// *** WARNING: this file was generated by the Coconut IDL Compiler (CIDLC).  ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as coconut from "@coconut/coconut";

import {RestAPI} from "./restAPI";

export class Resource extends coconut.Resource implements ResourceArgs {
    public readonly name: string;
    public readonly parent: Resource;
    public readonly pathPart: string;
    public readonly restAPI: RestAPI;

    constructor(args: ResourceArgs) {
        super();
        if (args.name === undefined) {
            throw new Error("Missing required argument 'name'");
        }
        this.name = args.name;
        if (args.parent === undefined) {
            throw new Error("Missing required argument 'parent'");
        }
        this.parent = args.parent;
        if (args.pathPart === undefined) {
            throw new Error("Missing required argument 'pathPart'");
        }
        this.pathPart = args.pathPart;
        if (args.restAPI === undefined) {
            throw new Error("Missing required argument 'restAPI'");
        }
        this.restAPI = args.restAPI;
    }
}

export interface ResourceArgs {
    readonly name: string;
    readonly parent: Resource;
    readonly pathPart: string;
    readonly restAPI: RestAPI;
}


