package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type TokenResponse struct {
    AccessToken string `json:"access_token"`
}

type Folder struct {
	Id string
	Name string
	ParentId string
}

type FolderItem struct {
	Id string
	Name string
	IsTrashed bool
}

type FolderContent struct {
	Items []FolderItem
}

type CreateFolderResponse struct {
	FolderId string
}



// Client holds all of the information required to connect to a server
type Client struct {
	url  string
	clientId string
	clientSecret string
	user string
	password string
	httpClient *http.Client
}

// NewClient returns a new client configured to communicate with Sony Ci
func NewClient(url string, clientId string, clientSecret string, user string, password string) *Client {
	return &Client{
		url:   url,
		clientId:  clientId,
		clientSecret: clientSecret,
		user: user,
		password: password,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetToken() (string, error) {

fmt.Println("Getting token")

			creds := c.user + ":" + c.password
            encodedStr := base64.StdEncoding.EncodeToString([]byte(creds))


 jsonStr := []byte(`{"client_id": "` + c.clientId + `", "client_secret": "` + c.clientSecret + `", grant_type: "password"}`)
req, err := http.NewRequest("POST", "https://api.cimediacloud.com/oauth2/token", bytes.NewBuffer(jsonStr))
req.Header.Add("Authorization", "Basic " + encodedStr)
 req.Header.Set("Content-Type", "application/json")

resp, err := c.httpClient.Do(req)

  if err != nil {
	  	fmt.Println(err)
        return "", err
	}	
		
	  body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    var t TokenResponse
    err = json.Unmarshal(body, &t)
    if err != nil {
        return "", err
	}

return t.AccessToken, nil

}

func (c *Client) Exists(parentFolderId string, name string) (bool, error) {
		
log.Printf("Checking folder " + name + " exists in parent folder " + parentFolderId)

resp, err := c.httpRequest("/folders/" + parentFolderId + "/contents?kind=folder&fields=name,isTrashed", "GET", bytes.Buffer{})

  if err != nil {
	  log.Printf("Error when checking" + err.Error())
        return false, err
    }	
	
	if(resp.StatusCode == 200)	{

 body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
	}

	log.Printf("[INFO] Found resource JSON: " + string(body));
	
    var t *FolderContent
    err = json.Unmarshal(body, &t)
    if err != nil {
        return false, err
	}

for i := range t.Items {
    if t.Items[i].Name == name && t.Items[i].IsTrashed == false {
		log.Printf("Found folder with id: " + t.Items[i].Id)
		return true, nil
    }
}
		
		return false, nil
}

	log.Printf("Folder " + name + " not found in parent folder " + parentFolderId)
	return false, nil

}

func (c *Client) Delete(id string) (bool, error) {

fmt.Println("Deleting folder with id " + id)

_, err := c.httpRequest("/folders/" + id, "DELETE", bytes.Buffer{})

  if err != nil {
        return false, err
	}

return true, nil
}


func (c *Client) Update(id string, name string) error {

fmt.Println("Updating folder with id " + id)

 jsonStr := []byte(`{"name": "` + name + `"}`)

_, err := c.httpRequest("/folders", "PUT", *bytes.NewBuffer(jsonStr))

  if err != nil {
        return err
	}	
	
	return nil
  }


func (c *Client) Get(id string) (*Folder, error) {
		   
fmt.Println("Getting folder with id " + id)

 resp, err := c.httpRequest("/folders/" + id, "GET", bytes.Buffer{})

  if err != nil {
        return nil, err
	}
	
	if(resp.StatusCode == 404)	{
		return nil, errors.New("Not found")
	}
	
	  body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
	}

	log.Printf("[INFO] Found resource JSON: " + string(body));
	
    var t *Folder
    err = json.Unmarshal(body, &t)
    if err != nil {
        return nil, err
	}

	return t, nil

}


func (c *Client) Create(workspaceId string, parentFolderId string, name string) (string, error) {
   
 jsonStr := []byte(`{"name": "` + name + `", "workspaceId": "` + workspaceId + `", parentFolderId: "` + parentFolderId + `"}`)
jsonByte:=  bytes.NewBuffer(jsonStr)

resp, err := c.httpRequest("/folders", "POST", *jsonByte)

  if err != nil {
        return "", err
    }	
	
	  body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    var t CreateFolderResponse
    err = json.Unmarshal(body, &t)
    if err != nil {
        return "", err
	}

	return t.FolderId, nil
}

func (c *Client) httpRequest(path string, method string, body bytes.Buffer) (response *http.Response, err error) {

newPath:=c.requestPath(path)

	req, err := http.NewRequest(method, newPath, &body)
	if err != nil {
		return nil, err
	}

token, _:= c.GetToken()

	req.Header.Add("Authorization", "Bearer " + token)
	switch method {
	case "GET":
	case "DELETE":
	default:
		req.Header.Add("Content-Type", "application/json")
	}

log.Printf("Making HTTP request to: %s",newPath)
	
	resp, err := c.httpClient.Do(req)

log.Printf("Returned status code: " + strconv.Itoa(resp.StatusCode))

	if err != nil {
		return nil, err
	}


	
	return resp, nil
}

func (c *Client) requestPath(path string) string {
	newPath:= fmt.Sprintf("%s/%s", c.url, path)
	fmt.Println("Generate path " + newPath)
	return newPath
}
