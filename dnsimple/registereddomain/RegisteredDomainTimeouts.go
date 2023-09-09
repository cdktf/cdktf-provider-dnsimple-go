// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package registereddomain


type RegisteredDomainTimeouts struct {
	// Create timeout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.0/docs/resources/registered_domain#create RegisteredDomain#create}
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Delete timeout (currently unused).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.0/docs/resources/registered_domain#delete RegisteredDomain#delete}
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Update timeout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.0/docs/resources/registered_domain#update RegisteredDomain#update}
	Update *string `field:"optional" json:"update" yaml:"update"`
}

