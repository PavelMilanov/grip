### Создание сервера
***
```
grip regru create \
  --image=
  --loc=
  --name=
  --plan=
  --bkp=
```

#### Флаги:
 - `--image=` - Название образа ОС ( по-умолчанию Debian 11 );
 - `--loc` - Местоположение сервера ( по-умолчанию МСК ); 
 - `--name` - Обязательный флаг. Название сервера;
 - `--plan` - Конфигурация сервера ( по-умолчанию base-1 );
 - `--bkp` - Автоматический запуск бэкапирования сервера ( по-умолчанию false ).

#### Пример:
```
grip regru create --name=vscale-server
```
> Server successfully created

![[grip/docs/regru/create.png]]