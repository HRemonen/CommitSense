package commit

type CommitInfo struct {
	CommitType                string
	CommitScope               string
	CommitDescription         string
	CommitBody                string
	IsBreakingChange          bool
	BreakingChangeDescription string
}
