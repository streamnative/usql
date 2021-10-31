// Package flinksql defines and registers usql's Flink SQL driver.
//
// See: https://github.com/streamnative/streamnative-sql-go
package flinksql

import (
	"context"
	"regexp"

	_ "github.com/streamnative/streamnative-sql-go/flinksql" // DRIVER
	"github.com/xo/usql/drivers"
)

func init() {
	endRE := regexp.MustCompile(`;?\s*$`)
	drivers.Register("flinksql", drivers.Driver{
		AllowMultilineComments: true,
		Process: func(prefix string, sqlstr string) (string, string, bool, error) {
			sqlstr = endRE.ReplaceAllString(sqlstr, "")
			typ, q := drivers.QueryExecType(prefix, sqlstr)
			return typ, sqlstr, q, nil
		},
		Version: func(ctx context.Context, db drivers.DB) (string, error) {
			return "StreamNative SQL Gateway (unknown)", nil
		},
	})
}
