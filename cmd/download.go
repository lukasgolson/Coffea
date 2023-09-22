/*
Copyright © 2023 Lukas G. Olson
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"main/internal/application"
	"os"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		filePath, _ := cmd.Flags().GetString("output")
		url, _ := cmd.Flags().GetString("URL")
		auth, _ := cmd.Flags().GetString("Auth")
		count, _ := cmd.Flags().GetInt("count")
		offset, _ := cmd.Flags().GetInt("offset")
		length, _ := cmd.Flags().GetInt("length")
		transform, _ := cmd.Flags().GetBool("transform")

		values, err := application.DownloadDiceValues(url, auth, count, offset)
		if err != nil {
			return err
		}

		if transform {
			values = application.TransformValues(values)
		}

		formattedValues := application.FormatNumbers(values, length)

		// Write the formatted values to a file.

		err = os.WriteFile(filePath, []byte(formattedValues), os.ModePerm)
		if err != nil {
			return err
		}

		fmt.Println("Successfully downloaded", len(values), "values.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	downloadCmd.Flags().StringP("output", "f", "output.txt", "output file")
	downloadCmd.Flags().IntP("count", "n", 32, "number of records to download")
	downloadCmd.Flags().IntP("offset", "o", 0, "offset of records to download")
	downloadCmd.Flags().IntP("length", "l", 16, "length of records per line")
	downloadCmd.Flags().BoolP("transform", "t", false, "transform the data to match the original random number generator output")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
