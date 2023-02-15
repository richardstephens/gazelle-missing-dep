package main

import "fmt"
import ocp "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"

func main() {
	fmt.Println(ocp.Span_SPAN_KIND_UNSPECIFIED)
}
