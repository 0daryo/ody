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
)

func Interprete() {
	flag.Parse()
	// ファイルをOpenする
	f, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)
	// 出力
	l := lexer.New(string(b))
	p := parser.New(l)

	program := p.ParseProgram()
	out := os.Stdout
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
	}
	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
