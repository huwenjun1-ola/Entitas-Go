package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"

	"github.com/Falldot/Entitas-Go/generator"
)

func main() {
	fset := token.NewFileSet()
	//inFileName := os.Getenv("GOFILE")
	inFileName := "components.go"
	//inFileName := "entitas/gen_Entitas.go"
	src, err := ioutil.ReadFile(inFileName)
	if err != nil {
		panic(err)
	}

	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	//generator.PackageName = f.Name.Name
	components := generator.FindComponents(f)

	generator.CreateEntitasLibFile()

	generator.InitContext(inFileName[:len(inFileName)-3])
	//常量单独写文件
	constText := generator.CreateEntitasContextFile(inFileName, components, src)
	genConstFile := fmt.Sprintf(generator.GetPath("const_" + inFileName))

	os.WriteFile(genConstFile, []byte(constText), os.ModePerm)
}
