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

const PLUGIN_NAME = "MongoAgent"

const PLUGIN_VERSION = "1.1.1"

const PLUGIN_USAGE = "MongoAgent for SmartAgent"

const ROUTING_KEY = "agentTopic"

const SERVICE_TYPE = "205"

const SERVICE_QUALIFIER_SLICE = "x"

const DEFAULT_IP = "127.0.0.1"

const DEFAULT_PORT = "27017"

const RESP_STATUS_OK = "1"
const RESP_STATUS_FAILD = "0"

const ENUM_CHANNEL_COUNT = 10

const SEND_PROXY_DEFAULT = "http://127.0.0.1:26789"
const SEND_PROXY_SEND_API = "send"

//每个协程间隔 毫秒
const ENUM_MILLISECOND_INTERVAL = 200

//距上次发送时间相隔多久则马上发送 秒
const DEFAULT_FREQUENCY = 60

//心跳频率
const HEATBEAT_FREQUENCY = 30

var (
	HostKey   string
	Frequency int
	HostIp    string
	Port      string
	SendProxy string
	Qualifier string
)
