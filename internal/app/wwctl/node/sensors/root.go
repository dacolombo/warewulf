package sensors

import (
	"github.com/spf13/cobra"
	"github.com/warewulf/warewulf/internal/pkg/node"
)

type variables struct {
	Showcmd bool
	Full    bool
	Fanout  int
}

func GetCommand() *cobra.Command {
	vars := variables{}

	powerCmd := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "sensors [OPTIONS] PATTERN",
		Short:                 "Show node IPMI sensor information",
		Long:                  "Show IPMI sensor information for nodes matching PATTERN.",
		Args:                  cobra.MinimumNArgs(1),
		RunE:                  CobraRunE(&vars),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) != 0 {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}

			nodeDB, _ := node.New()
			nodes, _ := nodeDB.FindAllNodes()
			var node_names []string
			for _, node := range nodes {
				node_names = append(node_names, node.Id())
			}
			return node_names, cobra.ShellCompDirectiveNoFileComp
		},
	}
	powerCmd.PersistentFlags().BoolVarP(&vars.Full, "full", "F", false, "show detailed output.")
	powerCmd.PersistentFlags().BoolVarP(&vars.Showcmd, "show", "s", false, "only show command which will be executed")
	powerCmd.PersistentFlags().IntVar(&vars.Fanout, "fanout", 50, "how many command should be executed in parallel")
	return powerCmd
}
