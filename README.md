Данное fullstack приложение позволяет пользователям вести учет потребляемых ими калорий, а также рассчитывать свой индекс массы тела и определять необходимую сучтоную норму калорий.

Минималистичный интерфейс. Возможность сохранять данные и просматривать их всегда, когда это необходимо.

Стек:
Back-end:  Golang + Gin, Docker, Docker-compose, Swagger<br>
Front-end: ReactJS + react-bootstrap<br>
DB:        MongoDB<br>

Для старта проекта необходимо перед первым запуском ввести комманду в терминале, находясь в корне проекта, для сборки контейнеров и установки необходимого окружения:<br>
```docker-compose up -d --build```<br><br>

После этого соберутся контейнеры, далее проект можно запускать с помощью команды:<br>
```docker-compose up -d```<br><br>

Notes: <br>
    1. После внесений изменений в проект необходимо пересобрать контейнеры, выполнив команду с флагом ```--build```<br>
    2. Для выведения логов контейнера в терминал необходимо ввести комманду:<br>
    ```docker-compose logs --tail=КОЛ-ВО_СТРОК НАЗВАНИЕ_КОНТЕЙНЕРА```<br><br>
    Структура back-end:
    ![strcuture](https://github.com/ADEXITUM/React-GO-Calorie-Tracker/assets/111490239/ff888910-0b7f-4139-bc82-5eef92fbe2ca)
