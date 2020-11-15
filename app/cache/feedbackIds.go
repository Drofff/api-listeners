package cache

/*
	Cached feedback IDs which are used for deduplication.
	Should be replaced with preCache values after each job run.
 */
var cachedIDs = make([]int64, 3)

/*
	Accumulates values to be cached during a job run.
	Should be moved to cachedIDs once job run is over.
 */
var preCache = make([]int64, 3)

/*
	Checks whether a feedback with the provided ID
	has been already processed by looking up cachedIDs field.
 */
func IsDuplicatedID(id int64) bool {
	for _, cachedID := range cachedIDs {
		if id == cachedID {
			return true
		}
	}
	return false
}

/*
	Cache an id to be used for deduplication purposes on
	the next job run.
 */
func SaveID(id int64) {
	preCache = append(preCache, id)
}

/*
	Moves saved ID values into cachedIDs so that the IDs
	could be used for deduplication on the next job run.
	Replaces old cachedIDs values and clears preCache.
 */
func Submit() {
	cachedIDs = preCache
	preCache = make([]int64, 3)
}
