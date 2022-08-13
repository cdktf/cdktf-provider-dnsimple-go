// Prebuilt dnsimple Provider for Terraform CDK (cdktf)
package dnsimple

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZoneRecordConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dnsimple/r/zone_record#name ZoneRecord#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dnsimple/r/zone_record#type ZoneRecord#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dnsimple/r/zone_record#value ZoneRecord#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dnsimple/r/zone_record#zone_name ZoneRecord#zone_name}.
	ZoneName *string `field:"required" json:"zoneName" yaml:"zoneName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dnsimple/r/zone_record#id ZoneRecord#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dnsimple/r/zone_record#priority ZoneRecord#priority}.
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dnsimple/r/zone_record#ttl ZoneRecord#ttl}.
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

