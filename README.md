# Overview
The Metrilcy Go Client allows you to push data to [Metricly](https://www.metricly.com) using [Go](https://golang.org).

## Package Documentation
To view package documentation, run the following `godoc` command in the cloned repo:
```shell
godoc -goroot=. -http=:8080
```
* [api doc](http://localhost:8080/pkg/api)
* [core model doc](http://localhost:8080/pkg/model/core)

## Examples
### Create a Client
```go
client := api.NewClient("https://api.app.netuitive.com/ingest", "43b6e3843e5db961fbc38cc24e796512")
````

### Create and Post Element
```go
element := core.NewElement("elementId", "Element Name", "SERVER", "use1a")
metric := core.NewMetric("metricId", "Metric Name", "bytes", "COUNTER", "None", core.Tag{"env", "prod"})
sample, _ := core.NewSample("metricId", time.Now(), 0.0)
element.AddAttribute("cpus", "4").AddTag("platform", "kubernetes").AddRelation("relatedElementId").AddMetric(metric).AddSample(sample)
elements := []core.Element{element}
client.PostElements(elements)
```

## Create and Post Event
```go
event := core.NewEvent("metric go client", "test post event", "INFO", time.Now(), core.ElementMessage{"elementId", "INFO", "test"})
event.AddTag("platform", "kubernetes")
client.PostEvents([]core.Event{event})
````

## Create and Post Check
```go
check := core.Check{"heartbeat", "elementId", 120}
client.PostCheck(check)
```