package replication


type Task string

const (
	EnableVolumeReplicationTaskName  Task = "Enable volume replication"
	DisableVolumeReplicationTaskName Task = "Disable volume replication"
	PromoteVolumeTaskName            Task = "Promote volume"
	DemoteVolumeTaskName             Task = "Demote volume"
	ResyncVolumeTaskName             Task = "Resync volume"
)
