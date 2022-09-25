def is_palindrome(node)
    reversed = reversed_and_clone(node)
    return is_equal(node, reversed)
end

def reversed_and_clone(node)
    head = nil
    until node.nil?
        n = LinkedList.new(node.value)
        n.next = head
        head = n
        node = node.next
    end
    head
end

def is_equal(one, two)
    until one.nil? && two.nil?
        if one.value != two.value
            return false
        end

        one = one.next
        two = two.next
    end
    return one.nil? && two.nil?
end
# Time complexity O(n)