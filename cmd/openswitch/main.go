package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	greet()
	ctx := withSignalHandler(context.Background())

	db, err := newDB("localhost", "5432", "postgres", "", "postgres")
	if err != nil {
		panic(err)
	}

	err = NewUpdater(db).Start(ctx)
	if err != nil {
		panic(err)
	}

	//NewAggregator().Start(ctx)
	goodbye()
}

func newDB(host, port, user, password, dbname string) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func greet() {
	fmt.Printf(asciiArt)
	fmt.Println("Version:", Version)
}

func goodbye() {
	fmt.Println("OpenSwitch shutted down")
}

//  returns a context that gets cancelled on SIGINT or SIGTERM.
func withSignalHandler(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Print("\r")
		cancel()
	}()
	return ctx
}

var Version = "dirty"

var asciiArt = `
   ___                  __          _ _       _
  /___\_ __   ___ _ __ / _\_      _(_) |_ ___| |__
 //  // '_ \ / _ \ '_ \\ \\ \ /\ / / | __/ __| '_ \
/ \_//| |_) |  __/ | | |\ \\ V  V /| | || (__| | | |
\___/ | .__/ \___|_| |_\__/ \_/\_/ |_|\__\___|_| |_|
      |_|
`
