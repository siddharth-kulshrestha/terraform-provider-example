package provider

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spaceapegames/terraform-provider-example/api/client"
)

func datasourceItems() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"item_names": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Item names of all items present on server",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"offset": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Offset for pagination",
			},
			"limit": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Limit for pagination",
			},
		},
		Read: resourceReadItemsWithPagination,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceReadItemsWithPagination(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	offset := d.Get("offset").(int)
	limit := d.Get("limit").(int)

	items, err := apiClient.GetFilteredItems(offset, limit)
	if err != nil {

		return fmt.Errorf("error in getting items list %s", err.Error())
	}
	log.Println("Returned list of all items in plugin: ")
	log.Println(items)
	names := []string{}
	if items != nil {
		for name, _ := range *items {
			names = append(names, name)
		}
	}
	log.Println("Returned list of all item names: ")
	log.Println(names)
	d.SetId(strings.Join(names, "-"))
	err = d.Set("item_names", names)
	return err
}
