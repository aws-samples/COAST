package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Base class for Vpn connections.
type VpnConnectionBase interface {
	awscdk.Resource
	IVpnConnection
	// The ASN of the customer gateway.
	CustomerGatewayAsn() *float64
	// The id of the customer gateway.
	CustomerGatewayId() *string
	// The ip address of the customer gateway.
	CustomerGatewayIp() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The id of the VPN connection.
	VpnId() *string
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Return the given named metric for this VPNConnection.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The bytes received through the VPN tunnel.
	//
	// Sum over 5 minutes.
	MetricTunnelDataIn(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The bytes sent through the VPN tunnel.
	//
	// Sum over 5 minutes.
	MetricTunnelDataOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The state of the tunnel. 0 indicates DOWN and 1 indicates UP.
	//
	// Average over 5 minutes.
	MetricTunnelState(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for VpnConnectionBase
type jsiiProxy_VpnConnectionBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IVpnConnection
}

func (j *jsiiProxy_VpnConnectionBase) CustomerGatewayAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customerGatewayAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionBase) CustomerGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionBase) CustomerGatewayIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnectionBase) VpnId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnId",
		&returns,
	)
	return returns
}


func NewVpnConnectionBase_Override(v VpnConnectionBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.VpnConnectionBase",
		[]interface{}{scope, id, props},
		v,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func VpnConnectionBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpnConnectionBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.VpnConnectionBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func VpnConnectionBase_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateVpnConnectionBase_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.VpnConnectionBase",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func VpnConnectionBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateVpnConnectionBase_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.VpnConnectionBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := v.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (v *jsiiProxy_VpnConnectionBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := v.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionBase) GetResourceNameAttribute(nameAttr *string) *string {
	if err := v.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionBase) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := v.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		v,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionBase) MetricTunnelDataIn(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := v.validateMetricTunnelDataInParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		v,
		"metricTunnelDataIn",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionBase) MetricTunnelDataOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := v.validateMetricTunnelDataOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		v,
		"metricTunnelDataOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionBase) MetricTunnelState(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := v.validateMetricTunnelStateParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		v,
		"metricTunnelState",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnectionBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

