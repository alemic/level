// +build ignore

package level

/*
	An implementation must include these functions:
*/

func newLRUCache(capacity int) UnderlyingCache
func destroyDatabase(name string, o options) error
func repairDatabase(name string, o options) error
func openDatabase(name string, o options) (database, error)
func newOptions() options
func newReadOptions() readOptions
func newWriteOptions() writeOptions
func newWriteBatch() writeBatch
