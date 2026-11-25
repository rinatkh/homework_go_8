package hashtable

// HashTable — собственная таблица с цепочками. В этом файле встроенная map
// запрещена. В соседних задачах frequencies/duplicate map разрешена.
// hasher==nil означает детерминированную FNV-1a; иначе данная функция определяет hash.
type entry struct {
	key   string
	value int
	next  *entry
}
type HashTable struct {
	buckets []*entry
	size    int
	hasher  func(string) uint64
}

func New(capacity int, hasher func(string) uint64) *HashTable { // TODO: 06.1, минимум 4 buckets.
	return &HashTable{}
}
func (t *HashTable) Set(key string, value int) { // TODO: update, collision, grow при load>0.75.
}
func (t *HashTable) Get(key string) (int, bool) { return 0, false } // TODO.
func (t *HashTable) Delete(key string) bool     { return false }    // TODO.
func (t *HashTable) Len() int                   { return 0 }        // TODO.
