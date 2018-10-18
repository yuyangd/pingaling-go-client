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

	pl "bitbucket.org/pingaling-monitoring/client/pkg/pingaling"
	"github.com/spf13/cobra"
)

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

// getHealthCmd represents the health command
var getHealthCmd = &cobra.Command{
	Use:   "health",
	Short: "List the health status summary",
	Example: `
 # list health status
 pingaling get health
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if h, err := session.GetHealthStatus(); err != nil {
			panic(err)
		} else {
			pl.TableHealth(h.Data)
		}
	},
}

// getEndpointCmd represents the endpoint command
var getEndpointCmd = &cobra.Command{
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

		if ep, err := session.GetEndpoints(args[0]); err != nil {
			panic(err)
		} else {
			pl.TableEndpoints(ep.Data)
		}
	},
}

// getEndpointsCmd represents the endpoints command
var getEndpointsCmd = &cobra.Command{
	Use:   "endpoints",
	Short: "List a health summary of all endpoints",
	Example: `
  # List all endpoints
	pingaling get endpoints
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Return health summary for now
		if h, err := session.GetHealthStatus(); err != nil {
			panic(err)
		} else {
			pl.TableHealth(h.Data)
		}
	},
}

// getIncidentsCmd represents the incidents command
var getIncidentsCmd = &cobra.Command{
	Use:   "incidents",
	Short: "List all incidents",
	Example: `
  # List all incidents
  pingaling get incidents`,

	Run: func(cmd *cobra.Command, args []string) {
		if i, err := session.GetIncidents(); err != nil {
			panic(err)
		} else {
			pl.TableIncidents(i.Data)
		}
	},
}

// getNotificationChannelsCmd represents the notification channel command
var getNotificationChannelsCmd = &cobra.Command{
	Use:     "notification-channels",
	Short:   "List all notification channels",
	Aliases: []string{"nc"},
	Example: `
  # List all notification channels
  pingaling get notification-channels
  pingaling get nc
 	`,

	Run: func(cmd *cobra.Command, args []string) {
		if nc, err := session.GetNotificationChannels(); err != nil {
			panic(err)
		} else {
			pl.TableNotificationChannels(nc.Data)
		}

	},
}

// getNotificationPoliciesCmd represents the notification policies command
var getNotificationPoliciesCmd = &cobra.Command{
	Use:     "notification-policies",
	Short:   "List all notification policies",
	Aliases: []string{"np"},
	Example: `
  # List all notification policies
  pingaling get notification-policies
  pingaling get np
 	`,

	Run: func(cmd *cobra.Command, args []string) {
		if nc, err := session.GetNotificationPolicies(); err != nil {
			panic(err)
		} else {
			pl.TableNotificationPolicies(nc.Data)
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getHealthCmd)
	getCmd.AddCommand(getEndpointCmd)
	getCmd.AddCommand(getEndpointsCmd)
	getCmd.AddCommand(getIncidentsCmd)
	getCmd.AddCommand(getNotificationChannelsCmd)
	getCmd.AddCommand(getNotificationPoliciesCmd)
}
