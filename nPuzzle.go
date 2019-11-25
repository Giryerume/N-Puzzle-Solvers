package main

import (
	"math/rand"
	"math"
	"time"
	"reflect"
	"log"
	"strconv"
)

type state struct {
	tiles  []int
	dad    *state
	blank  int
	lvl    int
	n      int
	man    int
}

func NewTest(t []int) state {
	b := indexBlank(t)
	n := int(math.Sqrt(float64(len(t))))
	test := state{tiles: t, blank: b, lvl: 0, n: n}
	return test
}

func (t *state) verifyV(up bool) bool {
	if up {
		return t.blank-t.n >= 0
	} else {
		return t.blank+t.n <= len(t.tiles)-1
	}
}

func (t *state) verifyH(right bool) bool {
	if right {
		i := ( t.blank + 1 ) % t.n
		return i != 0
	} else {
		j := t.blank % t.n
		return j != 0
	}
}

func (t *state) moveV(up bool) state {
	copy := copyState(*t)
	x := t.tiles[t.blank]
	if up {
		copy.tiles[t.blank] = t.tiles[t.blank-t.n]
		copy.tiles[t.blank-t.n] = x
		copy.blank = copy.blank - t.n
	} else {
		copy.tiles[t.blank] = t.tiles[t.blank+t.n]
		copy.tiles[t.blank+t.n] = x
		copy.blank = copy.blank + t.n
	}
	return copy
}

func (t *state) moveH(right bool) state {
	copy := copyState(*t)
	x := t.tiles[t.blank]
	if right {
		copy.tiles[t.blank] = t.tiles[t.blank+1]
		copy.tiles[t.blank+1] = x
		copy.blank = copy.blank + 1
	} else {
		copy.tiles[t.blank] = t.tiles[t.blank-1]
		copy.tiles[t.blank-1] = x
		copy.blank = copy.blank - 1
	}
	return copy
}

func (t *state) updateBlank() {
	for i, n := range t.tiles {
		if n == 0 {
			t.blank = i
			break
		}
	}
}

func printTiles(s state) {
	for i := 0; i < len(s.tiles); i++ {
		if i%s.n == 0 {
			print("\n")
		}
		print(s.tiles[i])
		print(" ")
	}
	println()
}

func printSolution(s *state){
	printTiles(*s)
	if(s.dad!=nil) {
		printSolution(s.dad)
	}
}

func indexBlank(tiles []int) int {
	for i, n := range tiles {
		if n == 0 {
			return i
		}
	}
	return -1
}

func randInt(min int, max int) int {
	var random = rand.New(rand.NewSource(time.Now().UnixNano()))
	return min+random.Intn(max-min)
}

func copyState(s state) state{
	var v []int
	for _, e := range s.tiles {
		v = append(v, e)
	}
	c := NewTest(v)
	return c
}

func randomState(s state, x int) state {
	t := copyState(s)
	aux := -1
	for i := 0; i < x; i++ {
		n := randInt(0, 4)
		switch n {
		case 0:
			if t.verifyH(false) && aux!=1 {
				t = t.moveH(false)
				aux = n
			} else {
				i--
			}
		case 1:
			if t.verifyH(true) && aux!=0 {
				t = t.moveH(true)
				aux = n
			} else {
				i--
			}
		case 2:
			if t.verifyV(false) && aux!=3 {
				t = t.moveV(false)
				aux = n
			} else {
				i--
			}
		case 3:
			if t.verifyV(true) && aux!=2 {
				t = t.moveV(true)
				aux = n
			} else {
				i--
			}
		}
	}
	return t
}

func manhattan(final state, atual state) int { //manhattan pronto
	var dist, x, y, j, valor int
	dist = 0
	for i := 0; i < len(final.tiles); i++ {
		j = 0
		valor = final.tiles[i]
		for k := 0; atual.tiles[k] != valor; k++ {
			j++
		}
		x=i/final.n
		y=j/final.n
		if x > y {
			dist = dist - y + x
		}
		if y > x {
			dist = dist - x + y
		}
		if i%final.n > j%final.n {
			dist = dist + i%final.n - j%final.n
		}
		if i%final.n < j%final.n {
			dist = dist + j%final.n - i%final.n
		}
	}
	return dist
}

func updateManhattan(final state, s []state) []state {
	for i := 1; i < len(s); i++ {
		s[i].man = manhattan(final, s[i])
	}
	return s
}

func bestManhattan(s []state) int {
	var min int
	m := s[0].man+s[0].lvl
	for i, e := range s {
		if x:=e.man+e.lvl; x<m {
			m = x
			min = i
		}
	}
	return min
}

func breeder(f state) []state {
	s := []state{}
	if f.verifyH(false ) {
		x := f.moveH(false )
		x.lvl = f.lvl + 1
		x.dad = &f
		s = append(s, x)
	}
	if f.verifyH(true ) {
		x := f.moveH(true )
		x.lvl = f.lvl + 1
		x.dad = &f
		s = append(s, x)
	}
	if f.verifyV(false ) {
		x := f.moveV(false )
		x.lvl = f.lvl + 1
		x.dad = &f
		s = append(s, x)
	}
	if f.verifyV(true ) {
		x := f.moveV(true )
		x.lvl = f.lvl + 1
		x.dad = &f
		s = append(s, x)
	}
	return s
}

func isIn(w state, s []state) bool {
	for _,e := range s {
		if reflect.DeepEqual(w.tiles, e.tiles) { return true }
	}
	return false
}

func keygen(t []int) string {
	var s string
	for _,e := range t {
		s += strconv.Itoa(e)
	}
	return s
}

func depth(final state, atual state) {
	start := time.Now()
	var depth int
	visited := []state{}
	visited = append(visited, atual)
	aux := copyState(atual)
	for ; !reflect.DeepEqual(atual.tiles, final.tiles); {
		for ; isIn(aux, visited) ; {
			aux = randomState(atual, 1)
			aux.lvl = atual.lvl + 1
		}
		atual = aux
		visited = append(visited, atual)
		//printTiles(atual)//Para imprimir a solução
		depth = atual.lvl
	}
	elapsed := time.Since(start)
	log.Printf("(Depth Search)")
	log.Printf("estados expandidos: %d", len(visited))
	log.Printf("tempo de execução: %s", elapsed)
	log.Printf("fator de ramificação: %d", 1)
	log.Printf("profundidade: %d", depth)
	println()
}

func depthFirst(start state, final state) {
	timer := time.Now()
	var depth  int
	var media, s, f float32
	stack := []state{}
	visited := make(map[string]state)
	stack = append(stack, start)
	for ; len(stack) != 0; {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		successors := breeder(v)
		s = s+float32(len(successors))
		f++
		media = s/f
		if reflect.DeepEqual(v.tiles, final.tiles) {
			//printSolution(&v)//Para imprimir a solução
			depth = v.lvl
			break
		}
		visited[keygen(v.tiles)] = v
		for _,w := range successors {
			_, t := visited[keygen(w.tiles)]
			if t==false {
				stack = append(stack, w)
			}
		}
	}
	elapsed := time.Since(timer)
	log.Printf("(Depth First Search)")
	log.Printf("estados expandidos: %d", len(visited))
	log.Printf("tempo de execução: %s", elapsed)
	log.Printf("fator de ramificação: %f", media)
	log.Printf("profundidade: %d", depth)
	println()
}

func breadthFirst(start state, final state) {
	timer := time.Now()
	var depth  int
	var media, s, f float32
	queue := []state{}
	visited := make(map[string]state)
	queue = append(queue, start)
	for ; len(queue) != 0; {
		v := queue[0]
		queue = queue[1:]
		successors := breeder(v)
		s = s+float32(len(successors))
		f++
		media = s/f
		if reflect.DeepEqual(v.tiles, final.tiles) {
			//printSolution(&v)//Para imprimir a solução
			depth = v.lvl
			break
		}
		visited[keygen(v.tiles)] = v
		for _,w := range successors {
			_, t := visited[keygen(w.tiles)]
			if t==false {
				queue = append(queue, w)
			}
		}
	}
	elapsed := time.Since(timer)
	log.Printf("(Breadth First Search)")
	log.Printf("estados expandidos: %d", len(visited))
	log.Printf("tempo de execução: %s", elapsed)
	log.Printf("fator de ramificação: %f", media)
	log.Printf("profundidade: %d", depth)
	println()
}

func aStar(start state, final state) {
	timer := time.Now()
	var depth  int
	var media, s, f float32
	var queue = []state{}
	visited := make(map[string]state)
	queue = append(queue, start)
	for ; len(queue) != 0; {
		i := bestManhattan(queue)
		v := queue[i]
		queue = append(queue[:i], queue[i+1:]...)
		successors := updateManhattan(final,breeder(v))
		s = s+float32(len(successors))
		f++
		media = s/f
		if reflect.DeepEqual(v.tiles, final.tiles) {
			//printSolution(&v)//Para imprimir a solução
			depth = v.lvl
			break
		}
		visited[keygen(v.tiles)] = v
		for _,w := range successors {
			_, t := visited[keygen(w.tiles)]
			if t==false {
				queue = append(queue, w)
			}
		}
	}
	elapsed := time.Since(timer)
	log.Printf("(A*)")
	log.Printf("estados expandidos: %d", len(visited))
	log.Printf("tempo de execução: %s", elapsed)
	log.Printf("fator de ramificação: %f", media)
	log.Printf("profundidade: %d", depth)
	println()
}

func main() {
	goal := NewTest([]int{1, 2, 3, 4, 5, 6, 7, 8, 0})
	//goal := NewTest([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0})
	//goal := NewTest([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 0})
	n := randInt(goal.n*goal.n, goal.n*goal.n+10)
	start := randomState(goal, n)
	print("Estado Objetivo")
	printTiles(goal)
	print("Estado Inicial, embaralhos: ",n)
	printTiles(start)
	println("\n+=+=+\n")
	aStar(goal, start)
	depth(goal, start)
	breadthFirst(goal, start)
	depthFirst(goal, start)
}
