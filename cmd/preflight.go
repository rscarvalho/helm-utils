package cmd

import (
	"fmt"
	"strings"

	"github.com/rscarvalho/helm-utils/pkg/helm"
	"github.com/spf13/cobra"
)

type helmArgs map[string]string

var preflightCmd = &cobra.Command{
	Use:   "preflight",
	Short: "Run a preflight deploy",
	Run:   runCommand,
	Args:  cobra.NoArgs,
}

var (
	releaseName   string
	chartPath     string
	helmExtraArgs helmArgs
)

func (args *helmArgs) String() string {
	out := make([]string, 0)

	for key, value := range *args {
		out = append(out, fmt.Sprintf("%s=%s", key, value))
	}

	return "{" + strings.Join(out, ",") + "}"
}

func (args *helmArgs) Set(value string) error {
	parts := strings.SplitN(value, "=", 2)

	if *args == nil {
		*args = make(map[string]string)
	}

	(*args)[parts[0]] = parts[1]
	return nil
}

func (args *helmArgs) Type() string {
	return "stringMap"
}

func runCommand(cmd *cobra.Command, args []string) {
	fmt.Printf("[PREFLIGH] release=%s, chartPath=%s, helmExtraArgs=%s\n",
		releaseName, chartPath, helmExtraArgs)
	fmt.Println("Command preflight called")

	status, err := helm.GetReleaseStatus(releaseName)

	if err != nil {
		panic(err)
	}
	fmt.Println("Status:", status)
}

func init() {
	RootCmd.AddCommand(preflightCmd)

	preflightCmd.Flags().
		StringVarP(&releaseName, "releaseName", "r", "", "Override the default release name")
	preflightCmd.Flags().
		StringVarP(&chartPath, "chart", "c", "", "path to the helm chart")
	preflightCmd.Flags().
		Var(&helmExtraArgs, "set", "Extra arguments for the helm command")

	preflightCmd.MarkFlagRequired("chart")
}
