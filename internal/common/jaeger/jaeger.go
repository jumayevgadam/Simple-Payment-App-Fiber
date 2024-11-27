package jaeger

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

// NewJaegerConn initializes and returns a new OpenTelemetry trace exporter for Jaeger.
// It takes a context and the Jaeger endpoint as arguments.
func NewJaegerConn(ctx context.Context, endpoint string) (*otlptrace.Exporter, error) {
	if endpoint == "" {
		return nil, fmt.Errorf("jaeger endpoint cannot be empty")
	}

	client := otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint(endpoint),
		otlptracehttp.WithInsecure(), // Use TLS if endpoint supports secure communication
	)

	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP trace exporter: %w", err)
	}

	return exporter, nil
}
