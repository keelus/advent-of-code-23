package day17

import (
	"2023/common/pair"
	"container/heap"
	"log"
)

type Visit struct {
	Pos pair.Pair
	// HeatLoss
	Accumulator int
	Dir         pair.Pair
}

var (
	ROWS = -1
	COLS = -1
)

type PriorityQueue []Task

type Task struct {
	V Visit
	H int
}

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].H < pq[j].H }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Task))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) YPush(visit Visit, heat int) {
	heap.Push(pq, Task{V: visit, H: heat})
}

func (pq *PriorityQueue) YPop() (Visit, int) {
	el := heap.Pop(pq).(Task)

	return el.V, el.H
}

func shortestPathCost(island [][]int, source pair.Pair, dest pair.Pair, minAccumulator, maxAccumulator int) int {
	ROWS = len(island)
	COLS = len(island[0])

	visited := make(map[Visit]int)
	pq := make(PriorityQueue, 0)

	pq.YPush(Visit{Pos: pair.Zero(), Dir: pair.Right(), Accumulator: 1}, 0)
	pq.YPush(Visit{Pos: pair.Zero(), Dir: pair.Down(), Accumulator: 1}, 0)

	for len(pq) > 0 {
		u, heat := pq.YPop()

		if !u.Pos.Zero() {
			heat += island[u.Pos.I][u.Pos.J]
		}

		if u.Pos.Eq(dest) && u.Accumulator >= minAccumulator {
			return heat
		}

		if _, ok := visited[u]; ok {
			if visited[u] <= heat {
				continue
			}
		}

		visited[u] = heat

		if u.Accumulator < maxAccumulator {
			nextPos := u.Pos.Add(u.Dir)
			if nextPos.InBounds(0, 0, ROWS, COLS) {
				newVisit := Visit{Pos: nextPos, Dir: u.Dir, Accumulator: u.Accumulator + 1}
				pq.YPush(newVisit, heat)
			}
		}

		if u.Accumulator >= minAccumulator {
			for _, dir := range neighbors[u.Dir] {
				pos := u.Pos.Add(dir)
				if pos.InBounds(0, 0, ROWS, COLS) {
					newVisit := Visit{Pos: pos, Dir: dir.Copy(), Accumulator: 1}
					pq.YPush(newVisit, heat)
				}
			}
		}
	}

	log.Print("We couldn't get to an end")
	return -1
}

var neighbors = map[pair.Pair][2]pair.Pair{
	pair.Up():    {pair.Left(), pair.Right()},
	pair.Down():  {pair.Left(), pair.Right()},
	pair.Left():  {pair.Up(), pair.Down()},
	pair.Right(): {pair.Up(), pair.Down()},
}
