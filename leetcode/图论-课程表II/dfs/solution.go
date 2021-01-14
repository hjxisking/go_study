package dfs

func FindOrder(numCourses int, prerequisites [][]int) []int {
    edges := make([][]int, numCourses)
    visited := make([]int, numCourses)
    result := []int{}

    for _, requisit := range prerequisites {
        edges[requisit[1]] = append(edges[requisit[1]], requisit[0])
    }

    for i := 0; i < numCourses; i ++ {
        if visited[i] == 0 && !dfs(edges, i, visited, &result) {
            return []int{}
        }
    }

    for i, j := 0, len(result) - 1; i < j; i, j = i + 1, j - 1 {
        result[i], result[j] = result[j], result[i]
    }

    return result
}

func dfs(edges [][]int, course int, visited []int, result *[]int) bool {
    visited[course] = 1
    nextCourses := edges[course]

    for _, nextCourse := range nextCourses {
        if visited[nextCourse] == 0 && !dfs(edges, nextCourse, visited, result) {
            return false
        }

        if visited[nextCourse] == 1 {
            return false
        }
    }

    visited[course] = 2
    *result = append(*result, course)
    return true
}