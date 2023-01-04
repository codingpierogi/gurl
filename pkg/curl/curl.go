package curl

import (
	"github.com/codingpierogi/gurl/pkg/fetch"
	"github.com/codingpierogi/gurl/pkg/print"
)

func Run(args []string, outputs []string, options fetch.Options) {
	for i, arg := range args {
		result := <-fetch.Body(arg, options)
		if i < len(outputs) {
			print.Body(outputs[i], result.Value)
		} else {
			print.Body("", result.Value)
		}
	}
}
