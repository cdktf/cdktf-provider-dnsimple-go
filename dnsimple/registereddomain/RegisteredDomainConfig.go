// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package registereddomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RegisteredDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/registered_domain#contact_id RegisteredDomain#contact_id}.
	ContactId *float64 `field:"required" json:"contactId" yaml:"contactId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/registered_domain#name RegisteredDomain#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/registered_domain#auto_renew_enabled RegisteredDomain#auto_renew_enabled}.
	AutoRenewEnabled interface{} `field:"optional" json:"autoRenewEnabled" yaml:"autoRenewEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/registered_domain#dnssec_enabled RegisteredDomain#dnssec_enabled}.
	DnssecEnabled interface{} `field:"optional" json:"dnssecEnabled" yaml:"dnssecEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/registered_domain#extended_attributes RegisteredDomain#extended_attributes}.
	ExtendedAttributes *map[string]*string `field:"optional" json:"extendedAttributes" yaml:"extendedAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/registered_domain#premium_price RegisteredDomain#premium_price}.
	PremiumPrice *string `field:"optional" json:"premiumPrice" yaml:"premiumPrice"`
	// Timeouts for operations, given as a parsable string as in `10m` or `30s`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/registered_domain#timeouts RegisteredDomain#timeouts}
	Timeouts *RegisteredDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/registered_domain#transfer_lock_enabled RegisteredDomain#transfer_lock_enabled}.
	TransferLockEnabled interface{} `field:"optional" json:"transferLockEnabled" yaml:"transferLockEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.2.1/docs/resources/registered_domain#whois_privacy_enabled RegisteredDomain#whois_privacy_enabled}.
	WhoisPrivacyEnabled interface{} `field:"optional" json:"whoisPrivacyEnabled" yaml:"whoisPrivacyEnabled"`
}

