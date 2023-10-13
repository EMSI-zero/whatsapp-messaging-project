package boot

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"whatapp-messaging/internal/logger"
	"whatapp-messaging/internal/repository/dbrepo"
	"whatapp-messaging/internal/server/httpserver"

	"github.com/joho/godotenv"
)

var env_dir_flag string

func Boot() error {
	if err := parseFlags(); err != nil {
		return err
	}

	if err := initEnvFile(".env"); err != nil {
		return err
	}

	if err := logger.InitLog(env_dir_flag); err != nil {
		return err
	}

	if err := dbrepo.NewDBConn(); err != nil {
		return err
	}

	if err := httpserver.StartServer(); err != nil {
		return err
	}

	return nil
}

func parseFlags() error {
	flags := flag.NewFlagSet("", flag.ContinueOnError)
	flags.StringVar(&env_dir_flag, "env_dir", ".", "")

	if err := flags.Parse(os.Args[1:]); err != nil {
		return err
	}

	return nil
}

func initEnvFile(fileName string) error {
	fullPath := filepath.Join(env_dir_flag, fileName)
	err := godotenv.Load(fullPath)
	if err != nil {
		log.Print(err)
		return fmt.Errorf("error loading .env file")
	}

	return nil
}
