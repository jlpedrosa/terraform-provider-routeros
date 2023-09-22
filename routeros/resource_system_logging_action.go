package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
    ".id": "*4",
    "action": "echo",
    "default": "true",
    "disabled": "false",
    "invalid": "false",
    "prefix": "",
    "topics": "critical"
}
*/

// ResourceSystemLoggingEmailAction defines the resource for the logging destiation email
// https://wiki.mikrotik.com/wiki/Manual:System/Log
func ResourceSystemLoggingEmailAction() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/logging/action"),
		MetaId:           PropId(Id),
		// TODO: this should not be exposed... but computed don't accept a default.
		"target": {
			Type:        schema.TypeString,
			Default:     "email",
			Optional:    true,
			Description: "Type of action, always email",
		},
		"email_start_tls": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether to use tls when sending email",
		},
		"email_to": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "email address where logs are sent",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name for the logging action",
			// TODO: pending validation here.
		},
	}
	return DefaultLoggingAction(resSchema)
}

func DefaultLoggingAction(resSchema map[string]*schema.Schema) *schema.Resource {
	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
