package loophole

//
// search for a ways and loops
//

type Node int64
type Weight float64

type Edge struct {
	From   Node
	To     Node
	Weight Weight
}

type Graph []Edge          // неупорядоченный набор переходов
type Path []Edge           // упорядоченный набор переходов
type NodeSet map[Node]bool // поле узлов, куда не надо заходить
type Nodes []Node          // упорядоченный набор узлов

func (g *Graph) _walk(start Node, finish Node, path *Path, seen *NodeSet, fn func(*Path) bool) bool {
	pix := len(*path)
	*path = append(*path, Edge{}) // добавим места для следующего шага
	for _, step := range *g {
		if step.From != start || (*seen)[step.To] {
			continue // "These Are Not the Droids You Are Looking For"
		}
		(*path)[pix] = step     // поставим переход
		(*seen)[step.To] = true // узел отметим как посещённый
		if step.To == finish {  // это узел, куда мы в итоге должны прийти
			if fn(path) {
				return true // коллбэк разрешил досрочный выход!
			}
		} else {
			if g._walk(step.To, finish, path, seen, fn) {
				return true // коллбэк в рекурсии разрешил досрочный выход!
			}
		}
		delete(*seen, step.To) // уберём отметку посещённости
	}
	*path = (*path)[:pix]
	return false // "Thank You, Mario! But Our Princess Is In Another Castle!"
}

// Обход графа и вызов callback-а на каждый найденный путь.
// если callback вернёт true обход завершится досрочно
func (g *Graph) Walk(start Node, finish Node, fn func(*Path) bool) bool {
	path := make(Path, 0, len(*g))
	seen := make(NodeSet)
	if start != finish {
		seen[start] = true
	}
	return g._walk(start, finish, &path, &seen, fn)
}
