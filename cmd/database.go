// Copyright © 2019 Alexandr Dubovikov <alexandr.dubovikov@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/spf13/cobra"
)


func checkErr(err error) {
	if err != nil {
	    log.Printf("%q:\n", err)
        }
}

func initSql(databaseName string) {
	//var created time.Time

	db, err := sql.Open("sqlite3", databaseName)
        checkErr(err)
	defer db.Close()

	sqlStmt := `
		CREATE TABLE IF NOT EXISTS userinfo (uid INTEGER PRIMARY KEY AUTOINCREMENT,username VARCHAR(64) NULL,departname VARCHAR(64) NULL,created DATE NULL);
	`		
        _, err = db.Exec(sqlStmt)
        checkErr(err)
          
  	/* dispatcher */      
        sqlStmt = `
        	CREATE TABLE IF NOT EXISTS dispatcher (
		    id INTEGER PRIMARY KEY NOT NULL,
		    setid INTEGER DEFAULT 0 NOT NULL,
		    destination VARCHAR(192) DEFAULT '' NOT NULL,
		    flags INTEGER DEFAULT 0 NOT NULL,
		    priority INTEGER DEFAULT 0 NOT NULL,
		    attrs VARCHAR(128) DEFAULT '' NOT NULL,
		    description VARCHAR(64) DEFAULT '' NOT NULL
		);			
	`
        _, err = db.Exec(sqlStmt)
        checkErr(err)
        
        //log.Printf("%q: %s\n", err, sqlStmt)


        // insert
        //stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
        //checkErr(err)
        
}


// databaseCmd represents the database command
var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {			

		if cmd.Flags().Changed("init") {
		
			initName, _:= cmd.Flags().GetString("init")
			fmt.Println("database called: "+ initName)
			initSql(initName)
		}

	},
}

func init() {
	rootCmd.AddCommand(databaseCmd)
	databaseCmd.Flags().StringP("init", "i", "sbc-database.db", "Init new database")

	//initSql();

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// databaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// databaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

