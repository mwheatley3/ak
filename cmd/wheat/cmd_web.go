package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

func webCmd() *cobra.Command {
	return &cobra.Command{
		Use: "web",
		Run: func(cmd *cobra.Command, args []string) {
			var (
				// 	c   = loadConfig()
				l = logrus.New()
				// 	db  = db.NewFromConfig(l, c.Postgres)
				// srv = web.Server{}
			)

			// if err := srv.HTTPServer.ListenAndServe(); err != nil {
			// 	l.Fatalf("server listen error: %s", err)
			// }
			fmt.Printf("^^^^^^^^^^^^^^^^^^^^")
		},
	}
}
