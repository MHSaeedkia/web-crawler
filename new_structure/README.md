# Telegram Bot - Web Crawler
This project is a web crawler specifically designed to scrape announcements related to buying, selling, and renting houses or apartments from popular Iranian websites. It focuses on crawling content from two major platforms: Diavar.

## Features
### Crawler
- Multi-site support: Crawls announcements from both Diavar and Sheypoor
- Specific focus on real estate listings (buying, selling, renting)
- Customizable crawling parameters
- Flexible output options for easy data processing

### Telegram Bot
- Different access levels - auth
- Creating a filter to display specific posts for each user
- Bookmark posts for each user
- Show bookmarks of users

### Project structure
- Modularity
- Custom migration definition
- Creating a seeder for default or fake data - filling the database with code
- Custom command for each module
- Separate test for each module
- Booting a separate env for tests, for example using `in-memory-sqlite` connections to run tests faster.
- Single or group testing of all modules
- The logic of every Telegram page can be tested, because we don't have any dependency on the logic of the Telegram bot
- You can have several drivers for how to send, edit and delete messages to Telegram
- You can define several database connections and dynamically change them in env, and all modules will use that connection.
- The isolation of each module
- According to SOLID principles and clean code as important as that part
- The design patterns used are:
  - Factory Method `To inject dependency from outside and change the strategy when writing tests and implementing design patterns like Singleton and Strategy Pattern`
  - Singleton / Lazy `Preventing the creation of multiple connections or other objects and increasing speed`
  -  Facade `Reduce dependencies between classes and use interfaces for direct use with classes`
  - Strategy `Changing the strategy while running the test and having several drivers to do the same job`
  - Repository `Not repeating queries, changing queries if necessary for other connections such as Redis`

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

4) Start the program

   Run each of these commands separately - you can use the supervisor to keep both commands up.
```
go run main.go telebot:start
```
```
go run main.go web-crawler:get-posts
```
> Note: You can get the program with the `go build main.go` build command and do not use the run command and run it directly as `main.exe telebot:start`.
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


# Commands
#### Show Commands:

Be sure to enter the path list of your module's test folders in this file
```
go run main.go command:list
```
### Start Bot Server

Be sure to enter the path list of your module's test folders in this file
```
go run main.go telebot:start
```
If you are using Windows, use the following commands to get v2ray proxy cmd.
```
set HTTP_PROXY=http://127.0.0.1:10809
set HTTPS_PROXY=http://127.0.0.1:10809
```
### Database
Generate tables
```
go run main.go database:migrate up -v
```

Delete tables
```
go run main.go database:migrate up -v
```

Seed fake data
```
go run main.go database:seed -v
```

Executes these commands `migrate up`, `migrate down`, `seed` in order - To rewrite the tables
```
go run main.go database:fresh
```

### Web Crawler
Start collecting posts and save in posts table
```
go run main.go web-crawler:get-posts
```