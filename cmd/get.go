// Copyright © 2018 The Pingaling Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	pl "github.com/spf13/pingaling/pkg/pingaling"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Display one or many resources",
	Long: `Prints a table of the most important information about the specified
resources.`,
	Example: `
 # List all health status
 pingaling get health

 # List all incidents
 pingaling get incidents
`,
}

// healthCmd represents the health command
var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "List the health status summary",
	Example: `
 # list health status
 pingaling get health
	`,
	Run: func(cmd *cobra.Command, args []string) {
		h, err := session.GetHealthStatus()
		checkError(err)
		pl.TableHealth(h.Data)
	},
}

// endpointCmd represents the endpoint command
var endpointCmd = &cobra.Command{
	Use:   "endpoint",
	Short: "List a specified endpoint",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires one endpoint resource")
		}
		return nil
	},
	Example: `
  # Describe a endpoint
  pingaling get endpoint foo-bar`,
	Run: func(cmd *cobra.Command, args []string) {

		ep, err := session.GetEndpoints(args[0])
		checkError(err)
		pl.TableEndpoints(ep.Data)
	},
}

// endpointsCmd represents the endpoints command
var endpointsCmd = &cobra.Command{
	Use:   "endpoints",
	Short: "List a health summary of all endpoints",
	Example: `
  # List all endpoints
	pingaling get endpoints
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Return health summary for now
		h, err := session.GetHealthStatus()
		checkError(err)
		pl.TableHealth(h.Data)
	},
}

// incidentsCmd represents the incidents command
var incidentsCmd = &cobra.Command{
	Use:   "incidents",
	Short: "List all incidents",
	Example: `
  # List all incidents
  pingaling get incidents`,

	Run: func(cmd *cobra.Command, args []string) {
		i, err := session.GetIncidents()
		checkError(err)
		pl.TableIncidents(i.Data)
	},
}

// notificationChannelsCmd represents the notification channel command
var notificationChannelsCmd = &cobra.Command{
	Use:     "notification-channels",
	Short:   "List all notification channels",
	Aliases: []string{"nc"},
	Example: `
  # List all notification channels
  pingaling get notification-channels
  pingaling get nc
 	`,

	Run: func(cmd *cobra.Command, args []string) {
		nc, err := session.GetNotificationChannels()
		checkError(err)
		pl.TableNotificationChannels(nc.Data)
	},
}

// notificationPoliciesCmd represents the notification policies command
var notificationPoliciesCmd = &cobra.Command{
	Use:     "notification-policies",
	Short:   "List all notification policies",
	Aliases: []string{"np"},
	Example: `
  # List all notification policies
  pingaling get notification-policies
  pingaling get np
 	`,

	Run: func(cmd *cobra.Command, args []string) {
		nc, err := session.GetNotificationPolicies()
		checkError(err)
		pl.TableNotificationPolicies(nc.Data)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(healthCmd)
	getCmd.AddCommand(endpointCmd)
	getCmd.AddCommand(endpointsCmd)
	getCmd.AddCommand(incidentsCmd)
	getCmd.AddCommand(notificationChannelsCmd)
	getCmd.AddCommand(notificationPoliciesCmd)
}
