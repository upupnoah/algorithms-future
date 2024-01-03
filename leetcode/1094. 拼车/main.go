package main

func carPooling(trips [][]int, capacity int) bool {
    toMax := 0
    for _, trip := range trips {
        toMax = max(toMax, trip[2])
    }
    
    diff := make([]int, toMax + 1)
    for _, trip := range trips {
        diff[trip[1]] += trip[0]
        diff[trip[2]] -= trip[0]
    }

    count := 0
    for i := 0; i < toMax; i++ {
        count += diff[i]
        if count > capacity {
            return false
        }
    }
    return true
}
