package generator

import (
	"fmt"
	"os"
	"path"
	"strings"
)

const contextsTemplate = `
const (
	{First} = iota
	//NEXT1
)

func CreateContexts() Contexts {
	return SetContexts(
		//NEXT2
	)
}`

const getContextTemplate = `
func (c *Contexts) {name}() EntityBase {
	return (*c)[{name}].(EntityBase)
}`

func InitContext(context string) {
	file, _ := os.Create(GetPath("Contexts.go"))
	defer file.Close()

	getter := strings.Replace(getContextTemplate, "{name}", context, -1)
	body := strings.Replace(contextsTemplate, "//NEXT2", "CreateEntityBase("+context+"ComponentTotal),\n//NEXT2", -1)
	body = strings.Replace(body, "{First}", context, -1)
	contextData := GetHeader() + body + getter

	file.WriteString(contextData)

}
func GetHeader() string {
	return fmt.Sprintf(header, PackageName)
}
func GetDir() string {
	return "./entitas"
}
func GetPath(fileName string) string {
	return path.Join(GetDir(), "gen_"+fileName)
}
