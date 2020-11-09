package provider

import (
	"hello/client"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceItem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"parentId": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The id of the parent folder",
			},
			"id": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The id of the folder",
				ForceNew:     true,
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
	
apiClient := m.(*client.Client)

 name :=d.Get("name").(string)
 workspaceId :=d.Get("workspaceId").(string)
 parentFolderId :=d.Get("parentFolderId").(string)

id, err:= apiClient.Create(workspaceId, parentFolderId, name)

 
    if err != nil {
        return err
	}

d.SetId(id)
return nil;
}

func resourceReadItem(d *schema.ResourceData, m interface{}) error {

apiClient := m.(*client.Client)
item, err:= apiClient.Get(d.Id())

if err != nil {
	d.SetId("")
	return nil
} 

		d.SetId(item.Id)
	d.Set("name", item.Name)
	d.Set("parentId", item.ParentId)
	return nil
}

func resourceUpdateItem(d *schema.ResourceData, m interface{}) error {

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

	id := d.Id()

	res, err := apiClient.Exists(id)

	if err != nil {
			return false, err
	}
	return res, nil

}