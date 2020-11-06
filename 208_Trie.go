package leetcode_300

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func TrieConstructor() Trie {
	return Trie{
		next:  [26]*Trie{},
		isEnd: false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.next[v-'a'] == nil {
			this.next[v-'a'] = new(Trie)
		} //>>
		this = this.next[v-'a']
	} //>
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if this.next[v-'a'] == nil {
			return false
		} //>>
		this = this.next[v-'a']
	} //>
	return this.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this.next[v-'a'] == nil {
			return false
		} //>>
		this = this.next[v-'a']
	} //>
	return true
}
