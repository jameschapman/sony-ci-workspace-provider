package provider

import (
	"hello/client"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_ADDRESS", ""),
			},
			"clientId": {
				Type:        schema.TypeInt,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_ID", ""),
			},
			"clientSecret": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_SECRET", ""),
			},
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("USER", ""),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PASSWORD", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource {
			"example_item": resourceItem(),
			
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	url := "https://api.cimediacloud.com"
	clientId := d.Get("clientId").(string)
	clientSecret := d.Get("clientSecret").(string)
	user := d.Get("user").(string)
	password := d.Get("password").(string)
	return client.NewClient(url, clientId, clientSecret, user, password), nil

}
