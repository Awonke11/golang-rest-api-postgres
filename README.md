## Golang Rest API with Python

A Golang Rest API project with Postgres as the Database. Used Python 
```requests``` 
to fetch the data from the API.

### Requirements to get started
1. You must have <a href="https://go.dev/doc/install">Golang installed</a> in your system.
2. You must also have <a href="https://www.python.org/downloads/">Python installed</a> in your system as well.
3. Lastly, you will need to have <a href="https://www.postgresql.org/download/">PostgreSQL installed</a>

### Project Setup
- <a href="https://www.freecodecamp.org/news/how-to-fork-a-github-repository/#:~:text=To%20follow%20along%2C%20browse%20to,created%20under%20your%20GitHub%20account.">Fork this repo</a> into your own Github
- Then clone the repo into your system
    ```git clone url <directory>```
- Run the following command on the terminal:
    ```go mod init [module-name]```
- Followed by this command which will install all the golang dependencies: ```go mod tidy```
- Create a ```.env``` file on the main folder and add the following code with your details:
    ```env
    DB_HOST=127.0.0.1
    DB_PORT=5432
    DB_PASSWORD=[your password]
    DB_USER=[username]
    DB_NAME=[database name]
    DB_SSLMODE=disable
    ```
