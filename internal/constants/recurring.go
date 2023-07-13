package constants

type RecurringOption int

const (
	ReccAutomaticCancel RecurringOption = 0
	ReccDoNotCancel     RecurringOption = 1
	ReccDoNotMakeToken  RecurringOption = 2
)
