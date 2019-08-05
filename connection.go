package mssqlx

import (
	"database/sql"
	"fmt"
	"net/url"

	//_ "github.com/denisenkom/go-mssqldb"
	"github.com/pkg/errors"
)

// Connection holds a connection string
type Connection struct {
	DriverPackage string // "mssql|odbc"
	//DBPackage     string // "sql|sqlx"

	FQDN     string // host:port
	Computer string
	Instance string
	Port     int32
	User     string
	Password string
	Database string
	AppName  string

	ODBCDriver string
}

// String returns a connection string
func (conn *Connection) String() (string, error) {
	if conn.DriverPackage == "" {
		conn.DriverPackage = DriverPackageMSSQL
	}

	switch conn.DriverPackage {
	case DriverPackageMSSQL:
		return conn.MSSQLString()
	case DriverPackageODBC:
		return conn.ODBCString()
	default:
		return "", errors.Errorf("mssqlx: unknown driver package: %s", DriverPackage)
	}
}

// Open a Connection returning a *sql.DB
func (conn *Connection) Open() (*sql.DB, error) {
	cxn, err := conn.String()
	if err != nil {
		return nil, err
	}
	return sql.Open(conn.DriverPackage, cxn)
}

// ODBCString returns a connection string an an "odbc" connection
func (conn *Connection) ODBCString() (string, error) {
	return "", errors.New("not implemented")
}

// MSSQLString returns a connection string for an "mssql" connection
func (conn *Connection) MSSQLString() (string, error) {
	if conn.Computer == "" {
		return "", errors.New("mssqlx: computer is required")
	}
	query := url.Values{}
	if conn.AppName != "" {
		query.Add("app name", conn.AppName)
	}

	if conn.Database != "" {
		query.Add("database", conn.Database)
	}

	u := &url.URL{
		Scheme:   "sqlserver",
		RawQuery: query.Encode(),
	}
	if conn.Port != 0 {
		u.Host = fmt.Sprintf("%s:%d", conn.Computer, conn.Port)
	} else {
		u.Host = conn.Computer
	}
	if conn.Instance != "" {
		u.Path = conn.Instance
	}
	if conn.User != "" || conn.Password != "" {
		u.User = url.UserPassword(conn.User, conn.Password)
	}
	return u.String(), nil
}
