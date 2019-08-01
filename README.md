# mssqlx
SQL Server helper library for GO
```
const (
  MSSQLDriver = "mssql"
  ODBCDriver = "odbc"
)

// Connection holds Native or MSSQL connections
type Connection struct {
  type "mssql", "odbc", "" -- will return either
  ODBCDriver string 
  Various settings...
}

type SQLServer struct {
  ComputerName string
  InstanceName string
  AtAtServer string
  MajorVersion int
  Build string (semver?)
  Edition string 
  Cores, RAM, etc.?
}

type SQLError struct {
  Number
  Severity
  State 
  ...
}

type Session struct {
  dm_exec_connection & dm_exec_session info
}
```
Connection.String("mssql") returns the connection string

mssqlx.Open("mssql", Connection)
mssqlx.OpenX(...) returns a SQLX connection 

Get SQL Server Version and other properties
Get session info (Kerberos, etc.)

mssql.GetVersion(db, Connection)
mssql.GetSession(db, Connection, "mssql|odbc")

// Parse SQL Server errors to number, state, line, etc.
// Parse @@VERSION



