package observability

import (
	"context"
	"log"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/log/global"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func InitLogger() func() {
	ctx := context.Background()

	exporter, err := otlploghttp.New(ctx,
		otlploghttp.WithEndpoint("otel-collector:4318"),
		otlploghttp.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}

	provider := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(
			sdklog.NewBatchProcessor(exporter),
		),
		sdklog.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("todo-app"),
		)),
	)

	global.SetLoggerProvider(provider)

	return func() {
		if err := provider.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}
}