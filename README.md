Базовая версия социальной сети с 4 ручками:
1. /login
2. /user/register
3. /user/get/{id}
4. /user/search?firstName={first}&secondName={second}

Запуск:
- docker-compose up
- проиницилизировать структуру БД запросами из папки `./migrations`
- для заполнения данными (1 млн записей): `make gen-csv`

Особенности:
- GRPC ручки проброшены на http (контракт `./api/soc-net.proto`)
- примеры запросов лежат в папке postman
- пароли пользователей хранятся в виде хэшей (с солью)