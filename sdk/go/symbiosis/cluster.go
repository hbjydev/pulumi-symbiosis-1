// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package symbiosis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages Kubernetes clusters.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/symbiosis-cloud/pulumi-symbiosis/sdk/go/symbiosis"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := symbiosis.NewCluster(ctx, "example", &symbiosis.ClusterArgs{
//				Region: pulumi.String("germany-1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Cluster struct {
	pulumi.CustomResourceState

	CaCertificate pulumi.StringOutput `pulumi:"caCertificate"`
	Certificate   pulumi.StringOutput `pulumi:"certificate"`
	// Cluster API server endpoint
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// When set to true it will deploy a highly available control plane with multiple replicas for redundancy.
	IsHighlyAvailable pulumi.BoolPtrOutput `pulumi:"isHighlyAvailable"`
	// Kubernetes version, see symbiosis.host for valid values or "latest" for the most recent supported version.
	KubeVersion pulumi.StringPtrOutput `pulumi:"kubeVersion"`
	// Cluster name. Changing the name forces re-creation.
	Name       pulumi.StringOutput `pulumi:"name"`
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	Region     pulumi.StringOutput `pulumi:"region"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"caCertificate",
		"certificate",
		"privateKey",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("symbiosis:index/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("symbiosis:index/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	CaCertificate *string `pulumi:"caCertificate"`
	Certificate   *string `pulumi:"certificate"`
	// Cluster API server endpoint
	Endpoint *string `pulumi:"endpoint"`
	// When set to true it will deploy a highly available control plane with multiple replicas for redundancy.
	IsHighlyAvailable *bool `pulumi:"isHighlyAvailable"`
	// Kubernetes version, see symbiosis.host for valid values or "latest" for the most recent supported version.
	KubeVersion *string `pulumi:"kubeVersion"`
	// Cluster name. Changing the name forces re-creation.
	Name       *string `pulumi:"name"`
	PrivateKey *string `pulumi:"privateKey"`
	Region     *string `pulumi:"region"`
}

type ClusterState struct {
	CaCertificate pulumi.StringPtrInput
	Certificate   pulumi.StringPtrInput
	// Cluster API server endpoint
	Endpoint pulumi.StringPtrInput
	// When set to true it will deploy a highly available control plane with multiple replicas for redundancy.
	IsHighlyAvailable pulumi.BoolPtrInput
	// Kubernetes version, see symbiosis.host for valid values or "latest" for the most recent supported version.
	KubeVersion pulumi.StringPtrInput
	// Cluster name. Changing the name forces re-creation.
	Name       pulumi.StringPtrInput
	PrivateKey pulumi.StringPtrInput
	Region     pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// When set to true it will deploy a highly available control plane with multiple replicas for redundancy.
	IsHighlyAvailable *bool `pulumi:"isHighlyAvailable"`
	// Kubernetes version, see symbiosis.host for valid values or "latest" for the most recent supported version.
	KubeVersion *string `pulumi:"kubeVersion"`
	// Cluster name. Changing the name forces re-creation.
	Name   *string `pulumi:"name"`
	Region string  `pulumi:"region"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// When set to true it will deploy a highly available control plane with multiple replicas for redundancy.
	IsHighlyAvailable pulumi.BoolPtrInput
	// Kubernetes version, see symbiosis.host for valid values or "latest" for the most recent supported version.
	KubeVersion pulumi.StringPtrInput
	// Cluster name. Changing the name forces re-creation.
	Name   pulumi.StringPtrInput
	Region pulumi.StringInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//	ClusterArray{ ClusterArgs{...} }
type ClusterArrayInput interface {
	pulumi.Input

	ToClusterArrayOutput() ClusterArrayOutput
	ToClusterArrayOutputWithContext(context.Context) ClusterArrayOutput
}

type ClusterArray []ClusterInput

func (ClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (i ClusterArray) ToClusterArrayOutput() ClusterArrayOutput {
	return i.ToClusterArrayOutputWithContext(context.Background())
}

func (i ClusterArray) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterArrayOutput)
}

// ClusterMapInput is an input type that accepts ClusterMap and ClusterMapOutput values.
// You can construct a concrete instance of `ClusterMapInput` via:
//
//	ClusterMap{ "key": ClusterArgs{...} }
type ClusterMapInput interface {
	pulumi.Input

	ToClusterMapOutput() ClusterMapOutput
	ToClusterMapOutputWithContext(context.Context) ClusterMapOutput
}

type ClusterMap map[string]ClusterInput

func (ClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (i ClusterMap) ToClusterMapOutput() ClusterMapOutput {
	return i.ToClusterMapOutputWithContext(context.Background())
}

func (i ClusterMap) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMapOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func (o ClusterOutput) CaCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.CaCertificate }).(pulumi.StringOutput)
}

func (o ClusterOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// Cluster API server endpoint
func (o ClusterOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// When set to true it will deploy a highly available control plane with multiple replicas for redundancy.
func (o ClusterOutput) IsHighlyAvailable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.IsHighlyAvailable }).(pulumi.BoolPtrOutput)
}

// Kubernetes version, see symbiosis.host for valid values or "latest" for the most recent supported version.
func (o ClusterOutput) KubeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.KubeVersion }).(pulumi.StringPtrOutput)
}

// Cluster name. Changing the name forces re-creation.
func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClusterOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

func (o ClusterOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].([]*Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].(map[string]*Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterArrayInput)(nil)).Elem(), ClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMapInput)(nil)).Elem(), ClusterMap{})
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}
