# Golang-CLI Lambda Test Assignment

### Requirements to run this Project in local
1. Install go in your system. You can download package from [here](https://golang.org/doc/install)
2. Install mysql in your system. You can download mysql [here](https://dev.mysql.com/doc/refman/8.0/en/macos-installation-pkg.html) 

### Steps to set up Database for this project
1. Open `terminal` and login to mysql  
`mysql -uroot` or `mysql -uroot -p` (if you have setup password for root)
2. Create local user to access database ->  
`CREATE USER 'testuser'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';`
3. Create database ->  
`Create database testDb;`
4. Grant Permission to testuser to access this db -> `GRANT ALL PRIVILEGES ON testDb.* TO 'testuser'@'localhost';`
5. Create table tasks ->  
`use testDb;`  
`CREATE TABLE tasks (
   id int unsigned NOT NULL AUTO_INCREMENT,
   description varchar(200) DEFAULT NULL,
   status varchar(100) DEFAULT NULL,
   PRIMARY KEY (id)
   );`  
`exit;`
### Steps to set up this project in local
1. Clone this repository in your local.
2. Open `terminal` and run `cd $PATH/TO/YOUR/REPO/Golang-CLI` 
3. Run `go mod download` command to install required packages.
4. Now run `go run .` to bring up the project.

### Commands supported
1. `task` -> Display the user manual
2. `task add` -> Add a task to your task list.
3. `task do "task_id"` -> Complete task with given id
4. `task list` -> List all remaining tasks

**Note:**
`Use "task [command] --help" for more information about a command.
`