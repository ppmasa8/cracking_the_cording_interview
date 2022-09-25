def is_palindrome(node)
    fast = node
    slow = node

    stack = []

    while fast != nil && fast.next != nil
        stack.push(slow.value)
        slow = slow.next
        fast = fast.next.next
    end

    # when fast pointer element is odd, slow pointer is next.
    if fast != nil
        slow = slow.next
    end

    while slow != nil
        top = stack.pop
        if top != slow.value
            return false
        end
        slow = slow.next
    end
    return true
end

# Time complexity O(n) half of 2.6.rb