package main

import (
    "bufio"
    "fmt"
    "github.com/jmoiron/sqlx"
    "log"
    "os"
    "strconv"
    "strings"
    db "golang-cli/v1/database"
)

const (
	InvalidCommandError = "Invalid command entered"
	Done = "DONE"
	Pending = "PENDING"
)

type Task struct {
    Id int `db:id`
    Description string `db:description`
    Status string `db:status`
}

func addTask(arg string, dbConn *sqlx.DB)  {
    if arg == "--help" {
        fmt.Println("  task add \"task_description\"         Add a new task to your TODO list")
    } else {
        addTaskEntry, _ := dbConn.Prepare("INSERT INTO `tasks` (`description`, `status`) VALUES (?,?)")
        _, err := addTaskEntry.Exec(arg, Pending)
        if err != nil {
            panic(err.Error())
        }
        fmt.Printf("Added \"%s\" to your task list.\n", arg)
    }
}

func doTask(arg string, dbConn *sqlx.DB) {
    if arg == "--help" {
        fmt.Println("  task do \"task_id\"         Mark task with \"task_id\" on your TODO list as complete")
    } else {
        taskId, err := strconv.Atoi(arg)
        if err != nil {
            panic(err.Error())
        }
        stmt, _ := dbConn.Prepare("Select * from tasks where id = ?")
        taskRes, err := stmt.Query(taskId)
        if err != nil {
            panic(err.Error())
        }
        task := Task{}
        if taskRes.Next() {
            err = taskRes.Scan(&task.Id, &task.Description, &task.Status)
            if err != nil {
                panic(err.Error())
            }
            if task.Status != Done {
                stmt, _ = dbConn.Prepare("UPDATE tasks set status = ? where id = ?")
                _, err = stmt.Exec("DONE", taskId)
                if err != nil {
                    panic(err.Error())
                }
                fmt.Printf("You have completed the \"%s\" task.\n", task.Description)
            } else {
                fmt.Printf("You have already completed the \"%s\" task.\n", task.Description)
            }
        } else {
            fmt.Printf("You have provided invalid task id.\n")
        }
    }
}

func listTask(dbConn *sqlx.DB) {
    fmt.Println("You have the following tasks:")
    list, err := dbConn.Query("Select * from tasks where status = 'PENDING'")
    for list.Next() {
        row := Task{}
        err := list.Scan(&row.Id, &row.Description, &row.Status)
        if err != nil {
            log.Fatalln(err)
        }
        fmt.Printf("%d. %s\n", row.Id, row.Description)
    }
    err = list.Err()
    if err != nil {
        panic(err.Error())
    }
}

func handleSingleArgCommand() {
    fmt.Println("task is a CLI for managing your TODOs.\n")
    fmt.Println("Usage:")
    fmt.Println("  task [command]\n")
    fmt.Println("Available Commands:")
    fmt.Println("  add         Add a new task to your TODO list")
    fmt.Println("  do          Mark a task on your TODO list as complete")
    fmt.Println("  list        List all of your incomplete tasks\n")
    fmt.Println("Use \"task [command] --help\" for more information about a command.")
    fmt.Println("Use \"exit\" for closing the CLI.\n")
}

func handleDoubleArgCommand(arg string, dbConn *sqlx.DB) {
    if arg == "list" {
        listTask(dbConn)
    } else {
        fmt.Println(InvalidCommandError)
    }
}

func handleTripleArgCommand(arg1 string, arg2 string, dbConn *sqlx.DB) {
    switch arg1 {
    case "add":
        addTask(arg2, dbConn)
    case "do":
        doTask(arg2, dbConn)
    default:
        fmt.Println(InvalidCommandError)
    }
}

func parseString(s string) string {
    return strings.Trim(s," \r\n")
}

func main() {
    handleSingleArgCommand()
    dbConn, err := db.GetSQLDbConnection()
    if err != nil {
        panic(err)
    }
    defer dbConn.Close()
    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("$ ")
        inputString, _ := reader.ReadString('\n')
        args := strings.SplitN(parseString(inputString), " ", 3)
        if args[0] == "task" {
            noOfArguments := len(args)
            switch noOfArguments {
            case 1:
                handleSingleArgCommand()
            case 2:
                handleDoubleArgCommand(args[1], dbConn)
            case 3:
                handleTripleArgCommand(args[1], args[2], dbConn)
            default:
                fmt.Println(InvalidCommandError)
            }
        } else if args[0] == "exit" {
            break
        } else if args[0] == "" {
            continue
        } else {
            fmt.Println(InvalidCommandError)
        }
    }
}
