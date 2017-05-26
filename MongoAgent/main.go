/*
  +------------------------------------------------------------------------------+
  | SmartAgent - MongoAgent                                                      |
  +------------------------------------------------------------------------------+
  | This source file is subject to GPL 3.0 of the GNU General Public License,    |
  | that is bundled with this package in the file LICENSE, and is available      |
  | through the world-wide-web at the following url:                             |
  | http://www.gnu.org/licenses/gpl-3.0.html                                     |
  | If you did not receive a copy of the GPL3.0 license and are unable to obtain |
  | it through the world-wide-web, please send a note to neeke@php.net           |
  | so we can mail you a copy immediately.                                       |
  +------------------------------------------------------------------------------+
  | Author: Neeke.Gao  <neeke@php.net>  <neeke.gao@cloudwise.com>                |
  +------------------------------------------------------------------------------+
  | Copyright: www.cloudwise.com                                                 |
  +------------------------------------------------------------------------------+
*/
package main

import (
	"github.com/CloudWise-OpenSource/GoCrab-Demo/MongoAgent/enums"
	"github.com/CloudWise-OpenSource/GoCrab-Demo/MongoAgent/tasks"
	"github.com/CloudWise-OpenSource/GoCrab/Cli"
)

func main() {
	GoCrab.SetLogger("file", `{"filename":"logs/error.log"}`)

	GoCrab.CrabApp.Name = enums.PLUGIN_NAME
	GoCrab.CrabApp.Usage = enums.PLUGIN_USAGE
	GoCrab.CrabApp.Version = enums.PLUGIN_VERSION

	GoCrab.CrabApp.Commands = []GoCrab.Command{
		{
			Name:  "start",
			Usage: "start MongoAgent",
			Action: func(c *GoCrab.Context) {
				tasks.TaskInit()
			},
		},
	}

	GoCrab.Run()
}
