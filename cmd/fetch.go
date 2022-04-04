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

	execute "github.com/alexellis/go-execute/pkg/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var imageFromLocal string
var taggedImage string

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch docker image from configured local repo and tag in the local repo",
	Long:  `Fetch docker image from configured local repo and tag in the local repo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fetch called")
		localRegistry := viper.GetString("registry")
		if localRegistry == "" {
			fmt.Fprintln(os.Stderr, "Your local docker registry is not configured. Run 'syncdock config -r your-local.repo' to configure it.")
			os.Exit(0)
		} else {
			localImage := localRegistry + "/" + imageFromLocal
			// Final tagged image
			if taggedImage == "" {
				taggedImage = imageFromLocal
			}
			fmt.Fprintln(os.Stderr, "Image '", localImage, "' will be tagged to '", taggedImage, "'")

			// Docker pull local image
			cmd := execute.ExecTask{
				Command:     "docker",
				Args:        []string{"pull", localImage},
				StreamStdio: true,
			}
			_, err := cmd.Execute()
			if err != nil {
				panic(err)
			}

			// Docker tag pulled image
			cmd = execute.ExecTask{
				Command:     "docker",
				Args:        []string{"tag", localImage, taggedImage},
				StreamStdio: true,
			}
			_, err = cmd.Execute()
			if err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	fetchCmd.Flags().StringVarP(&imageFromLocal, "image-name", "i", "", "[Required] The image with tag to fetch from local repo")
	fetchCmd.MarkFlagRequired("image-name")
	fetchCmd.Flags().StringVarP(&taggedImage, "tagged-image", "t", "", "[Optional] The final tagged image.If not set, the final image will be tagged without local repo")
}
