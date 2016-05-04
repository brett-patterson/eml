package main

import "github.com/brett-patterson/eml/eml"
import "gopkg.in/alecthomas/kingpin.v2"

var (
	namespace = kingpin.Flag("namespace", "Optional namespace to include").Short('n').Default("").String()
	envFile   = kingpin.Arg("env", "The file to read the environment from").Required().String()
	command   = eml.CmdList(kingpin.Arg("command", "Command to run"))
)

func main() {
	kingpin.Parse()

	args := *command
	if len(args) > 0 {
		env, err := eml.LoadEnv(*envFile, *namespace)
		if err != nil {
			println(err.Error())
			return
		}

		err = eml.RunInEnv(args[0], args[1:], env)
		if err != nil {
			println(err.Error())
			return
		}
	}
}
