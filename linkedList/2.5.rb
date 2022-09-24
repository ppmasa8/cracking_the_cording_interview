def sum_lists(node1, node2)
    if node1.nil? && node2.nil?
        return
    end

    carry = 0
    result = LinkedList.new

    until node1.nil? && node2.nil?
        val1 = node1.nil? ? 0 : node1.value
        val2 = node2.nil? ? 0 : node2.value

        sum = val1 + val2 + carry
        result.append(sum%10)
        carry = sum / 10

        node1 = node1.next unless node1.nil?
        node2 = node2.next unless node2.nil?
    end
    result.append(carry) if carry > 0
    result
end

# Time complexity O(n) n = [node1.size, node2.size].max