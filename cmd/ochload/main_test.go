package main

import (
    "testing"
    //"github.com/HenryVolkmer/di"
    //"github.com/ochload/core/premiumaccounts"
    "github.com/ochload/core/pkg/models"
)

func TestCanParseConfig(t *testing.T) {
    cfg,err := models.NewAppConfig("../../fixtures/appconfig_fixture.json")
    if err != nil {
        t.Fatalf(err.Error())
    }
    acc,err := cfg.FindAccountConfigByTitle("rapidgator")
    if err != nil {

        t.Fatalf("%#v",cfg.AccountConfigs)

    	t.Fatalf(err.Error())
    }
    assertEquals("123456",acc.Token,t)
}

func assertEquals(expected,actual string,t *testing.T) {
    if expected != actual {
        t.Fatalf("Failed asserting that two strings are equal.\n--- Expected\n+++ Actual\n@@ @@\n-'%s'\n+'%s'", expected, actual)
    }
}