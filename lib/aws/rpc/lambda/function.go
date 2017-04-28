// *** WARNING: this file was generated by the Coconut IDL Compiler (CIDLC).  ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/coconut/pkg/resource"
    "github.com/pulumi/coconut/pkg/tokens"
    "github.com/pulumi/coconut/pkg/util/contract"
    "github.com/pulumi/coconut/pkg/util/mapper"
    "github.com/pulumi/coconut/sdk/go/pkg/cocorpc"

    __aws "github.com/pulumi/coconut/lib/aws/rpc"
)

/* Marshalable DeadLetterConfig structure(s) */

// DeadLetterConfig is a marshalable representation of its corresponding IDL type.
type DeadLetterConfig struct {
    Target *resource.ID `json:"target"`
}

// DeadLetterConfig's properties have constants to make dealing with diffs and property bags easier.
const (
    DeadLetterConfig_Target = "target"
)

/* RPC stubs for Function resource provider */

// FunctionToken is the type token corresponding to the Function package type.
const FunctionToken = tokens.Type("aws:lambda/function:Function")

// FunctionProviderOps is a pluggable interface for Function-related management functionality.
type FunctionProviderOps interface {
    Check(ctx context.Context, obj *Function) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *Function) (string, *FunctionOuts, error)
    Get(ctx context.Context, id string) (*Function, error)
    InspectChange(ctx context.Context,
        id string, old *Function, new *Function, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id string, old *Function, new *Function, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id string) error
}

// FunctionProvider is a dynamic gRPC-based plugin for managing Function resources.
type FunctionProvider struct {
    ops FunctionProviderOps
}

// NewFunctionProvider allocates a resource provider that delegates to a ops instance.
func NewFunctionProvider(ops FunctionProviderOps) cocorpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &FunctionProvider{ops: ops}
}

func (p *FunctionProvider) Check(
    ctx context.Context, req *cocorpc.CheckRequest) (*cocorpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(FunctionToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr == nil || len(decerr.Failures()) == 0 {
        failures, err := p.ops.Check(ctx, obj)
        if err != nil {
            return nil, err
        }
        if len(failures) > 0 {
            decerr = mapper.NewDecodeErr(failures)
        }
    }
    return resource.NewCheckResponse(decerr), nil
}

func (p *FunctionProvider) Name(
    ctx context.Context, req *cocorpc.NameRequest) (*cocorpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(FunctionToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        return nil, errors.New("Name property cannot be empty")
    }
    return &cocorpc.NameResponse{Name: obj.Name}, nil
}

func (p *FunctionProvider) Create(
    ctx context.Context, req *cocorpc.CreateRequest) (*cocorpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(FunctionToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    id, outs, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &cocorpc.CreateResponse{
        Id:   id,
        Outputs: resource.MarshalProperties(
            nil, resource.NewPropertyMap(outs), resource.MarshalOptions{},
        ),
    }, nil
}

func (p *FunctionProvider) Get(
    ctx context.Context, req *cocorpc.GetRequest) (*cocorpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(FunctionToken))
    id := req.GetId()
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &cocorpc.GetResponse{
        Properties: resource.MarshalProperties(
            nil, resource.NewPropertyMap(obj), resource.MarshalOptions{}),
    }, nil
}

func (p *FunctionProvider) InspectChange(
    ctx context.Context, req *cocorpc.ChangeRequest) (*cocorpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(FunctionToken))
    old, oldprops, decerr := p.Unmarshal(req.GetOlds())
    if decerr != nil {
        return nil, decerr
    }
    new, newprops, decerr := p.Unmarshal(req.GetNews())
    if decerr != nil {
        return nil, decerr
    }
    diff := oldprops.Diff(newprops)
    var replaces []string
    if diff.Changed("name") {
        replaces = append(replaces, "name")
    }
    more, err := p.ops.InspectChange(ctx, req.GetId(), old, new, diff)
    if err != nil {
        return nil, err
    }
    return &cocorpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *FunctionProvider) Update(
    ctx context.Context, req *cocorpc.ChangeRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(FunctionToken))
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, req.GetId(), old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *FunctionProvider) Delete(
    ctx context.Context, req *cocorpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(FunctionToken))
    if err := p.ops.Delete(ctx, req.GetId()); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *FunctionProvider) Unmarshal(
    v *pbstruct.Struct) (*Function, resource.PropertyMap, mapper.DecodeError) {
    var obj Function
    props := resource.UnmarshalProperties(v)
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable Function structure(s) */

// Function is a marshalable representation of its corresponding IDL type.
type Function struct {
    Name string `json:"name"`
    Code *resource.Asset `json:"code"`
    Handler string `json:"handler"`
    Role *resource.ID `json:"role"`
    Runtime Runtime `json:"runtime"`
    FunctionName *string `json:"functionName,omitempty"`
    DeadLetterConfig *DeadLetterConfig `json:"deadLetterConfig,omitempty"`
    Description *string `json:"description,omitempty"`
    Environment *Environment `json:"environment,omitempty"`
    KMSKey *resource.ID `json:"kmsKey,omitempty"`
    MemorySize *float64 `json:"memorySize,omitempty"`
    Timeout *float64 `json:"timeout,omitempty"`
    VPCConfig *VPCConfig `json:"vpcConfig,omitempty"`
    ARN __aws.ARN `json:"arn,omitempty"`
}

// FunctionOuts is a marshalable representation of its IDL type's output properties.
type FunctionOuts struct {
    ARN __aws.ARN `json:"arn"`
}

// Function's properties have constants to make dealing with diffs and property bags easier.
const (
    Function_Name = "name"
    Function_Code = "code"
    Function_Handler = "handler"
    Function_Role = "role"
    Function_Runtime = "runtime"
    Function_FunctionName = "functionName"
    Function_DeadLetterConfig = "deadLetterConfig"
    Function_Description = "description"
    Function_Environment = "environment"
    Function_KMSKey = "kmsKey"
    Function_MemorySize = "memorySize"
    Function_Timeout = "timeout"
    Function_VPCConfig = "vpcConfig"
    Function_ARN = "arn"
)

/* Marshalable VPCConfig structure(s) */

// VPCConfig is a marshalable representation of its corresponding IDL type.
type VPCConfig struct {
    SecurityGroups []*resource.ID `json:"securityGroups"`
    Subnets []*resource.ID `json:"subnets"`
}

// VPCConfig's properties have constants to make dealing with diffs and property bags easier.
const (
    VPCConfig_SecurityGroups = "securityGroups"
    VPCConfig_Subnets = "subnets"
)

/* Typedefs */

type (
    Environment map[string]string
    Runtime string
)

/* Constants */

const (
    DotnetCore1d0Runtime Runtime = "dotnetcore1.0"
    Java8Runtime Runtime = "java8"
    NodeJS4d3EdgeRuntime Runtime = "nodejs4.3-edge"
    NodeJS4d3Runtime Runtime = "nodejs4.3"
    NodeJS6d10Runtime Runtime = "nodejs6.10"
    NodeJSRuntime Runtime = "nodejs"
    Python2d7Runtime Runtime = "python2.7"
)


