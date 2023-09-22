// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadnsimpleregistrantchangecheck

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-dnsimple-go/dnsimple/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-dnsimple-go/dnsimple/v7/datadnsimpleregistrantchangecheck/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.3.0/docs/data-sources/registrant_change_check dnsimple_registrant_change_check}.
type DataDnsimpleRegistrantChangeCheck interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContactId() *string
	SetContactId(val *string)
	ContactIdInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DomainId() *string
	SetDomainId(val *string)
	DomainIdInput() *string
	ExtendedAttributes() DataDnsimpleRegistrantChangeCheckExtendedAttributesList
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RegistryOwnerChange() cdktf.IResolvable
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataDnsimpleRegistrantChangeCheck
type jsiiProxy_DataDnsimpleRegistrantChangeCheck struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) ContactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) ContactIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) DomainIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) ExtendedAttributes() DataDnsimpleRegistrantChangeCheckExtendedAttributesList {
	var returns DataDnsimpleRegistrantChangeCheckExtendedAttributesList
	_jsii_.Get(
		j,
		"extendedAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) RegistryOwnerChange() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"registryOwnerChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.3.0/docs/data-sources/registrant_change_check dnsimple_registrant_change_check} Data Source.
func NewDataDnsimpleRegistrantChangeCheck(scope constructs.Construct, id *string, config *DataDnsimpleRegistrantChangeCheckConfig) DataDnsimpleRegistrantChangeCheck {
	_init_.Initialize()

	if err := validateNewDataDnsimpleRegistrantChangeCheckParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDnsimpleRegistrantChangeCheck{}

	_jsii_.Create(
		"@cdktf/provider-dnsimple.dataDnsimpleRegistrantChangeCheck.DataDnsimpleRegistrantChangeCheck",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.3.0/docs/data-sources/registrant_change_check dnsimple_registrant_change_check} Data Source.
func NewDataDnsimpleRegistrantChangeCheck_Override(d DataDnsimpleRegistrantChangeCheck, scope constructs.Construct, id *string, config *DataDnsimpleRegistrantChangeCheckConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-dnsimple.dataDnsimpleRegistrantChangeCheck.DataDnsimpleRegistrantChangeCheck",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck)SetContactId(val *string) {
	if err := j.validateSetContactIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactId",
		val,
	)
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck)SetDomainId(val *string) {
	if err := j.validateSetDomainIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainId",
		val,
	)
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDnsimpleRegistrantChangeCheck)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
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
func DataDnsimpleRegistrantChangeCheck_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDnsimpleRegistrantChangeCheck_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-dnsimple.dataDnsimpleRegistrantChangeCheck.DataDnsimpleRegistrantChangeCheck",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDnsimpleRegistrantChangeCheck_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDnsimpleRegistrantChangeCheck_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-dnsimple.dataDnsimpleRegistrantChangeCheck.DataDnsimpleRegistrantChangeCheck",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDnsimpleRegistrantChangeCheck_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDnsimpleRegistrantChangeCheck_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-dnsimple.dataDnsimpleRegistrantChangeCheck.DataDnsimpleRegistrantChangeCheck",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDnsimpleRegistrantChangeCheck_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-dnsimple.dataDnsimpleRegistrantChangeCheck.DataDnsimpleRegistrantChangeCheck",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDnsimpleRegistrantChangeCheck) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

