package cli

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"unity.service/apps/v1/entity/v_quake/quake"
	"unity.service/apps/v1/entity/v_quake/ent"
)

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "生成表结构",
	Long:  "生成表结构",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		schema(args)
		fmt.Println(`--- ok`)
	},
}

/**
 * generate
 * @Description:
 * @param args
 */
func schema(args []string) {

	err := quake.NewRepo().Call(&mysql{}).M.Schema.Create(context.TODO())
	fmt.Println(`schema.error`, err)

}

type mysql struct {

}

func (m mysql) Link() (client *ent.Client) {
	open, _ := ent.Open(`mysql`, `root:root@tcp(127.0.0.1:3306)/test?charset=utf8`)
	fmt.Println(open)
	return open
}

func init() {
	rootCmd.AddCommand(schemaCmd)
}
