def kth_to_last(node, k)
    fast_pointer = node
    slow_pointer = node

    k.times do
        fast_pointer = fast_pointer.next
    end

    until fast_pointer.nil?
        fast_pointer = fast_pointer.next
        slow_pointer = slow_pointer.next
    end

    slow_pointer.value
end

# Time complexity O(n)