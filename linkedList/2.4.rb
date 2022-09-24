def partition(node, x)
    head = node

    until node.nil?
        if !node.next.nil? && node.next.value < x
            tmp = node.next
            node.next = node.next.next
            tmp.next = head
            head = tmp
        else
            node = node.next
        end
    end
    head
end

# Time complexity O(n)