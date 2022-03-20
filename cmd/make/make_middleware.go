package make

import (
	"fmt"
	"os"

	"github.com/goer-project/goer/config"
	"github.com/spf13/cobra"
)

var CmdMakeMiddleware = &cobra.Command{
	Use:   "middleware",
	Short: "Create a new middleware",
	Run:   runMiddleware,
	Args:  cobra.ExactArgs(1),
}

func runMiddleware(cmd *cobra.Command, args []string) {

	model := makeModelFromString(args[0])

	dir := fmt.Sprintf("%s/", config.NewDir.Middleware)

	// mkdir -p, 0777
	_ = os.MkdirAll(dir, os.ModePerm)

	createFileFromStub(dir+model.PackageName+".go", "middleware", model)
}
