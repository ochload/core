package models

import (
    "fmt"
    "testing"
    "github.com/HenryVolkmer/di"
    "github.com/ochload/core/premiumaccounts"
)

func TestCanInitPremiumAccount(t *testing.T) {
    cfg,err := NewAppConfig("../../fixtures/appconfig_fixture.json")
    if err != nil {
        t.Fatalf(err.Error())
    }
    c := di.NewContainer()
    c.Add("AppConfig",cfg)
    c.Add("rapidgator",&premiumaccounts.Rapidgator{}).Tag("premium.account")
    c.AddCompilerPass(cfg.GetCompilerPass())
    c.Compile()
    cfgImpl := c.Get("AppConfig").(*AppConfig)
    accs,err := cfgImpl.GetConfiguredAccounts(true)
    if len(accs) != 1 {
        t.Fatalf(fmt.Sprintf("%d accounts configured ",len(accs)))
    }
}

