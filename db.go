/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package main

import (
    "database/sql"
    "fmt"
)

type db sql.DB

func connect(driver, username, pw, host string, port int) (*db, error) {
    DB, err := sql.Open(driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, pw, host, port, "information_schema"))
    if err != nil {
        return nil, err
    }
    return (*db)(DB), nil
}

func close(db *db) {
    (*sql.DB)(db).Close()
}
