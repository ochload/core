package main

import (
    "flag"
    "net/http"
    "fmt"
    "log"
    "github.com/HenryVolkmer/di"
    "github.com/ochload/core/premiumaccounts"
    "github.com/ochload/core/pkg/models"
    "github.com/ochload/core/pkg/interfaces"
)

var (
    port = flag.String("port", "8080", "http-port for serving")
    configFile = flag.String("config", "config.yaml", "Path to config.yaml")
    login = flag.String("login", nil, "Login given account")
    version = "dev"
    commit  = "none"
    date    = ""
)

func main() {

    c := di.NewContainer()
    cfg,err := models.NewAppConfig(*configFile)
    Check(err)

    // configure container
    c.Add("AppConfig",cfg)
    c.Add("rapidgator",&premiumaccounts.Rapidgator{}).Tag("premium.account")
    c.Add("HttpController",&HttpController{})
    c.AddCompilerPass(fcfg.GetCompilerPass())
    

    if login {
        log.Printf("Login %s ",login)
        cfgImpl := c.Get("AppConfig").(*models.AppConfig)
        accs := cfgImpl.GetConfiguredAccounts()

        err := accs[0].Login()

        if err != nil {
            panic(err.Error())
        }

        return
    }

    // serve App
    ctrl := c.Get("HttpController").(*HttpController)
    addr := fmt.Sprintf(":%s", *port)
    log.Printf("starting http-server at %s",addr)
    mux := http.NewServeMux()
    mux.HandleFunc("/", ctrl.App)
    mux.HandleFunc("/assets/", ctrl.Assets)
    server := &http.Server{Addr: addr, Handler: mux}
    log.Printf(server.ListenAndServe().Error())
}

func Check(e error) {
    if e != nil {
        panic(e)
    }
}