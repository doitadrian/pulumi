// *** WARNING: this file was generated by the Coconut IDL Compiler (CIDLC).  ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/coconut/pkg/resource"
    "github.com/pulumi/coconut/pkg/tokens"
    "github.com/pulumi/coconut/pkg/util/contract"
    "github.com/pulumi/coconut/pkg/util/mapper"
    "github.com/pulumi/coconut/sdk/go/pkg/cocorpc"
)

/* RPC stubs for Object resource provider */

// ObjectToken is the type token corresponding to the Object package type.
const ObjectToken = tokens.Type("aws:s3/object:Object")

// ObjectProviderOps is a pluggable interface for Object-related management functionality.
type ObjectProviderOps interface {
    Check(ctx context.Context, obj *Object) ([]mapper.FieldError, error)
    Name(ctx context.Context, obj *Object) (string, error)
    Create(ctx context.Context, obj *Object) (string, error)
    Get(ctx context.Context, id string) (*Object, error)
    InspectChange(ctx context.Context,
        id string, old *Object, new *Object, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id string, old *Object, new *Object, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id string) error
}

// ObjectProvider is a dynamic gRPC-based plugin for managing Object resources.
type ObjectProvider struct {
    ops ObjectProviderOps
}

// NewObjectProvider allocates a resource provider that delegates to a ops instance.
func NewObjectProvider(ops ObjectProviderOps) cocorpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &ObjectProvider{ops: ops}
}

func (p *ObjectProvider) Check(
    ctx context.Context, req *cocorpc.CheckRequest) (*cocorpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(ObjectToken))
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

func (p *ObjectProvider) Name(
    ctx context.Context, req *cocorpc.NameRequest) (*cocorpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(ObjectToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    name, err := p.ops.Name(ctx, obj)
    return &cocorpc.NameResponse{Name: name}, err
}

func (p *ObjectProvider) Create(
    ctx context.Context, req *cocorpc.CreateRequest) (*cocorpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(ObjectToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &cocorpc.CreateResponse{
        Id:   id,
    }, nil
}

func (p *ObjectProvider) Get(
    ctx context.Context, req *cocorpc.GetRequest) (*cocorpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(ObjectToken))
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

func (p *ObjectProvider) InspectChange(
    ctx context.Context, req *cocorpc.ChangeRequest) (*cocorpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(ObjectToken))
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
    if diff.Changed("key") {
        replaces = append(replaces, "key")
    }
    if diff.Changed("bucket") {
        replaces = append(replaces, "bucket")
    }
    if diff.Changed("source") {
        replaces = append(replaces, "source")
    }
    more, err := p.ops.InspectChange(ctx, req.GetId(), old, new, diff)
    if err != nil {
        return nil, err
    }
    return &cocorpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *ObjectProvider) Update(
    ctx context.Context, req *cocorpc.ChangeRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(ObjectToken))
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

func (p *ObjectProvider) Delete(
    ctx context.Context, req *cocorpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(ObjectToken))
    if err := p.ops.Delete(ctx, req.GetId()); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *ObjectProvider) Unmarshal(
    v *pbstruct.Struct) (*Object, resource.PropertyMap, mapper.DecodeError) {
    var obj Object
    props := resource.UnmarshalProperties(v)
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable Object structure(s) */

// Object is a marshalable representation of its corresponding IDL type.
type Object struct {
    Key string `json:"key"`
    Bucket *resource.ID `json:"bucket"`
    Source *resource.Asset `json:"source"`
}

// Object's properties have constants to make dealing with diffs and property bags easier.
const (
    Object_Key = "key"
    Object_Bucket = "bucket"
    Object_Source = "source"
)


