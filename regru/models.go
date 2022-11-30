package regru

type RegruServer struct {
	Image    string `json:"image"`
	Size     string `json:"size"`
	Name     string `json:"name,omitempty"`
	Keys     []int  `json:"ssh_keys,omitempty"` // заменит keys пустым значением, если мы его не передаем
	Backups  bool   `json:"backups,omitempty"`
	Location string `json:"region_slug,omitempty"`
}

type ServerConfig struct {
	Server ServerReglet `json:"reglet"`
}

type ServerReglet struct {
	Backup               bool        `json:"backups_enabled,omitempty"`
	CreatedAt            string      `json:"created_at"`
	Disk                 int         `json:"disk"`
	External_application string      `json:"external_application,omitempty"`
	Id                   int         `json:"id"`
	Image                ServerImage `json:"image"`
	ImageId              int         `json:"image_id"`
	Ip                   string      `json:"ip,omitempty"`
	LastBackupDate       string      `json:"last_backup_date,omitempty"`
	Locked               int         `json:"locked"`
	Memory               int         `json:"memory"`
	Name                 string      `json:"name"`
	Ptr                  string      `json:"ptr"`
	Region               string      `json:"region_slug"`
	SizeSlug             string      `json:"size_slug"`
	Status               string      `json:"status"`
	SubStatus            string      `json:"sub_status,omitempty"`
}

type ServerImage struct {
	CreatedAt    string `json:"created_at"`
	Distribution string `json:"distribution"`
	Id           int    `json:"id"`
	DiskSize     string `json:"min_disk_size"`
	Name         string `json:"name"`
	Private      bool   `json:"private"`
	Region       string `json:"region_slug"`
	Size         string `json:"size_gigabytes"`
	Slug         string `json:"slug"`
	Type         string `json:"type"`
}
