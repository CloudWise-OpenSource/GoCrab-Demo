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
	"github.com/CloudWise-OpenSource/GoCrab/Cli"
	"github.com/CloudWise-OpenSource/GoCrab-Demo/MongoAgent/enums"
	"github.com/CloudWise-OpenSource/GoCrab/Libs/dbdrive/mgo.v2"
)

type States struct {
	Host           string
	Version        string
	Process        string
	Uptime         int64
	UptimeEstimate int64
	LocalTime      string
	GlobalLock     struct {
		TotalTime    int     `bson:"totalTime"`
		LockTime     int     `bson:"lockTime"`
		Ratio        float64 `bson:"ratio"`
		CurrentQueue struct {
			Total   int `bson:"total"`
			Readers int `bson:"readers"`
			Writers int `bson:"writers"`
		}
		ActiveClients struct {
			Total   int `bson:"total"`
			Readers int `bson:"readers"`
			Writers int `bson:"writers"`
		}
	}

	Mem struct {
		Bits              int `bson:"bits"`
		Resident          int `bson:"resident"`
		Virtual           int `bson:"virtual"`
		Supported         int `bson:"supported"`
		Mapped            int `bson:"mapped"`
		MappedWithJournal int `bson:"mappedWithJournal"`
	}

	Connections struct {
		Current   int `bson:"current"`
		Available int `bson:"available"`
	}

	Cursors struct {
		TotalOpen int `bson:"totalOpen"`
		TimedOut  int `bson:"timedOut"`
	}

	BackgroundFlushing struct {
		Flushes    int `bson:"flushes"`
		Total_ms   int `bson:"total_ms"`
		Average_ms int `bson:"average_ms"`
	}

	IndexCounters struct {
		Btree struct {
			Accesses  int `bson:"accesses"`
			Hits      int `bson:"hits"`
			Misses    int `bson:"misses"`
			Resets    int `bson:"resets"`
			MissRatio int `bson:"missRatio"`
		}
	}

	Network struct {
		BytesIn     int `bson:"bytesIn"`
		BytesOut    int `bson:"bytesOut"`
		NumRequests int `bson:"numRequests"`
	}
	Opcounters struct {
		Insert  int `bson:"insert"`
		Query   int `bson:"query"`
		Update  int `bson:"update"`
		Delete  int `bson:"delete"`
		Getmore int `bson:"getmore"`
		Command int `bson:"command"`
	}
	WriteBacksQueued bool `bson:"writeBacksQueued"`
	Ok               int  `bson:"ok"`
}

func GetStats() (maps map[string]interface{}, err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			GoCrab.Error(panicErr)
		}
	}()

	session, err := mgo.Dial(enums.GetHostIp() + ":" + enums.GetPort())
	if err != nil {
		GoCrab.Error(err)
	}
	defer session.Close()

	//var states map[string]interface{}
	var states States
	err = session.Run("serverStatus", &states)
	if err != nil {
		GoCrab.Error(err)
	}

	result := make(map[string]interface{})
	result["globalLock_currentQueue_writers"] = states.GlobalLock.CurrentQueue.Writers
	result["globalLock_currentQueue_total"] = states.GlobalLock.CurrentQueue.Total
	result["globalLock_ratio"] = states.GlobalLock.Ratio
	result["globalLock_currentQueue_readers"] = states.GlobalLock.CurrentQueue.Readers

	result["opcounters_getmore"] = states.Opcounters.Getmore
	result["opcounters_command"] = states.Opcounters.Command
	result["opcounters_update"] = states.Opcounters.Update
	result["opcounters_delete"] = states.Opcounters.Delete
	result["opcounters_insert"] = states.Opcounters.Insert
	result["opcounters_query"] = states.Opcounters.Query

	result["connections_current"] = states.Connections.Current
	result["connections_available"] = states.Connections.Available

	result["nest_sub"] = make(map[string]string)
	result["uptime"] = states.Uptime

	if states.IndexCounters.Btree.Accesses == 0 {
		result["indexCounters_usefrequency"] = 0
	} else {
		result["indexCounters_usefrequency"] = states.IndexCounters.Btree.Hits / states.IndexCounters.Btree.Accesses
	}

	result["indexCounters_access"] = states.IndexCounters.Btree.Accesses
	result["indexCounters_missRatio"] = states.IndexCounters.Btree.MissRatio
	result["indexCounters_resets"] = states.IndexCounters.Btree.Resets

	result["mem_resident"] = states.Mem.Resident
	result["mem_virtual"] = states.Mem.Virtual
	result["mem_supported"] = states.Mem.Supported

	result["host"] = states.Host
	result["version"] = states.Version

	return result, err
}
