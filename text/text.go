package text

// цвета текста
// const colorReset = "\033[0m"
const RED = "\033[31m"
const GREEN = "\033[32m"
const YELLOW = "\033[33m"
const BLUE = "\033[34m"
const CYAN = "\033[36m"
const WHITE = "\033[37m"

const GRIP_MENU = `
grip init	- add prodvider token. (vscale, regru, ruvds)
grip vscale	- menu interaction of vscale-provider.
grip regru	- menu interaction of regru-provider.
grip ruvds	- menu interaction of ruvds-provider.  
`

const VSCALE_MENU = `
grip vscale ls		- view servers.
grip vscale create	- create new server.
grip vscale inspect	- inspect server config by name.
grip vscale rm		- remove server by name.
grip vscale stop	- stop server.
grip vscale start	- start server.
grip vscale restart	- restart server.
grip vscale ssh		- ssh connection to server by alias.
`
const REGRU_MENU = `
grip regru ls		- view servers.
grip regru create	- create new server.
grip regru inspect	- inspect server config by name.
grip regru rm		- remove server by name.
grip regru stop		- stop server.
grip regru start	- start server.
grip regru restart	- restart server.
grip regru ssh		- ssh connection to server by alias.
`

const RUVDS_MENU = `
grip ruvds ls		- view servers.
grip ruvds create	- create new server.
grip ruvds inspect	- inspect server config by name.
grip ruvds rm		- remove server by name.
grip ruvds stop		- stop server.
grip ruvds start	- start server.
grip ruvds restart	- restart server.
grip ruvds ssh		- ssh connection to server by alias.
`

const INIT_MENU = `
grip init -provider=<provider> -token=<provider token> (vscale, regru)
grip init -provider=<provider> -token=<provider token> -username=<username> -password=<password> (ruvds)
`
