//go:build baremetal
// +build baremetal

package main

func getFullRefreshes() uint {
	return display.FullRefreshes()
}
