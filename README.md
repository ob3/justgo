# justgo
Simple Api Provider

### usage:

~~~~
justgo.AddRoute(http.MethodGet, "/with-middleware", headerPrinterHandler, middleWareDummyOne, otherAuthHandler)
justgo.AddRoute(http.MethodGet, "/no-middleware", headerPrinterHandler)
justgo.Start()
~~~~

### Todo:
* [x] config
* [x] metric
    * [x] newrelic
    * [x] custom hooks
* [ ] sentry 
* [ ] i18n 
* [ ] DB migration
* [ ] api client with hystrix
* [ ] shell application interface
