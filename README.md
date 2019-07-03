# justgo
Simple Api Provider

### usage:

~~~~
justgo.AddRoute(http.MethodGet, "/with-middleware", headerPrinterHandler, middleWareDummyOne, otherAuthHandler)
justgo.AddRoute(http.MethodGet, "/no-middleware", headerPrinterHandler)
justgo.Start()
~~~~

### Features:
* [x] extensible app interface
* [x] default app interface
* [x] config
* [x] metric
    * [x] newrelic
    * [x] custom hooks
* [x] sentry reporting 
* [ ] i18n 
* [ ] api client with hystrix
* [ ] shell application interface
