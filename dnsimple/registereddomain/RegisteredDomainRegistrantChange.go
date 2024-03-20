// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package registereddomain


type RegisteredDomainRegistrantChange struct {
	// State of the registrant change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.5.0/docs/resources/registered_domain#state RegisteredDomain#state}
	State *string `field:"optional" json:"state" yaml:"state"`
}

