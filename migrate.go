package migrate

import (
	"fmt"
	"github.com/mattes/migrate/migrate"
	_ "github.com/mattes/migrate/driver/postgres"
	_ "github.com/mattes/migrate/driver/cassandra"
	"os"
	"strconv"
)

type migrationsData struct {
	path string
	url string
}

func IsExecMigrate() bool {
	return len(os.Args) >= 3 && os.Args[1] == "migrate"
}

func ExecMigrate(dir string, url string) {

	md := &migrationsData {
		path: dir,
		url: url,
	}

	switch os.Args[2] {
	case "up":
		if len(os.Args)>3 {
			n, err := strconv.Atoi(os.Args[3])
			if err != nil {
				md.migrateUp()
			} else {
				md.migrateSync(n)
			}
		} else {
			md.migrateUp()
		}
	case "down":
		if len(os.Args)>3 {
			n, err := strconv.Atoi(os.Args[3])
			if err != nil {
				md.migrateDown()
			} else {
				md.migrateSync(0 - n)
			}
		} else {
			md.migrateDown()
		}
	case "create":
		f, err := migrate.Create(md.url, md.path, os.Args[3])
		if err != nil {
			panic("[MIGRATE] Error:  " + err.Error())
		}
		fmt.Printf("[MIGRATE] New migration files: up: %s, down: %s created;\n", f.UpFile.FileName, f.DownFile.FileName)
	case "version":
		md.migrateVersion()
	case "redo":
		md.migrateRedo()
	default:
		fmt.Println("[MIGRATE] Error: command " + os.Args[2] + " dose not exist;\n")
	}
}

func (md *migrationsData) migrateUp() {
	md.migrateVersion()
	if errs, ok := migrate.UpSync(md.url, md.path); !ok {
		for _, err := range errs {
			fmt.Println("[MIGRATE] Error: " + err.Error())
		}
		panic("[MIGRATE] Up command terminated;\n")
	}
	fmt.Println("[MIGRATE] Up command completed successfully;\n")
	md.migrateVersion()
}

func (md *migrationsData) migrateDown() {
	md.migrateVersion()
	if errs, ok := migrate.DownSync(md.url, md.path); !ok {
		for _, err := range errs {
			fmt.Println("[MIGRATE] Error: " + err.Error())
		}
		panic("[MIGRATE] Down command terminated;\n")
	}
	fmt.Println("[MIGRATE] Down command completed successfully;")
	md.migrateVersion()
}

func (md *migrationsData) migrateSync(n int) {
	md.migrateVersion()

	l := "Up"
	if n < 0 {
		l = "Down"
	}
	if errs, ok := migrate.MigrateSync(md.url, md.path, n); !ok {
		for _, err := range errs {
			fmt.Println("[MIGRATE] Error: " + err.Error())
		}
		panic(fmt.Sprintf("[MIGRATE] %s command terminated;\n", l))
	}
	fmt.Printf("[MIGRATE] %s command completed successfully;", l)
	md.migrateVersion()
}

func (md *migrationsData) migrateRedo() {
	md.migrateVersion()
	if errs, ok := migrate.RedoSync(md.url, md.path); !ok {
		for _, err := range errs {
			fmt.Println("[MIGRATE] Error: " + err.Error())
		}
		panic("[MIGRATE] Redo command terminated;\n")
	}
	fmt.Println("[MIGRATE] Redo command completed successfully;")
	md.migrateVersion()
}

func (md *migrationsData) migrateVersion() {
	version, err := migrate.Version(md.url, md.path)
	if err != nil {
		panic("[MIGRATE] Error: " + err.Error())
	}
	fmt.Printf("[MIGRATE] Current version is %d;\n", version)
}