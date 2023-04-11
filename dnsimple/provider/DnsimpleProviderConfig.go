package provider


type DnsimpleProviderConfig struct {
	// The account for API operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dnsimple#account DnsimpleProvider#account}
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dnsimple#alias DnsimpleProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Flag to enable the prefetch of zone records.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dnsimple#prefetch DnsimpleProvider#prefetch}
	Prefetch interface{} `field:"optional" json:"prefetch" yaml:"prefetch"`
	// Flag to enable the sandbox API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dnsimple#sandbox DnsimpleProvider#sandbox}
	Sandbox interface{} `field:"optional" json:"sandbox" yaml:"sandbox"`
	// The API v2 token for API operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dnsimple#token DnsimpleProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Custom string to append to the user agent used for sending HTTP requests to the API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dnsimple#user_agent DnsimpleProvider#user_agent}
	UserAgent *string `field:"optional" json:"userAgent" yaml:"userAgent"`
}

