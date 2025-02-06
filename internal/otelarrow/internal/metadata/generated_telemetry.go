// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"context"
	"errors"
	"sync"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/embedded"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                           metric.Meter
	mu                              sync.Mutex
	registrations                   []metric.Registration
	OtelarrowAdmissionInFlightBytes metric.Int64ObservableUpDownCounter
	// TODO: Remove in v0.119.0 when remove deprecated funcs.
	observeOtelarrowAdmissionInFlightBytes func(context.Context, metric.Observer) error
	OtelarrowAdmissionWaitingBytes         metric.Int64ObservableUpDownCounter
	// TODO: Remove in v0.119.0 when remove deprecated funcs.
	observeOtelarrowAdmissionWaitingBytes func(context.Context, metric.Observer) error
}

// TelemetryBuilderOption applies changes to default builder.
type TelemetryBuilderOption interface {
	apply(*TelemetryBuilder)
}

type telemetryBuilderOptionFunc func(mb *TelemetryBuilder)

func (tbof telemetryBuilderOptionFunc) apply(mb *TelemetryBuilder) {
	tbof(mb)
}

// Deprecated: [v0.119.0] use RegisterOtelarrowAdmissionInFlightBytesCallback.
func WithOtelarrowAdmissionInFlightBytesCallback(cb func() int64, opts ...metric.ObserveOption) TelemetryBuilderOption {
	return telemetryBuilderOptionFunc(func(builder *TelemetryBuilder) {
		builder.observeOtelarrowAdmissionInFlightBytes = func(_ context.Context, o metric.Observer) error {
			o.ObserveInt64(builder.OtelarrowAdmissionInFlightBytes, cb(), opts...)
			return nil
		}
	})
}

// RegisterOtelarrowAdmissionInFlightBytesCallback sets callback for observable OtelarrowAdmissionInFlightBytes metric.
func (builder *TelemetryBuilder) RegisterOtelarrowAdmissionInFlightBytesCallback(cb metric.Int64Callback) error {
	reg, err := builder.meter.RegisterCallback(func(ctx context.Context, o metric.Observer) error {
		cb(ctx, &observerInt64{inst: builder.OtelarrowAdmissionInFlightBytes, obs: o})
		return nil
	}, builder.OtelarrowAdmissionInFlightBytes)
	if err != nil {
		return err
	}
	builder.mu.Lock()
	defer builder.mu.Unlock()
	builder.registrations = append(builder.registrations, reg)
	return nil
}

// Deprecated: [v0.119.0] use RegisterOtelarrowAdmissionWaitingBytesCallback.
func WithOtelarrowAdmissionWaitingBytesCallback(cb func() int64, opts ...metric.ObserveOption) TelemetryBuilderOption {
	return telemetryBuilderOptionFunc(func(builder *TelemetryBuilder) {
		builder.observeOtelarrowAdmissionWaitingBytes = func(_ context.Context, o metric.Observer) error {
			o.ObserveInt64(builder.OtelarrowAdmissionWaitingBytes, cb(), opts...)
			return nil
		}
	})
}

// RegisterOtelarrowAdmissionWaitingBytesCallback sets callback for observable OtelarrowAdmissionWaitingBytes metric.
func (builder *TelemetryBuilder) RegisterOtelarrowAdmissionWaitingBytesCallback(cb metric.Int64Callback) error {
	reg, err := builder.meter.RegisterCallback(func(ctx context.Context, o metric.Observer) error {
		cb(ctx, &observerInt64{inst: builder.OtelarrowAdmissionWaitingBytes, obs: o})
		return nil
	}, builder.OtelarrowAdmissionWaitingBytes)
	if err != nil {
		return err
	}
	builder.mu.Lock()
	defer builder.mu.Unlock()
	builder.registrations = append(builder.registrations, reg)
	return nil
}

type observerInt64 struct {
	embedded.Int64Observer
	inst metric.Int64Observable
	obs  metric.Observer
}

func (oi *observerInt64) Observe(value int64, opts ...metric.ObserveOption) {
	oi.obs.ObserveInt64(oi.inst, value, opts...)
}

// Shutdown unregister all registered callbacks for async instruments.
func (builder *TelemetryBuilder) Shutdown() {
	builder.mu.Lock()
	defer builder.mu.Unlock()
	for _, reg := range builder.registrations {
		reg.Unregister()
	}
}

// NewTelemetryBuilder provides a struct with methods to update all internal telemetry
// for a component
func NewTelemetryBuilder(settings component.TelemetrySettings, options ...TelemetryBuilderOption) (*TelemetryBuilder, error) {
	builder := TelemetryBuilder{}
	for _, op := range options {
		op.apply(&builder)
	}
	builder.meter = Meter(settings)
	var err, errs error
	builder.OtelarrowAdmissionInFlightBytes, err = builder.meter.Int64ObservableUpDownCounter(
		"otelcol_otelarrow_admission_in_flight_bytes",
		metric.WithDescription("Number of bytes that have started processing but are not finished."),
		metric.WithUnit("By"),
	)
	errs = errors.Join(errs, err)
	if builder.observeOtelarrowAdmissionInFlightBytes != nil {
		reg, err := builder.meter.RegisterCallback(builder.observeOtelarrowAdmissionInFlightBytes, builder.OtelarrowAdmissionInFlightBytes)
		errs = errors.Join(errs, err)
		if err == nil {
			builder.registrations = append(builder.registrations, reg)
		}
	}
	builder.OtelarrowAdmissionWaitingBytes, err = builder.meter.Int64ObservableUpDownCounter(
		"otelcol_otelarrow_admission_waiting_bytes",
		metric.WithDescription("Number of items waiting to start processing."),
		metric.WithUnit("By"),
	)
	errs = errors.Join(errs, err)
	if builder.observeOtelarrowAdmissionWaitingBytes != nil {
		reg, err := builder.meter.RegisterCallback(builder.observeOtelarrowAdmissionWaitingBytes, builder.OtelarrowAdmissionWaitingBytes)
		errs = errors.Join(errs, err)
		if err == nil {
			builder.registrations = append(builder.registrations, reg)
		}
	}
	return &builder, errs
}
