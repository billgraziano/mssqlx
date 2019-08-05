# mssqlx
SQL Server helper library for GO

This is a work-in-progress.  Please don't use.
```
type Server struct {
  ComputerName string
  InstanceName string
  AtAtServer string
  AtAtVersion
  MajorVersion int
  Build string (semver?)
  Edition string 
  Cores, RAM, etc.?
}

type Error struct {
  Number
  Severity
  State 
  Message
  RawMessage
  ...
}

func ParseError(e string) (Error, error) {
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

// Define interfaces that I need for the DB
mssql.GetVersion(db, Connection)
mssql.GetSession(db, Connection, "mssql|odbc")

// Parse SQL Server errors to number, state, line, etc.
// Parse @@VERSION
