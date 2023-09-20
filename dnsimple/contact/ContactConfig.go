// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package contact

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContactConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/contact#address1 Contact#address1}.
	Address1 *string `field:"required" json:"address1" yaml:"address1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/contact#city Contact#city}.
	City *string `field:"required" json:"city" yaml:"city"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/contact#country Contact#country}.
	Country *string `field:"required" json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/contact#email Contact#email}.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/contact#first_name Contact#first_name}.
	FirstName *string `field:"required" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/contact#last_name Contact#last_name}.
	LastName *string `field:"required" json:"lastName" yaml:"lastName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/contact#phone Contact#phone}.
	Phone *string `field:"required" json:"phone" yaml:"phone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/contact#postal_code Contact#postal_code}.
	PostalCode *string `field:"required" json:"postalCode" yaml:"postalCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/contact#state_province Contact#state_province}.
	StateProvince *string `field:"required" json:"stateProvince" yaml:"stateProvince"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/contact#address2 Contact#address2}.
	Address2 *string `field:"optional" json:"address2" yaml:"address2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/contact#fax Contact#fax}.
	Fax *string `field:"optional" json:"fax" yaml:"fax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/contact#job_title Contact#job_title}.
	JobTitle *string `field:"optional" json:"jobTitle" yaml:"jobTitle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/contact#label Contact#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/contact#organization_name Contact#organization_name}.
	OrganizationName *string `field:"optional" json:"organizationName" yaml:"organizationName"`
}

