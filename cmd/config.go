/*
Copyright © 2022 James Kokou GAGLO <freemanpolys@gmail.com>

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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var registry string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure your local docker repository",
	Long: `Provide your local repository.
	All the image pull from docker will tagged and push to the provided repository.
	Example:
	syncdock config --registry mylocal.repo
	`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("registry", registry)
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Fprintln(os.Stderr, "Local docker resistry '", registry, "' is configured")
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	configCmd.Flags().StringVarP(&registry, "registry", "r", "", "Your local docker registry")
	configCmd.MarkFlagRequired("registry")
}
