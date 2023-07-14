package main

import (
        "fmt"

        "github.com/valrusu/util/oracle/getdb"
)

func main() {
        db, dtype, err := oracle.GetDB("oracle://user:password@server:port/service")
        fmt.Println(err, db, dtype)
}

