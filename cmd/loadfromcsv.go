package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

var fileName string

// loadfromcsvCmd represents the loadfromcsv command
var loadfromcsvCmd = &cobra.Command{
	Use:   "loadfromcsv",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runLoadFromCSVCommand(fileName)
	},
}

func init() {
	rootCmd.AddCommand(loadfromcsvCmd)

	loadfromcsvCmd.Flags().StringVarP(&fileName, "fileName", "f", "", "File Name of CSV File")
}

func runLoadFromCSVCommand(fileName string) {
	records := readCsvFile(fileName)

	for _, record := range records {

		// fmt.Println(record[0])

		// Load Websites
		err := exec.Command("open", record[0]).Start()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(2000)
	}

}

func readCsvFile(fileName string) [][]string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Unable to read input file "+fileName, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+fileName, err)
	}

	return records
}
