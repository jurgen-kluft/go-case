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

const (
	StateClean     = 1
	StateAdded     = 2
	StateOutOfDate = 4
	StateModified  = 8
	StateMissing   = 16
	StateUntracked = 32
)
