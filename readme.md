<h2>Русский (Russian)</h2>
<b>FTP (File Transfer Protocol)</b> - протокол передачи файлов.

В стеке TCP/IP этот протокол находится на прикладном уровне (рядом с HTTP, DNS, ...).<br>
FTP использует протокол транспортного уровня TCP с портом 21 управляющего соединения.
Соединение для передачи данных может быть установленно в двух режимах: активном и пассивном. 
При этом используются разные номера портов.

Часть 1 (Part 1)
https://betterprogramming.pub/how-to-write-a-concurrent-ftp-server-in-go-part-1-3904f2e3a9e5

Часть 2 (Part 2)
https://betterprogramming.pub/how-to-write-a-concurrent-ftp-server-in-go-part-2-4a59f4216639

| Command        | Действие           | Action |
| :-------------: |:-------------:| :---: |
|USER|Указать имя пользователя|Set username|
|PASS|Указать пароль|Send password|
|LIST|Просмотр содержимого каталога|Show everything in current directory|
|PWD|Возвращает текущий каталог|Returns current directory|
|CWD|Смена текущего каталога|Move to directory|
|CDUP|Сменить каталог на вышестоящий|Move up directory|
|MKD|Создать каталог|Create directory|
|RMD|Удалить каталог|Remove directory|
|RETR|Передать файл с сервера на клиент|Download file from server to client|
|STOR|Передать файл с клиента на сервер|Upload file from client to server|
|TYPE|Установить режим передачи|Set transfer type|
|DELE|Удалить файл|Remove file|
|PASV|Использовать пассивный режим|Use passive mode|
|REIN|Реинициализировать подключение|Connection reinitialize|
|QUIT|Выход и разрыв соединения|End


Коды ответов FTP<br>
Сокращенно:

| Код        | Расшифровка           |
| :-------------: |:-------------:|
|2xx|Успешный ответ|
|4xx/5xx|команда не может быть выполнена|
|1xx/3xx|Ошибка или неполный ответ|

sad

| link | as|
| :---: | :---: |
|1 | 2|

| https://www.serv-u.com/resource/tutorial/110-120-125-150-ftp-response-codes ||
|https://www.serv-u.com/resource/tutorial/202-211-212-213-ftp-response-codes|
https://www.serv-u.com/resource/tutorial/214-215-220-221-ftp-response-codes
https://www.serv-u.com/resource/tutorial/225-226-227-230-ftp-response-codes
