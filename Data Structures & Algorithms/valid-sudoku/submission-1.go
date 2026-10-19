func isValidSudoku(board [][]byte) bool {
    var rows, cols, boxes [9][9]bool
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            c := board[i][j]
            if c == '.' {
                continue
            }
            d := c - '1'
            b := (i/3)*3 + j/3
            if rows[i][d] || cols[j][d] || boxes[b][d] {
                return false
            }
            rows[i][d] = true
            cols[j][d] = true
            boxes[b][d] = true
        }
    }
    return true
}