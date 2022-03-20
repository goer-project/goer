package make

import (
	"fmt"
	"os"

	"github.com/goer-project/goer/config"
	"github.com/spf13/cobra"
)

var CmdMakeCMD = &cobra.Command{
	Use:   "cmd",
	Short: "Create a new command",
	Run:   runMakeCMD,
	Args:  cobra.ExactArgs(1),
}

func runMakeCMD(cmd *cobra.Command, args []string) {

	model := makeModelFromString(args[0])
	if model.Directory == "" {
		model.PackageName = "cmd"
	}

	dir := fmt.Sprintf("%s/%s", config.NewDir.Cmd, model.Directory)

	// mkdir -p, 0777
	_ = os.MkdirAll(dir, os.ModePerm)

	createFileFromStub(dir+model.VariableName+".go", "cmd", model)
}
