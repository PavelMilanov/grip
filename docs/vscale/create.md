### Создание сервера
***
```
grip vscale create \
  --image=
  --loc=
  --name=
  --plan=
  --pwd=
  --start=
```

#### Флаги:
 - `--image=` - Название образа ОС ( по-умолчанию Debian 11 );
 - `--loc` - Местоположение сервера ( по-умолчанию МСК ); 
 - `--name` - Обязательный флаг. Название сервера;
 - `--plan` - Конфигурация сервера ( по-умолчанию small );
 - `--pwd` - Обязательный флаг. Root-пароль;
 - `--start` - Автоматический запуск сервера после создания ( по-умолчанию true ).

#### Пример:
```
grip vscale create --name=vscale-server --pwd=password
```
> Server successfully created

![[vscale create.png]]