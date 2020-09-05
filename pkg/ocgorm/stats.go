package ocgorm

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// Measures
var (
	QueryCount = stats.Int64("opencensus.io/ocgorm/query_count", "Number of queries started", stats.UnitDimensionless)
)

// Tags applied to measures
var (
	// Operation is the type of query (SELECT, INSERT, UPDATE, DELETE)
	Operation, _ = tag.NewKey("ocgorm.operation")

	// Table name of the target database table
	Table, _ = tag.NewKey("ocgorm.table")
)

var (
	QueryCountView = &view.View{
		Name:        "opencensus.io/ocgorm/query_count",
		Description: "Count of queries started",
		TagKeys:     []tag.Key{Operation, Table},
		Measure:     QueryCount,
		Aggregation: view.Count(),
	}
)
