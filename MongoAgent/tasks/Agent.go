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
package tasks

import (
	"encoding/json"
	"github.com/CloudWise-OpenSource/GoCrab-Demo/MongoAgent/enums"
	"github.com/CloudWise-OpenSource/GoCrab-Demo/MongoAgent/models"
	"github.com/CloudWise-OpenSource/GoCrab/Cli"
	"github.com/CloudWise-OpenSource/GoCrab/Core/channel"
	"github.com/CloudWise-OpenSource/GoCrab/Libs/httplib"
	"time"
)

var (
	ChannelSharded channel.Sharded
	SendUrl        string
	LastSendTime   int64
)

func init() {
	LastSendTime = time.Now().Unix()
}

func writer() bool {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			GoCrab.Error(panicErr)
		}
	}()

	timeSpace := time.Now().Unix() - LastSendTime

	GoCrab.Debug(enums.PLUGIN_NAME, " tasks writer called")

	if int(timeSpace) >= enums.GetFrequency() {
		valueMap, err := models.GetStats()
		if err != nil {
			GoCrab.Debug(enums.PLUGIN_NAME, " tasks err", err)
			return false
		}

		valueMap["resp_status"] = enums.RESP_STATUS_OK
		valueMap["service_qualifier"] = enums.GetQualifier()
		valueMap["service_type"] = enums.SERVICE_TYPE

		valueStr, err := json.Marshal(valueMap)
		if err != nil {
			GoCrab.Debug("json.Marshal error", err, valueMap)
			panic(err.Error())
		}

		dataMap := make(map[string]string)
		dataMap["agentTopic"] = enums.GetHostKey() + "_" + enums.SERVICE_TYPE + "_" + enums.GetQualifier()
		dataMap["value"] = string(valueStr)

		dataStr, err := json.Marshal(dataMap)
		if err != nil {
			GoCrab.Debug("json.Marshal error", err, dataMap)
			panic(err.Error())
		}

		sendMap := make(map[string]string)
		sendMap["RoutingKey"] = enums.ROUTING_KEY
		sendMap["Content"] = string(dataStr)
		sendStr, err := json.Marshal(sendMap)
		if err != nil {
			GoCrab.Debug("json.Marshal error", err, sendMap)
		}

		SendUrl = enums.GetSendProxy()

		GoCrab.Debug("Memcache tasks -- data send -- Begin")
		GoCrab.Debug("SendProxy is ", SendUrl)

		req := httplib.Post(SendUrl)

		if GoCrab.RunMode == GoCrab.RUNMODE_DEV {
			req.Debug(true)
		}

		req.SetTimeout(5*time.Second, 5*time.Second)
		req.Header("Content-Type", "text/plain")
		req.Body(string(sendStr))

		_, repErr := req.Bytes()

		GoCrab.Debug("DataMQ tasks -- data send -- Content is ", string(sendStr))

		if repErr != nil {
			models.Status.SendDataErrorCount++
			GoCrab.Error("api error ", repErr)
		}
		LastSendTime = time.Now().Unix()
		models.Status.SendDataCount++

		GoCrab.Debug("DataMQ tasks -- data send -- End")

		time.Sleep(time.Millisecond * enums.ENUM_MILLISECOND_INTERVAL)
	}

	return true
}

func reader() bool {
	//fmt.Printf("reader  %v\n", models.GetCount())
	return true
}

func TaskInit() {
	GoCrab.Info("Agent Running ", enums.PLUGIN_NAME)

	channel.SetChannelCount(enums.ENUM_CHANNEL_COUNT)
	ChannelSharded = channel.Sharded{make(chan int), make(chan int)}
	channel.ShardedWatching(ChannelSharded, writer, reader)

	for {

		go func() {
			WriteWatcher()
		}()

		time.Sleep(time.Second * time.Duration(enums.GetFrequency()))
	}
}

func WriteWatcher() {
	ChannelSharded.Writer <- 1
}

func ReadWatcher() {
	<-ChannelSharded.Reader
}
