package server

// CommandPayload the structure expected from any http requested for delivering
// the info associated with the command that needs to be executed
type CommandPayload struct {
	Key         string        `json:"key"`
	CommandName string        `json:"command_name"`
	CommandArgs []interface{} `json:"command_args"`
}

// ResponsePayload structure for the respone containing the command result or error
type ResponsePayload struct {
	ErrorLog      string `json:"error_log"`
	CommandResult string `json:"command_result"`
}
