package tutorial30daysofcode

import (
	"fmt"
	"sort"
	"strings"

	"1d1go/utils/str"
)

func filter(arr [][]string) {
	var output []string
	for _, ss := range arr {
		firstName, emailID := ss[0], ss[1]
		//noinspection SpellCheckingInspection
		if strings.HasSuffix(emailID, "@gmail.com") {
			output = append(output, firstName)
		}
	}

	sort.Sort(str.SortAdapter(output))
	for _, s := range output {
		fmt.Println(s)
	}
}
