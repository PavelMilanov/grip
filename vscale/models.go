package vscale

type VscaleServer struct {
	Image    string `json:"make_from"`
	Size     string `json:"rplan"`
	Do_start bool   `json:"do_start"`
	Name     string `json:"name"`
	Keys     []int  `json:"keys,omitempty"` // заменит keys пустым значением, если мы его не передаем
	Password string `json:"password"`
	Location string `json:"location"`
}

type ServerConfig struct {
	Ctid        int               `json:"ctid"`
	Name        string            `json:"name"`
	Status      string            `json:"status"`
	Location    string            `json:"location"`
	Size        string            `json:"rplan"`
	Keys        []int             `json:"keys,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	PublicAddr  map[string]string `json:"public_address,omitempty"`
	PrivateAddr map[string]string `json:"private_address,omitempty"`
	Image       string            `json:"made_from,omitempty"`
	CreateTime  string            `json:"created,omitempty"`
	Active      bool              `json:"active"`
	Loced       bool              `json:"loced"`
	Deleted     bool              `json:"deleted,omitempty"`
}
