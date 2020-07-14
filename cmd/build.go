/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build website in public directory",
	// REVIEW - add Long description
	Long: `Longer description here.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildSite()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func buildSite() {
	err := os.Mkdir("public", os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	
}
