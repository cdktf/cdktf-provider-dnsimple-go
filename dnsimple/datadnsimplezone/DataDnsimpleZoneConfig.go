package datadnsimplezone

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDnsimpleZoneConfig struct {
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
	// Zone Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/dnsimple/dnsimple/1.1.2/docs/data-sources/zone#name DataDnsimpleZone#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

