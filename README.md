# Arifmetic_Server

Привет тебе путник!

Ну а если серьезно.

Что бы запустить сервер выполни команду make, что бы запустьть makefile находящийся в главной директории. Если не знаешь что это такое или нет программы для их запуска введи в консоле команду:

go run services/cmd/server/main.go --config=config_file/configuration.json 

Итак, а теперь давай объясню как весь мой бредовый код работает.
В директории есть файл Sheme. Там немного визуализированно то как сейчас это работает.

Собственно, что бы отправит мат выражение открой браузер и введи:
http://localhost:8080/
(Ps. Можно естественно отправлять запросы с поощью Postman тогда адрес отправки должен быть http://localhost:8080/arifmetic)

Там откроется простенький интерфейс. (Я во фронте ничего не смыслю так, что как получилось)

В первом окне вводится выражение. Можно ввести неколько разделяя их знаком ;.

Нажав кнопку отправить в строке "Ваш ID зароса" выведиться id. К этому ID привязаны все отправленные выражения.  

Во втором окне можно ввести полученный ID и отправить его на сервер. 
(Ps. Отправлять можно только один ID, но отправить на сервер можно несколько выражений и соответственно иметь несколько ID зароса)

В нижней части находится таблица где выводится информация об отправленном ID.
Если ID на сервере не существует то в таблице выведится запись например:
123 NULL 0 NOT FOUND

Если ID существует то выведится вся информация об связанных с ним выражениях.
У каждого выражения на сервере свой ID и выводится в таблице будет он.
Если выражение посчитано то выведится примерно такая запись:
12 1 + 1 2 OK
Если выражение еще в работе то выведится такая запись:
12 NULL 0 WORKING
(Ps. Пока что не успел реализовать вывод если выражение с ошибкой)
(Ps. Что бы обновлять ответ нужно еще раз нажимать кнопку отправить. Если все выражения посчитались ответом должна быть строка: 
ваш ID NULL 0 NOT FOUND)
(Ps. При количестве отправленных выражений больше 10, есть вероятность что поледнее выражение зависнет. Пока не смог пофиксить)


Теперь посмотрим что проиходит на стороне сервера.

Вся основная логика это пакет pkg/agent. В будующем планирую разнести логику по более мелким пакетам.

Функция MainOrchestrator принимает запросы и отправляет на них ответы. 
(Ps. Я так и не до конца понял что называть оркестратором, демоном и агентом , так что здесь все уловно)

Наша функция принимает два запроса POST и GET.

На первый запрос функция формирует список структур UserTask и отправляет их в структуру AgentServiceInput. В ответ отправляет сгенерированный ID к которому привязаны ID уже конкретного выражения.

На второй запрос функция берет Header c нашим ID и формирует свой ответ в зависиоти есть ли ID в базе сервера и от статуса привязанных к нему выражений. (Ps. Эта часть делана на скорую руку так что позже планирую ее переделать) 

При запуске сервера с в горутинах создается несколько функций Demon. У себя они имеют AgentServiceInput откуда берется очередное выражение и  AgentServiceOutput куда посчитанные выражения отправляются. (Ps. По дефолту функций создается 5 но в файле конфигурации в параметре count_agent их количество можно поменять)

Каждая функця считает одно целое выражение с омощью модуля arifm. Данный одуль работает на основе готового парсера.

Так же в горутине оздается функция Output которая берет из структуры AgentServiceOutput посчитанное выражение и заносит его в мапу в структуре MainOrchestratorService.

В данной структуре находится другая труктура DataIndex где находится несколько мап в которых информация к какому ID зароса привязаны ID выражения отправлялось ли это выражения и количество оставшихся выражений. (Ps. Эта часть собственно и сделана на скорую руку и в требует переделки)

В общих чертах как то так, возможно что то упустил.

Что насчет примеров. Можно слать любые в том чиле и со скобками ответ будет правильным. Количетво одновременно отправляемых примеров до 20, больше я не отправлял. На выражения вида 1 + 2 89 будет ответ -1. Если в выражении есть остаток от деления ответом будет целая часть. (Ps. Надо будет пофиксить)

Примеры выражений:
1 + 2 
4 * 3 
6 - (8 + 2)
89 / 8
678 + 877 
8877 / 88 
7 * 65
0 - 877 * 76
1 + 234 * (88766 - 888) 
76 - 8 
45 + 866 
7546 - 87777 
788 / 88 
766 / 5 
90 - 888 
890 / 7 
788 * 777


PS. Возможно нужно будет установить парсер выражении командой.
go get github.com/mgenware/go-shunting-yard

PS. Что хочется сделать:
1.Интегрировать базу данных в проект. Если посмотрет файлах уже есть сама база в storage и пакет cmd/migrations с выполненнием миграций базы. Сами файлы миграций находяться в корне в папке migration. (Ps. Если не знаешь что такое миграции почитай довольно полезная вещь).

2. Создать Docker контейнер. Сам Dockerfile есть но не успел создать сам контейнер.

3. Переписать все костыли в проекте. И переделать проект под полные требования задачи.

4. Возможно сделать более удобный интерфейс.

Спасибо за внимание! 
Tg: @Nikolasff
