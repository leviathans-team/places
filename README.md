# Структура проекта
## Цимис
Чтобы всем было понятно в любом проекте, где какая хуйня хранится, где искать логику, допустим, на исправление бага, мы должны придерживаться одной архитектуры приложения. Поэтому будет этот репозиторий с базовой информацией и шаблонной структурой проекта чисто с папками и пустыми файлами для понимания ситации

## Общий концепт

Существуют общие пакеты (папки простым языком), в которых мы работаем, и которые должны быть в каждом проекте (от языка может меняться названия, но точно не корневых).
1. `cmd` - это папка со всеми исполняемыми файлами, то есть файлами, где есть функция, запускающая программу (в джаве класс с мейн методом, в го мейн функция и мейн пакет, в питоне `__init__ == '__main__': app.run()`). Сама функция/класс должны лишь создавать инстанс сервера/консьюмера и запускать его. Может мелкая предвадительная конфигурация, её логгирование и всё. Больше ничего.
2. `config` - там лежит `.go`/`.java`/`.py` файл, парсящий конфиги и отдающий их для сервера + `.yml`, в котором прописаны конфиги (если совсем не получается в парсинг yml можете и из `.env` файла, но лучше не надо, пожалуйста). Пример конфиги и её оформления лежит в соответствующем пакете.
3. `internal` - главный по сути пакет приложения. Там находится вся бизнес-логика. То есть вся ваша работа будет проходить именно там. На нём остановимся поподробнее:
    1. Итак, мы решаем, допустим, сделать сервис с ивентами. У нас обязательно должна быть авторизация (об этом чуть позже подробнее). Тогда мы создаём пакет `auth`. Для авторизации нам надо как-то проверить данные, которые к нам пришли - значит нам нужен доступ к БД - это дело _репозитория_.
        * **Репозиторий** - это класс (структура, мб вы просто функции нахерачите, но лучше класс/структура, выполняющий контракт какого-то интерфейса), который имеет при себе соединение с БД и активно пользуется им. То есть его цель - только закинуть SQL (или NoSQL) запрос к БД, вытащить данные и отдать дальше.
        * Дальше мы должны эти данные обработать - это работа _юзкейса_ (или сервиса в логике Java). **Юзкейс** - это так же класс, который имеет при себе репозиторий и другие средства для выполнения бизнес-логики. То есть именно здесь мы имеем функции/методы, которые обеспечивают непосредственно работу приложения. При этом стоит учесть, что он может иметь при себе репозитории или юзкейсы из других пакетов (главное без циклического импорта), так как ему может потребоваться дополнительная внутрення логика, которую вы выделили в отдельный пакет (допустим, авторизационные юзкейсы могут быть нужны в других для получения данных для авторизации своих запросов на другие сервисы).
        * **Хэндлеры** (_контроллеры_ в Java) - классы, просто функции, которые отвечают за работы непосредственно с сетью (по любому протоколу, вообще сюда можно отнести и функции консьюминга). То есть в пакете `delivery` мы можем разделить их на `http`, `websocket`, `rabbit` и т.д. Хрестоматийный пример в данном случае - обработчик запроса по роуту `auth/sign_in`. Цель хэндлера - получить и обработать данные в специфике своего протокола передачи данных и отдать в _юзкейсы_, которые уже совершат всё действо, а, получив от них ответы, отдать по своему же протоколу.
    2. Дальше у нас будет пакет `event`, в котором будет аналогичный набор пакетов и файлов, но уже связанный непосредственно с логикой мероприятий. То есть будут репозитории, работающие с данными в БД о мероприятиях, юзкейсы, реализующую логику работы с этими и поступающими из хэндлеров данных, и хэндлеры, принимающие данные от пользователя и других сервисов.
    3. Пакет `middleware` отвечает за _мидлварки_. Что это такое? **Middleware** простыми словами в конкретно данном случае (и только в данном в упрощённом объяснении для понимания вами логики необходимости этого пакета) являет собой промежуточный этап между запросом и хэндлером. Допустим, он может доставать авторизационные данные из хэдеров запроса или кукисов и валидировать их. Если мы будем говорить о `Go` с `Fiber`, то там это такой же инстанс _fiber.Handler_, но имеющий в конце не ответ на запрос, а вызов от контекста метода _Next()_ (а-ля `c.Next()`), то есть перекидывающий на следующий запрос, либо возвращающий ошибку, если валидация не прошла. Хрестоматийный пример - авторизация.
    4. Пакет `server` - пакет, отвечающий за создание и настройку сервера. Можно выделить в отдельный пакет, конечно, но каждое приложение настраивается по-разному, поэтому считается так же частью бизнес-логики. Можно сделать класс сервера со статическим методом-фабрикой, который будет возвращать готовый объект из конфига, и затем от него запускать метод `Run()`. Можно сделать конструктор, маппер хэндлеров, а затем метод запуска. На ваше усмотрение, так сказать.
4. `migrations` - папка с миграциями. Миграции простым языком это перенос базы данных с места на места. Чтобы не ебаться каждый раз, легче создать систему миграций, которая для каждого переноса будет подтягивать схемы таблиц, возможно какие-то критические дампы и так далее. Может быть просто файлом с SQL запросами, может быть реализовано через какую-то систему миграция вроде yoyo в python (здесь в примерном docker-compose подтягивается из простого sql файла). Вы можете даже не автоматизировать этот процесс, но не иметь sql файла хотя бы, чтобы даже вручную быстро поднять, - это преступление
5. `pkg` - пакет, переводящийся с английского как пакет :) По сути здесь находятся все пакеты и файлы, которые можно без модификации перекидывать из сервиса в сервис, для каких-то удобных фич: общий логгер компании, создания подключений к СУБД, создание паблишера и консьюмера для брокеров сообщений, http-клиент, обработчик ошибок и так далее.

## Выводы

Гоферы могут в принципе склонить этот репозиторий и с него начать работу уже, так как тут есть те же докер файлы, которые позволят вам не ебаться много с ним. Но в целом вы будете его наполнять сами (кроме pkg, его я вам как смогу накидаю, чтобы гемора убрать). Так же парсер конфига можете написать сами (в го очень удобная библа для этого viper), можете подождать и пока хардкодить, а потом я дам вам красивый парсер.
Надеюсь, вы хоть что-то поняли из этого полотна. Пишите, звоните, телеграфируйте. Удачи!

P.S. Только в своих репозиториях реальные конфиги и .env файлы не палите. Сразу ваши репозитории буду сносить к хуям.