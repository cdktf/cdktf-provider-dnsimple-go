// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type DnsimpleProviderConfig struct {
	// The account for API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.5.0/docs#account DnsimpleProvider#account}
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.5.0/docs#alias DnsimpleProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Flag to enable the prefetch of zone records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.5.0/docs#prefetch DnsimpleProvider#prefetch}
	Prefetch interface{} `field:"optional" json:"prefetch" yaml:"prefetch"`
	// Flag to enable the sandbox API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.5.0/docs#sandbox DnsimpleProvider#sandbox}
	Sandbox interface{} `field:"optional" json:"sandbox" yaml:"sandbox"`
	// The API v2 token for API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.5.0/docs#token DnsimpleProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Custom string to append to the user agent used for sending HTTP requests to the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.5.0/docs#user_agent DnsimpleProvider#user_agent}
	UserAgent *string `field:"optional" json:"userAgent" yaml:"userAgent"`
}

