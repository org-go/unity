package cli

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"unity.service/cli/tool"
)

var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "生成表结构",
	Long:  "生成表结构",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New(`--- - or table`)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		Generate(args)
		fmt.Println(`--- ok`)
	},
}

/**
 * generate
 * @Description:
 * @param args
 */
func Generate(args []string) {

	_path_, _ := os.Getwd()


	_ = tool.Generate().SavePath(fmt.Sprintf(`%s/%s`, _path_, `apps/v1/entity/model/`)).
		Dsn("root:root@tcp(127.0.0.1:3306)/unity?charset=utf8").
		RealNameMethod("TableName").
		Table(args[0]).
		EnableJsonTag(true).
		Config(&tool.T2tConfig{
			StructNameToHump: true,
			TagToLower:       true,
			UcFirstOnly:      true,
			SeperatFile:      true,
		}).
		TagKey("gorm").
		TagCallBack(func(columnName string, isPrimaryKey bool) string {
			var tag string
			// if isPrimaryKey {
			// 	tag = "primary_key;"
			// }
			tag += columnName
			return tag
		}).
		Run()

}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func Ddl()  {
	path, _ := os.Getwd()
	path = path + "/apps/v1/entity/ddl/"
	ddls := tool.Generate().Dsn("root:root@tcp(127.0.0.1:3306)/unity?charset=utf8").DDL()
	for table, ddl := range ddls {
		create, _ := os.Create(path + table + ".sql")
		writeString, err := create.WriteString(ddl)
		fmt.Println(writeString, err)
	}
}
