package ipc

// type 'Command' is a DTO struct for IPC between todo daemon <-> todo program
type Command struct {
	Command string `json:"command"`
	TaskID  string `json:"task_id,omitempty"`
}
