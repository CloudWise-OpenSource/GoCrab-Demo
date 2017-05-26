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
package models

import (
	"github.com/CloudWise-OpenSource/GoCrab-Demo/MongoAgent/enums"
	"time"
)

var (
	Status status
)

type status struct {
	AppName            string
	Version            string
	States             string
	ChannelCount       int
	StartTime          int64
	Duration           int64
	GetDataCount       int64
	SendDataCount      int64
	SendDataErrorCount int64
}

func init() {
	appName := enums.PLUGIN_NAME
	version := enums.PLUGIN_VERSION
	states := "ok"

	//startTime := strconv.FormatInt(time.Now().Unix(), 10)
	startTime := time.Now().Unix()
	channelCount := enums.ENUM_CHANNEL_COUNT

	Status = status{appName, version, states, channelCount, startTime, 0, 0, 0, 0}
}

func GetStatus() status {
	Status.Duration = time.Now().Unix() - Status.StartTime
	return Status
}
