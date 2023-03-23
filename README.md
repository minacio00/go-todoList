# go-todoList
this is a simple rest backend make in go and postgres as a database but the main focus here was to use docker compose.

## Running the project:
  You should be able to run it just by cloning the repo and then running 'docker compose up' in the repository's root.
  docker will wait until the database is ready to start the server, so when you see something like: 

  ![image](https://user-images.githubusercontent.com/43221251/224573798-dc8ae154-db37-46b3-883c-62e6b46682b7.png)


  you are good to go.
  
  The REST backend will be exposed to the port 8080, a get request on '/' will query all the rows in a table called lul, a get on '/todos' you can guess
