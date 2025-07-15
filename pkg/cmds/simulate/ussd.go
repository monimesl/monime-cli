package simulate

import (
	"fmt"
	"github.com/monimesl/monime-cli/pkg/resource/account"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

const (
	script = `
export MONIME_CLI_TOKEN='{TOKEN}';
'apps/ussd-simulator/build/bin/Monime 715.app/Contents/MacOS/monime-715'
`
)

var ussdSimulator = &cobra.Command{
	Use:   "ussd",
	Short: "Open the USSD simulator",
	RunE: func(cmd *cobra.Command, args []string) error {
		token, err := account.CheckActiveToken(cmd.Context())
		if err != nil {
			return err
		}
		fmt.Println("✅️ Opening the USSD simulator")
		path := strings.ReplaceAll(script, "{TOKEN}", token)
		c := exec.Command("bash", "-c", path)
		return c.Run()
	},
}
