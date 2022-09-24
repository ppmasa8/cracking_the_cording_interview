module Util
    class LinkedList
        attr_accessor :head

        class Node
            attr_accessor :value, :next

            def initialize(value)
                @value = value
            end

            def to_a
                node = self
                array = []
                until node.nil?
                    array << node.value
                    node = node.next
                end
                array
            end

            def append(value)
                new_node = Node.new(value)
                if @head.nil?
                    @head = new_node
                else
                    node = @head
                    node = node.next until node.next.nil?
                    node.next = new_node
                end
            end

            def to_a
                @head.to_a
            end
        end
    end
end