package pkg

import (
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

func GetTracer() *zipkin.Tracer {

	return getTracer("go-server", "127.0.0.1")
}

func getTracer(serviceName string, ip string) *zipkin.Tracer {
	// create a reporter to be used by the tracer
	reporter := zipkinhttp.NewReporter("http://127.0.0.1:9411/api/v2/spans")
	// set-up the local endpoint for our service
	endpoint, _ := zipkin.NewEndpoint(serviceName, ip)
	// set-up our sampling strategy
	sampler := zipkin.NewModuloSampler(1)
	// initialize the tracer
	tracer, _ := zipkin.NewTracer(
		reporter,
		zipkin.WithLocalEndpoint(endpoint),
		zipkin.WithSampler(sampler),
	)
	return tracer
}
