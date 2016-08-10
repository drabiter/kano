package utils

import (
	"github.com/drabiter/kano/hummingbird"
)

type RecordLastWatchedDescSorter []hummingbird.Record

func (r RecordLastWatchedDescSorter) Len() int      { return len(r) }
func (r RecordLastWatchedDescSorter) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r RecordLastWatchedDescSorter) Less(i, j int) bool {
	return r[i].LastWatched.Unix() > r[j].LastWatched.Unix()
}
