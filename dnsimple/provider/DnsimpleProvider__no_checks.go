// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DnsimpleProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_DnsimpleProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateDnsimpleProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDnsimpleProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDnsimpleProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsimpleProvider) validateSetPrefetchParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsimpleProvider) validateSetSandboxParameters(val interface{}) error {
	return nil
}

func validateNewDnsimpleProviderParameters(scope constructs.Construct, id *string, config *DnsimpleProviderConfig) error {
	return nil
}

