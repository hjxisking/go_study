package dfs

func CanFinish(numCourses int, prerequisites [][]int) bool {
    edges := make([][]int, numCourses)
    visited := make([]int, numCourses)

    for _, requisit := range prerequisites {
        edges[requisit[1]] = append(edges[requisit[1]], requisit[0])
    }

    for i := 0; i < numCourses; i ++ {
        if visited[i] == 0 {
            if !dfs(edges, i, visited) {
                return false
            }
        }
    }

    return true
}

func dfs(edges [][]int, course int, visited []int) bool {
    visited[course] = 1
    nextCourses := edges[course]
    if len(nextCourses) == 0 {
        visited[course] = 2
        return true
    } else {
        for _, nextCourse := range nextCourses {
            if visited[nextCourse] == 0 && !dfs(edges, nextCourse, visited) {
                return false
            }

            if visited[nextCourse] == 1 {
                return false
            }
        }
        visited[course] = 2
    }
    return true
}

