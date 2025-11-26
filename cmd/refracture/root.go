/*
Copyright © 2025 Davezant 0.2.0
*/
package cmd

import (
	"fmt"
	"os"
	"refracture/internal/core"
	"refracture/pkg/shutils"

	"github.com/spf13/cobra"
)

var (
	srcPath string
	outPath string
)
var copyAddons = false
var rootCmd = &cobra.Command{
	Use:   "refracture",
	Short: "Godot tool for cleaning mess for lazy devs",
	Long: `
	Refactors and organizes Godot project files, cleaning and exporting them to a designated folder.
                                        
        ((((      ((((      (((((       
        (((((((((((((((((((((((        
   ((    (((((((((((((((((((((    ((  
   (((((((((((((((((((((((((((((((((  
   (((((((((((((((((((((((((((((((((  
     (((((    (((((((((((((    ((((   
     (((( %%%%% (((   ((( %%%%% ((((  
     ((((  %%% ((((   (((( %%% ((((   
     ((((((((((((((((((((((((((((((   
         ((((((       ((((((          
     ((((( (((((  ((((( .((((( (((((  
      ((((((((((((((((((((((((((((    
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if srcPath == "" || outPath == "" {
			fmt.Println("Error: both --src and --out flags must be provided")
			cmd.Help()
			os.Exit(1)
		}
		if shutils.HasProjectFile(srcPath) {
			fmt.Println("Starting Refactoring at '" + srcPath + "'")
			pm := core.GetRecursivePaths(srcPath)
			s := core.NewStructure(outPath)
			d := pm.DesignateFiles(s, srcPath)

			s.CreateFolders()

			if err := s.CopyFilesToDesignates(d); err != nil {
				fmt.Println("Error copying files:", err)

			}

			if err := s.CopyRawSpecialFolders(copyAddons, srcPath); err != nil {
				fmt.Println("Error copying special folders:", err)
			}

			reader := core.NewReader(d, srcPath, outPath+"/refractureOut")
			if err := reader.ReplaceInFiles(); err != nil {
				fmt.Println("Error replacing strings:", err)
			}

			fmt.Println("Project successfully refractured!")
		} else {
			fmt.Println("No project.godot found in '" + srcPath + "'")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&srcPath, "src", "s", "", "Path to the Godot project to refactor")
	rootCmd.Flags().StringVarP(&outPath, "out", "o", "", "Path where the refactored project will be exported")
	rootCmd.Flags().BoolVarP(&copyAddons, "addons", "a", false, "Should copy addons folder")

}
