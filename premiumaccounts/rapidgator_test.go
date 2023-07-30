package premiumaccounts

import (
    "testing"
    "github.com/ochload/core/pkg/models"
    "github.com/ochload/core/premiumaccounts"
)


func TestCanParse(t *testing.T) {
    rg := &premiumaccounts.Rapidgator{}
    cfg,_ := models.NewAppConfig("acc.json")
    cfg.AddAccount(rg)
    rgArmed,_ := cfg.GetConfiguredAccounts(true)
    rg2 := rgArmed[0].(*premiumaccounts.Rapidgator)
    rg2.Login()
    t.Fatalf("Token: %s, isPremium: %t",rg2.Token,rg2.IsPremium)
}