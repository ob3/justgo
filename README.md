# justgo
Simple Api Provider

### usage:

```
justgo.AddRoute(http.MethodGet, "/with-middleware", headerPrinterHandler, middleWareDummyOne, otherAuthHandler)
justgo.AddRoute(http.MethodGet, "/no-middleware", headerPrinterHandler)
justgo.Start()
```	

### Todo:
* config
* metric
* shell application interface
