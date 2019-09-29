package xap

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/vinchauhan/terraform-provider-xap/xap/xapapi"
)

func resourceSpace() *schema.Resource {

	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"partition": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
			},
			"backups": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
			},
			"requiresIsolation": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
			},
		},
		SchemaVersion:      0,
		MigrateState:       nil,
		StateUpgraders:     nil,
		Create:             resourceSpaceCreate,
		Read:               nil,
		Update:             nil,
		Delete:             nil,
		Exists:             nil,
		CustomizeDiff:      nil,
		Importer:           nil,
		DeprecationMessage: "",
		Timeouts:           nil,
	}
}

func resourceSpaceCreate(d *schema.ResourceData, meta interface{}) error {

	////grab the session
	session := meta.(*xapapi.Session)

	if session == nil {
		return fmt.Errorf("Http Client is not instantiated")
	}

	//Collect resource parameters
	name := d.Get("name").(string)
	partitions := d.Get("partitions").(int)
	backups := d.Get("backups").(bool)
	requiresIsolation := d.Get("requiresIsolation").(bool)
	sm := xapapi.NewSpaceManager(name, partitions, backups, requiresIsolation)
	//var (name string
	//	partitions int
	//	backups bool
	//	requiresIsolation bool)
	//name = d.Get("name").(string)
	//xapClient := newXapClient()
	//id, err := sm.CreateSpace(name, partitions, backups, requiresIsolation);
	//if err != nil {
	//	return err
	//}
	//d.Set(id)
	//return resourceSpaceUpdate(d, NewResourceMeta{meta})

}
