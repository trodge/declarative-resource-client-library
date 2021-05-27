// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	monitoringpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/monitoring_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring"
)

// Server implements the gRPC interface for Dashboard.
type DashboardServer struct{}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(n[len("DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(n[len("DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(n[len("DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(n[len("DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(n[len("DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(n[len("DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(n[len("DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(n[len("DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(n[len("DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(n[len("DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(n[len("DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(n[len("DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(n[len("DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(n[len("DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum(n[len("DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartDataSetsPlotTypeEnum converts a DashboardWidgetXyChartDataSetsPlotTypeEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum(e monitoringpb.MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum) *monitoring.DashboardWidgetXyChartDataSetsPlotTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartDataSetsPlotTypeEnum(n[len("DashboardWidgetXyChartDataSetsPlotTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartThresholdsColorEnum converts a DashboardWidgetXyChartThresholdsColorEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartThresholdsColorEnum(e monitoringpb.MonitoringDashboardWidgetXyChartThresholdsColorEnum) *monitoring.DashboardWidgetXyChartThresholdsColorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartThresholdsColorEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartThresholdsColorEnum(n[len("DashboardWidgetXyChartThresholdsColorEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartThresholdsDirectionEnum converts a DashboardWidgetXyChartThresholdsDirectionEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartThresholdsDirectionEnum(e monitoringpb.MonitoringDashboardWidgetXyChartThresholdsDirectionEnum) *monitoring.DashboardWidgetXyChartThresholdsDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartThresholdsDirectionEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartThresholdsDirectionEnum(n[len("DashboardWidgetXyChartThresholdsDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartXAxisScaleEnum converts a DashboardWidgetXyChartXAxisScaleEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartXAxisScaleEnum(e monitoringpb.MonitoringDashboardWidgetXyChartXAxisScaleEnum) *monitoring.DashboardWidgetXyChartXAxisScaleEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartXAxisScaleEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartXAxisScaleEnum(n[len("DashboardWidgetXyChartXAxisScaleEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartYAxisScaleEnum converts a DashboardWidgetXyChartYAxisScaleEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartYAxisScaleEnum(e monitoringpb.MonitoringDashboardWidgetXyChartYAxisScaleEnum) *monitoring.DashboardWidgetXyChartYAxisScaleEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartYAxisScaleEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartYAxisScaleEnum(n[len("DashboardWidgetXyChartYAxisScaleEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetXyChartChartOptionsModeEnum converts a DashboardWidgetXyChartChartOptionsModeEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartChartOptionsModeEnum(e monitoringpb.MonitoringDashboardWidgetXyChartChartOptionsModeEnum) *monitoring.DashboardWidgetXyChartChartOptionsModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetXyChartChartOptionsModeEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetXyChartChartOptionsModeEnum(n[len("DashboardWidgetXyChartChartOptionsModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(n[len("DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(n[len("DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(n[len("DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(n[len("DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(n[len("DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(n[len("DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(n[len("DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(n[len("DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(n[len("DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(n[len("DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(n[len("DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(n[len("DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(n[len("DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(n[len("DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryApiSourceEnum converts a DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryApiSourceEnum(e monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryApiSourceEnum) *monitoring.DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryApiSourceEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum(n[len("DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum converts a DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(e monitoringpb.MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum) *monitoring.DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(n[len("DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardThresholdsColorEnum converts a DashboardWidgetScorecardThresholdsColorEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardThresholdsColorEnum(e monitoringpb.MonitoringDashboardWidgetScorecardThresholdsColorEnum) *monitoring.DashboardWidgetScorecardThresholdsColorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardThresholdsColorEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardThresholdsColorEnum(n[len("DashboardWidgetScorecardThresholdsColorEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetScorecardThresholdsDirectionEnum converts a DashboardWidgetScorecardThresholdsDirectionEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardThresholdsDirectionEnum(e monitoringpb.MonitoringDashboardWidgetScorecardThresholdsDirectionEnum) *monitoring.DashboardWidgetScorecardThresholdsDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetScorecardThresholdsDirectionEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetScorecardThresholdsDirectionEnum(n[len("DashboardWidgetScorecardThresholdsDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardWidgetTextFormatEnum converts a DashboardWidgetTextFormatEnum enum from its proto representation.
func ProtoToMonitoringDashboardWidgetTextFormatEnum(e monitoringpb.MonitoringDashboardWidgetTextFormatEnum) *monitoring.DashboardWidgetTextFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringDashboardWidgetTextFormatEnum_name[int32(e)]; ok {
		e := monitoring.DashboardWidgetTextFormatEnum(n[len("DashboardWidgetTextFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToDashboardGridLayout converts a DashboardGridLayout resource from its proto representation.
func ProtoToMonitoringDashboardGridLayout(p *monitoringpb.MonitoringDashboardGridLayout) *monitoring.DashboardGridLayout {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardGridLayout{
		Columns: dcl.Int64OrNil(p.Columns),
	}
	for _, r := range p.GetWidgets() {
		obj.Widgets = append(obj.Widgets, *ProtoToMonitoringDashboardWidget(r))
	}
	return obj
}

// ProtoToDashboardMosaicLayout converts a DashboardMosaicLayout resource from its proto representation.
func ProtoToMonitoringDashboardMosaicLayout(p *monitoringpb.MonitoringDashboardMosaicLayout) *monitoring.DashboardMosaicLayout {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardMosaicLayout{
		Columns: dcl.Int64OrNil(p.Columns),
	}
	for _, r := range p.GetTiles() {
		obj.Tiles = append(obj.Tiles, *ProtoToMonitoringDashboardMosaicLayoutTiles(r))
	}
	return obj
}

// ProtoToDashboardMosaicLayoutTiles converts a DashboardMosaicLayoutTiles resource from its proto representation.
func ProtoToMonitoringDashboardMosaicLayoutTiles(p *monitoringpb.MonitoringDashboardMosaicLayoutTiles) *monitoring.DashboardMosaicLayoutTiles {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardMosaicLayoutTiles{
		XPos:   dcl.Int64OrNil(p.XPos),
		YPos:   dcl.Int64OrNil(p.YPos),
		Width:  dcl.Int64OrNil(p.Width),
		Height: dcl.Int64OrNil(p.Height),
		Widget: ProtoToMonitoringDashboardWidget(p.GetWidget()),
	}
	return obj
}

// ProtoToDashboardRowLayout converts a DashboardRowLayout resource from its proto representation.
func ProtoToMonitoringDashboardRowLayout(p *monitoringpb.MonitoringDashboardRowLayout) *monitoring.DashboardRowLayout {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardRowLayout{}
	for _, r := range p.GetRows() {
		obj.Rows = append(obj.Rows, *ProtoToMonitoringDashboardRowLayoutRows(r))
	}
	return obj
}

// ProtoToDashboardRowLayoutRows converts a DashboardRowLayoutRows resource from its proto representation.
func ProtoToMonitoringDashboardRowLayoutRows(p *monitoringpb.MonitoringDashboardRowLayoutRows) *monitoring.DashboardRowLayoutRows {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardRowLayoutRows{
		Weight: dcl.Int64OrNil(p.Weight),
	}
	for _, r := range p.GetWidgets() {
		obj.Widgets = append(obj.Widgets, *ProtoToMonitoringDashboardWidget(r))
	}
	return obj
}

// ProtoToDashboardColumnLayout converts a DashboardColumnLayout resource from its proto representation.
func ProtoToMonitoringDashboardColumnLayout(p *monitoringpb.MonitoringDashboardColumnLayout) *monitoring.DashboardColumnLayout {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardColumnLayout{}
	for _, r := range p.GetColumns() {
		obj.Columns = append(obj.Columns, *ProtoToMonitoringDashboardColumnLayoutColumns(r))
	}
	return obj
}

// ProtoToDashboardColumnLayoutColumns converts a DashboardColumnLayoutColumns resource from its proto representation.
func ProtoToMonitoringDashboardColumnLayoutColumns(p *monitoringpb.MonitoringDashboardColumnLayoutColumns) *monitoring.DashboardColumnLayoutColumns {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardColumnLayoutColumns{
		Weight: dcl.Int64OrNil(p.Weight),
	}
	for _, r := range p.GetWidgets() {
		obj.Widgets = append(obj.Widgets, *ProtoToMonitoringDashboardWidget(r))
	}
	return obj
}

// ProtoToDashboardWidget converts a DashboardWidget resource from its proto representation.
func ProtoToMonitoringDashboardWidget(p *monitoringpb.MonitoringDashboardWidget) *monitoring.DashboardWidget {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidget{
		Title:     dcl.StringOrNil(p.Title),
		XyChart:   ProtoToMonitoringDashboardWidgetXyChart(p.GetXyChart()),
		Scorecard: ProtoToMonitoringDashboardWidgetScorecard(p.GetScorecard()),
		Text:      ProtoToMonitoringDashboardWidgetText(p.GetText()),
		Blank:     ProtoToMonitoringDashboardWidgetBlank(p.GetBlank()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChart converts a DashboardWidgetXyChart resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChart(p *monitoringpb.MonitoringDashboardWidgetXyChart) *monitoring.DashboardWidgetXyChart {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChart{
		TimeshiftDuration: dcl.StringOrNil(p.TimeshiftDuration),
		XAxis:             ProtoToMonitoringDashboardWidgetXyChartXAxis(p.GetXAxis()),
		YAxis:             ProtoToMonitoringDashboardWidgetXyChartYAxis(p.GetYAxis()),
		ChartOptions:      ProtoToMonitoringDashboardWidgetXyChartChartOptions(p.GetChartOptions()),
	}
	for _, r := range p.GetDataSets() {
		obj.DataSets = append(obj.DataSets, *ProtoToMonitoringDashboardWidgetXyChartDataSets(r))
	}
	for _, r := range p.GetThresholds() {
		obj.Thresholds = append(obj.Thresholds, *ProtoToMonitoringDashboardWidgetXyChartThresholds(r))
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSets converts a DashboardWidgetXyChartDataSets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSets) *monitoring.DashboardWidgetXyChartDataSets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSets{
		TimeSeriesQuery:    ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQuery(p.GetTimeSeriesQuery()),
		PlotType:           ProtoToMonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum(p.GetPlotType()),
		LegendTemplate:     dcl.StringOrNil(p.LegendTemplate),
		MinAlignmentPeriod: dcl.StringOrNil(p.MinAlignmentPeriod),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQuery converts a DashboardWidgetXyChartDataSetsTimeSeriesQuery resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQuery(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQuery) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQuery {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQuery{
		TimeSeriesFilter:        ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(p.GetTimeSeriesFilter()),
		TimeSeriesFilterRatio:   ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(p.GetTimeSeriesFilterRatio()),
		TimeSeriesQueryLanguage: dcl.StringOrNil(p.TimeSeriesQueryLanguage),
		ApiSource:               ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum(p.GetApiSource()),
		UnitOverride:            dcl.StringOrNil(p.UnitOverride),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{
		Filter:               dcl.StringOrNil(p.Filter),
		Aggregation:          ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(p.GetAggregation()),
		SecondaryAggregation: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(p.GetSecondaryAggregation()),
		PickTimeSeriesFilter: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(p.GetPickTimeSeriesFilter()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{
		AlignmentPeriod:              dcl.StringOrNil(p.AlignmentPeriod),
		PerSeriesAligner:             ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer:           ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
		ReduceFractionLessThanParams: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams(p.GetReduceFractionLessThanParams()),
		ReduceMakeDistributionParams: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams(p.GetReduceMakeDistributionParams()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams{
		Threshold: dcl.Float64OrNil(p.Threshold),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams{
		BucketOptions:    ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions(p.GetBucketOptions()),
		ExemplarSampling: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling(p.GetExemplarSampling()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		Width:            dcl.Float64OrNil(p.Width),
		Offset:           dcl.Float64OrNil(p.Offset),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		GrowthFactor:     dcl.Float64OrNil(p.GrowthFactor),
		Scale:            dcl.Float64OrNil(p.Scale),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.Float64OrNil(p.MinimumValue),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{
		AlignmentPeriod:              dcl.StringOrNil(p.AlignmentPeriod),
		PerSeriesAligner:             ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer:           ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
		ReduceFractionLessThanParams: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams(p.GetReduceFractionLessThanParams()),
		ReduceMakeDistributionParams: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams(p.GetReduceMakeDistributionParams()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams{
		Threshold: dcl.Float64OrNil(p.Threshold),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams{
		BucketOptions:    ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions(p.GetBucketOptions()),
		ExemplarSampling: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling(p.GetExemplarSampling()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		Width:            dcl.Float64OrNil(p.Width),
		Offset:           dcl.Float64OrNil(p.Offset),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		GrowthFactor:     dcl.Float64OrNil(p.GrowthFactor),
		Scale:            dcl.Float64OrNil(p.Scale),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.Float64OrNil(p.MinimumValue),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{
		RankingMethod: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(p.GetRankingMethod()),
		NumTimeSeries: dcl.Int64OrNil(p.NumTimeSeries),
		Direction:     ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{
		Numerator:            ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(p.GetNumerator()),
		Denominator:          ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(p.GetDenominator()),
		SecondaryAggregation: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(p.GetSecondaryAggregation()),
		PickTimeSeriesFilter: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(p.GetPickTimeSeriesFilter()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{
		Filter:      dcl.StringOrNil(p.Filter),
		Aggregation: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(p.GetAggregation()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{
		AlignmentPeriod:              dcl.StringOrNil(p.AlignmentPeriod),
		PerSeriesAligner:             ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer:           ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
		ReduceFractionLessThanParams: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams(p.GetReduceFractionLessThanParams()),
		ReduceMakeDistributionParams: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams(p.GetReduceMakeDistributionParams()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams{
		Threshold: dcl.Float64OrNil(p.Threshold),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams{
		BucketOptions:    ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions(p.GetBucketOptions()),
		ExemplarSampling: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling(p.GetExemplarSampling()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		Width:            dcl.Float64OrNil(p.Width),
		Offset:           dcl.Float64OrNil(p.Offset),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		GrowthFactor:     dcl.Float64OrNil(p.GrowthFactor),
		Scale:            dcl.Float64OrNil(p.Scale),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.Float64OrNil(p.MinimumValue),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{
		Filter:      dcl.StringOrNil(p.Filter),
		Aggregation: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(p.GetAggregation()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{
		AlignmentPeriod:              dcl.StringOrNil(p.AlignmentPeriod),
		PerSeriesAligner:             ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer:           ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
		ReduceFractionLessThanParams: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams(p.GetReduceFractionLessThanParams()),
		ReduceMakeDistributionParams: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams(p.GetReduceMakeDistributionParams()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams{
		Threshold: dcl.Float64OrNil(p.Threshold),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams{
		BucketOptions:    ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions(p.GetBucketOptions()),
		ExemplarSampling: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling(p.GetExemplarSampling()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		Width:            dcl.Float64OrNil(p.Width),
		Offset:           dcl.Float64OrNil(p.Offset),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		GrowthFactor:     dcl.Float64OrNil(p.GrowthFactor),
		Scale:            dcl.Float64OrNil(p.Scale),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.Float64OrNil(p.MinimumValue),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{
		AlignmentPeriod:              dcl.StringOrNil(p.AlignmentPeriod),
		PerSeriesAligner:             ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer:           ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
		ReduceFractionLessThanParams: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams(p.GetReduceFractionLessThanParams()),
		ReduceMakeDistributionParams: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams(p.GetReduceMakeDistributionParams()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams{
		Threshold: dcl.Float64OrNil(p.Threshold),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams{
		BucketOptions:    ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions(p.GetBucketOptions()),
		ExemplarSampling: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling(p.GetExemplarSampling()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		Width:            dcl.Float64OrNil(p.Width),
		Offset:           dcl.Float64OrNil(p.Offset),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		GrowthFactor:     dcl.Float64OrNil(p.GrowthFactor),
		Scale:            dcl.Float64OrNil(p.Scale),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.Float64OrNil(p.MinimumValue),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(p *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{
		RankingMethod: ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(p.GetRankingMethod()),
		NumTimeSeries: dcl.Int64OrNil(p.NumTimeSeries),
		Direction:     ProtoToMonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartThresholds converts a DashboardWidgetXyChartThresholds resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartThresholds(p *monitoringpb.MonitoringDashboardWidgetXyChartThresholds) *monitoring.DashboardWidgetXyChartThresholds {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartThresholds{
		Label:     dcl.StringOrNil(p.Label),
		Value:     dcl.Float64OrNil(p.Value),
		Color:     ProtoToMonitoringDashboardWidgetXyChartThresholdsColorEnum(p.GetColor()),
		Direction: ProtoToMonitoringDashboardWidgetXyChartThresholdsDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartXAxis converts a DashboardWidgetXyChartXAxis resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartXAxis(p *monitoringpb.MonitoringDashboardWidgetXyChartXAxis) *monitoring.DashboardWidgetXyChartXAxis {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartXAxis{
		Label: dcl.StringOrNil(p.Label),
		Scale: ProtoToMonitoringDashboardWidgetXyChartXAxisScaleEnum(p.GetScale()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartYAxis converts a DashboardWidgetXyChartYAxis resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartYAxis(p *monitoringpb.MonitoringDashboardWidgetXyChartYAxis) *monitoring.DashboardWidgetXyChartYAxis {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartYAxis{
		Label: dcl.StringOrNil(p.Label),
		Scale: ProtoToMonitoringDashboardWidgetXyChartYAxisScaleEnum(p.GetScale()),
	}
	return obj
}

// ProtoToDashboardWidgetXyChartChartOptions converts a DashboardWidgetXyChartChartOptions resource from its proto representation.
func ProtoToMonitoringDashboardWidgetXyChartChartOptions(p *monitoringpb.MonitoringDashboardWidgetXyChartChartOptions) *monitoring.DashboardWidgetXyChartChartOptions {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetXyChartChartOptions{
		Mode: ProtoToMonitoringDashboardWidgetXyChartChartOptionsModeEnum(p.GetMode()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecard converts a DashboardWidgetScorecard resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecard(p *monitoringpb.MonitoringDashboardWidgetScorecard) *monitoring.DashboardWidgetScorecard {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecard{
		TimeSeriesQuery: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQuery(p.GetTimeSeriesQuery()),
		GaugeView:       ProtoToMonitoringDashboardWidgetScorecardGaugeView(p.GetGaugeView()),
		SparkChartView:  ProtoToMonitoringDashboardWidgetScorecardSparkChartView(p.GetSparkChartView()),
	}
	for _, r := range p.GetThresholds() {
		obj.Thresholds = append(obj.Thresholds, *ProtoToMonitoringDashboardWidgetScorecardThresholds(r))
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQuery converts a DashboardWidgetScorecardTimeSeriesQuery resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQuery(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQuery) *monitoring.DashboardWidgetScorecardTimeSeriesQuery {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQuery{
		TimeSeriesFilter:        ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(p.GetTimeSeriesFilter()),
		TimeSeriesFilterRatio:   ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(p.GetTimeSeriesFilterRatio()),
		TimeSeriesQueryLanguage: dcl.StringOrNil(p.TimeSeriesQueryLanguage),
		ApiSource:               ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryApiSourceEnum(p.GetApiSource()),
		UnitOverride:            dcl.StringOrNil(p.UnitOverride),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{
		Filter:               dcl.StringOrNil(p.Filter),
		Aggregation:          ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(p.GetAggregation()),
		SecondaryAggregation: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(p.GetSecondaryAggregation()),
		PickTimeSeriesFilter: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(p.GetPickTimeSeriesFilter()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{
		AlignmentPeriod:              dcl.StringOrNil(p.AlignmentPeriod),
		PerSeriesAligner:             ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer:           ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
		ReduceFractionLessThanParams: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams(p.GetReduceFractionLessThanParams()),
		ReduceMakeDistributionParams: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams(p.GetReduceMakeDistributionParams()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams{
		Threshold: dcl.Float64OrNil(p.Threshold),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams{
		BucketOptions:    ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions(p.GetBucketOptions()),
		ExemplarSampling: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling(p.GetExemplarSampling()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		Width:            dcl.Float64OrNil(p.Width),
		Offset:           dcl.Float64OrNil(p.Offset),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		GrowthFactor:     dcl.Float64OrNil(p.GrowthFactor),
		Scale:            dcl.Float64OrNil(p.Scale),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.Float64OrNil(p.MinimumValue),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{
		AlignmentPeriod:              dcl.StringOrNil(p.AlignmentPeriod),
		PerSeriesAligner:             ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer:           ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
		ReduceFractionLessThanParams: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams(p.GetReduceFractionLessThanParams()),
		ReduceMakeDistributionParams: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams(p.GetReduceMakeDistributionParams()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams{
		Threshold: dcl.Float64OrNil(p.Threshold),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams{
		BucketOptions:    ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions(p.GetBucketOptions()),
		ExemplarSampling: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling(p.GetExemplarSampling()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		Width:            dcl.Float64OrNil(p.Width),
		Offset:           dcl.Float64OrNil(p.Offset),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		GrowthFactor:     dcl.Float64OrNil(p.GrowthFactor),
		Scale:            dcl.Float64OrNil(p.Scale),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.Float64OrNil(p.MinimumValue),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{
		RankingMethod: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(p.GetRankingMethod()),
		NumTimeSeries: dcl.Int64OrNil(p.NumTimeSeries),
		Direction:     ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{
		Numerator:            ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(p.GetNumerator()),
		Denominator:          ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(p.GetDenominator()),
		SecondaryAggregation: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(p.GetSecondaryAggregation()),
		PickTimeSeriesFilter: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(p.GetPickTimeSeriesFilter()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{
		Filter:      dcl.StringOrNil(p.Filter),
		Aggregation: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(p.GetAggregation()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{
		AlignmentPeriod:              dcl.StringOrNil(p.AlignmentPeriod),
		PerSeriesAligner:             ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer:           ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
		ReduceFractionLessThanParams: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams(p.GetReduceFractionLessThanParams()),
		ReduceMakeDistributionParams: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams(p.GetReduceMakeDistributionParams()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams{
		Threshold: dcl.Float64OrNil(p.Threshold),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams{
		BucketOptions:    ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions(p.GetBucketOptions()),
		ExemplarSampling: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling(p.GetExemplarSampling()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		Width:            dcl.Float64OrNil(p.Width),
		Offset:           dcl.Float64OrNil(p.Offset),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		GrowthFactor:     dcl.Float64OrNil(p.GrowthFactor),
		Scale:            dcl.Float64OrNil(p.Scale),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.Float64OrNil(p.MinimumValue),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{
		Filter:      dcl.StringOrNil(p.Filter),
		Aggregation: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(p.GetAggregation()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{
		AlignmentPeriod:              dcl.StringOrNil(p.AlignmentPeriod),
		PerSeriesAligner:             ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer:           ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
		ReduceFractionLessThanParams: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams(p.GetReduceFractionLessThanParams()),
		ReduceMakeDistributionParams: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams(p.GetReduceMakeDistributionParams()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams{
		Threshold: dcl.Float64OrNil(p.Threshold),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams{
		BucketOptions:    ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions(p.GetBucketOptions()),
		ExemplarSampling: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling(p.GetExemplarSampling()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		Width:            dcl.Float64OrNil(p.Width),
		Offset:           dcl.Float64OrNil(p.Offset),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		GrowthFactor:     dcl.Float64OrNil(p.GrowthFactor),
		Scale:            dcl.Float64OrNil(p.Scale),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.Float64OrNil(p.MinimumValue),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{
		AlignmentPeriod:              dcl.StringOrNil(p.AlignmentPeriod),
		PerSeriesAligner:             ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(p.GetPerSeriesAligner()),
		CrossSeriesReducer:           ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(p.GetCrossSeriesReducer()),
		ReduceFractionLessThanParams: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams(p.GetReduceFractionLessThanParams()),
		ReduceMakeDistributionParams: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams(p.GetReduceMakeDistributionParams()),
	}
	for _, r := range p.GetGroupByFields() {
		obj.GroupByFields = append(obj.GroupByFields, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams{
		Threshold: dcl.Float64OrNil(p.Threshold),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams{
		BucketOptions:    ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions(p.GetBucketOptions()),
		ExemplarSampling: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling(p.GetExemplarSampling()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		Width:            dcl.Float64OrNil(p.Width),
		Offset:           dcl.Float64OrNil(p.Offset),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.NumFiniteBuckets),
		GrowthFactor:     dcl.Float64OrNil(p.GrowthFactor),
		Scale:            dcl.Float64OrNil(p.Scale),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.Float64OrNil(p.MinimumValue),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter(p *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{
		RankingMethod: ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(p.GetRankingMethod()),
		NumTimeSeries: dcl.Int64OrNil(p.NumTimeSeries),
		Direction:     ProtoToMonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardGaugeView converts a DashboardWidgetScorecardGaugeView resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardGaugeView(p *monitoringpb.MonitoringDashboardWidgetScorecardGaugeView) *monitoring.DashboardWidgetScorecardGaugeView {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardGaugeView{
		LowerBound: dcl.Float64OrNil(p.LowerBound),
		UpperBound: dcl.Float64OrNil(p.UpperBound),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardSparkChartView converts a DashboardWidgetScorecardSparkChartView resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardSparkChartView(p *monitoringpb.MonitoringDashboardWidgetScorecardSparkChartView) *monitoring.DashboardWidgetScorecardSparkChartView {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardSparkChartView{
		SparkChartType:     ProtoToMonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(p.GetSparkChartType()),
		MinAlignmentPeriod: dcl.StringOrNil(p.MinAlignmentPeriod),
	}
	return obj
}

// ProtoToDashboardWidgetScorecardThresholds converts a DashboardWidgetScorecardThresholds resource from its proto representation.
func ProtoToMonitoringDashboardWidgetScorecardThresholds(p *monitoringpb.MonitoringDashboardWidgetScorecardThresholds) *monitoring.DashboardWidgetScorecardThresholds {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetScorecardThresholds{
		Label:     dcl.StringOrNil(p.Label),
		Value:     dcl.Float64OrNil(p.Value),
		Color:     ProtoToMonitoringDashboardWidgetScorecardThresholdsColorEnum(p.GetColor()),
		Direction: ProtoToMonitoringDashboardWidgetScorecardThresholdsDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToDashboardWidgetText converts a DashboardWidgetText resource from its proto representation.
func ProtoToMonitoringDashboardWidgetText(p *monitoringpb.MonitoringDashboardWidgetText) *monitoring.DashboardWidgetText {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetText{
		Content: dcl.StringOrNil(p.Content),
		Format:  ProtoToMonitoringDashboardWidgetTextFormatEnum(p.GetFormat()),
	}
	return obj
}

// ProtoToDashboardWidgetBlank converts a DashboardWidgetBlank resource from its proto representation.
func ProtoToMonitoringDashboardWidgetBlank(p *monitoringpb.MonitoringDashboardWidgetBlank) *monitoring.DashboardWidgetBlank {
	if p == nil {
		return nil
	}
	obj := &monitoring.DashboardWidgetBlank{}
	return obj
}

// ProtoToDashboard converts a Dashboard resource from its proto representation.
func ProtoToDashboard(p *monitoringpb.MonitoringDashboard) *monitoring.Dashboard {
	obj := &monitoring.Dashboard{
		Name:         dcl.StringOrNil(p.Name),
		DisplayName:  dcl.StringOrNil(p.DisplayName),
		GridLayout:   ProtoToMonitoringDashboardGridLayout(p.GetGridLayout()),
		MosaicLayout: ProtoToMonitoringDashboardMosaicLayout(p.GetMosaicLayout()),
		RowLayout:    ProtoToMonitoringDashboardRowLayout(p.GetRowLayout()),
		ColumnLayout: ProtoToMonitoringDashboardColumnLayout(p.GetColumnLayout()),
		Project:      dcl.StringOrNil(p.Project),
		Etag:         dcl.StringOrNil(p.Etag),
	}
	return obj
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(0)
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnumToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum_value["DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum(0)
}

// DashboardWidgetXyChartDataSetsPlotTypeEnumToProto converts a DashboardWidgetXyChartDataSetsPlotTypeEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnumToProto(e *monitoring.DashboardWidgetXyChartDataSetsPlotTypeEnum) monitoringpb.MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum_value["DashboardWidgetXyChartDataSetsPlotTypeEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnum(0)
}

// DashboardWidgetXyChartThresholdsColorEnumToProto converts a DashboardWidgetXyChartThresholdsColorEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartThresholdsColorEnumToProto(e *monitoring.DashboardWidgetXyChartThresholdsColorEnum) monitoringpb.MonitoringDashboardWidgetXyChartThresholdsColorEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartThresholdsColorEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartThresholdsColorEnum_value["DashboardWidgetXyChartThresholdsColorEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartThresholdsColorEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartThresholdsColorEnum(0)
}

// DashboardWidgetXyChartThresholdsDirectionEnumToProto converts a DashboardWidgetXyChartThresholdsDirectionEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartThresholdsDirectionEnumToProto(e *monitoring.DashboardWidgetXyChartThresholdsDirectionEnum) monitoringpb.MonitoringDashboardWidgetXyChartThresholdsDirectionEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartThresholdsDirectionEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartThresholdsDirectionEnum_value["DashboardWidgetXyChartThresholdsDirectionEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartThresholdsDirectionEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartThresholdsDirectionEnum(0)
}

// DashboardWidgetXyChartXAxisScaleEnumToProto converts a DashboardWidgetXyChartXAxisScaleEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartXAxisScaleEnumToProto(e *monitoring.DashboardWidgetXyChartXAxisScaleEnum) monitoringpb.MonitoringDashboardWidgetXyChartXAxisScaleEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartXAxisScaleEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartXAxisScaleEnum_value["DashboardWidgetXyChartXAxisScaleEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartXAxisScaleEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartXAxisScaleEnum(0)
}

// DashboardWidgetXyChartYAxisScaleEnumToProto converts a DashboardWidgetXyChartYAxisScaleEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartYAxisScaleEnumToProto(e *monitoring.DashboardWidgetXyChartYAxisScaleEnum) monitoringpb.MonitoringDashboardWidgetXyChartYAxisScaleEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartYAxisScaleEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartYAxisScaleEnum_value["DashboardWidgetXyChartYAxisScaleEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartYAxisScaleEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartYAxisScaleEnum(0)
}

// DashboardWidgetXyChartChartOptionsModeEnumToProto converts a DashboardWidgetXyChartChartOptionsModeEnum enum to its proto representation.
func MonitoringDashboardWidgetXyChartChartOptionsModeEnumToProto(e *monitoring.DashboardWidgetXyChartChartOptionsModeEnum) monitoringpb.MonitoringDashboardWidgetXyChartChartOptionsModeEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetXyChartChartOptionsModeEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetXyChartChartOptionsModeEnum_value["DashboardWidgetXyChartChartOptionsModeEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetXyChartChartOptionsModeEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetXyChartChartOptionsModeEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum_value["DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(0)
}

// DashboardWidgetScorecardTimeSeriesQueryApiSourceEnumToProto converts a DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryApiSourceEnumToProto(e *monitoring.DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum) monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryApiSourceEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryApiSourceEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryApiSourceEnum_value["DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryApiSourceEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryApiSourceEnum(0)
}

// DashboardWidgetScorecardSparkChartViewSparkChartTypeEnumToProto converts a DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnumToProto(e *monitoring.DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum) monitoringpb.MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum_value["DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(0)
}

// DashboardWidgetScorecardThresholdsColorEnumToProto converts a DashboardWidgetScorecardThresholdsColorEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardThresholdsColorEnumToProto(e *monitoring.DashboardWidgetScorecardThresholdsColorEnum) monitoringpb.MonitoringDashboardWidgetScorecardThresholdsColorEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardThresholdsColorEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardThresholdsColorEnum_value["DashboardWidgetScorecardThresholdsColorEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardThresholdsColorEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardThresholdsColorEnum(0)
}

// DashboardWidgetScorecardThresholdsDirectionEnumToProto converts a DashboardWidgetScorecardThresholdsDirectionEnum enum to its proto representation.
func MonitoringDashboardWidgetScorecardThresholdsDirectionEnumToProto(e *monitoring.DashboardWidgetScorecardThresholdsDirectionEnum) monitoringpb.MonitoringDashboardWidgetScorecardThresholdsDirectionEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetScorecardThresholdsDirectionEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetScorecardThresholdsDirectionEnum_value["DashboardWidgetScorecardThresholdsDirectionEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetScorecardThresholdsDirectionEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetScorecardThresholdsDirectionEnum(0)
}

// DashboardWidgetTextFormatEnumToProto converts a DashboardWidgetTextFormatEnum enum to its proto representation.
func MonitoringDashboardWidgetTextFormatEnumToProto(e *monitoring.DashboardWidgetTextFormatEnum) monitoringpb.MonitoringDashboardWidgetTextFormatEnum {
	if e == nil {
		return monitoringpb.MonitoringDashboardWidgetTextFormatEnum(0)
	}
	if v, ok := monitoringpb.MonitoringDashboardWidgetTextFormatEnum_value["DashboardWidgetTextFormatEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringDashboardWidgetTextFormatEnum(v)
	}
	return monitoringpb.MonitoringDashboardWidgetTextFormatEnum(0)
}

// DashboardGridLayoutToProto converts a DashboardGridLayout resource to its proto representation.
func MonitoringDashboardGridLayoutToProto(o *monitoring.DashboardGridLayout) *monitoringpb.MonitoringDashboardGridLayout {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardGridLayout{
		Columns: dcl.ValueOrEmptyInt64(o.Columns),
	}
	for _, r := range o.Widgets {
		p.Widgets = append(p.Widgets, MonitoringDashboardWidgetToProto(&r))
	}
	return p
}

// DashboardMosaicLayoutToProto converts a DashboardMosaicLayout resource to its proto representation.
func MonitoringDashboardMosaicLayoutToProto(o *monitoring.DashboardMosaicLayout) *monitoringpb.MonitoringDashboardMosaicLayout {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardMosaicLayout{
		Columns: dcl.ValueOrEmptyInt64(o.Columns),
	}
	for _, r := range o.Tiles {
		p.Tiles = append(p.Tiles, MonitoringDashboardMosaicLayoutTilesToProto(&r))
	}
	return p
}

// DashboardMosaicLayoutTilesToProto converts a DashboardMosaicLayoutTiles resource to its proto representation.
func MonitoringDashboardMosaicLayoutTilesToProto(o *monitoring.DashboardMosaicLayoutTiles) *monitoringpb.MonitoringDashboardMosaicLayoutTiles {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardMosaicLayoutTiles{
		XPos:   dcl.ValueOrEmptyInt64(o.XPos),
		YPos:   dcl.ValueOrEmptyInt64(o.YPos),
		Width:  dcl.ValueOrEmptyInt64(o.Width),
		Height: dcl.ValueOrEmptyInt64(o.Height),
		Widget: MonitoringDashboardWidgetToProto(o.Widget),
	}
	return p
}

// DashboardRowLayoutToProto converts a DashboardRowLayout resource to its proto representation.
func MonitoringDashboardRowLayoutToProto(o *monitoring.DashboardRowLayout) *monitoringpb.MonitoringDashboardRowLayout {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardRowLayout{}
	for _, r := range o.Rows {
		p.Rows = append(p.Rows, MonitoringDashboardRowLayoutRowsToProto(&r))
	}
	return p
}

// DashboardRowLayoutRowsToProto converts a DashboardRowLayoutRows resource to its proto representation.
func MonitoringDashboardRowLayoutRowsToProto(o *monitoring.DashboardRowLayoutRows) *monitoringpb.MonitoringDashboardRowLayoutRows {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardRowLayoutRows{
		Weight: dcl.ValueOrEmptyInt64(o.Weight),
	}
	for _, r := range o.Widgets {
		p.Widgets = append(p.Widgets, MonitoringDashboardWidgetToProto(&r))
	}
	return p
}

// DashboardColumnLayoutToProto converts a DashboardColumnLayout resource to its proto representation.
func MonitoringDashboardColumnLayoutToProto(o *monitoring.DashboardColumnLayout) *monitoringpb.MonitoringDashboardColumnLayout {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardColumnLayout{}
	for _, r := range o.Columns {
		p.Columns = append(p.Columns, MonitoringDashboardColumnLayoutColumnsToProto(&r))
	}
	return p
}

// DashboardColumnLayoutColumnsToProto converts a DashboardColumnLayoutColumns resource to its proto representation.
func MonitoringDashboardColumnLayoutColumnsToProto(o *monitoring.DashboardColumnLayoutColumns) *monitoringpb.MonitoringDashboardColumnLayoutColumns {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardColumnLayoutColumns{
		Weight: dcl.ValueOrEmptyInt64(o.Weight),
	}
	for _, r := range o.Widgets {
		p.Widgets = append(p.Widgets, MonitoringDashboardWidgetToProto(&r))
	}
	return p
}

// DashboardWidgetToProto converts a DashboardWidget resource to its proto representation.
func MonitoringDashboardWidgetToProto(o *monitoring.DashboardWidget) *monitoringpb.MonitoringDashboardWidget {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidget{
		Title:     dcl.ValueOrEmptyString(o.Title),
		XyChart:   MonitoringDashboardWidgetXyChartToProto(o.XyChart),
		Scorecard: MonitoringDashboardWidgetScorecardToProto(o.Scorecard),
		Text:      MonitoringDashboardWidgetTextToProto(o.Text),
		Blank:     MonitoringDashboardWidgetBlankToProto(o.Blank),
	}
	return p
}

// DashboardWidgetXyChartToProto converts a DashboardWidgetXyChart resource to its proto representation.
func MonitoringDashboardWidgetXyChartToProto(o *monitoring.DashboardWidgetXyChart) *monitoringpb.MonitoringDashboardWidgetXyChart {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChart{
		TimeshiftDuration: dcl.ValueOrEmptyString(o.TimeshiftDuration),
		XAxis:             MonitoringDashboardWidgetXyChartXAxisToProto(o.XAxis),
		YAxis:             MonitoringDashboardWidgetXyChartYAxisToProto(o.YAxis),
		ChartOptions:      MonitoringDashboardWidgetXyChartChartOptionsToProto(o.ChartOptions),
	}
	for _, r := range o.DataSets {
		p.DataSets = append(p.DataSets, MonitoringDashboardWidgetXyChartDataSetsToProto(&r))
	}
	for _, r := range o.Thresholds {
		p.Thresholds = append(p.Thresholds, MonitoringDashboardWidgetXyChartThresholdsToProto(&r))
	}
	return p
}

// DashboardWidgetXyChartDataSetsToProto converts a DashboardWidgetXyChartDataSets resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsToProto(o *monitoring.DashboardWidgetXyChartDataSets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSets{
		TimeSeriesQuery:    MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryToProto(o.TimeSeriesQuery),
		PlotType:           MonitoringDashboardWidgetXyChartDataSetsPlotTypeEnumToProto(o.PlotType),
		LegendTemplate:     dcl.ValueOrEmptyString(o.LegendTemplate),
		MinAlignmentPeriod: dcl.ValueOrEmptyString(o.MinAlignmentPeriod),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQuery resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQuery) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQuery {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQuery{
		TimeSeriesFilter:        MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterToProto(o.TimeSeriesFilter),
		TimeSeriesFilterRatio:   MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioToProto(o.TimeSeriesFilterRatio),
		TimeSeriesQueryLanguage: dcl.ValueOrEmptyString(o.TimeSeriesQueryLanguage),
		ApiSource:               MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnumToProto(o.ApiSource),
		UnitOverride:            dcl.ValueOrEmptyString(o.UnitOverride),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{
		Filter:               dcl.ValueOrEmptyString(o.Filter),
		Aggregation:          MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationToProto(o.Aggregation),
		SecondaryAggregation: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto(o.SecondaryAggregation),
		PickTimeSeriesFilter: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto(o.PickTimeSeriesFilter),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{
		AlignmentPeriod:              dcl.ValueOrEmptyString(o.AlignmentPeriod),
		PerSeriesAligner:             MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner),
		CrossSeriesReducer:           MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer),
		ReduceFractionLessThanParams: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParamsToProto(o.ReduceFractionLessThanParams),
		ReduceMakeDistributionParams: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsToProto(o.ReduceMakeDistributionParams),
	}
	for _, r := range o.GroupByFields {
		p.GroupByFields = append(p.GroupByFields, r)
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParamsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParamsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams{
		Threshold: dcl.ValueOrEmptyDouble(o.Threshold),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams{
		BucketOptions:    MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsToProto(o.BucketOptions),
		ExemplarSampling: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o.ExemplarSampling),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o.LinearBuckets),
		ExponentialBuckets: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets),
		ExplicitBuckets:    MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		Width:            dcl.ValueOrEmptyDouble(o.Width),
		Offset:           dcl.ValueOrEmptyDouble(o.Offset),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		GrowthFactor:     dcl.ValueOrEmptyDouble(o.GrowthFactor),
		Scale:            dcl.ValueOrEmptyDouble(o.Scale),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range o.Bounds {
		p.Bounds = append(p.Bounds, r)
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSamplingToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.ValueOrEmptyDouble(o.MinimumValue),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{
		AlignmentPeriod:              dcl.ValueOrEmptyString(o.AlignmentPeriod),
		PerSeriesAligner:             MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner),
		CrossSeriesReducer:           MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer),
		ReduceFractionLessThanParams: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParamsToProto(o.ReduceFractionLessThanParams),
		ReduceMakeDistributionParams: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsToProto(o.ReduceMakeDistributionParams),
	}
	for _, r := range o.GroupByFields {
		p.GroupByFields = append(p.GroupByFields, r)
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParamsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParamsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams{
		Threshold: dcl.ValueOrEmptyDouble(o.Threshold),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams{
		BucketOptions:    MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsToProto(o.BucketOptions),
		ExemplarSampling: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o.ExemplarSampling),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o.LinearBuckets),
		ExponentialBuckets: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets),
		ExplicitBuckets:    MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		Width:            dcl.ValueOrEmptyDouble(o.Width),
		Offset:           dcl.ValueOrEmptyDouble(o.Offset),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		GrowthFactor:     dcl.ValueOrEmptyDouble(o.GrowthFactor),
		Scale:            dcl.ValueOrEmptyDouble(o.Scale),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range o.Bounds {
		p.Bounds = append(p.Bounds, r)
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSamplingToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.ValueOrEmptyDouble(o.MinimumValue),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{
		RankingMethod: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto(o.RankingMethod),
		NumTimeSeries: dcl.ValueOrEmptyInt64(o.NumTimeSeries),
		Direction:     MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto(o.Direction),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{
		Numerator:            MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto(o.Numerator),
		Denominator:          MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto(o.Denominator),
		SecondaryAggregation: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto(o.SecondaryAggregation),
		PickTimeSeriesFilter: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto(o.PickTimeSeriesFilter),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{
		Filter:      dcl.ValueOrEmptyString(o.Filter),
		Aggregation: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto(o.Aggregation),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{
		AlignmentPeriod:              dcl.ValueOrEmptyString(o.AlignmentPeriod),
		PerSeriesAligner:             MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner),
		CrossSeriesReducer:           MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer),
		ReduceFractionLessThanParams: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParamsToProto(o.ReduceFractionLessThanParams),
		ReduceMakeDistributionParams: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsToProto(o.ReduceMakeDistributionParams),
	}
	for _, r := range o.GroupByFields {
		p.GroupByFields = append(p.GroupByFields, r)
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParamsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParamsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams{
		Threshold: dcl.ValueOrEmptyDouble(o.Threshold),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams{
		BucketOptions:    MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsToProto(o.BucketOptions),
		ExemplarSampling: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o.ExemplarSampling),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o.LinearBuckets),
		ExponentialBuckets: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets),
		ExplicitBuckets:    MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		Width:            dcl.ValueOrEmptyDouble(o.Width),
		Offset:           dcl.ValueOrEmptyDouble(o.Offset),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		GrowthFactor:     dcl.ValueOrEmptyDouble(o.GrowthFactor),
		Scale:            dcl.ValueOrEmptyDouble(o.Scale),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range o.Bounds {
		p.Bounds = append(p.Bounds, r)
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSamplingToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.ValueOrEmptyDouble(o.MinimumValue),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{
		Filter:      dcl.ValueOrEmptyString(o.Filter),
		Aggregation: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto(o.Aggregation),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{
		AlignmentPeriod:              dcl.ValueOrEmptyString(o.AlignmentPeriod),
		PerSeriesAligner:             MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner),
		CrossSeriesReducer:           MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer),
		ReduceFractionLessThanParams: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParamsToProto(o.ReduceFractionLessThanParams),
		ReduceMakeDistributionParams: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsToProto(o.ReduceMakeDistributionParams),
	}
	for _, r := range o.GroupByFields {
		p.GroupByFields = append(p.GroupByFields, r)
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParamsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParamsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams{
		Threshold: dcl.ValueOrEmptyDouble(o.Threshold),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams{
		BucketOptions:    MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsToProto(o.BucketOptions),
		ExemplarSampling: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o.ExemplarSampling),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o.LinearBuckets),
		ExponentialBuckets: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets),
		ExplicitBuckets:    MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		Width:            dcl.ValueOrEmptyDouble(o.Width),
		Offset:           dcl.ValueOrEmptyDouble(o.Offset),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		GrowthFactor:     dcl.ValueOrEmptyDouble(o.GrowthFactor),
		Scale:            dcl.ValueOrEmptyDouble(o.Scale),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range o.Bounds {
		p.Bounds = append(p.Bounds, r)
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSamplingToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.ValueOrEmptyDouble(o.MinimumValue),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{
		AlignmentPeriod:              dcl.ValueOrEmptyString(o.AlignmentPeriod),
		PerSeriesAligner:             MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner),
		CrossSeriesReducer:           MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer),
		ReduceFractionLessThanParams: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParamsToProto(o.ReduceFractionLessThanParams),
		ReduceMakeDistributionParams: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsToProto(o.ReduceMakeDistributionParams),
	}
	for _, r := range o.GroupByFields {
		p.GroupByFields = append(p.GroupByFields, r)
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParamsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParamsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams{
		Threshold: dcl.ValueOrEmptyDouble(o.Threshold),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams{
		BucketOptions:    MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsToProto(o.BucketOptions),
		ExemplarSampling: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o.ExemplarSampling),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o.LinearBuckets),
		ExponentialBuckets: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets),
		ExplicitBuckets:    MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		Width:            dcl.ValueOrEmptyDouble(o.Width),
		Offset:           dcl.ValueOrEmptyDouble(o.Offset),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		GrowthFactor:     dcl.ValueOrEmptyDouble(o.GrowthFactor),
		Scale:            dcl.ValueOrEmptyDouble(o.Scale),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range o.Bounds {
		p.Bounds = append(p.Bounds, r)
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSamplingToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.ValueOrEmptyDouble(o.MinimumValue),
	}
	return p
}

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto converts a DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter resource to its proto representation.
func MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto(o *monitoring.DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{
		RankingMethod: MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto(o.RankingMethod),
		NumTimeSeries: dcl.ValueOrEmptyInt64(o.NumTimeSeries),
		Direction:     MonitoringDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto(o.Direction),
	}
	return p
}

// DashboardWidgetXyChartThresholdsToProto converts a DashboardWidgetXyChartThresholds resource to its proto representation.
func MonitoringDashboardWidgetXyChartThresholdsToProto(o *monitoring.DashboardWidgetXyChartThresholds) *monitoringpb.MonitoringDashboardWidgetXyChartThresholds {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartThresholds{
		Label:     dcl.ValueOrEmptyString(o.Label),
		Value:     dcl.ValueOrEmptyDouble(o.Value),
		Color:     MonitoringDashboardWidgetXyChartThresholdsColorEnumToProto(o.Color),
		Direction: MonitoringDashboardWidgetXyChartThresholdsDirectionEnumToProto(o.Direction),
	}
	return p
}

// DashboardWidgetXyChartXAxisToProto converts a DashboardWidgetXyChartXAxis resource to its proto representation.
func MonitoringDashboardWidgetXyChartXAxisToProto(o *monitoring.DashboardWidgetXyChartXAxis) *monitoringpb.MonitoringDashboardWidgetXyChartXAxis {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartXAxis{
		Label: dcl.ValueOrEmptyString(o.Label),
		Scale: MonitoringDashboardWidgetXyChartXAxisScaleEnumToProto(o.Scale),
	}
	return p
}

// DashboardWidgetXyChartYAxisToProto converts a DashboardWidgetXyChartYAxis resource to its proto representation.
func MonitoringDashboardWidgetXyChartYAxisToProto(o *monitoring.DashboardWidgetXyChartYAxis) *monitoringpb.MonitoringDashboardWidgetXyChartYAxis {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartYAxis{
		Label: dcl.ValueOrEmptyString(o.Label),
		Scale: MonitoringDashboardWidgetXyChartYAxisScaleEnumToProto(o.Scale),
	}
	return p
}

// DashboardWidgetXyChartChartOptionsToProto converts a DashboardWidgetXyChartChartOptions resource to its proto representation.
func MonitoringDashboardWidgetXyChartChartOptionsToProto(o *monitoring.DashboardWidgetXyChartChartOptions) *monitoringpb.MonitoringDashboardWidgetXyChartChartOptions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetXyChartChartOptions{
		Mode: MonitoringDashboardWidgetXyChartChartOptionsModeEnumToProto(o.Mode),
	}
	return p
}

// DashboardWidgetScorecardToProto converts a DashboardWidgetScorecard resource to its proto representation.
func MonitoringDashboardWidgetScorecardToProto(o *monitoring.DashboardWidgetScorecard) *monitoringpb.MonitoringDashboardWidgetScorecard {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecard{
		TimeSeriesQuery: MonitoringDashboardWidgetScorecardTimeSeriesQueryToProto(o.TimeSeriesQuery),
		GaugeView:       MonitoringDashboardWidgetScorecardGaugeViewToProto(o.GaugeView),
		SparkChartView:  MonitoringDashboardWidgetScorecardSparkChartViewToProto(o.SparkChartView),
	}
	for _, r := range o.Thresholds {
		p.Thresholds = append(p.Thresholds, MonitoringDashboardWidgetScorecardThresholdsToProto(&r))
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryToProto converts a DashboardWidgetScorecardTimeSeriesQuery resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQuery) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQuery {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQuery{
		TimeSeriesFilter:        MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterToProto(o.TimeSeriesFilter),
		TimeSeriesFilterRatio:   MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioToProto(o.TimeSeriesFilterRatio),
		TimeSeriesQueryLanguage: dcl.ValueOrEmptyString(o.TimeSeriesQueryLanguage),
		ApiSource:               MonitoringDashboardWidgetScorecardTimeSeriesQueryApiSourceEnumToProto(o.ApiSource),
		UnitOverride:            dcl.ValueOrEmptyString(o.UnitOverride),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{
		Filter:               dcl.ValueOrEmptyString(o.Filter),
		Aggregation:          MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationToProto(o.Aggregation),
		SecondaryAggregation: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto(o.SecondaryAggregation),
		PickTimeSeriesFilter: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto(o.PickTimeSeriesFilter),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{
		AlignmentPeriod:              dcl.ValueOrEmptyString(o.AlignmentPeriod),
		PerSeriesAligner:             MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner),
		CrossSeriesReducer:           MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer),
		ReduceFractionLessThanParams: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParamsToProto(o.ReduceFractionLessThanParams),
		ReduceMakeDistributionParams: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsToProto(o.ReduceMakeDistributionParams),
	}
	for _, r := range o.GroupByFields {
		p.GroupByFields = append(p.GroupByFields, r)
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParamsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParamsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams{
		Threshold: dcl.ValueOrEmptyDouble(o.Threshold),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams{
		BucketOptions:    MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsToProto(o.BucketOptions),
		ExemplarSampling: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o.ExemplarSampling),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o.LinearBuckets),
		ExponentialBuckets: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets),
		ExplicitBuckets:    MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		Width:            dcl.ValueOrEmptyDouble(o.Width),
		Offset:           dcl.ValueOrEmptyDouble(o.Offset),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		GrowthFactor:     dcl.ValueOrEmptyDouble(o.GrowthFactor),
		Scale:            dcl.ValueOrEmptyDouble(o.Scale),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range o.Bounds {
		p.Bounds = append(p.Bounds, r)
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSamplingToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.ValueOrEmptyDouble(o.MinimumValue),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{
		AlignmentPeriod:              dcl.ValueOrEmptyString(o.AlignmentPeriod),
		PerSeriesAligner:             MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner),
		CrossSeriesReducer:           MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer),
		ReduceFractionLessThanParams: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParamsToProto(o.ReduceFractionLessThanParams),
		ReduceMakeDistributionParams: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsToProto(o.ReduceMakeDistributionParams),
	}
	for _, r := range o.GroupByFields {
		p.GroupByFields = append(p.GroupByFields, r)
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParamsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParamsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams{
		Threshold: dcl.ValueOrEmptyDouble(o.Threshold),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams{
		BucketOptions:    MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsToProto(o.BucketOptions),
		ExemplarSampling: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o.ExemplarSampling),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o.LinearBuckets),
		ExponentialBuckets: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets),
		ExplicitBuckets:    MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		Width:            dcl.ValueOrEmptyDouble(o.Width),
		Offset:           dcl.ValueOrEmptyDouble(o.Offset),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		GrowthFactor:     dcl.ValueOrEmptyDouble(o.GrowthFactor),
		Scale:            dcl.ValueOrEmptyDouble(o.Scale),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range o.Bounds {
		p.Bounds = append(p.Bounds, r)
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSamplingToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.ValueOrEmptyDouble(o.MinimumValue),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{
		RankingMethod: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumToProto(o.RankingMethod),
		NumTimeSeries: dcl.ValueOrEmptyInt64(o.NumTimeSeries),
		Direction:     MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumToProto(o.Direction),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{
		Numerator:            MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto(o.Numerator),
		Denominator:          MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto(o.Denominator),
		SecondaryAggregation: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto(o.SecondaryAggregation),
		PickTimeSeriesFilter: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto(o.PickTimeSeriesFilter),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{
		Filter:      dcl.ValueOrEmptyString(o.Filter),
		Aggregation: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto(o.Aggregation),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{
		AlignmentPeriod:              dcl.ValueOrEmptyString(o.AlignmentPeriod),
		PerSeriesAligner:             MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner),
		CrossSeriesReducer:           MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer),
		ReduceFractionLessThanParams: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParamsToProto(o.ReduceFractionLessThanParams),
		ReduceMakeDistributionParams: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsToProto(o.ReduceMakeDistributionParams),
	}
	for _, r := range o.GroupByFields {
		p.GroupByFields = append(p.GroupByFields, r)
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParamsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParamsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams{
		Threshold: dcl.ValueOrEmptyDouble(o.Threshold),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams{
		BucketOptions:    MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsToProto(o.BucketOptions),
		ExemplarSampling: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o.ExemplarSampling),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o.LinearBuckets),
		ExponentialBuckets: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets),
		ExplicitBuckets:    MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		Width:            dcl.ValueOrEmptyDouble(o.Width),
		Offset:           dcl.ValueOrEmptyDouble(o.Offset),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		GrowthFactor:     dcl.ValueOrEmptyDouble(o.GrowthFactor),
		Scale:            dcl.ValueOrEmptyDouble(o.Scale),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range o.Bounds {
		p.Bounds = append(p.Bounds, r)
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSamplingToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.ValueOrEmptyDouble(o.MinimumValue),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{
		Filter:      dcl.ValueOrEmptyString(o.Filter),
		Aggregation: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto(o.Aggregation),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{
		AlignmentPeriod:              dcl.ValueOrEmptyString(o.AlignmentPeriod),
		PerSeriesAligner:             MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner),
		CrossSeriesReducer:           MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer),
		ReduceFractionLessThanParams: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParamsToProto(o.ReduceFractionLessThanParams),
		ReduceMakeDistributionParams: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsToProto(o.ReduceMakeDistributionParams),
	}
	for _, r := range o.GroupByFields {
		p.GroupByFields = append(p.GroupByFields, r)
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParamsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParamsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams{
		Threshold: dcl.ValueOrEmptyDouble(o.Threshold),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams{
		BucketOptions:    MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsToProto(o.BucketOptions),
		ExemplarSampling: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o.ExemplarSampling),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o.LinearBuckets),
		ExponentialBuckets: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets),
		ExplicitBuckets:    MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		Width:            dcl.ValueOrEmptyDouble(o.Width),
		Offset:           dcl.ValueOrEmptyDouble(o.Offset),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		GrowthFactor:     dcl.ValueOrEmptyDouble(o.GrowthFactor),
		Scale:            dcl.ValueOrEmptyDouble(o.Scale),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range o.Bounds {
		p.Bounds = append(p.Bounds, r)
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSamplingToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.ValueOrEmptyDouble(o.MinimumValue),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{
		AlignmentPeriod:              dcl.ValueOrEmptyString(o.AlignmentPeriod),
		PerSeriesAligner:             MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumToProto(o.PerSeriesAligner),
		CrossSeriesReducer:           MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumToProto(o.CrossSeriesReducer),
		ReduceFractionLessThanParams: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParamsToProto(o.ReduceFractionLessThanParams),
		ReduceMakeDistributionParams: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsToProto(o.ReduceMakeDistributionParams),
	}
	for _, r := range o.GroupByFields {
		p.GroupByFields = append(p.GroupByFields, r)
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParamsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParamsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams{
		Threshold: dcl.ValueOrEmptyDouble(o.Threshold),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams{
		BucketOptions:    MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsToProto(o.BucketOptions),
		ExemplarSampling: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o.ExemplarSampling),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions{
		LinearBuckets:      MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o.LinearBuckets),
		ExponentialBuckets: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets),
		ExplicitBuckets:    MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBucketsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		Width:            dcl.ValueOrEmptyDouble(o.Width),
		Offset:           dcl.ValueOrEmptyDouble(o.Offset),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBucketsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.ValueOrEmptyInt64(o.NumFiniteBuckets),
		GrowthFactor:     dcl.ValueOrEmptyDouble(o.GrowthFactor),
		Scale:            dcl.ValueOrEmptyDouble(o.Scale),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBucketsToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{}
	for _, r := range o.Bounds {
		p.Bounds = append(p.Bounds, r)
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSamplingToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSamplingToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling{
		MinimumValue: dcl.ValueOrEmptyDouble(o.MinimumValue),
	}
	return p
}

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto converts a DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter resource to its proto representation.
func MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterToProto(o *monitoring.DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) *monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{
		RankingMethod: MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumToProto(o.RankingMethod),
		NumTimeSeries: dcl.ValueOrEmptyInt64(o.NumTimeSeries),
		Direction:     MonitoringDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumToProto(o.Direction),
	}
	return p
}

// DashboardWidgetScorecardGaugeViewToProto converts a DashboardWidgetScorecardGaugeView resource to its proto representation.
func MonitoringDashboardWidgetScorecardGaugeViewToProto(o *monitoring.DashboardWidgetScorecardGaugeView) *monitoringpb.MonitoringDashboardWidgetScorecardGaugeView {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardGaugeView{
		LowerBound: dcl.ValueOrEmptyDouble(o.LowerBound),
		UpperBound: dcl.ValueOrEmptyDouble(o.UpperBound),
	}
	return p
}

// DashboardWidgetScorecardSparkChartViewToProto converts a DashboardWidgetScorecardSparkChartView resource to its proto representation.
func MonitoringDashboardWidgetScorecardSparkChartViewToProto(o *monitoring.DashboardWidgetScorecardSparkChartView) *monitoringpb.MonitoringDashboardWidgetScorecardSparkChartView {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardSparkChartView{
		SparkChartType:     MonitoringDashboardWidgetScorecardSparkChartViewSparkChartTypeEnumToProto(o.SparkChartType),
		MinAlignmentPeriod: dcl.ValueOrEmptyString(o.MinAlignmentPeriod),
	}
	return p
}

// DashboardWidgetScorecardThresholdsToProto converts a DashboardWidgetScorecardThresholds resource to its proto representation.
func MonitoringDashboardWidgetScorecardThresholdsToProto(o *monitoring.DashboardWidgetScorecardThresholds) *monitoringpb.MonitoringDashboardWidgetScorecardThresholds {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetScorecardThresholds{
		Label:     dcl.ValueOrEmptyString(o.Label),
		Value:     dcl.ValueOrEmptyDouble(o.Value),
		Color:     MonitoringDashboardWidgetScorecardThresholdsColorEnumToProto(o.Color),
		Direction: MonitoringDashboardWidgetScorecardThresholdsDirectionEnumToProto(o.Direction),
	}
	return p
}

// DashboardWidgetTextToProto converts a DashboardWidgetText resource to its proto representation.
func MonitoringDashboardWidgetTextToProto(o *monitoring.DashboardWidgetText) *monitoringpb.MonitoringDashboardWidgetText {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetText{
		Content: dcl.ValueOrEmptyString(o.Content),
		Format:  MonitoringDashboardWidgetTextFormatEnumToProto(o.Format),
	}
	return p
}

// DashboardWidgetBlankToProto converts a DashboardWidgetBlank resource to its proto representation.
func MonitoringDashboardWidgetBlankToProto(o *monitoring.DashboardWidgetBlank) *monitoringpb.MonitoringDashboardWidgetBlank {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringDashboardWidgetBlank{}
	return p
}

// DashboardToProto converts a Dashboard resource to its proto representation.
func DashboardToProto(resource *monitoring.Dashboard) *monitoringpb.MonitoringDashboard {
	p := &monitoringpb.MonitoringDashboard{
		Name:         dcl.ValueOrEmptyString(resource.Name),
		DisplayName:  dcl.ValueOrEmptyString(resource.DisplayName),
		GridLayout:   MonitoringDashboardGridLayoutToProto(resource.GridLayout),
		MosaicLayout: MonitoringDashboardMosaicLayoutToProto(resource.MosaicLayout),
		RowLayout:    MonitoringDashboardRowLayoutToProto(resource.RowLayout),
		ColumnLayout: MonitoringDashboardColumnLayoutToProto(resource.ColumnLayout),
		Project:      dcl.ValueOrEmptyString(resource.Project),
		Etag:         dcl.ValueOrEmptyString(resource.Etag),
	}

	return p
}

// ApplyDashboard handles the gRPC request by passing it to the underlying Dashboard Apply() method.
func (s *DashboardServer) applyDashboard(ctx context.Context, c *monitoring.Client, request *monitoringpb.ApplyMonitoringDashboardRequest) (*monitoringpb.MonitoringDashboard, error) {
	p := ProtoToDashboard(request.GetResource())
	res, err := c.ApplyDashboard(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DashboardToProto(res)
	return r, nil
}

// ApplyDashboard handles the gRPC request by passing it to the underlying Dashboard Apply() method.
func (s *DashboardServer) ApplyMonitoringDashboard(ctx context.Context, request *monitoringpb.ApplyMonitoringDashboardRequest) (*monitoringpb.MonitoringDashboard, error) {
	cl, err := createConfigDashboard(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyDashboard(ctx, cl, request)
}

// DeleteDashboard handles the gRPC request by passing it to the underlying Dashboard Delete() method.
func (s *DashboardServer) DeleteMonitoringDashboard(ctx context.Context, request *monitoringpb.DeleteMonitoringDashboardRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDashboard(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDashboard(ctx, ProtoToDashboard(request.GetResource()))

}

// ListDashboard handles the gRPC request by passing it to the underlying DashboardList() method.
func (s *DashboardServer) ListMonitoringDashboard(ctx context.Context, request *monitoringpb.ListMonitoringDashboardRequest) (*monitoringpb.ListMonitoringDashboardResponse, error) {
	cl, err := createConfigDashboard(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDashboard(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*monitoringpb.MonitoringDashboard
	for _, r := range resources.Items {
		rp := DashboardToProto(r)
		protos = append(protos, rp)
	}
	return &monitoringpb.ListMonitoringDashboardResponse{Items: protos}, nil
}

func createConfigDashboard(ctx context.Context, service_account_file string) (*monitoring.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return monitoring.NewClient(conf), nil
}
