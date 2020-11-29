/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/0daryo/ody/evaluator"
	"github.com/0daryo/ody/lexer"
	"github.com/0daryo/ody/object"
	"github.com/0daryo/ody/parser"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		// ファイルをOpenする
		f, err := os.Open(fmt.Sprintf("%s.od", flag.Arg(1)))
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		b, _ := ioutil.ReadAll(f)
		l := lexer.New(string(b))
		p := parser.New(l)

		program := p.ParseProgram()
		out := os.Stdout
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
		}
		env := object.NewEnvironment()
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil && evaluated.Inspect() != "null" {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
