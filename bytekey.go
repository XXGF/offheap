package offheap

import xxh64 "github.com/glycerine/xxhash-64"

// the byte-key (BK) interface to the hash table

// use the 64-bit implimentation of XXHash for speed.
// see
//   https://github.com/OneOfOne/xxhash (github.com/glycerine/xxhash-64 version-locks)
//   http://fastcompression.blogspot.com/2014/07/xxhash-wider-64-bits.html
//
var xxHasher64 = xxh64.New64()

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*

func (t *HashTable) InsertBK(bytekey []byte, value interface{}) bool {
	xxHasher64.Reset()
	min := minimum(len(key_t{}), len(bytekey))
	_, err := xxHasher64.Write(bytekey[:min])
	if err != nil {
		panic(err)
	}
	hashkey := xxHasher64.Sum64()
	cell, ok := t.Insert(hashkey)
	copy(cell.ByteKey[:], bytekey)
	cell.SetValue(value)
	return ok
}

func (t *HashTable) LookupBK(bytekey []byte) (Val_t, bool) {
	xxHasher64.Reset()
	min := minimum(len(key_t{}), len(bytekey))
	_, err := xxHasher64.Write(bytekey[:min])
	if err != nil {
		panic(err)
	}
	hashkey := xxHasher64.Sum64()
	cell := t.Lookup(hashkey)
	if cell == nil {
		return Val_t{}, false
	}
	return cell.Value, true
}

func (t *HashTable) DeleteBK(bytekey key_t) bool {
	xxHasher64.Reset()
	min := minimum(len(key_t{}), len(bytekey))
	_, err := xxHasher64.Write(bytekey[:min])
	if err != nil {
		panic(err)
	}
	hashkey := xxHasher64.Sum64()
	cell := t.Lookup(hashkey)
	if cell == nil {
		return false
	}

	t.DeleteCell(cell)
	return true
}

func (t *HashTable) InsertStringKey(strkey string, value interface{}) bool {
	xxHasher64.Reset()

	min := minimum(len(key_t{}), len(strkey))
	var bytekey key_t
	copy(bytekey[:], []byte(strkey))
	_, err := xxHasher64.Write(bytekey[:min])
	if err != nil {
		panic(err)
	}
	hashkey := xxHasher64.Sum64()
	cell, ok := t.Insert(hashkey)
	cell.ByteKey = bytekey
	cell.SetValue(value)
	//fmt.Printf("assigned value : '%v'  to key: '%v', with strkey: '%v'\n", value, hashkey, strkey)

	return ok
}

func HashString(s string) uint64 {
	xxHasher64.Reset()
	u := []byte(s)
	min := minimum(len(key_t{}), len(u))
	_, err := xxHasher64.Write(u[:min])
	if err != nil {
		panic(err)
	}
	hashkey := xxHasher64.Sum64()
	return hashkey
}

func (t *HashTable) LookupStringKey(strkey string) (capn.Object, bool) {
	xxHasher64.Reset()
	min := minimum(len(key_t{}), len(strkey))
	var bytekey key_t
	copy(bytekey[:], []byte(strkey))
	_, err := xxHasher64.Write(bytekey[:min])
	if err != nil {
		panic(err)
	}
	hashkey := xxHasher64.Sum64()
	cell := t.Lookup(hashkey)
	if cell == nil {
		return capn.Object{}, false
	}
	return cell.Value, true
}

func (t *HashTable) DeleteStringKey(strkey string) bool {
	xxHasher64.Reset()
	min := minimum(len(key_t{}), len(strkey))
	var bytekey key_t
	copy(bytekey[:], []byte(strkey))
	_, err := xxHasher64.Write(bytekey[:min])
	if err != nil {
		panic(err)
	}
	hashkey := xxHasher64.Sum64()
	cell := t.Lookup(hashkey)
	if cell == nil {
		return false
	}

	t.DeleteCell(cell)
	return true
}

func (t *HashTable) DumpStringKey() {

	fmt.Printf(" DumpStringKey(): (table ArraySize: %d\n", t.ArraySize)
	for it := NewIterator(t); it.Cur != nil; it.Next() {
		fmt.Printf("  '%v' -> %v\n", string(it.Cur.ByteKey[:]), it.Cur.Value)
	}

}
*/
