package main

import (
        "fmt"

        "github.com/valrusu/util/oracle/getdb"
)

func main() {
        db, dtype, err := oracle.GetDB("oracle://vrusu:vali20@toroda02-scan:1521/TRPSUAT")
        fmt.Println(err, db, dtype)
}

