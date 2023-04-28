package main

import (
	"nft_platform/core"
	"nft_platform/initialization"
)

func main() {
	initialization.New()
	core.Run()
}
