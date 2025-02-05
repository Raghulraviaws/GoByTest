package iteration

import "testing"

func TestRepeat(t *testing.T) {

	got := Repeat("bangaari")
	want := "bangaaribangaaribangaaribangaaribangaari"

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("bangaari")
	}
}
