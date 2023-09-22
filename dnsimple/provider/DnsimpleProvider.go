// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-dnsimple-go/dnsimple/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-dnsimple-go/dnsimple/v7/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.3.0/docs dnsimple}.
type DnsimpleProvider interface {
	cdktf.TerraformProvider
	Account() *string
	SetAccount(val *string)
	AccountInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	Prefetch() interface{}
	SetPrefetch(val interface{})
	PrefetchInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	Sandbox() interface{}
	SetSandbox(val interface{})
	SandboxInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	UserAgent() *string
	SetUserAgent(val *string)
	UserAgentInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccount()
	ResetAlias()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrefetch()
	ResetSandbox()
	ResetToken()
	ResetUserAgent()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DnsimpleProvider
type jsiiProxy_DnsimpleProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_DnsimpleProvider) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) AccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) Prefetch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prefetch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) PrefetchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prefetchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) Sandbox() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sandbox",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) SandboxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sandboxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) UserAgent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsimpleProvider) UserAgentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAgentInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.3.0/docs dnsimple} Resource.
func NewDnsimpleProvider(scope constructs.Construct, id *string, config *DnsimpleProviderConfig) DnsimpleProvider {
	_init_.Initialize()

	if err := validateNewDnsimpleProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DnsimpleProvider{}

	_jsii_.Create(
		"@cdktf/provider-dnsimple.provider.DnsimpleProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.3.0/docs dnsimple} Resource.
func NewDnsimpleProvider_Override(d DnsimpleProvider, scope constructs.Construct, id *string, config *DnsimpleProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-dnsimple.provider.DnsimpleProvider",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DnsimpleProvider)SetAccount(val *string) {
	_jsii_.Set(
		j,
		"account",
		val,
	)
}

func (j *jsiiProxy_DnsimpleProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_DnsimpleProvider)SetPrefetch(val interface{}) {
	if err := j.validateSetPrefetchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefetch",
		val,
	)
}

func (j *jsiiProxy_DnsimpleProvider)SetSandbox(val interface{}) {
	if err := j.validateSetSandboxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sandbox",
		val,
	)
}

func (j *jsiiProxy_DnsimpleProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_DnsimpleProvider)SetUserAgent(val *string) {
	_jsii_.Set(
		j,
		"userAgent",
		val,
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
func DnsimpleProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDnsimpleProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-dnsimple.provider.DnsimpleProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DnsimpleProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDnsimpleProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-dnsimple.provider.DnsimpleProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DnsimpleProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDnsimpleProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-dnsimple.provider.DnsimpleProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DnsimpleProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-dnsimple.provider.DnsimpleProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DnsimpleProvider) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DnsimpleProvider) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DnsimpleProvider) ResetAccount() {
	_jsii_.InvokeVoid(
		d,
		"resetAccount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsimpleProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		d,
		"resetAlias",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsimpleProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsimpleProvider) ResetPrefetch() {
	_jsii_.InvokeVoid(
		d,
		"resetPrefetch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsimpleProvider) ResetSandbox() {
	_jsii_.InvokeVoid(
		d,
		"resetSandbox",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsimpleProvider) ResetToken() {
	_jsii_.InvokeVoid(
		d,
		"resetToken",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsimpleProvider) ResetUserAgent() {
	_jsii_.InvokeVoid(
		d,
		"resetUserAgent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsimpleProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsimpleProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsimpleProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsimpleProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

