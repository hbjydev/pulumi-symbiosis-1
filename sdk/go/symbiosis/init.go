// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package symbiosis

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "symbiosis:index/cluster:Cluster":
		r = &Cluster{}
	case "symbiosis:index/clusterServiceAccount:ClusterServiceAccount":
		r = &ClusterServiceAccount{}
	case "symbiosis:index/nodePool:NodePool":
		r = &NodePool{}
	case "symbiosis:index/teamMember:TeamMember":
		r = &TeamMember{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:symbiosis" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"symbiosis",
		"index/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"symbiosis",
		"index/clusterServiceAccount",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"symbiosis",
		"index/nodePool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"symbiosis",
		"index/teamMember",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"symbiosis",
		&pkg{version},
	)
}
