# yagoll (Yet Another Go Logging Library)

There are many logging libraries for Go and this is one of them. Project was created without any ambitions or higher goals. 
I found no logging lib that I was comfortable to work with and has all features I want, so I've created my own.

### Goals
 * easy to migrate from standard Go `log` by changing `import "log"` to `import log "github.com/sparkoo/yagoll"`
 * coming from Java world, I'm used to 5 levels -> TRACE, DEBUG, INFO, WARN, ERROR
 * ability to filter messages by levels
 * print file and line of log message

### Getting started
`go get github.com/sparkoo/yagoll`
```
import "github.com/sparkoo/yagoll"

func main() {
  yagoll.Debug("Hello World")
}
```

### Migrating from native Go `log`
```
/// current source
import "log"

func main() {
  log.Println("Hello World")
}
```

`go get github.com/sparkoo/yagoll`

```
/// new source with yagoll logging
import log "github.com/sparkoo/yagoll"  /// just changed import line

func main() {
  log.Println("Hello World")
}
```
