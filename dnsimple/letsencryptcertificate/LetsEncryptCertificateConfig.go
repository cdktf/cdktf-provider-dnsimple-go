// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package letsencryptcertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LetsEncryptCertificateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.6.0/docs/resources/lets_encrypt_certificate#auto_renew LetsEncryptCertificate#auto_renew}.
	AutoRenew interface{} `field:"required" json:"autoRenew" yaml:"autoRenew"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.6.0/docs/resources/lets_encrypt_certificate#domain_id LetsEncryptCertificate#domain_id}.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.6.0/docs/resources/lets_encrypt_certificate#name LetsEncryptCertificate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.6.0/docs/resources/lets_encrypt_certificate#alternate_names LetsEncryptCertificate#alternate_names}.
	AlternateNames *[]*string `field:"optional" json:"alternateNames" yaml:"alternateNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.6.0/docs/resources/lets_encrypt_certificate#signature_algorithm LetsEncryptCertificate#signature_algorithm}.
	SignatureAlgorithm *string `field:"optional" json:"signatureAlgorithm" yaml:"signatureAlgorithm"`
}

