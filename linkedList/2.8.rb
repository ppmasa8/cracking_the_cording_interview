def detect_loop(node)
    fast = node
    slow = node

    if fast.nil? || slow.nil?
        return nil
    end

    until fast && fast.next
        fast = fast.next.next
        slow = slow.next
        break if fast == slow
    end

    slow = node
    until slow == fast
        slow = slow.next
        fast = fast.next
    end
    slow
end
