package models

// GlobalConfig for nginx
type GlobalConfig struct {
	// nginx user
	User string
	// worker_processes
	WorkerProcesses int
	// is daemon
	Daemon string
}
