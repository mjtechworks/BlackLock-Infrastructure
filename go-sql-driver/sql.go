package main

import (
_ "github.com/go-sql-driver/mysql"
"database/sql"
"fmt"
)

func main(){
db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/employeedb")
if err != nil {
fmt.Println(err)
}else{
fmt.Println("Connection Established")
}
defer db.Close()
}
