package "main"

import (
	"io"
	"bytes"
	"net/http"
	"encoding/json"
	"net/url"
	"mime/multipart"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

const (
	STATE_PENDING DownloadState iota
	STATE_PROGRESS
	STATE_ERROR
	STATE_SUSPENDED
)

type DownloadState uint4 {
}
func (ds DownloadState) String() string {
	return [...]string{"pending","progress","error","suspended"}[ds]
}

// DownloadJob is the Result of decrypting a DLC
// A DownloadJob consists of many DownloadItems which
// are created too.
type DownloadJob struct {
	gorm.Model
}

// DecryptDLC decrypts a DLC by passing its content to
// POST http://dcrypt.it/decrypt/paste
// Result is a JSON:
// {"success": {"message": "Success! Here are your links.", "links": ["https://rapidgator.net/file/1a8018efc8e7962bbd91746a766d1eb1", "https://rapidgator.net/file/058ce5afd855d2169270bf016b34f3f6", "https://rapidgator.net/file/d7f8555313ec2da906ff2fddcd0006ba", "https://rapidgator.net/file/c08199567cb358d316f50109086faf1e", "https://rapidgator.net/file/59e7aa5edb4576e9ab038481cb578807", "https://rapidgator.net/file/9aa16e7bcdf8b17982cbde4187581c36", "https://rapidgator.net/file/16736daf321a42f1220502801f5b9ed2", "https://rapidgator.net/file/f4e814a164e9a4f30e68b7b37d213804", "https://rapidgator.net/file/91b2f147c6844bd52c232629d9e8f49b", "https://rapidgator.net/file/aa8df84e910733dd5ab7e8afb42d6e30"]}}
func DecryptDLC(d *DownloadJob) (dlc []bytes) ([]url.URL,error) {

	/* As Fileupload
	dlcReader := bytes.NewReader(dlc)
	var buffer bytes.Buffer
	multipartWriter := multipart.NewWriter(&buffer)

	fileField, err := multipartWriter.CreateFormFile("dlcfile", "dlcfile.dlc")
	if err != nil {
		return err
	}
	_, err = io.Copy(fileField, dlcReader)
	if err != nil {
		return err
	}

	multipartWriter.Close()

	request, err := http.NewRequest("POST", "http://dcrypt.it/decrypt/upload", &buffer)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	client := http.DefaultClient
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	*/

	this.client = &http.Client{}
	payload := url.Values{}
	payload.Set("content", string(dlc))
	resp,err := this.client.PostForm("http://dcrypt.it/decrypt/paste",payload)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Could not decrypt DLC, status code: %s", response.Status))
	}
	defer resp.Body.Close()
	var objmap map[string]any
	err = json.NewDecoder(resp.Body).Decode(&objmap)
	if err != nil {
		return err
	}
	strUrls := objmap["success"].(map[string]interface{})["links"].([]string)
	var urls []url.URL
	for _,url := range strUrls {
		newUrl,ok := url.URL.Parse(url)
		if (ok) {
			urls = append(urls, newUrl)
		}
	}
	return urls,nil
}

// NewDownloadJob creates a new DownloadJob by given encrypted DLC
func NewDownloadJob(dlc []bytes) (DownloadJob,error) {

	job := &DownloadJob{}

	var links []url.URL

	links,err = job.DecryptDLC(dlc)
	if err != nil {
		return err
	}

	for _,link := range links {



	}
	

}


type DownloadItem struct {
	gorm.Model
	Pid uint64
	State DownloadState
	TargetFile string
	FileSize uint64
}