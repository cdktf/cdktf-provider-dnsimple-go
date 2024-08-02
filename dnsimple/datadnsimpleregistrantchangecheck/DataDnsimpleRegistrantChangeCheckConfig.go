// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadnsimpleregistrantchangecheck

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDnsimpleRegistrantChangeCheckConfig struct {
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
	// DNSimple contact ID for which the registrant change check is being performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.7.0/docs/data-sources/registrant_change_check#contact_id DataDnsimpleRegistrantChangeCheck#contact_id}
	ContactId *string `field:"required" json:"contactId" yaml:"contactId"`
	// DNSimple domain ID for which the registrant change check is being performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.7.0/docs/data-sources/registrant_change_check#domain_id DataDnsimpleRegistrantChangeCheck#domain_id}
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
}

