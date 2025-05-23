package main

import (
	"github.com/spf13/cobra"
	"log"
)

var params struct {
	TargetPath  string
	SamplePaths []string
	Confirm     bool
}

func InitCommand() error {
	rootCmd := &cobra.Command{
		Use:   "duplicate checker",
		Short: "a tool used to check and delete duplicate files in different folders.",
		Run:   rootFunc,
	}
	rootCmd.Flags().StringP("target", "t", "", "target folder path")
	rootCmd.Flags().StringArrayP("sample", "s", []string{}, "sample folder path (multiple)")
	rootCmd.Flags().BoolP("yes", "y", false, "skip confirmation")

	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}

func rootFunc(cmd *cobra.Command, args []string) {
	// get params
	targetPath, err := cmd.Flags().GetString("target")
	if err != nil {
		log.Fatal(err)
	}
	samplePaths, err := cmd.Flags().GetStringArray("sample")
	if err != nil {
		log.Fatal(err)
	}
	skipConfirm, err := cmd.Flags().GetBool("yes")
	if err != nil {
		log.Fatal(err)
	}

	if targetPath == "" || len(samplePaths) == 0 {
		log.Fatal("target folder path or sample folder path are required")
	}

	params.TargetPath = targetPath
	params.SamplePaths = samplePaths
	params.Confirm = skipConfirm

}
