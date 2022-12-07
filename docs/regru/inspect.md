#### Просмотр конфигураци инфраструктуры
***
```
grip regru inspect <server>
```
#### Пример:
```
grip regru inspect vscale-server
```

>{
>	"reglet": {
>		"created_at": "2022-12-07 23:50:11",
>		"disk": 15,
>		"id": 2285845,
>		"image": {
>			"created_at": "2021-08-18 09:34:14",
>			"distribution": "debian-11",
>			"id": 975939,
>			"min_disk_size": "5",
>			"name": "Debian 11",
>			"private": false,
>			"region_slug": "msk1",
>			"size_gigabytes": "1",
>			"slug": "debian-11-amd64",
>			"type": "distribution"
>		},
>	"image_id": 975939,
>	"locked": 0,
>	"memory": 1024,
>	"name": "regru-server",
>	"ptr": "",
>	"region_slug": "msk1",
>	"size_slug": "base-1",
>	"status": "ordered"
>	}
>}