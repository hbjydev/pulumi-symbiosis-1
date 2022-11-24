# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ClusterServiceAccountArgs', 'ClusterServiceAccount']

@pulumi.input_type
class ClusterServiceAccountArgs:
    def __init__(__self__, *,
                 cluster_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a ClusterServiceAccount resource.
        :param pulumi.Input[str] cluster_name: Cluster name. Changing the name forces re-creation.
        """
        pulumi.set(__self__, "cluster_name", cluster_name)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[str]:
        """
        Cluster name. Changing the name forces re-creation.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_name", value)


@pulumi.input_type
class _ClusterServiceAccountState:
    def __init__(__self__, *,
                 cluster_ca_certificate: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ClusterServiceAccount resources.
        :param pulumi.Input[str] cluster_ca_certificate: Cluster CA certificate
        :param pulumi.Input[str] cluster_name: Cluster name. Changing the name forces re-creation.
        :param pulumi.Input[str] token: Service account token
        """
        if cluster_ca_certificate is not None:
            pulumi.set(__self__, "cluster_ca_certificate", cluster_ca_certificate)
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="clusterCaCertificate")
    def cluster_ca_certificate(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster CA certificate
        """
        return pulumi.get(self, "cluster_ca_certificate")

    @cluster_ca_certificate.setter
    def cluster_ca_certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_ca_certificate", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster name. Changing the name forces re-creation.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        Service account token
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)


class ClusterServiceAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages Kubernetes clusters service accounts.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_symbiosis as symbiosis

        example = symbiosis.ClusterServiceAccount("example", cluster_name=symbiosis_cluster["example"]["name"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: Cluster name. Changing the name forces re-creation.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterServiceAccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Kubernetes clusters service accounts.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_symbiosis as symbiosis

        example = symbiosis.ClusterServiceAccount("example", cluster_name=symbiosis_cluster["example"]["name"])
        ```

        :param str resource_name: The name of the resource.
        :param ClusterServiceAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterServiceAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClusterServiceAccountArgs.__new__(ClusterServiceAccountArgs)

            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            __props__.__dict__["cluster_ca_certificate"] = None
            __props__.__dict__["token"] = None
        super(ClusterServiceAccount, __self__).__init__(
            'symbiosis:index/clusterServiceAccount:ClusterServiceAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_ca_certificate: Optional[pulumi.Input[str]] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None) -> 'ClusterServiceAccount':
        """
        Get an existing ClusterServiceAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_ca_certificate: Cluster CA certificate
        :param pulumi.Input[str] cluster_name: Cluster name. Changing the name forces re-creation.
        :param pulumi.Input[str] token: Service account token
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClusterServiceAccountState.__new__(_ClusterServiceAccountState)

        __props__.__dict__["cluster_ca_certificate"] = cluster_ca_certificate
        __props__.__dict__["cluster_name"] = cluster_name
        __props__.__dict__["token"] = token
        return ClusterServiceAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterCaCertificate")
    def cluster_ca_certificate(self) -> pulumi.Output[str]:
        """
        Cluster CA certificate
        """
        return pulumi.get(self, "cluster_ca_certificate")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        Cluster name. Changing the name forces re-creation.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        Service account token
        """
        return pulumi.get(self, "token")

