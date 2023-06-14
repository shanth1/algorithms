package algorithms

func BFS(graph map[string][]string, start string, check func(name string) bool) bool {

	var queue []string
	queue = append(queue, start)
	queue = append(queue, graph[start]...)

	searched := map[string]bool{}

	for len(queue) > 0 {
		person := queue[0]
		if _, ok := searched[person]; !ok {
			if check(person) {
				return true
			} else {
				queue = append(queue, graph[person]...)
				searched[person] = true
			}
		}
		queue = queue[1:]

	}
	return false
}
