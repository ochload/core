package models

import (
    "os"
    "path"
    "errors"
    "fmt"
    "encoding/json"
    "github.com/HenryVolkmer/di"
    "github.com/ochload/core/pkg/interfaces"
)

type AppConfig struct {
    // Existing Account configs (accounts.json)
    AccountConfigs []AccountConfig `json:"Accounts"`
    // Accounts supported by ochload (native and plugins)
    accountsSupported []interfaces.PremiumAccount
    // Accounts configured and ready for useage (done in runtime)
    accountsArmed []interfaces.PremiumAccount
}

type AccountConfig struct {
    Title string `json`
    Token string `json`
    Username string `json`
    Password string `json`
}
func (this *AccountConfig) GetTitle() string {
    return this.Title
}
func (this *AccountConfig) GetToken() string {
    return this.Token
}
func (this *AccountConfig) GetUsername() string {
    return this.Username
}
func (this *AccountConfig) GetPassword() string {
    return this.Password
}

func NewAppConfig(cfgFile string) (*AppConfig,error) {
    cfg := &AppConfig{}
    cfgRaw, err := os.ReadFile(path.Join(cfgFile))
    if err != nil {
        return cfg,err
    }
    err = json.Unmarshal(cfgRaw, cfg)
    return cfg,err
}

func (this *AppConfig) AddAccount(a interfaces.PremiumAccount) {
    this.accountsSupported = append(this.accountsSupported,a)
}

func (this *AppConfig) FindAccountConfigByTitle(title string) (*AccountConfig,error) {
    for _,v := range this.AccountConfigs {
        if v.Title == title {
            return &v,nil
        }
    }
    return &AccountConfig{},errors.New(fmt.Sprintf("Account %s not configured!",title))
}

func (this *AppConfig) GetConfiguredAccounts(refresh bool) ([]interfaces.PremiumAccount,error) {
    if !refresh && len(this.accountsArmed) > 0 {
        return this.accountsArmed,nil
    }
    for _,cfg := range this.AccountConfigs {
        for _,acc := range this.accountsSupported {
            if ok := acc.Init(&cfg); ok {
                this.accountsArmed = append(this.accountsArmed,acc)
            }
        }
    }

    return this.accountsArmed,nil
}

func (this *AppConfig) GetCompilerPass() func(c *di.Container) {
    return func(c *di.Container) {
        accounts,ok := c.GetTaggedServices("premium.account")
        if !ok {
            panic("No Services with Tag premium.account found!")
        }
        config,_ := c.Get("AppConfig").(*AppConfig)
        for _,acc := range accounts {
            AccountService := acc.(interfaces.PremiumAccount)
            config.AddAccount(AccountService)
        }
    }
}