package scheduler

func Init() {
	scheduleUpdates()
	cleanStaleInformed()
	cleanFWCode()
}
