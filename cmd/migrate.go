package cmd

import (
	"fmt"
	"log"

	"github.com/rodiond26/gomigrator/config"
	"github.com/rodiond26/gomigrator/db"
	"github.com/rodiond26/gomigrator/logger"
	"github.com/rodiond26/gomigrator/migrations"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	c   *config.Config
	l   *zap.Logger
	err error
)

func init() {
	c, err = config.GetConfig()
	if err != nil {
		log.Fatalf("unable to load config: %v", err)
	}
	fmt.Printf(">>> config = [%+v]\n", c)

	l, err = logger.GetLogger(c)
	if err != nil {
		log.Fatalf("unable to load logger: %v", err)
	}
	fmt.Printf(">>> logger = [%+v]\n", l)

	upCmd.Flags().IntP("step", "s", 0, "Number of migrations to execute")
	downCmd.Flags().IntP("step", "s", 0, "Number of migrations to execute")
	redoCmd.Flags().IntP("step", "s", 0, "Number of migrations to execute")

	rootCmd.AddCommand(createCmd, upCmd, downCmd, redoCmd, statusCmd)
}

var createCmd = &cobra.Command{
	Use:                   "create <migration_name>",
	Short:                 "create a new empty migrations file",
	DisableFlagsInUseLine: true,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("Need migration name")
			return
		}
		if err = migrations.Create(args[0]); err != nil {
			fmt.Println("Unable to create migration", err.Error())
			return
		}
	},
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "run up migrations",
	Run: func(cmd *cobra.Command, args []string) {

		step, err := cmd.Flags().GetInt("step")
		if err != nil {
			fmt.Println("Unable to read flag `step`")
			return
		}

		db := db.NewDB(c)

		migrator, err := migrations.Init(db)
		if err != nil {
			fmt.Println("Unable to fetch migrator")
			return
		}

		err = migrator.Up(step)
		if err != nil {
			fmt.Println("Unable to run `up` migrations")
			return
		}
	},
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "run down migrations",
	Run: func(cmd *cobra.Command, args []string) {

		step, err := cmd.Flags().GetInt("step")
		if err != nil {
			fmt.Println("Unable to read flag `step`")
			return
		}

		db := db.NewDB(c)

		migrator, err := migrations.Init(db)
		if err != nil {
			fmt.Println("Unable to fetch migrator")
			return
		}

		err = migrator.Down(step)
		if err != nil {
			fmt.Println("Unable to run `down` migrations")
			return
		}
	},
}

var redoCmd = &cobra.Command{
	Use:   "redo",
	Short: "run redo migrations",
	Run: func(cmd *cobra.Command, args []string) {

		step, err := cmd.Flags().GetInt("step")
		if err != nil {
			fmt.Println("Unable to read flag `step`")
			return
		}

		db := db.NewDB(c)

		migrator, err := migrations.Init(db)
		if err != nil {
			fmt.Println("Unable to fetch migrator")
			return
		}

		err = migrator.Down(step)
		if err != nil {
			fmt.Println("Unable to run `down` migrations")
			return
		}
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "display status of each migrations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(">>> status = [%+v]\n", c)

		db := db.NewDB(c)

		migrator, err := migrations.Init(db)
		if err != nil {
			fmt.Println("Unable to fetch migrator")
			return
		}

		if err := migrator.MigrationStatus(); err != nil {
			fmt.Println("Unable to fetch migration status")
			return
		}

		return
	},
}
