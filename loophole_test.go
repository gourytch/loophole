package loophole

import "testing"

func testWalk_empty(t *testing.T) {
	G := Graph{}
	(&G).Walk(1, 9, func(p *Path) bool {
		t.Errorf("bad path: %v", *p)
		return false
	})
}

func testWalk_single(t *testing.T) {
	G := Graph{}
	G = append(G, Edge{1, 9, 1.23})
	visited := false
	(&G).Walk(1, 9, func(p *Path) bool {
		if visited {
			t.Errorf("duplicate visit: %v", *p)
		}
		visited = true
		L := len(*p)
		if L != 1 {
			t.Errorf("path with bad length: %v", *p)
		}
		if (*p)[0].From != 1 {
			t.Errorf("bad path start: %v", *p)
		}
		if (*p)[L-1].To != 9 {
			t.Errorf("bad path finish: %v", *p)
		}
		return false
	})
	if !visited {
		t.Error("path not found")
	}
}

func testWalk_chain(t *testing.T) {
	G := Graph{}
	G = append(G, Edge{1, 2, 1.23})
	G = append(G, Edge{2, 3, 2.34})
	G = append(G, Edge{3, 4, 3.45})
	G = append(G, Edge{4, 5, 4.56})
	G = append(G, Edge{5, 6, 5.67})
	G = append(G, Edge{6, 7, 6.78})
	G = append(G, Edge{7, 8, 7.89})
	G = append(G, Edge{8, 9, 8.90})
	visited := false
	(&G).Walk(1, 9, func(p *Path) bool {
		if visited {
			t.Errorf("duplicate visit: %v", *p)
		}
		visited = true
		L := len(*p)
		if L != 8 {
			t.Errorf("path with bad length: %v", *p)
		}
		if (*p)[0].From != 1 {
			t.Errorf("bad path start: %v", *p)
		}
		if (*p)[L-1].To != 9 {
			t.Errorf("bad path finish: %v", *p)
		}
		return false
	})
	if !visited {
		t.Error("path not found")
	}
}

func testWalk_multiple(t *testing.T) {
	G := Graph{}
	G = append(G, Edge{1, 2, 1.2})
	G = append(G, Edge{1, 3, 1.3})
	G = append(G, Edge{1, 4, 1.4})
	G = append(G, Edge{2, 4, 2.4})
	G = append(G, Edge{2, 9, 2.9})
	G = append(G, Edge{3, 9, 3.9})
	count := 0
	(&G).Walk(1, 9, func(p *Path) bool {
		L := len(*p)
		if L != 2 {
			t.Errorf("path with bad length: %v", *p)
		}
		if (*p)[0].From != 1 {
			t.Errorf("bad path start: %v", *p)
		}
		if (*p)[L-1].To != 9 {
			t.Errorf("bad path finish: %v", *p)
		}
		count++
		return false
	})
	if count != 2 {
		t.Error("bad number of paths")
	}
}

func testWalk_bidi(t *testing.T) {
	G := Graph{}
	G = append(G, Edge{1, 2, 1.2})
	G = append(G, Edge{1, 3, 1.3})
	G = append(G, Edge{1, 4, 1.4})
	G = append(G, Edge{2, 4, 2.4})
	G = append(G, Edge{2, 3, 2.3})
	G = append(G, Edge{3, 2, 3.2})
	G = append(G, Edge{2, 9, 2.9})
	G = append(G, Edge{3, 9, 3.9})
	count := 0
	(&G).Walk(1, 9, func(p *Path) bool {
		L := len(*p)
		if L != 2 && L != 3 {
			t.Errorf("path with bad length: %v", *p)
		}
		if (*p)[0].From != 1 {
			t.Errorf("bad path start: %v", *p)
		}
		if (*p)[L-1].To != 9 {
			t.Errorf("bad path finish: %v", *p)
		}
		count++
		return false
	})
	if count != 4 {
		t.Error("bad number of paths")
	}
}

func TestWalk(t *testing.T) {
	testWalk_empty(t)
	testWalk_single(t)
	testWalk_chain(t)
	testWalk_multiple(t)
	testWalk_bidi(t)
}
