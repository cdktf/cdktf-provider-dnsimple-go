// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package registereddomain


type RegisteredDomainDomainRegistration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.4.0/docs/resources/registered_domain#state RegisteredDomain#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
}

