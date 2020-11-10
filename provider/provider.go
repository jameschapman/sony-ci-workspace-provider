package provider

import (
	"sonyciprovider/client"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_ID", ""),
			},
			"client_secret": {
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
			"ci_folder": resourceItem(),
			
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	url := "https://api.cimediacloud.com"
	clientId := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)
	user := d.Get("user").(string)
	password := d.Get("password").(string)
	return client.NewClient(url, clientId, clientSecret, user, password), nil

}
