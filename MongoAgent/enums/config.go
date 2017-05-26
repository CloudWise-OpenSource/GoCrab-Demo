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
package enums

import (
	"github.com/CloudWise-OpenSource/GoCrab/Cli"
	"strconv"
)

func GetHostKey() string {
	if HostKey != "" {
		return HostKey
	}

	if hostKey := GoCrab.AppConfig.String("HostKey"); hostKey != "" {
		HostKey = hostKey
	}

	return HostKey
}

func GetFrequency() int {
	if Frequency != 0 {
		return Frequency
	}

	if frequency, err := GoCrab.AppConfig.Int("Frequency"); err == nil {
		Frequency = frequency
	} else {
		Frequency = DEFAULT_FREQUENCY
	}

	return Frequency
}

func GetHostIp() string {
	if HostIp != "" {
		return HostIp
	}

	if hostip := GoCrab.AppConfig.String("HostIp"); hostip != "" {
		HostIp = hostip
	} else {
		HostIp = DEFAULT_IP
	}

	return HostIp
}

func GetPort() string {
	if Port != "" {
		return Port
	}

	if port, err := GoCrab.AppConfig.Int("Port"); err == nil {
		Port = strconv.Itoa(port)
	} else {
		Port = DEFAULT_PORT
	}

	return Port
}

func GetSendProxy() string {
	if SendProxy != "" {
		return SendProxy
	}

	if sendProxy := GoCrab.AppConfig.String("SendProxy"); sendProxy != "" {
		SendProxy = sendProxy + "/" + SEND_PROXY_SEND_API
	} else {
		SendProxy = SEND_PROXY_DEFAULT + "/" + SEND_PROXY_SEND_API
	}

	return SendProxy
}

func GetQualifier() string {
	if Qualifier != "" {
		return Qualifier
	}

	if qualifier := GoCrab.AppConfig.String("Qualifier"); qualifier != "" {
		Qualifier = qualifier
	} else {
		Qualifier = SERVICE_TYPE + SERVICE_QUALIFIER_SLICE + GetPort()
	}

	return Qualifier
}
