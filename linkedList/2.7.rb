def intersection(node1, node2)
    length1 = length(node1)
    length2 = length(node2)

    if length1 > length2
        (length1 - length2).times {node1 = node1.next}
    else
        (length2 - length1).times {node2 = node2.next}
    end

    until node1
        return node1 if node1 == node2
        node1 = node1.next
        node2 = node2.next
    end
    return nil
end

def length(node)
    i = 0
    until node
        i += 1
        node = node.nextd
    end
    i
end

# Time complexity O(n)