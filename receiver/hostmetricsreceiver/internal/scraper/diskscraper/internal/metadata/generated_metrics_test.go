// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper/scrapertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testDataSet int

const (
	testDataSetDefault testDataSet = iota
	testDataSetAll
	testDataSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name        string
		metricsSet  testDataSet
		resAttrsSet testDataSet
		expectEmpty bool
	}{
		{
			name: "default",
		},
		{
			name:        "all_set",
			metricsSet:  testDataSetAll,
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "none_set",
			metricsSet:  testDataSetNone,
			resAttrsSet: testDataSetNone,
			expectEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := scrapertest.NewNopSettings(scrapertest.NopType)
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, tt.name), settings, WithStartTime(start))

			expectedWarnings := 0

			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSystemDiskIoDataPoint(ts, 1, "device-val", AttributeDirectionRead)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSystemDiskIoTimeDataPoint(ts, 1, "device-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSystemDiskMergedDataPoint(ts, 1, "device-val", AttributeDirectionRead)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSystemDiskOperationTimeDataPoint(ts, 1, "device-val", AttributeDirectionRead)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSystemDiskOperationsDataPoint(ts, 1, "device-val", AttributeDirectionRead)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSystemDiskPendingOperationsDataPoint(ts, 1, "device-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSystemDiskWeightedIoTimeDataPoint(ts, 1, "device-val")

			res := pcommon.NewResource()
			metrics := mb.Emit(WithResource(res))

			if tt.expectEmpty {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if tt.metricsSet == testDataSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if tt.metricsSet == testDataSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "system.disk.io":
					assert.False(t, validatedMetrics["system.disk.io"], "Found a duplicate in the metrics slice: system.disk.io")
					validatedMetrics["system.disk.io"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Disk bytes transferred.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("device")
					assert.True(t, ok)
					assert.EqualValues(t, "device-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "read", attrVal.Str())
				case "system.disk.io_time":
					assert.False(t, validatedMetrics["system.disk.io_time"], "Found a duplicate in the metrics slice: system.disk.io_time")
					validatedMetrics["system.disk.io_time"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Time disk spent activated. On Windows, this is calculated as the inverse of disk idle time.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("device")
					assert.True(t, ok)
					assert.EqualValues(t, "device-val", attrVal.Str())
				case "system.disk.merged":
					assert.False(t, validatedMetrics["system.disk.merged"], "Found a duplicate in the metrics slice: system.disk.merged")
					validatedMetrics["system.disk.merged"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of disk reads/writes merged into single physical disk access operations.", ms.At(i).Description())
					assert.Equal(t, "{operations}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("device")
					assert.True(t, ok)
					assert.EqualValues(t, "device-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "read", attrVal.Str())
				case "system.disk.operation_time":
					assert.False(t, validatedMetrics["system.disk.operation_time"], "Found a duplicate in the metrics slice: system.disk.operation_time")
					validatedMetrics["system.disk.operation_time"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Time spent in disk operations.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("device")
					assert.True(t, ok)
					assert.EqualValues(t, "device-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "read", attrVal.Str())
				case "system.disk.operations":
					assert.False(t, validatedMetrics["system.disk.operations"], "Found a duplicate in the metrics slice: system.disk.operations")
					validatedMetrics["system.disk.operations"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Disk operations count.", ms.At(i).Description())
					assert.Equal(t, "{operations}", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("device")
					assert.True(t, ok)
					assert.EqualValues(t, "device-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("direction")
					assert.True(t, ok)
					assert.EqualValues(t, "read", attrVal.Str())
				case "system.disk.pending_operations":
					assert.False(t, validatedMetrics["system.disk.pending_operations"], "Found a duplicate in the metrics slice: system.disk.pending_operations")
					validatedMetrics["system.disk.pending_operations"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The queue size of pending I/O operations.", ms.At(i).Description())
					assert.Equal(t, "{operations}", ms.At(i).Unit())
					assert.False(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("device")
					assert.True(t, ok)
					assert.EqualValues(t, "device-val", attrVal.Str())
				case "system.disk.weighted_io_time":
					assert.False(t, validatedMetrics["system.disk.weighted_io_time"], "Found a duplicate in the metrics slice: system.disk.weighted_io_time")
					validatedMetrics["system.disk.weighted_io_time"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Time disk spent activated multiplied by the queue length.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					assert.True(t, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("device")
					assert.True(t, ok)
					assert.EqualValues(t, "device-val", attrVal.Str())
				}
			}
		})
	}
}
