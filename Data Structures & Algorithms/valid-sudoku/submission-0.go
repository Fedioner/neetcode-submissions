func isValidSudoku(board [][]byte) bool {
    tmp := make(map[byte]int, 9)

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            c := board[i][j]
            if c == '.' {
                continue
            }
            tmp[c]++
            if tmp[c] > 1 {
                return false
            }
        }
        clear(tmp)
    }

    for j := 0; j < 9; j++ {
        for i := 0; i < 9; i++ {
            c := board[i][j]
            if c == '.' {
                continue
            }
            tmp[c]++
            if tmp[c] > 1 {
                return false
            }
        }
        clear(tmp)
    }

    for bi := 0; bi < 3; bi++ {
        for bj := 0; bj < 3; bj++ {
            for ii := 0; ii < 3; ii++ {
                for jj := 0; jj < 3; jj++ {
                    c := board[bi*3+ii][bj*3+jj]
                    if c == '.' {
                        continue
                    }
                    tmp[c]++
                    if tmp[c] > 1 {
                        return false
                    }
                }
            }
            clear(tmp)
        }
    }

    return true
}