package main

import (
	"GinSample/v4GetParam"
)

func main() {

	// r1 := v1HelloGin.SetupRouter()
	// r2 := v2Router.SetupRouter()
	// r1.Run(":8088")
	// r2.Run(":8089")

	// r3 := v3JsonResp.SetupRouter()
	// r3.Run(":8080")

	r4 := v4GetParam.SetupRouter()

	r4.Run()

}