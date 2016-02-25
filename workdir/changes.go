package workdir

//	Use 2 WorkTree objects to find changes
//	The states of a file
//	# Clean
//	  - A.Size == B.Size
//	  - A.ModificationTime == B.ModificationTime
//	# Modified
//	  - A.Size != B.Size
//    or
//	  - A.ModificationTime != B.ModificationTime
//	# Out-Of-Date
//
//  # Missing
//    - A has this file - B doesn't have this file
//  # Added
//    - A doesn't have this file - B has this file

type State int

const (
	StateClean     State = 1
	StateAdded     State = 2
	StateOutOfDate State = 4
	StateModified  State = 8
	StateMissing   State = 16
	StateUntracked State = 32
)

func DetectChanges(latest WorkTree, current WorkTree) map[WorkItemID]State {
	return map[WorkItemID]State{}
}
