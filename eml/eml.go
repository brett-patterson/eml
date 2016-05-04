package eml

import "fmt"
import "io/ioutil"
import "os"
import "os/exec"

import "gopkg.in/yaml.v2"

// LoadEnv loads the environment from the specified YAML file
func LoadEnv(envFilename string, namespace string) []string {
	data, err := ioutil.ReadFile(envFilename)
	if err != nil {
		panic(fmt.Sprintf("Error reading %s", envFilename))
	}

	env := make(map[string]map[string]string)
	err = yaml.Unmarshal(data, &env)
	if err != nil {
		panic(err.Error())
	}

	out := os.Environ()
	if global, ok := env["global"]; ok {
		for k, v := range global {
			out = append(out, fmt.Sprintf("%s=%s", k, v))
		}
	}

	if namespace != "" {
		if ns, ok := env[namespace]; ok {
			for k, v := range ns {
				out = append(out, fmt.Sprintf("%s=%s", k, v))
			}
		}
	}

	return out
}

// RunInEnv runs a command with given arguments and environment
func RunInEnv(command string, args []string, env []string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
