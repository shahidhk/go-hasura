package hasura

// Constants defined for SQL response
const (
	TuplesOK  = "TuplesOk"
	CommandOK = "CommandOk"
)

// Query is the Hasura Query object
type Query struct {
	Type string `json:"type"`
	Args `json:"args"`
}

// Args for a query
type Args struct {
	SQL        string        `json:"sql,omitempty"`
	Table      interface{}   `json:"table,omitempty"`
	Columns    interface{}   `json:"columns,omitempty"`
	Where      interface{}   `json:"where,omitempty"`
	OrderBy    interface{}   `json:"order_by,omitempty"`
	Objects    []interface{} `json:"objects,omitempty"`
	Limit      int           `json:"limit,omitempty"`
	Returning  []string      `json:"returning,omitempty"`
	Set        interface{}   `json:"$set,omitempty"`
	OnConflict `json:"on_conflict,omitempty"`
}

// OnConflict argument
type OnConflict struct {
	// action: one of update or ignore
	Action       string   `json:"action"`
	Constraint   string   `json:"constraint,omitempty"`
	ConstraintOn []string `json:"constraint_on,omitempty"`
}

// Bulk is a query in which multiple queries can be executed.
type Bulk struct {
	Type string  `json:"type"`
	Args []Query `json:"args"`
}

// OrderBy is the Hasura order_by expression
type OrderBy struct {
	Column string `json:"column,omitempty"`
	Type   string `json:"type,omitempty"`
	Nulls  string `json:"nulls,omitempty"`
}

// RelatedColumn is a relationship expression in a select query
type RelatedColumn struct {
	Name    string      `json:"name"`
	Columns interface{} `json:"columns,omitempty"`
}

// Error is a Hasura erro response structure;e
type Error struct {
	Path     string         `json:"path"`
	Err      string         `json:"error"`
	Internal *InternalError `json:"internal,omitempty"`
	Message  string         `json:"message,omitempty"`
	Code     string         `json:"code"`
}

// InternalError is thrown when SQL execution fails
type InternalError struct {
	Arguments []string      `json:"arguments"`
	Error     PostgresError `json:"error"`
	Prepared  bool          `json:"prepared"`
	Statement string        `json:"statement"`
}

// PostgresError is the error thrown by Postgres
type PostgresError struct {
	StatusCode  string `json:"status_code"`
	ExecStatus  string `json:"exec_status"`
	Message     string `json:"message"`
	Description string `json:"description"`
	Hint        string `json:"hint"`
}

// Error returns the error message
func (e Error) Error() string {
	return e.Err
}

// RunSQLResponse is the structured response obtained when SQL is executed
type RunSQLResponse struct {
	ResultType string     `json:"result_type"`
	Result     [][]string `json:"result"`
}
