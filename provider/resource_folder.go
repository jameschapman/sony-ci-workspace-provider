package provider

import (
	"log"
	"sonyciprovider/client"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceItem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"parent_id": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The id of the parent folder",
			},
			"workspace_id": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The id of the workspace",
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the folder",
				ForceNew:     true,
			},
		},
		Create: resourceCreateItem,
		Read:   resourceReadItem,
		Update: resourceUpdateItem,
		Delete: resourceDeleteItem,
		Exists: resourceExistsItem,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateItem(d *schema.ResourceData, m interface{}) error {
	
	log.Printf("[INFO] Creating resource")

apiClient := m.(*client.Client)

 name :=d.Get("name").(string)
 workspaceId :=d.Get("workspace_id").(string)
 parentFolderId :=d.Get("parent_id").(string)

id, err:= apiClient.Create(workspaceId, parentFolderId, name)

 
    if err != nil {
        return err
	}

d.SetId(id)
return nil;
}

func resourceReadItem(d *schema.ResourceData, m interface{}) error {

	log.Printf("[INFO] Reading resource with id " + d.Id())
apiClient := m.(*client.Client)
item, err:= apiClient.Get(d.Id())

if err != nil {
	d.SetId("")
	return nil
} else {
		log.Printf("[INFO] Found resource")
		log.Printf("[INFO] Resource name: " + item.Name)
	d.SetId(d.Id())
	d.Set("name", item.Name)
	d.Set("parent_id", item.ParentId)
	d.Set("workspace_id", d.Get("workspace_id").(string))
	return nil
}
}

func resourceUpdateItem(d *schema.ResourceData, m interface{}) error {

	log.Printf("[INFO] Updating resource")

apiClient := m.(*client.Client)
 name :=d.Get("name").(string)
err:= apiClient.Update(d.Id(), name)

if err != nil {
		return err
	}
	return nil

}

func resourceDeleteItem(d *schema.ResourceData, m interface{}) error {
apiClient := m.(*client.Client)

log.Printf("[INFO] Deleting resource")


 id := d.Id()
_, err:= apiClient.Delete(id)

if err != nil {
		return err
	}

d.SetId("")
	return nil
}

func resourceExistsItem(d *schema.ResourceData, m interface{}) (bool, error) {
apiClient := m.(*client.Client)

log.Printf("[INFO] Checking resource exists")

	name :=d.Get("name").(string)
 parentFolderId :=d.Get("parent_id").(string)

	res, err := apiClient.Exists(parentFolderId, name)

	if err != nil {
			return false, err
	}
	return res, nil

}