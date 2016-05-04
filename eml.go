package main

import "github.com/brett-patterson/eml/eml"
import "gopkg.in/alecthomas/kingpin.v2"

var (
	namespace = kingpin.Flag("namespace", "Optional namespace to include").Short('n').Default("").String()
	command   = eml.CmdList(kingpin.Arg("command", "Command to run"))
)

func main() {
	kingpin.Parse()

	args := *command
	env := eml.LoadEnv(".env.yml", *namespace)
	eml.RunInEnv(args[0], args[1:], env)
}
