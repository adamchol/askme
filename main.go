/*
Copyright © 2024 ADAM CHOLEWIŃSKI
*/
package main

import (
	"github.com/adamchol/askme/cmd"
	"github.com/adamchol/askme/internal/utils"
)

func main() {
	utils.LoadEnvVars()

	cmd.Execute()
}
