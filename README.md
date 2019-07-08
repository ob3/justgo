# justgo
Simple Api Provider

### usage:

#### custom config file
as default, justgo scan app folder for application.yml for config. use this command to override where to get the config file
~~~~
justgo.Config.ConfigFile("./sample/anything.yml")
~~~~

#### in code config
~~~~
justgo.Config.Add("APP_NAME", "My Volatile Config")
~~~~


#### http
http interface is enabled by default. if you only want to use http server only, simply add route and start justgo
~~~~
justgo.AddRoute(http.MethodGet, "/with-middleware", headerPrinterHandler, middleWareDummyOne, otherAuthHandler)
justgo.AddRoute(http.MethodGet, "/no-middleware", headerPrinterHandler)
justgo.Start()
~~~~

#### Cli 
adding another interface in same app
~~~~
// add cli handler via telnet
cliInterface := &justgo.CliInterface{Address: ":12345"}

// add command on telnet command
cliInterface.AddCommand("test", echo)
cliInterface.AddCommand("panic", panicCommand)
cliInterface.AddCommand("fatal", fatalCommand)

// register interface
justgo.RegisterInterface(cliInterface)

justgo.Start()
~~~~

### Features:
* [x] app interface
    * [x] http
    * [x] telnet
    * [ ] kafka
* [x] metric
    * [x] newrelic
    * [x] custom hooks
    * [x] StatsD integration
    * [x] sentry reporting 
* [ ] i18n 
* [ ] api client with hystrix
