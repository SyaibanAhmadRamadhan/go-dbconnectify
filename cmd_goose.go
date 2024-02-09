package go_dbconnectify

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pressly/goose/v3"
)

// ConnForMigrate opens a new database connection using the supplied driver and URL.
// It uses the goose library to open the connection and panics in case of any errors.
func ConnForMigrate(url string, driver string) *sql.DB {
	db, err := goose.OpenDBWithDriver(driver, url)
	if err != nil {
		panic(err)
	}

	return db
}

// Migrate accepts a database connection (db), command to perform migration (cmd), argument for the command (arg), a directory (dir) in which migration files are located and a boolean value (seq)
// to specify whether the migrations need to be executed sequentially. The function closes the database connection after performing the operations.
func Migrate(cmd string, arg string, dir string, seq bool, db *sql.DB) {
	goose.SetSequential(seq)
	if cmd == "" && arg == "" {
		cmd, arg = cmdMigrateMaster()
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("failed close goose db")
		}
	}()

	var err error

	switch cmd {
	case "up":
		err = goose.Up(db, dir)
	case "up-by-one":
		err = goose.UpByOne(db, dir)
	case "status":
		err = goose.Status(db, dir)
	case "down":
		err = goose.Down(db, dir)
	case "redo":
		err = goose.Redo(db, dir)
	case "reset":
		err = goose.Reset(db, dir)
	case "version":
		err = goose.Version(db, dir)
	case "fix":
		err = goose.Fix(dir)
	case "create":
		err = goose.Create(db, dir, arg, "sql")
	case "create-go":
		err = goose.Create(db, dir, arg, "go")
	case "up-to":
		argInt, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		err = goose.UpTo(db, dir, int64(argInt))
	case "down-to":
		argInt, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		err = goose.DownTo(db, dir, int64(argInt))
	default:
		err = goose.Status(db, dir)
	}
	if err != nil {
		panic(err)
	}
}

// cmdMigrateMaster provides an interface to the user to enter migration commands.
// It returns a command and an argument as string.
// Calling this function prints a list of possible commands and waits for the user to enter a command.
// It then validates the input and returns a command and possibly an argument.
func cmdMigrateMaster() (cmd string, arg string) {
	fmt.Println("\nSelect command migrate:")
	fmt.Println("up\t\t\tMigrate the DB to the most recent arg available")
	fmt.Println("up-by-one\t\tMigrate the DB up by 1")
	fmt.Println("up-to VERSION\t\tMigrate the DB to a specific VERSION")
	fmt.Println("down\t\t\tRoll back the migration by 1. desc migration priority")
	fmt.Println("down-to VERSION\t\tRoll back the DB to a specific VERSION")
	fmt.Println("redo\t\t\tRe-run the latest migration")
	fmt.Println("reset\t\t\tRoll back all migrations")
	fmt.Println("status\t\t\tDump the migration status for the current DB")
	fmt.Println("version\t\t\tPrint the current version of the database")
	fmt.Println("create NAME\t\tCreates new migration file with the increment current migration")
	fmt.Println("create-go NAME\t\tCreates new migration file go with the increment current migration")
	fmt.Print("Masukkan perintah yang ingin dijalankan any key for stop: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")
	cmd = parts[0]

	switch cmd {
	case "status", "up", "down", "reset", "up-by-one", "redo", "version":
		fmt.Println("\n=== your command valid. start migration ===")
		return cmd, arg
	case "create", "create-go", "up-to", "down-to":
		if len(parts) < 2 {
			panic("argument 2 is empty")
		}
		arg = parts[1]
		return cmd, arg
	default:
		panic("invalid command")
	}
}
