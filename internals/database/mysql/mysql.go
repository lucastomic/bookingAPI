package mysql

func ExecStmt(stmt string, params []any) error {
	db := GetInstance()
	sqlStmt, err := db.Prepare(stmt)
	if err == nil {
		defer sqlStmt.Close()
		_, err = sqlStmt.Exec(params...)
	}
	return err
}
