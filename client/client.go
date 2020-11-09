package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TokenResponse struct {
    AccessToken string `json:"access_token"`
}

type Folder struct {
	Id string
	Name string
	ParentId string
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
			creds := c.user + ":" + c.password
            encodedStr := base64.StdEncoding.EncodeToString([]byte(creds))


 jsonStr := []byte(`{"client_id": "` + c.clientId + `", "client_secret": "` + c.clientSecret + `", grant_type: "password"}`)
req, err := http.NewRequest("POST", "https://api.cimediacloud.com/oauth2/token", bytes.NewBuffer(jsonStr))
req.Header.Add("Authorization", "Basic " + encodedStr)
 req.Header.Set("Content-Type", "application/json")

resp, err := c.httpClient.Do(req)

  if err != nil {
        panic(err)
	}	
		
	  body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    var t TokenResponse
    err = json.Unmarshal(body, &t)
    if err != nil {
        panic(err)
	}

return t.AccessToken, nil

}

func (c *Client) Exists(id string) (bool, error) {
		
resp, err := c.httpRequest("GET", "/folders/" + id, bytes.Buffer{})

  if err != nil {
        return false, err
    }	
	
	// should try and improve this...
	if(resp.StatusCode == 200)	{
		return true, nil
	}

	return false, nil

}

func (c *Client) Delete(id string) (bool, error) {

_, err := c.httpRequest("DELETE", "/folders/" + id, bytes.Buffer{})

  if err != nil {
        return false, err
	}

return true, nil
}


func (c *Client) Update(id string, name string) error {
 jsonStr := []byte(`{"name": "` + name + `"}`)

_, err := c.httpRequest("PUT", "https://api.cimediacloud.com/folders", *bytes.NewBuffer(jsonStr))

  if err != nil {
        return err
	}	
	
	return nil
  }


func (c *Client) Get(id string) (*Folder, error) {
		   
 resp, err := c.httpRequest("/folders" + id, "GET", bytes.Buffer{})

  if err != nil {
        panic(err)
	}
	
	if(resp.StatusCode == 404)	{
		return nil, nil
	}
	
	  body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
	}
	
    fmt.Println((string(body)))
    var t *Folder
    err = json.Unmarshal(body, &t)
    if err != nil {
        panic(err)
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

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (response *http.Response, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
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

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s/%s", c.url, path)
}