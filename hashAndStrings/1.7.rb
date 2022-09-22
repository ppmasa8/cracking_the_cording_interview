def main()
    matrix = [
        [1,1,1,1],
        [2,2,2,2],
        [3,3,3,3],
        [4,4,4,4]
    ]
    p matrix
    rotation(matrix)
    p matrix
end

def rotation(matrix)
    if matrix.size == 0 || matrix.size != matrix[0].size
        false
    end

    n = matrix.size

    (0...n/2).each do |layer|
        first = layer
        last = n - 1 - layer
        (first...last).each do |i|
            offset = i - first
            top = matrix[first][i]

            #left -> upper
            matrix[first][i] = matrix[last-offset][first]

            #lower -> left
            matrix[last-offset][first] = matrix[last][last-offset]

            #right -> lower
            matrix[last][last-offset] = matrix[i][last]

            #upper -> right
            matrix[i][last] = top
        end
    end
    true
end

main()
