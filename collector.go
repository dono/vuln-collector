package collector

import (
	"database/sql"

	_ "modernc.org/sqlite"

	gocd "github.com/vulsio/go-cve-dictionary/commands"
	gost "github.com/vulsio/gost/cmd"
)

type DB struct {
	*sql.DB
}

func Open(path string) (*DB, error) {
	db, err := sql.Open(`sqlite`, path)
	return &DB{DB: db}, err
}

func DumpDB(sqlitePath string) error {
	// redhat security tracker
	gost.RootCmd.SetArgs([]string{"fetch", "redhat", "--dbpath", sqlitePath})
	if err := gost.RootCmd.Execute(); err != nil {
		return err
	}

	// debian security tracker
	gost.RootCmd.SetArgs([]string{"fetch", "debian", "--dbpath", sqlitePath})
	if err := gost.RootCmd.Execute(); err != nil {
		return err
	}

	// ubuntu security tracker
	gost.RootCmd.SetArgs([]string{"fetch", "ubuntu", "--dbpath", sqlitePath})
	if err := gost.RootCmd.Execute(); err != nil {
		return err
	}

	// nvd database
	gocd.RootCmd.SetArgs([]string{"fetch", "nvd", "--dbpath", sqlitePath})
	if err := gocd.RootCmd.Execute(); err != nil {
		return err
	}

	return nil
}

