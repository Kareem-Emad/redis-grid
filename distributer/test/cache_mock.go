package distributertest

type mockCache struct {
	ExecutionCount       int
	GlobalExecutionCount *int
}

// InitWithDefault intialize a connection to redis server
func (rc *mockCache) InitWithDefault() error {
	return nil
}

// Init intialize a connection to redis server
func (rc *mockCache) Init(connURL string) error {
	return nil
}

// ExecuteRawCommand sends a command string to redis
func (rc *mockCache) ExecuteRawCommand(commandName string, args ...interface{}) (string, error) {
	rc.ExecutionCount++
	*rc.GlobalExecutionCount++

	return "", nil
}
