package main

/* Задача Реализуйте префиксное дерево (Trie)
Префиксное дерево (произносится как «три») — это древовидная структура данных, используемая для эффективного хранения и извлечения ключей из набора строковых данных. Эта структура данных находит различные применения, например, для автодополнения и проверки орфографии.
Реализуйте класс Trie:
Trie() Инициализирует объект префиксного дерева.
void insert(String word) Вставляет строку word в префиксное дерево.
boolean search(String word) Возвращает true, если строка word присутствует в префиксном дереве (т.е. была вставлена ​​ранее), и false в противном случае.
boolean startingWith(String prefix) Возвращает true, если ранее была вставлена ​​строка word с заданным префиксом prefix, и false в противном случае. */

import (
	"fmt"
)

// ─────────────╮
// TrieNode представляет узел в префиксном дереве (Trie)
type TrieNode struct {
	children map[rune]*TrieNode // карта потомков: ключ - символ, значение - указатель на следующий узел
	isEnd    bool               // флаг, указывающий, что узел является концом слова
}

// Trie представляет само префиксное дерево
type Trie struct {
	root *TrieNode // корень дерева
}

// NewTrieNode создает новый узел Trie
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

// Constructor создает новый объект Trie
func Constructor() Trie {
	return Trie{root: NewTrieNode()}
}

// Insert вставляет слово в Trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word { // проходим по каждому символу слова
		if _, ok := node.children[ch]; !ok {
			// если символа нет среди потомков, создаем новый узел
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch] // переходим к следующему узлу
	}
	node.isEnd = true // помечаем конец слова
}

// Search проверяет, есть ли слово в Trie
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			return false // если символа нет, слово не найдено
		}
		node = node.children[ch]
	}
	return node.isEnd // слово найдено только если конец слова помечен
}

// StartsWith проверяет, есть ли слово с заданным префиксом
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, ok := node.children[ch]; !ok {
			return false // если символа нет, префикс не найден
		}
		node = node.children[ch]
	}
	return true // все символы префикса найдены
}

// ─────────────╯
/*
(root)
 ├── a
 │    └── p
 │         ├── p*
 │         │    └── l
 │         │         └── e*
 │         └── e*
 └── b
      └── a
           └── t*
           └── d*
*/
// ─────────────╮╰──>
func main() {
	fmt.Println()
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // true
	fmt.Println(trie.Search("app"))     // false
	fmt.Println(trie.StartsWith("app")) // true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // true
	fmt.Println()
}
