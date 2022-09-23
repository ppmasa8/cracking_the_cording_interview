def main()
    matrix = [
        [1,2,3],
        [4,5,6],
        [0,3,2],
        [3,4,3]
    ]
    p matrix
    setZeros(matrix)
    p matrix
end

def setZeros(matrix)
    row, column = Array.new(matrix.size, false), Array.new(matrix[0].size, false)

    for i in 0...matrix.size
        for j in 0...matrix[0].size
            if matrix[i][j] == 0
                row[i] = true
                column[j] = true
            end
        end
    end

    for i in 0...row.size
        nullifyRow(matrix, i) if row[i]
    end

    for j in 0...column.size
        nullifyColumn(matrix, j) if column[j]
    end
end

def nullifyRow(matrix, row)
    for j in 0...matrix[0].size
        matrix[row][j] = 0
    end
end

def nullifyColumn(matrix, column)
    for i in 0...matrix.size
        matrix[i][column] = 0
    end
end

main()
