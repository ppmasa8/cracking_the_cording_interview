def delete_middle_node(node)
    if node.nil? || node.next.nil?
        return false
    end

    node.value = node.next.value
    node.next = node.next.next
end

# Time complexity O(1)