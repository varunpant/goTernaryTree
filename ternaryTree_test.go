package goTernaryTree

import (
	"bufio"
	"os"
	"testing"
)

func TestCreateTernaryTree(t *testing.T) {
	ttree := New()
	inputKeys := [12]string{"aback", "abacus", "abalone", "abandon", "abase", "abash", "abate", "abbas", "abbe", "abbey", "abbot", "Abbott"}

	for _, key := range inputKeys {
		ttree.add(key, key)
	}

	if ttree == nil {
		t.Error("Failed to create terenaryTree")
	}

	if ttree.size != 12 {
		t.Errorf("Expected size to be %d, but got %d", 12, ttree.size)
	}

	for _, key := range inputKeys {

		value, _ := ttree.get(key)
		if key != value {
			t.Errorf("Failed to get the valid value for key %s from tree, recieved %s", key, value)
		}
	}
}

func TestAddingDuplicateKeysInTree(t *testing.T) {
	ttree := New()
	ttree.add("Key1", "value1");
	if ttree.size != 1 {
		t.Errorf("Expected size to be %d, but got %d", 1, ttree.size)
	}
	val, _ := ttree.get("Key1")
	if val != "value1" {
		t.Errorf("Failed to get the valid value ,expected %s, recieved %s", "value1", val)
	}
	ttree.add("Key1", "value1");
	if ttree.size != 1 {
		t.Errorf("Expected size to be %d, but got %d", 1, ttree.size)
	}
}

func TestAddingNullValuesInTree(t *testing.T) {
	ttree := New()
	ttree.add("Key1", nil);
	if ttree.size != 1 {
		t.Errorf("Expected size to be %d, but got %d", 1, ttree.size)
	}
	val, _ := ttree.get("Key1")

	if val != nil {
		t.Errorf("Expected value to be nil, but got %d", val)
	}

}

func TestAddingNullKeysInTree(t *testing.T) {
	ttree := New()
	ttree.add("", "Empty Key");
	if ttree.size != 0 {
		t.Errorf("Expected size to be %d, but got %d", 1, ttree.size)
	}

	ttree.add(" ", "White space");
	if ttree.size != 0 {
		t.Errorf("Expected size to be %d, but got %d", 1, ttree.size)
	}

}

func TestPrefixSearch(t *testing.T) {
	ttree := New()
	inputKeys := [12]string{"aback", "abacus", "abalone", "abandon", "abase", "abash", "abate", "abbas", "abbe", "abbey", "abbot", "Abbott"}

	for _, key := range inputKeys {
		ttree.add(key, key)
	}

	hits := ttree.prefixMatch("aba")

	if len(hits) != 7 {
		t.Errorf("incorect prefix search hits from tree, expected 7, but recieved %d", len(hits))
	}

	expectedHits := [7]string{"aback", "abacus", "abalone", "abandon", "abase", "abash", "abate"}
	for idx, hit := range hits {
		if hit != expectedHits[idx] {
			t.Errorf("expected search hit %s, but got %s", hit, expectedHits[idx])
		}
	}
}

func TestWildcardMatch(t *testing.T) {
	ttree := New()
	inputKeys := [12]string{"aback", "abacus", "abalone", "abandon", "abase", "abash", "abate", "abbas", "abbe", "abbey", "abbot", "abbott"}

	for _, key := range inputKeys {
		ttree.add(key, key)
	}

	{
		matches := ttree.wildcardMatch("...cus")
		if len(matches) != 1 {
			t.Errorf("Incorrect match count, expected %d, but got %d", 1, len(matches))
		}
		if matches[0] != "abacus" {
			t.Errorf("Incorrect match, expected %s, but got %s", "abacus", matches[0])
		}
	}

	{
		matches := ttree.wildcardMatch("....y")
		if len(matches) != 1 {
			t.Errorf("Incorrect match count, expected %d, but got %d", 1, len(matches))
		}
		if matches[0] != "abbey" {
			t.Errorf("Incorrect match, expected %s, but got %s", "abbey", matches[0])
		}
	}

	{
		matches := ttree.wildcardMatch("a.b.t.")
		if len(matches) != 1 {
			t.Errorf("Incorrect match count, expected %d, but got %d", 1, len(matches))
		}
		if matches[0] != "abbott" {
			t.Errorf("Incorrect match, expected %s, but got %s", "abbott", matches[0])
		}
	}

	{
		matches := ttree.wildcardMatch("aba...")
		if len(matches) != 1 {
			t.Errorf("Incorrect match count, expected %d, but got %d", 1, len(matches))
		}
		if matches[0] != "abacus" {
			t.Errorf("Incorrect match, expected %s, but got %s", "abacus", matches[0])
		}
	}

	{
		matches := ttree.wildcardMatch("..a..")
		if len(matches) != 4 {
			t.Errorf("Incorrect match count, expected %d, but got %d", 4, len(matches))
		}
		if matches[0] != "aback" {
			t.Errorf("Incorrect match, expected %s, but got %s", "aback", matches[0])
		}
		if matches[3] != "abate" {
			t.Errorf("Incorrect match, expected %s, but got %s", "abate", matches[0])
		}
	}

}

func BenchmarkTernaryTreeInsert(b *testing.B) {

	//370099 keys
	words, err := readWords("words_alpha.txt")
	if err != nil {
		panic("file not found")
	}

	ttree := New()
	b.ResetTimer()
	b.Run(`InsertKeys`, func(b *testing.B) {
		for _, key := range words {
			ttree.add(key, key)
		}

	})

	b.Run(`RetrieveKeys`, func(b *testing.B) {
		for _, key := range words {
			value, err := ttree.get(key)
			if err != nil {
				b.Errorf("Recieved error while getting %s,%s", key, err)
			}
			if key != value {
				b.Errorf("Failed to get the valid value for the key %s ,recieved %s", key, value)
			}
		}

	})

}

func readWords(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
