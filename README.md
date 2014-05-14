fswatcher
=========
A simple file watcher written in Go.

### config example
```
[
  {
    "path" : "/path/to/your/project",
    "pattern" : "*.go",
    "ignore" : ".*",
    "command" : "go test -v ./...", 
    "recursive" : true
  },
  {
    "path" : "/path/to/another/project", 
    "pattern" : "*.rb",
    "command" : "rake"
  }
]
```

The configuration file must be named ".fswatcher" and be located in current directry.
"path", "pattern" and "command" are necessary. "ignore" and "recursive" are optional.
The working directory of the command will be the directry specified by "path".

