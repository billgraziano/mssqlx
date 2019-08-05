package mssqlx

import "github.com/pkg/errors"

// DriverPackage sets the default driver for the package
var DriverPackage string

// DBPackage sets the default package for the DB pool
var DBPackage string

const (
	// DriverPackageODBC is the constant for an ODBC connection
	DriverPackageODBC = "odbc"
	// DriverPackageMSSQL is the constant for an MSSQL connection
	DriverPackageMSSQL = "mssql"

	// DBPackageSQL returns a standard connection pool
	DBPackageSQL = "sql"
	// DBPackageSQLX returns a conenction pool of sqlx connections
	DBPackageSQLX = "sqlx"
)

// func init() {
// 	DriverType = DriverTypeMSSQL
// }

// MSSQLX lets you build a new "mssqlx" variable
type MSSQLX struct {
	DriverPackage string
	//DBPackage     string
}

// New returns a new MSSQLX with the driver set
func New(driver string) (*MSSQLX, error) {
	if driver != DriverPackageMSSQL && driver != DriverPackageODBC {
		return nil, errors.Errorf("mssqlx: invalid driver package: %s", driver)
	}
	// if dbpkg != DBPackageSQL && dbpkg != DBPackageSQLX {
	// 	return nil, errors.Errorf("mssqlx: invalid db connection package: ", dbpkg)
	// }
	return &MSSQLX{DriverPackage: driver}, nil
}


