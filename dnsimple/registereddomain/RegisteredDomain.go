// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package registereddomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-dnsimple-go/dnsimple/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-dnsimple-go/dnsimple/v9/registereddomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.4.0/docs/resources/registered_domain dnsimple_registered_domain}.
type RegisteredDomain interface {
	cdktf.TerraformResource
	AccountId() *float64
	AutoRenewEnabled() interface{}
	SetAutoRenewEnabled(val interface{})
	AutoRenewEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContactId() *float64
	SetContactId(val *float64)
	ContactIdInput() *float64
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DnssecEnabled() interface{}
	SetDnssecEnabled(val interface{})
	DnssecEnabledInput() interface{}
	DomainRegistration() RegisteredDomainDomainRegistrationOutputReference
	ExpiresAt() *string
	ExtendedAttributes() *map[string]*string
	SetExtendedAttributes(val *map[string]*string)
	ExtendedAttributesInput() *map[string]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PremiumPrice() *string
	SetPremiumPrice(val *string)
	PremiumPriceInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RegistrantChange() RegisteredDomainRegistrantChangeOutputReference
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() RegisteredDomainTimeoutsOutputReference
	TimeoutsInput() interface{}
	TransferLockEnabled() interface{}
	SetTransferLockEnabled(val interface{})
	TransferLockEnabledInput() interface{}
	UnicodeName() *string
	WhoisPrivacyEnabled() interface{}
	SetWhoisPrivacyEnabled(val interface{})
	WhoisPrivacyEnabledInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *RegisteredDomainTimeouts)
	ResetAutoRenewEnabled()
	ResetDnssecEnabled()
	ResetExtendedAttributes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPremiumPrice()
	ResetTimeouts()
	ResetTransferLockEnabled()
	ResetWhoisPrivacyEnabled()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for RegisteredDomain
type jsiiProxy_RegisteredDomain struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RegisteredDomain) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) AutoRenewEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRenewEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) AutoRenewEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRenewEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) ContactId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"contactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) ContactIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"contactIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) DnssecEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnssecEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) DnssecEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnssecEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) DomainRegistration() RegisteredDomainDomainRegistrationOutputReference {
	var returns RegisteredDomainDomainRegistrationOutputReference
	_jsii_.Get(
		j,
		"domainRegistration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) ExpiresAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiresAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) ExtendedAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extendedAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) ExtendedAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extendedAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) PremiumPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"premiumPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) PremiumPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"premiumPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) RegistrantChange() RegisteredDomainRegistrantChangeOutputReference {
	var returns RegisteredDomainRegistrantChangeOutputReference
	_jsii_.Get(
		j,
		"registrantChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) Timeouts() RegisteredDomainTimeoutsOutputReference {
	var returns RegisteredDomainTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) TransferLockEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transferLockEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) TransferLockEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transferLockEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) UnicodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unicodeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) WhoisPrivacyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"whoisPrivacyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegisteredDomain) WhoisPrivacyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"whoisPrivacyEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.4.0/docs/resources/registered_domain dnsimple_registered_domain} Resource.
func NewRegisteredDomain(scope constructs.Construct, id *string, config *RegisteredDomainConfig) RegisteredDomain {
	_init_.Initialize()

	if err := validateNewRegisteredDomainParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RegisteredDomain{}

	_jsii_.Create(
		"@cdktf/provider-dnsimple.registeredDomain.RegisteredDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.4.0/docs/resources/registered_domain dnsimple_registered_domain} Resource.
func NewRegisteredDomain_Override(r RegisteredDomain, scope constructs.Construct, id *string, config *RegisteredDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-dnsimple.registeredDomain.RegisteredDomain",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RegisteredDomain)SetAutoRenewEnabled(val interface{}) {
	if err := j.validateSetAutoRenewEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRenewEnabled",
		val,
	)
}

func (j *jsiiProxy_RegisteredDomain)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RegisteredDomain)SetContactId(val *float64) {
	if err := j.validateSetContactIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactId",
		val,
	)
}

func (j *jsiiProxy_RegisteredDomain)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RegisteredDomain)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RegisteredDomain)SetDnssecEnabled(val interface{}) {
	if err := j.validateSetDnssecEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnssecEnabled",
		val,
	)
}

func (j *jsiiProxy_RegisteredDomain)SetExtendedAttributes(val *map[string]*string) {
	if err := j.validateSetExtendedAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendedAttributes",
		val,
	)
}

func (j *jsiiProxy_RegisteredDomain)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RegisteredDomain)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RegisteredDomain)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_RegisteredDomain)SetPremiumPrice(val *string) {
	if err := j.validateSetPremiumPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"premiumPrice",
		val,
	)
}

func (j *jsiiProxy_RegisteredDomain)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RegisteredDomain)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RegisteredDomain)SetTransferLockEnabled(val interface{}) {
	if err := j.validateSetTransferLockEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transferLockEnabled",
		val,
	)
}

func (j *jsiiProxy_RegisteredDomain)SetWhoisPrivacyEnabled(val interface{}) {
	if err := j.validateSetWhoisPrivacyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whoisPrivacyEnabled",
		val,
	)
}

// Generates CDKTF code for importing a RegisteredDomain resource upon running "cdktf plan <stack-name>".
func RegisteredDomain_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRegisteredDomain_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-dnsimple.registeredDomain.RegisteredDomain",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func RegisteredDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRegisteredDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-dnsimple.registeredDomain.RegisteredDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RegisteredDomain_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRegisteredDomain_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-dnsimple.registeredDomain.RegisteredDomain",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RegisteredDomain_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRegisteredDomain_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-dnsimple.registeredDomain.RegisteredDomain",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RegisteredDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-dnsimple.registeredDomain.RegisteredDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RegisteredDomain) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RegisteredDomain) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RegisteredDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegisteredDomain) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegisteredDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegisteredDomain) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegisteredDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegisteredDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegisteredDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegisteredDomain) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegisteredDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegisteredDomain) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegisteredDomain) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RegisteredDomain) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegisteredDomain) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RegisteredDomain) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RegisteredDomain) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RegisteredDomain) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RegisteredDomain) PutTimeouts(value *RegisteredDomainTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RegisteredDomain) ResetAutoRenewEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetAutoRenewEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegisteredDomain) ResetDnssecEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetDnssecEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegisteredDomain) ResetExtendedAttributes() {
	_jsii_.InvokeVoid(
		r,
		"resetExtendedAttributes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegisteredDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegisteredDomain) ResetPremiumPrice() {
	_jsii_.InvokeVoid(
		r,
		"resetPremiumPrice",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegisteredDomain) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegisteredDomain) ResetTransferLockEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetTransferLockEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegisteredDomain) ResetWhoisPrivacyEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetWhoisPrivacyEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegisteredDomain) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegisteredDomain) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegisteredDomain) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegisteredDomain) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegisteredDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegisteredDomain) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

