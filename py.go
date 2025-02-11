package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
	"text/template"
)

func CommandPyFunc(cmd *cobra.Command, args []string) {
	if funcTmpl.ProblemName == "" {
		log.Println("You need to at least provide a problem name")
		cmd.Help()
		os.Exit(1)
	}

	Tmpls.FuncTmpl = &funcTmpl
	createProblemDir(funcTmpl.ProblemName)
}

func CommandPyClass(cmd *cobra.Command, args []string) {
	if classTmpl.ProblemName == "" {
		log.Println("You need to at least provide a problem name")
		cmd.Help()
		os.Exit(1)
	}

	Tmpls.ClassTmpl = &classTmpl
	createProblemDir(classTmpl.ProblemName)
}

func createProblemDir(assignment_name string) {
	name_parts := strings.Split(assignment_name, "-")
	name := name_parts[len(name_parts)-1]
	if Tmpls.FuncTmpl != nil {
		Tmpls.FuncTmpl.ParsedName = name
	} else if Tmpls.ClassTmpl != nil {
		Tmpls.ClassTmpl.ParsedName = name
	}

	py_file := fmt.Sprintf("%s.py", name)
	starter_file := fmt.Sprintf("_starter/%s", py_file)
	test_file := fmt.Sprintf("tests/test_%s", py_file)
	dirs := []string{"_starter", "doc", "tests"}
	files := []string{"doc/doc.md", "problem.cfg", starter_file, test_file, py_file}

	for i := range dirs {
		dirs[i] = fmt.Sprintf("%s/%s", assignment_name, dirs[i])
	}

	for i := range files {
		files[i] = fmt.Sprintf("%s/%s", assignment_name, files[i])
	}

	os.Mkdir(assignment_name, 0755)
	for i := range dirs {
		os.Mkdir(dirs[i], 0755)
	}
	for i := range files {
		os.Create(files[i])
	}

	err := buildTemplate(assignment_name, test_file)
	if err != nil {
		os.Exit(1)
	}

	problemCfg := fmt.Sprintf("%s/problem.cfg", assignment_name)
	fobj, err := os.OpenFile(problemCfg, os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening problem.cfg for writing: ", err)
		fobj.Close()
		os.Exit(1)
	}
	str := fmt.Sprintf("[problem]\ntype = python3unittest\nunique = %s\nnote = \ntag = \n", assignment_name)
	_, err = fobj.WriteString(str)
	if err != nil {
		log.Println("error writing to problem.cfg: ", err)
		fobj.Close()
		os.Exit(1)
	}
	fobj.Close()
}

func buildTemplate(dir string, test_name string) error {
	home, _ := os.UserHomeDir()
	var tmpl_path string
	if Tmpls.FuncTmpl != nil {
		tmpl_path = fmt.Sprintf("%s/.config/scaffold/test_template_function.py", home)
		log.Println("Using function template")
	} else if Tmpls.ClassTmpl != nil {
		log.Println("Using class template")
		tmpl_path = fmt.Sprintf("%s/.config/scaffold/test_template_class.py", home)
	}
	fobj, err := os.ReadFile(tmpl_path)
	if err != nil {
		log.Println("Error opening template: ", err)
		return err
	}
	tmpl, _ := template.New("test").Parse(string(fobj))

	test_file := fmt.Sprintf("%s/%s", dir, test_name)
	outputFile, err := os.Create(test_file)
	if err != nil {
		log.Println("Error creating test output: ", err)
		outputFile.Close()
		return err
	}
	defer outputFile.Close()

	var tmplData map[string]string
	if Tmpls.FuncTmpl != nil {
		tmplData = map[string]string{
			"FileName":       Tmpls.FuncTmpl.ParsedName,
			"TestClassName":  Tmpls.FuncTmpl.TestName,
			"ParamTypeHint":  Tmpls.FuncTmpl.ParamType,
			"ReturnTypeHint": Tmpls.FuncTmpl.ReturnType,
			"FuncToTest":     Tmpls.FuncTmpl.FuncToTest,
		}
	} else if Tmpls.ClassTmpl != nil {
		tmplData = map[string]string{
			"FileName":      Tmpls.ClassTmpl.ParsedName,
			"TestClassName": Tmpls.ClassTmpl.TestName,
			"StudentClass":  Tmpls.ClassTmpl.ClassName,
		}
	}

	_ = tmpl.Execute(outputFile, tmplData)
	return nil
}
