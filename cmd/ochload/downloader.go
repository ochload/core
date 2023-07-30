package main

import (
    "net/http"
    "net/url"
    "fmt"
    "strconv"
)

type Downloader() {
    Progress float32
}

// Start/Resumes Download
func (dl *Downloader) Start(job DownloadJob) error {

    


}


// 1. HEAD File: get the Size
// 2. create 0,5MB Chunks
// 3. Download each chunk, measure Time and calculate bandwith
func (dl *Downloader) download(url *url.URL , bandwith <-chan, progress <-chan) error {
    
    finfo, err := http.Head(fmt.Sprintf("%s",url))
    if err != nil {
        return nil
    }
    if !finfo.IsOk {
        return errors.New(fmt.Sprintf("Could not HEAD URL %s: %s",url,finfo.Status))
    }

    length := finfo.Header.Get("Content-Length")

    var totalBytes int64
    totalBytes,err = strconv.ParseInt(length,10,64)
    if err != nil {
        return err
    }
}

