package interfaces

import (
	"net/url"
	"net/http"
)

type PremiumAccount interface {
	// Init the Account by given Config
	// if the provided Config is not suitable, false is returned.
	Init(cfg AccountConfigInterface) bool
	// Filesize in bytes
	GetFileSize() uint16
	GetFilename() string
	// Whether this Account should be used for Download
	// the given Download-URL.
	// A Plugin may implement this by checking the URL for 
	// known Fragments.
	ProbeResponsibility(url *url.URL) bool
	// GetClient gets a Client Instance for given File URL
	GetClient(file *url.URL) (&http.Client,error)
}

type AccountConfigInterface interface {
    GetTitle() string
    GetUsername() string
    GetPassword() string
}