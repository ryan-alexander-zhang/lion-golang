
```shell
 curl -v \
  -H 'traceparent: 00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01' \
  http://localhost:8181/ping
* Host localhost:8181 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8181...
* Connected to localhost (::1) port 8181
> GET /ping HTTP/1.1
> Host: localhost:8181
> User-Agent: curl/8.7.1
> Accept: */*
> traceparent: 00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Server-Timing: traceparent;desc="00-4bf92f3577b34da6a3ce929d0e0e4736-5b706537397df6db-01"
< Traceparent: 00-4bf92f3577b34da6a3ce929d0e0e4736-5b706537397df6db-01
< X-Request-Id: 5a8b8fe3-8316-4061-bf98-fdb622e88773
< Date: Thu, 30 Oct 2025 09:29:47 GMT
< Content-Length: 18
< 
* Connection #0 to host localhost left intact
{"message":"pong"}%                                                                                                                                                             ➜  lion-golang git:(main) ✗ 

```


```shell
{"level":"info","ts":"2025-10-30T09:29:47.502419Z","caller":"gin/main.go:62","msg":"ping coming"}
{"level":"info","ts":"2025-10-30T09:29:47.502761Z","caller":"zap@v1.1.5/zap.go:125","msg":"/ping","status":200,"method":"GET","path":"/ping","query":"","ip":"::1","user-agent":"curl/8.7.1","latency":0.000317417,"time":"2025-10-30T09:29:47Z","request_id":"5a8b8fe3-8316-4061-bf98-fdb622e88773","trace_id":"4bf92f3577b34da6a3ce929d0e0e4736","span_id":"5b706537397df6db","body":""}
```

```json
{
	"Name": "GET /ping",
	"SpanContext": {
		"TraceID": "4bf92f3577b34da6a3ce929d0e0e4736",
		"SpanID": "5b706537397df6db",
		"TraceFlags": "01",
		"TraceState": "",
		"Remote": false
	},
	"Parent": {
		"TraceID": "4bf92f3577b34da6a3ce929d0e0e4736",
		"SpanID": "00f067aa0ba902b7",
		"TraceFlags": "01",
		"TraceState": "",
		"Remote": true
	},
	"SpanKind": 2,
	"StartTime": "2025-10-30T17:29:47.502297+08:00",
	"EndTime": "2025-10-30T17:29:47.50280125+08:00",
	"Attributes": [
		{
			"Key": "server.address",
			"Value": {
				"Type": "STRING",
				"Value": "your-service-name"
			}
		},
		{
			"Key": "http.request.method",
			"Value": {
				"Type": "STRING",
				"Value": "GET"
			}
		},
		{
			"Key": "url.scheme",
			"Value": {
				"Type": "STRING",
				"Value": "http"
			}
		},
		{
			"Key": "server.port",
			"Value": {
				"Type": "INT64",
				"Value": 8181
			}
		},
		{
			"Key": "network.peer.address",
			"Value": {
				"Type": "STRING",
				"Value": "::1"
			}
		},
		{
			"Key": "network.peer.port",
			"Value": {
				"Type": "INT64",
				"Value": 62827
			}
		},
		{
			"Key": "user_agent.original",
			"Value": {
				"Type": "STRING",
				"Value": "curl/8.7.1"
			}
		},
		{
			"Key": "client.address",
			"Value": {
				"Type": "STRING",
				"Value": "::1"
			}
		},
		{
			"Key": "url.path",
			"Value": {
				"Type": "STRING",
				"Value": "/ping"
			}
		},
		{
			"Key": "network.protocol.version",
			"Value": {
				"Type": "STRING",
				"Value": "1.1"
			}
		},
		{
			"Key": "http.route",
			"Value": {
				"Type": "STRING",
				"Value": "/ping"
			}
		},
		{
			"Key": "http.response.body.size",
			"Value": {
				"Type": "INT64",
				"Value": 18
			}
		},
		{
			"Key": "http.response.status_code",
			"Value": {
				"Type": "INT64",
				"Value": 200
			}
		}
	],
	"Events": null,
	"Links": null,
	"Status": {
		"Code": "Unset",
		"Description": ""
	},
	"DroppedAttributes": 0,
	"DroppedEvents": 0,
	"DroppedLinks": 0,
	"ChildSpanCount": 0,
	"Resource": [
		{
			"Key": "service.name",
			"Value": {
				"Type": "STRING",
				"Value": "your-service-name"
			}
		}
	],
	"InstrumentationScope": {
		"Name": "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin",
		"Version": "0.63.0",
		"SchemaURL": "",
		"Attributes": null
	},
	"InstrumentationLibrary": {
		"Name": "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin",
		"Version": "0.63.0",
		"SchemaURL": "",
		"Attributes": null
	}
}
```