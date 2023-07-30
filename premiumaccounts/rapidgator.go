package premiumaccounts

import (
	"time"
	"context"
	"fmt"
	"encoding/json"
	"github.com/ochload/core/pkg/interfaces"
	"net/http"
	"net/url"
	"errors"
)

const (
	PROVIDER_NAME = "rapidgator"
	API = "https://rapidgator.net/api/v2"
)


type Rapidgator struct {
	config interfaces.AccountConfigInterface
	client *http.Client
	token string
	IsPremium bool
	validTill *time.Time
	isLoggedIn bool
}

func (this *Rapidgator) Init(cfg interfaces.AccountConfigInterface) bool {
	if cfg.GetTitle() != PROVIDER_NAME {
		return false
	}
	this.isLoggedIn = false
	this.config = cfg
	return true;
}

func (this *Rapidgator) GetFileSize() uint16 {
	return 1
}

func (this *Rapidgator) GetFilename() string {
	return "foo"
}

func (this *Rapidgator) ProbeResponsibility(url *url.URL) bool {
	return true
}

func (this *Rapidgator) Download(file *url.URL) (func(ctx ctx.Context) error,error) {
	token, err := this.getToken()
	if err != nil {
		return nil,err
	}
	return func(ctx ctx.Context) error {
		this.client = &http.Client{}
		payload := url.Values{}
		payload.Set("token", token)
		payload.Set("file_id", fileId)
		resp,err := this.client.PostForm(fmt.Sprintf("%s/file/download",API),payload)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return errors.New(fmt.Sprintf("%s Could not fetch Downloadlink!",resp.Status))
		}
		var objmap map[string]any
		err = json.NewDecoder(resp.Body).Decode(&objmap)
		
		if err != nil {
        	return err
    	}
    	
    	dlUrl := objmap["response"].(map[string]interface{})["download_url"].(string)

	}

}

func (this *Rapidgator) getToken() (string,error) {
	if !this.isLoggedIn || len(this.token) == 0 {
		err := this.login
		if err != nil {
			return _,err
		}
	}
	return this.token
}

func (this *Rapidgator) login() error {
	this.client = &http.Client{}
	payload := url.Values{}
	payload.Set("login", this.config.GetUsername())
	payload.Set("password", this.config.GetPassword())
	resp,err := this.client.PostForm(fmt.Sprintf("%s/user/login",API),payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("%s Login failed!",resp.Status))
	}
	var objmap map[string]any
	err = json.NewDecoder(resp.Body).Decode(&objmap)
	if err != nil {
        return err
    }
    this.IsPremium = objmap["response"].(map[string]interface{})["user"].(map[string]interface{})["is_premium"].(bool)
    this.token = objmap["response"].(map[string]interface{})["token"].(string)
    this.isLoggedIn = true
    return nil
}
