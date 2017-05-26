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
	"github.com/CloudWise-OpenSource/GoCrab-Demo/MongoAgent/enums"
	"github.com/CloudWise-OpenSource/GoCrab/Core/heatbeat"
)

func init() {
	heatbeat.FileInit(enums.GetFrequency())
}
