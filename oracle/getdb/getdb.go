package oracle

import (
	"database/sql"
	"errors"
	"strings"

	_ "github.com/godror/godror"
	_ "github.com/sijms/go-ora/v2"
)

func GetDBGodror(connStr string) (pdb *sql.DB, perr error) {
	defer func() {
		if rec := recover(); rec != nil {
			pdb = nil

			switch trec := rec.(type) {
			case string:
				perr = errors.New(trec)
			case error:
				perr = trec
			default:
				perr = errors.New("Unknown panic in getOraDBGodror")
			}
		}
	}()

	db, err := sql.Open("godror", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	//var check string
	//err = db.QueryRow("select 'ok' from dual").Scan(&check)
	//if err != nil {
	//	return nil, err
	//}

	return db, err
}

func GetDBGoora(connStr string) (pdb *sql.DB, perr error) {
	defer func() {
		if rec := recover(); rec != nil {
			pdb = nil

			switch trec := rec.(type) {
			case string:
				perr = errors.New(trec)
			case error:
				perr = trec
			default:
				perr = errors.New("Unknown panic in getOraDBGo_ora")
			}
		}
	}()

	db, err := sql.Open("oracle", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	//	var check string
	//	err = db.QueryRow("select 'ok' from dual").Scan(&check)
	//if err != nil {
	//	return nil, err
	//}

	return db, err
}

func getOraDBGodror(connStr string) (pdb *sql.DB, perr error) {
	defer func() {
		if rec := recover(); rec != nil {
			pdb = nil

			switch trec := rec.(type) {
			case string:
				perr = errors.New(trec)
			case error:
				perr = trec
			default:
				perr = errors.New("Unknown panic in getOraDBGodror")
			}
		}
	}()

	db, err := sql.Open("godror", connStr)
	if err != nil {
		return nil, err
	}
	return db, err
}

func getOraDBGoora(connStr string) (pdb *sql.DB, perr error) {
	defer func() {
		if rec := recover(); rec != nil {
			pdb = nil

			switch trec := rec.(type) {
			case string:
				perr = errors.New(trec)
			case error:
				perr = trec
			default:
				perr = errors.New("Unknown panic in getOraDBGoora")
			}
		}
	}()

	db, err := sql.Open("oracle", connStr)
	if err != nil {
		return nil, err
	}
	return db, err
}

// getDB returns a DB connection by trying both Oracle drivers, godror and goora, and a string containing which one of them was used
func GetDB(connStr string) (pdb *sql.DB, ptype string, perr error) {

	var (
		order string = "goora godror"
		db    *sql.DB
		err   error
	)

	// goora does not work with OS-auth or usr/pwd only
	if strings.HasPrefix(connStr, "/") || strings.Count(connStr, "/") == 1 {
		order = "godror goora"
	}

	for _, x := range strings.Fields(order) {

		switch x {
		case "godror":
			db, err = GetDBGodror(connStr)
		case "goora":
			db, err = GetDBGoora(connStr)
		}

		if err == nil {
			return db, x, nil
		}
	}

	return nil, "none", err
}
