package main

import (
	"github.com/spf13/cobra"
)

type Func struct {
	ProblemName string
	TestName    string
	ParamType   string
	ReturnType  string
	FuncToTest  string
	ParsedName  string
}

type Class struct {
	ProblemName string
	TestName    string
	ClassName   string
	ParsedName  string
}

type Templates struct {
	FuncTmpl  *Func
	ClassTmpl *Class
}

var funcTmpl = Func{}
var classTmpl = Class{}
var Tmpls Templates

func main() {

	cli := &cobra.Command{
		Use:   "scaffold",
		Short: "Command-line interface for templating codegrinder assignments",
	}

	cliPy := &cobra.Command{
		Use:   "py",
		Short: "all Python related options",
	}
	cli.AddCommand(cliPy)

	cliPyFunc := &cobra.Command{
		Use:   "function",
		Short: "function-based unit tests",
		Run:   CommandPyFunc,
	}
	cliPyFunc.Flags().StringVarP(&funcTmpl.ProblemName, "problem", "n", "", "name of the problem directory")
	cliPyFunc.Flags().StringVarP(&funcTmpl.TestName, "test", "t", "", "name of the unittest class")
	cliPyFunc.Flags().StringVarP(&funcTmpl.ParamType, "params", "p", "", "function parameter types")
	cliPyFunc.Flags().StringVarP(&funcTmpl.ReturnType, "return", "r", "", "function return type")
	cliPyFunc.Flags().StringVarP(&funcTmpl.FuncToTest, "function", "f", "", "function to test")
	cliPy.AddCommand(cliPyFunc)

	cliPyClass := &cobra.Command{
		Use:   "class",
		Short: "class-based unit tests",
		Run:   CommandPyClass,
	}
	cliPyClass.Flags().StringVarP(&classTmpl.ProblemName, "problem", "n", "", "name of the problem directory")
	cliPyClass.Flags().StringVarP(&classTmpl.TestName, "test", "t", "", "name of the unittest class")
	cliPyClass.Flags().StringVarP(&classTmpl.ClassName, "class", "c", "", "name of the student's class to test")
	cliPy.AddCommand(cliPyClass)

	cli.Execute()
}
