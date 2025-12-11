package task

import "time"

type Status string

const (
    StatusAssigned  Status = "ASIGNADO"
    StatusInProcess Status = "EN_PROCESO"
    StatusCompleted Status = "COMPLETADO"
    StatusExpired   Status = "VENCIDO"
)

type Task struct {
    ID          string
    Title       string
    Description string
    DueDate     time.Time
    AssignedTo  string
    Status      Status
    Comments    []string
}
