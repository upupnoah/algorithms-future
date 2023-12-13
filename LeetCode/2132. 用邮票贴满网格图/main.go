package possible_to_stamp

func possibleToStamp(grid [][]int, stampHeight, stampWidth int) bool {
    m, n := len(grid), len(grid[0])

    // 1. 计算 grid 的二维前缀和
    s := make([][]int, m+1)
    s[0] = make([]int, n+1)
    for i, row := range grid {
        s[i+1] = make([]int, n+1)
        for j, v := range row {
            s[i+1][j+1] = s[i+1][j] + s[i][j+1] - s[i][j] + v
        }
    }

    // 2. 计算二维差分
    // 为方便第 3 步的计算，在 d 数组的最上面和最左边各加了一行（列），所以下标要 +1
    d := make([][]int, m+2)
    for i := range d {
        d[i] = make([]int, n+2)
    }
    for i2 := stampHeight; i2 <= m; i2++ {
        for j2 := stampWidth; j2 <= n; j2++ {
            i1 := i2 - stampHeight + 1
            j1 := j2 - stampWidth + 1
            if s[i2][j2]-s[i2][j1-1]-s[i1-1][j2]+s[i1-1][j1-1] == 0 {
                d[i1][j1]++
                d[i1][j2+1]--
                d[i2+1][j1]--
                d[i2+1][j2+1]++
            }
        }
    }

    // 3. 还原二维差分矩阵对应的计数矩阵（原地计算）
    for i, row := range grid {
        for j, v := range row {
            d[i+1][j+1] += d[i+1][j] + d[i][j+1] - d[i][j]
            if v == 0 && d[i+1][j+1] == 0 {
                return false
            }
        }
    }
    return true
}