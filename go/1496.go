type IntPair struct {
    x int
    y int
}

func isPathCrossing(path string) bool {
    visited := make(map[IntPair]bool)

    location := IntPair{0,0}

    visited[location] = true

    for _,direction := range path {
         if direction == 'N' {
             location.y--
         } else if direction == 'S' {
             location.y++
         } else if direction == 'W' {
             location.x--
         } else if direction == 'E' {
             location.x++
         }

         _,ok := visited[location]
         if ok {
           return true
         } else {
             visited[location] = true
         }
    }
    return false
}
