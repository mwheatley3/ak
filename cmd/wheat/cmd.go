package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "wheat",
	}

	cmd.AddCommand(
		webCmd(),
		// createUserCmd(),
		// internal.PsqlCmd(pgLoadConfig),
	)

	if err := cmd.Execute(); err != nil {
		fmt.Printf("Run error: %s\n", err)
		os.Exit(1)
	}
}
