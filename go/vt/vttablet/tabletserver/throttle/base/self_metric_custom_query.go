/*
Copyright 2024 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package base

import (
	"context"

	"vitess.io/vitess/go/vt/vttablet/tabletserver/connpool"
)

var _ SelfMetric = registerSelfMetric(&CustomQuerySelfMetric{})

type CustomQuerySelfMetric struct {
}

func (m *CustomQuerySelfMetric) Name() MetricName {
	return CustomMetricName
}

func (m *CustomQuerySelfMetric) DefaultScope() Scope {
	return SelfScope
}

func (m *CustomQuerySelfMetric) DefaultThreshold() float64 {
	return 0
}

func (m *CustomQuerySelfMetric) RequiresConn() bool {
	return true
}

func (m *CustomQuerySelfMetric) Read(ctx context.Context, throttler ThrottlerMetricsPublisher, conn *connpool.Conn) *ThrottleMetric {
	return ReadSelfMySQLThrottleMetric(ctx, conn, throttler.GetCustomMetricsQuery())
}
