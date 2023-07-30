package main

import (
    "os"
    "testing"
)

func TestCanDecryptDlc(t *testing.T) {
    
    file,err := os.Open("../../fixtures/dlc.dlc")
    if err != nil {
        t.Fatalf(err.Error())
    }
    
    
    
}
    