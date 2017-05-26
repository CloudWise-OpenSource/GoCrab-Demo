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

### ./MongoAgent -h

### ./MongoAgent -v

### ./MongoAgent start
