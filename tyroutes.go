package tyroutes

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

var Analyzer = &analysis.Analyzer{
	Name:     "tyroutes",
	Doc:      "find routes used and defined by a typhon service",
	Run:      run,
	Requires: []*analysis.Analyzer{buildssa.Analyzer},
}

var service string

func init() {
	Analyzer.Flags.StringVar(&service, "service", service, "path to the root of the service")
}

func run(pass *analysis.Pass) (interface{}, error) {
	pkgssa := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	for _, fn := range pkgssa.SrcFuncs {
		findTyphonRoute(pass, fn)
	}
	return nil, nil
}

func findTyphonRoute(pass *analysis.Pass, fn *ssa.Function) {
}
