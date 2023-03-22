package init_run

import (
	"fmt"
	"github.com/spf13/cobra"
)

var fundOutRunCmd = &cobra.Command{
	Use:   "fund-init",
	Short: "version subcommand show git version info.",
	Run: func(cmd *cobra.Command, args []string) {
		// 调用远端接口
		fmt.Println("调用成功")
	},
}
