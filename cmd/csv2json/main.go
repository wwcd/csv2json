package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wwcd/csv2json"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "csv2json CSVFILE",
		Short: "convert csv to json",
		RunE: func(cmd *cobra.Command, args []string) error {
			fromRow, _ := cmd.Flags().GetInt("from-row")
			toRow, _ := cmd.Flags().GetInt("to-row")
			fromCol, _ := cmd.Flags().GetInt("from-col")
			toCol, _ := cmd.Flags().GetInt("to-col")
			header, _ := cmd.Flags().GetInt("header")

			input, err := os.Open(args[0])
			if err != nil {
				return err
			}

			output := &bytes.Buffer{}
			err = csv2json.Conv(input, output, csv2json.WithHeader(header), csv2json.WithRow(fromRow, toRow), csv2json.WithCol(fromCol, toCol))
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", output.String())

			return nil
		},
		Args: cobra.ExactArgs(1),
	}

	rootCmd.Flags().Int("from-row", 0, "start row")
	rootCmd.Flags().Int("to-row", 0x7fffffff, "end row")
	rootCmd.Flags().Int("from-col", 0, "start col")
	rootCmd.Flags().Int("to-col", 0x7fffffff, "end col")
	rootCmd.Flags().Int("header", 0, "end col")

	rootCmd.Execute()
}
