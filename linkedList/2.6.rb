def is_palindrome(node)
    reversed = reversed_and_clone(node)
    until node && reversed
        if node.value != reversed.value
            return false
        end

        node = node.next
        reversed = reversed.next
    end
    return node.nil? && reversed.nil?
end

def reversed_and_clone(node)
    head = node
    until node
        n = LinkedList.new(node.value)
        n.next = head
        head = n
        node = node.next
    end
    head
end

# Time complexity O(n)