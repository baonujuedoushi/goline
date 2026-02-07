package main

import (
	"fmt"

	"github.com/samber/lo"
)

func anyPkg() {
	names := []string{"t1", "b2", "t1"}
	uniqueName := lo.Uniq(names)
	fmt.Println(uniqueName)
}
