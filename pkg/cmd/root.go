package cmd

import (
	"github.com/jenkins-x/jx-helpers/v3/pkg/cobras"
	"github.com/jenkins-x/jx-logging/v3/pkg/log"
	"github.com/jenkins-x/jx-verify/pkg/cmd/ingress"
	"github.com/jenkins-x/jx-verify/pkg/cmd/tls"
	"github.com/jenkins-x/jx-verify/pkg/cmd/version"
	"github.com/jenkins-x/jx-verify/pkg/rootcmd"
	"github.com/spf13/cobra"
)

// Main creates the new command
func Main() *cobra.Command {
	cmd := &cobra.Command{
		Use:   rootcmd.TopLevelCommand,
		Short: "commands for verifying Jenkins X environments",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Logger().Errorf(err.Error())
			}
		},
	}
	cmd.AddCommand(cobras.SplitCommand(tls.NewCmdVerifyTLS()))
	cmd.AddCommand(cobras.SplitCommand(version.NewCmdVersion()))
	cmd.AddCommand(cobras.SplitCommand(ingress.CmdVerifyIngress()))

	return cmd
}
