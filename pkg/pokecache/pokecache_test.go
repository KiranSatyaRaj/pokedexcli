package pokecache

import (
	"fmt"
	"testing"
	"time"
	"github.com/KiranSatyaRaj/pokedexcli/pkg/results"
)

func TestAddGet(t *testing.T) {
	const interval = time.Duration(5 * time.Second)
	cases := []struct {
		key string
		val []results.LocationArea
	}{
		{
			key: "https://example.com",
			val: []results.LocationArea{{"testing"}},
		},
		{
			key: "https://example.com/path",
			val: []results.LocationArea{{"moretestdata"}},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func (t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			for i, data := range val {
				if data.Name != c.val[i].Name {
					t.Errorf("expected to find value")
					return
				}
			}
		})

	}

}


func TestReapLoop(t *testing.T) {
	const baseTime = time.Duration(5 * time.Millisecond)
	const waitTime = baseTime + 5 * time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []results.LocationArea{{"testdata"}})

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)
	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
