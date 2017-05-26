# GoCrab-Demo
## MongoAgent

### MongoAgent Main
```
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

```

### Directories and files
```
.
├── README.md
├── bin
│   └── touch
├── conf
│   └── app.ini
├── enums
│   ├── config.go
│   └── enum.go
├── logs
│   └── touch
├── main.go
├── models
│   ├── MongoDB.go
│   └── Status.go
├── pics
│   ├── mongo_agent_help.png
│   ├── mongo_agent_start.png
│   └── mongo_agent_version.png
└── tasks
    ├── Agent.go
    └── Heatbeat.go

7 directories, 14 files
```

### ./MongoAgent -h
![MongoAgentHelp](https://raw.githubusercontent.com/CloudWise-OpenSource/GoCrab-Demo/master/MongoAgent/pics/mongo_agent_help.png)

### ./MongoAgent -v
![MongoAgentVersion](https://raw.githubusercontent.com/CloudWise-OpenSource/GoCrab-Demo/master/MongoAgent/pics/mongo_agent_version.png)

### ./MongoAgent start
![MongoAgentStart](https://raw.githubusercontent.com/CloudWise-OpenSource/GoCrab-Demo/master/MongoAgent/pics/mongo_agent_start.png)

