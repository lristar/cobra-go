package init_run

import (
	"github.com/spf13/cobra"
)

var InitRunCmd = &cobra.Command{
	Use:   "init",
	Short: "init run ",
	Long:  `The 'init' run run`,
}

func init() {
	InitRunCmd.AddCommand(fundOutRunCmd)
}
