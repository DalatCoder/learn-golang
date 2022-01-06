package newfeeds

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{Title: "Hieu", Post: "Ha"})

	if len(feed.Items) == 0 {
		t.Error("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{Title: "Hieu", Post: "Ha"})
	feed.Add(Item{Title: "Ha", Post: "Hieu"})

	results := feed.GetAll()

	if len(results) != 2 {
		t.Error("Item was not added")
	}
}
