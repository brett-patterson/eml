package eml

import "gopkg.in/alecthomas/kingpin.v2"

type cmdList []string

func (c *cmdList) Set(value string) error {
	*c = append(*c, value)
	return nil
}

func (c *cmdList) String() string {
	return ""
}

func (c *cmdList) IsCumulative() bool {
	return true
}

// CmdList builds a cmdList flag from a kingpin constructor
func CmdList(s kingpin.Settings) *[]string {
	target := new([]string)
	s.SetValue((*cmdList)(target))
	return target
}
