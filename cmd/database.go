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

	/* sbc trunk */
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS sbc_interface_list (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			if_name VARCHAR(50) NOT NULL,
			mac VARCHAR(40) NOT NULL,
			created DATE NULL,
			username VARCHAR(40) NOT NULL
		);	
	`		
    _, err = db.Exec(sqlStmt)
    checkErr(err)
        
    /* sbc rtpengine */
    sqlStmt = `
		CREATE TABLE IF NOT EXISTS sbc_ip_address (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			grp INTEGER DEFAULT 1 NOT NULL,
			ip_addr VARCHAR(50) NOT NULL,
			mask INTEGER DEFAULT 32 NOT NULL,
			tag VARCHAR(64),
			created DATE NULL,
			username VARCHAR(40) NOT NULL
		);
	`		
    _, err = db.Exec(sqlStmt)
	checkErr(err)
	
	/* sbc */
    sqlStmt = `
		CREATE TABLE IF NOT EXISTS sbc_trunk (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(50) NOT NULL,
			grp INTEGER DEFAULT 1 NOT NULL,
			grp_sig_untrust_dst VARCHAR(100) NOT NULL,
			grp_sig_untrust_src VARCHAR(100) NOT NULL,
			grp_sig_trust_dst VARCHAR(100) NOT NULL,
			grp_sig_trust_src VARCHAR(100) NOT NULL,
			grp_media_trust VARCHAR(100) NOT NULL,
			grp_media_untrust VARCHAR(100) NOT NULL,
			description VARCHAR(200) NOT NULL,
			topohide SMALLINT NOT NULL, 
			type SMALLINT NOT NULL,
			active SMALLINT NOT NULL,
			created DATE NULL,
			username VARCHAR(40) NOT NULL
		);	
	`		
    _, err = db.Exec(sqlStmt)
    checkErr(err)
			  
	/* sbc */
    sqlStmt = `
		CREATE TABLE IF NOT EXISTS sbc_signaling (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			ip_addr VARCHAR(50) NOT NULL,
			port INTEGER DEFAULT 5060 NOT NULL,
			proto SMALLINT NOT NULL,
			description VARCHAR(100) NOT NULL,
			created DATE NULL,
			username VARCHAR(40) NOT NULL
		);		
	`		
    _, err = db.Exec(sqlStmt)
	checkErr(err)
	
	/* sbc */
    sqlStmt = `
		CREATE TABLE IF NOT EXISTS sbc_media (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			ip_addr VARCHAR(50) NOT NULL,
			description VARCHAR(100) NOT NULL,
			created DATE NULL,
			username VARCHAR(40) NOT NULL
		);`		
    _, err = db.Exec(sqlStmt)
	checkErr(err)
	
    /* KAMAILIO/OPENSIPS */
    /* version */      
    sqlStmt = `
        	CREATE TABLE IF NOT EXISTS version (
		    table_name VARCHAR(32) NOT NULL,
		    table_version INTEGER DEFAULT 0 NOT NULL,
		    CONSTRAINT version_table_name_idx UNIQUE (table_name)
	);

	INSERT INTO version (table_name, table_version) values ('version','1');
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

		INSERT INTO version (table_name, table_version) values ('dispatcher','4');
	`
        _, err = db.Exec(sqlStmt)
        checkErr(err)
        
        /* trusted */      
        sqlStmt = `
        	CREATE TABLE IF NOT EXISTS trusted (
  		  id INTEGER PRIMARY KEY NOT NULL,
  		  src_ip VARCHAR(50) NOT NULL,
  		  proto VARCHAR(4) NOT NULL,
  		  from_pattern VARCHAR(64) DEFAULT NULL,
  		  ruri_pattern VARCHAR(64) DEFAULT NULL,
  		  tag VARCHAR(64),
  		  priority INTEGER DEFAULT 0 NOT NULL
  		  );

  		  CREATE INDEX trusted_peer_idx ON trusted (src_ip);

  		  INSERT INTO version (table_name, table_version) values ('trusted','6');
	`
        _, err = db.Exec(sqlStmt)
        checkErr(err)
                	
	/* address */      
        sqlStmt = `        
	       	CREATE TABLE IF NOT EXISTS address (
		    id INTEGER PRIMARY KEY NOT NULL,
		    grp INTEGER DEFAULT 1 NOT NULL,
		    ip_addr VARCHAR(50) NOT NULL,
		    mask INTEGER DEFAULT 32 NOT NULL,
		    port SMALLINT DEFAULT 0 NOT NULL,
		    tag VARCHAR(64)
		);

		INSERT INTO version (table_name, table_version) values ('address','6');
	`
        _, err = db.Exec(sqlStmt)
        checkErr(err)
        
        /* pl_pipelimit  */      
        sqlStmt = `     
		CREATE TABLE IF NOT EXISTS pl_pipes (
		    id INTEGER PRIMARY KEY NOT NULL,
		    pipeid VARCHAR(64) DEFAULT '' NOT NULL,
		    algorithm VARCHAR(32) DEFAULT '' NOT NULL,
		    plimit INTEGER DEFAULT 0 NOT NULL
		);

		INSERT INTO version (table_name, table_version) values ('pl_pipes','1');
	`
        _, err = db.Exec(sqlStmt)
        checkErr(err)
        
        /* domain  */      
        sqlStmt = `     
        	CREATE TABLE IF NOT EXISTS domain (
  		  id INTEGER PRIMARY KEY NOT NULL,
  		  domain VARCHAR(64) NOT NULL,
  		  did VARCHAR(64) DEFAULT NULL,
  		  last_modified TIMESTAMP WITHOUT TIME ZONE DEFAULT '2000-01-01 00:00:01' NOT NULL,
  		  CONSTRAINT domain_domain_idx UNIQUE (domain)
  		);
  		
  		INSERT INTO version (table_name, table_version) values ('domain','2');
	`
        _, err = db.Exec(sqlStmt)
        checkErr(err)
        
        /* domain_attrs  */      
        sqlStmt = `             	
        	CREATE TABLE IF NOT EXISTS domain_attrs (
		    id INTEGER PRIMARY KEY NOT NULL,
		    did VARCHAR(64) NOT NULL,
		    name VARCHAR(32) NOT NULL,
		    type INTEGER NOT NULL,
		    value VARCHAR(255) NOT NULL,
		    last_modified TIMESTAMP WITHOUT TIME ZONE DEFAULT '2000-01-01 00:00:01' NOT NULL
		);

		CREATE INDEX domain_attrs_domain_attrs_idx ON domain_attrs (did, name);
		INSERT INTO version (table_name, table_version) values ('domain_attrs','1');
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
	Long: `A longer description.`,
	
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

