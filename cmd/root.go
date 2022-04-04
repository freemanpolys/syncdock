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

	execute "github.com/alexellis/go-execute/pkg/v1"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var image string
var imageFullUrl string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "syncdock",
	Short: "Push docker image to custom registry",
	Long: `
	This cli pull an image from docker and push it to a local docker registry
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		localRegistry := viper.GetString("registry")
		if localRegistry == "" {
			fmt.Fprintln(os.Stderr, "Your local docker registry is not configured. Run 'syncdock config -r your-local.repo' to configure it.")
			os.Exit(0)
		} else {
			localImage := localRegistry + "/" + image
			fmt.Fprintln(os.Stderr, "Your image '", image, "' will be pushed to '", localImage, "'")
			// Docker pull image
			if imageFullUrl != "" {
				image = imageFullUrl
			}
			cmd := execute.ExecTask{
				Command:     "docker",
				Args:        []string{"pull", image},
				StreamStdio: true,
			}
			_, err := cmd.Execute()
			if err != nil {
				panic(err)
			}

			// Docker tag image
			cmd = execute.ExecTask{
				Command:     "docker",
				Args:        []string{"tag", image, localImage},
				StreamStdio: true,
			}
			_, err = cmd.Execute()
			if err != nil {
				panic(err)
			}

			// Docker  image pushed to local registry
			cmd = execute.ExecTask{
				Command:     "docker",
				Args:        []string{"push", localImage},
				StreamStdio: true,
			}
			_, err = cmd.Execute()
			if err != nil {
				panic(err)
			}

		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.syncdock.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&image, "image-name", "i", "", "[Required] The image name and tag to push to local registry")
	rootCmd.MarkFlagRequired("image-name")
	rootCmd.Flags().StringVarP(&imageFullUrl, "image-full-name", "f", "", "[Optional] The full url image to push to local registry")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".syncdock" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".syncdock")
		configFile := home + "/.syncdock.yaml"
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			_, err = os.Create(configFile) // Create your file
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error when creation config file:", err)
			}
		}

	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
