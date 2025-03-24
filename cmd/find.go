package cmd

import (
	"duplicate-finder/pkg/dupfinder"
	"github.com/spf13/cobra"
	"log/slog"
)

func init() {
	var (
		file       string
		verbose    bool
		typeOfScan string
	)
	findCMD.Flags().StringVarP(&file, "output", "o", "", "output file name")
	findCMD.Flags().StringVarP(&typeOfScan, "type", "t", "value", "type of scanning of file by value (default) | key")
	findCMD.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output of work")
	rootCMD.AddCommand(findCMD)
}

var findCMD = &cobra.Command{
	Use:   "find",
	Short: "find <arg1> <arg2> <arg3> ...",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			slog.Error("Need one or more args")
			return
		}

		typeOfScan := dupfinder.ScanType(cmd.Flag("type").Value.String())

		finder := &dupfinder.DupFinder{}

		switch typeOfScan {
		case dupfinder.ByValue:
			finder = dupfinder.NewScanDupFinder(args, dupfinder.ByValue)
		case dupfinder.ByKey:
			finder = dupfinder.NewScanDupFinder(args, dupfinder.ByKey)
		default:
			finder = dupfinder.NewScanDupFinder(args, dupfinder.ByValue)

		}

		resOfScan, err := finder.FindDuplicates()
		if err != nil {
			slog.Error("Error of scanning")
		}
		report := finder.Report(resOfScan)
		if err != nil {
			slog.Error("error of prepare output")
		}

		verbose := cmd.Flag("verbose").Value.String()
		if verbose == "true" {
			dupfinder.Verbose(resOfScan, typeOfScan)
		}

		// prepare report
		if v := cmd.Flag("output").Value.String(); v != "" {
			err := report.ReportToJSONFile(v)
			if err != nil {
				slog.Error(err.Error())
			}
		}
	},
}
