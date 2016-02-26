package workdir

import (
	"github.com/jurgen-kluft/Case/hashing"
)

//	Use 2 WorkTree objects to find changes
//	The states of a file
//	# Clean
//	  - A.Size == B.Size
//	  - A.ModificationTime == A.CreationTime == B.ModificationTime
//	# Modified
//	  - (A.Size != B.Size) or (A.ModificationTime != B.ModificationTime)
//	# Out-Of-Date
//    - (A.ModificationTime == A.CreationTime) != B.ModicationTime
//  # Missing
//    - A has this file - B doesn't have this file
//  # Added
//    - A doesn't have this file - B has this file

type State int

const (
	StateClean     State = 0
	StateAdded     State = 1
	StateOutOfDate State = 2
	StateModified  State = 4
	StateMissing   State = 8
	StateUntracked State = 16
)

type workItemID [16]byte

func CompareWorkItems(latest WorkItem, current WorkItem) (state State) {
	if current.size != latest.size {
		return StateModified
	}
	if current.btime == current.mtime {
		// Not modified
		if current.btime != latest.btime {
			return StateOutOfDate
		}
	}
	return StateClean
}

func DetectChanges(latest []WorkItem, current []WorkItem) (clean, added, outofdate, modified, missing, untracked []WorkItem) {
	hasher := hashing.NewHasher(hashing.SHA256)
	latestItems := make(map[workItemID]int)
	for i, latestitem := range latest {
		latestItemID := workItemID{}
		hasher.Hash([]byte(latestitem.filepath), latestItemID[:])
		latestItems[latestItemID] = i
	}

	clean = make([]WorkItem, 8)
	added = make([]WorkItem, 8)
	outofdate = make([]WorkItem, 8)
	modified = make([]WorkItem, 8)
	missing = make([]WorkItem, 8)
	untracked = make([]WorkItem, 8)

	for _, currentItem := range current {
		currentItemID := workItemID{}
		hasher.Hash([]byte(currentItem.filepath), currentItemID[:])

		latestItemIndex, ok := latestItems[currentItemID]
		if ok {
			// Remove it from the map so that later we can detect 'missing' items
			delete(latestItems, currentItemID)

			// Obtain the latestItem and compare it with the currentItem to retrieve the State
			latestItem := latest[latestItemIndex]
			state := CompareWorkItems(latestItem, currentItem)
			switch state {
			case StateClean:
				clean = append(clean, currentItem)
			case StateOutOfDate:
				outofdate = append(outofdate, currentItem)
			case StateModified:
				modified = append(modified, currentItem)
			}
		} else {
			// This current item doesn't exist as 'latest' so it hasn't been 'added
			untracked = append(untracked, currentItem)
		}
	}

	// Whatever there is left in the map are items that are 'Missing'
	for _, i := range latestItems {
		missingItem := latest[i]
		missing = append(missing, missingItem)
	}
	return clean, added, outofdate, modified, missing, untracked
}
