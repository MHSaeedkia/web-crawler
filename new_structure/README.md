# TelegramBot-Boilerplate
Project structure for the Telegram bot in golang and similar to Laravel.


## Getting Started
1) install packages
```
go mod tidy
```

2) copy env file - set config db
```
cp example.env .env
```
```
cp example.env .env.test
```

3) generate DB 
```
go run main.go database:fresh
```

## Run tests

Be sure to enter the path list of your module's test folders in this file
```
sh run-tests.sh
```
Custom test file
```
sh run-single-test.sh ./modules/auth/Tests/auth_page_test.go
```

If you are on Windows and want to use the sqlite database for tests, you must download the (MinGW-w64 based) software from the link below (make sure to close and reopen cmd after installation)
```
https://jmeubank.github.io/tdm-gcc/
```
If you don't want to use sqlite, set the `CGO_ENABLED` value to `0` in the `run-tests.sh` file
#### Suggestion
Make sure you have the `.env.test` file and the `DB_CONNECTION=in_memory_sqlite` value is set


## Commands
#### Show Commands:

Be sure to enter the path list of your module's test folders in this file
```
go run main.go command:list
```
#### Start Bot Server

Be sure to enter the path list of your module's test folders in this file
```
go run main.go telebot:start
```
If you are using Windows, use the following commands to get v2ray proxy cmd.
```
set HTTP_PROXY=http://127.0.0.1:10809
set HTTPS_PROXY=http://127.0.0.1:10809
```

Generate tables
```
go run main.go database:migrate up -v
```

Delete tables
```
go run main.go database:migrate up -v
```

seed fake data
```
go run main.go database:seed -v
```