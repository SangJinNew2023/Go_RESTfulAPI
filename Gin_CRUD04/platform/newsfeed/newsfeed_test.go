package newsfeed

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	fmt.Println(feed.Items)
	fmt.Println(len(feed.Items))
	if len(feed.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	results := feed.GetAll()
	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}
