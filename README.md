### Клиент-серверный файловый сервис на базе протокола TCP
Клиент *client.go* шифрует свои  файлы  алгоритмом AES с режимом CBC и помещает их на сервер *server.go*. Также, скачивает с сервера и расшифровывает.

#### Порядок работы:
- [x] Пользователь запускает client, при успешном соединении клиента с сервером вводит свой логин/пароль для входа на сервер. Появляется строка приглашения **ftp>**.
- [x] Для копирования файла на сервер вводит `upload <имя своего файла>`, затем пароль (не меньше 5ти символов).
- [x] Происходит шифрование файла алгоритмом AES с режимом CBC, затем закачка файла на сервер. Сервер сообщает новое имя файла в формате UUID.
- [x] Пользователь может просмотреть список файлов на сервере при помощи *ls*.
- [x] Пользователь вводит имя (на сервере все имена файлов в формате UUID) нужного ему на сервере файла `upload <имя файла на сервере>`, затем пароль.
- [x] Происходит передача файла с сервера на клиентскую часть, расшифровка и сохранение с серверным именем. 


#### Список команд:
- **upload** upload file from `filestore/clientDir` to `filestore`
- **download** download file from server
- **ls** list files on server
-  **close/exit** close connection


#### Config and run
Fill `credential.json` your usernames/passwords.
Open two terminals on project folder
1st terminal:
```bash
go run cmd/server.go
```
2nd terminal:
```bash
#Multiple clients can be attached 
go run cmd/client/client.go
```
#### Code review:
Основной проблемой при передаче данных по сокету являлось разделение потока байтов на команды (string) и сам файл ([]byte). `net.Conn`, к сожалению, имеет очень ограниченный набор команд. В частности, не хватает io.ReadRune. Поэтому выходом стало скачивание потока побайтно `io.CopyN(,,1)` с сравнением каждого байта с '\n'. До '\n' передается размер файла в байтах, после '\n' скачиватся этот размер файла (`client/getfile.go`).

Обнаружилось, что для блочного шифрования файла его исходный разиер в байтах должен быть кратным 16. Поэтому выполнено дописывание в конец файла \x01 и затем нулей \x00, соответственно которые убираются при дешифровке (`client/crypto.go`).

Проект выполнен встроенным набором Golang без сторонних библиотек, кроме получения UUID при помощи [google/uuid](github.com/google/uuid).

Асинхронность достигается выделением каждого



*Make sure to allow traffic to the port specified on your VPC firewall.*

Thanks for samples [bisakhmondal/FTP-Go](https://github.com/bisakhmondal/FTP-Go), [kdama/gopl](https://github.com/kdama/gopl/tree/master/ch08/ex02), [fclairamb/ftpserver](https://github.com/fclairamb/ftpserver).
