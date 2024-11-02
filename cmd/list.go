package cmd

import (
	"doem/internal/db"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todos",
	Long:  "List todos",
	RunE: func(cmd *cobra.Command, args []string) error {
		query := `
            SELECT
                id,
                name
            FROM todo
            WHERE is_completed = 0
            ORDER BY created_at ASC
        `
		rows, err := db.GetConnection().Query(query)
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var name string

			err := rows.Scan(&id, &name)
			if err != nil {
				return err
			}

			fmt.Printf("%d\t%s\n", id, name)
		}
		return nil
	},
}
