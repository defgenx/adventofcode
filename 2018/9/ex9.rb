require '../libs/common'

module Adventofcode
  module List
    class LinkedList
      private

      def to_end(node)
        while node.next_node != @tail
          node = node.next_node
        end
        node
      end

      public

      def initialize
        @tail = nil
        @head = @tail
        @length = 0
      end

      def append(name, value)
        if @head == @tail
          @head = Node.new(value, name, @tail)
        else
          current_node = @head
          current_node = to_end(current_node)
          current_node.next_node = Node.new(value, name, @tail)
        end
        @length += 1
      end

      def prepend(name, value)
        if @length == 0
          @head = Node.new(value, name, @tail)
        else
          node = Node.new(value, name, @head)
          @head = node
        end
        @length += 1
      end

      def size
        @length
      end

      def head
        @head
      end

      def tail
        current_node = @head
        current_node = to_end(current_node)
        current_node
      end

      def at_index(index)
        if index < 0 || index > @length - 1
          nil
        else
          current_node = @head
          (index).times do
            current_node = current_node.next_node
          end
          current_node.value
        end
      end

      def pop
        if @length == 0
          nil
        else
          node = @head
          while node.next_node.next_node != @tail
            node = node.next_node
          end
          pop_node = node.next_node
          node.next_node = @tail
          @length -= 1
          pop_node
        end
      end

      def contains?(value)
        if @length == 0
          false
        else
          if find(value)
            return true
          else
            return false
          end
        end
      end

      def find(data)
        #returns index at which a node exists that contains 'data',
        #else returns nil
        case
        when @head == @tail
          return nil
        when @head.value == data
          return 0
        else
          current_node = @head
          index = 0
          while current_node.value != data && current_node.next_node != @tail
            index += 1
            current_node = current_node.next_node
          end
          if current_node.value == data
            return index
          else
            return nil
          end
        end
      end

      def to_s
        s = ''
        current_node = @head
        (@length).times do
          s += current_node.value.to_s + " -> "
          current_node = current_node.next_node
        end
      end

      def insert_at(index, name = nil, value = nil)
        if index < 0 || index > (@length - 1)
          p index
          nil
        else
          current_node = @head
          case index
          when 0
            prepend(name, value)
          when (@length - 1)
            append(name, value)
          else
            while (index - 1) > 0
              current_node = current_node.next_node
              index -= 1
            end
            right_hand_node = current_node.next_node
            current_node.next_node = Node.new(value, name, right_hand_node)
          end
          @length += 1
        end
      end

      def remove_at(index)
        if index < 0 || index > (@length - 1)
          nil
        else
          unless @length == 0 #can't remove nodes from empty list
            current_node = @head
            case index
            when 0
              @head = @head.next_node
              @length -= 1
            when (@length - 1)
              pop
            else
              while (index) > 0
                prev_node = current_node
                current_node = current_node.next_node
                index -= 1
              end
              prev_node.next_node = current_node.next_node
              current_node.next_node = nil
              @length -= 1
              return current_node
            end
          end
        end
      end
    end

    class Node
      attr_reader :value
      attr_accessor :next_node

      def initialize(value = nil, name = nil, next_node = nil)
        @value = value
        @name = name
        @next_node = next_node
      end
    end
  end
  module FileParser
    include Adventofcode::List
    include Adventofcode::StreamFile

    def parse(file)
      reader = readFullContent(file)
      /^(\d+) players; last marble is worth (\d+) points/.match(reader.chomp)
    end

    def build(parsedString)
      score = Array.new
      list = LinkedList.new
      loopCounter = 1
      while (score.max.nil? || score.max == parsedString[2].to_i)
        1 * loopCounter.upto(parsedString[1].to_i * loopCounter) do |index|
          if (index % 23) == 0
            score[index / loopCounter] = (score[index / loopCounter] || 0)
            score[index / loopCounter] += index
            p list
            score[index / loopCounter] += list.at_index(list.size() - 7)
            list.remove_at(list.size() - 7)
          else
            list.insert_at(index, index, index)
          end
        end
        loopCounter += loopCounter
      end
      score
    end
  end
end

class PartOne
  include Adventofcode::FileParser
  include Adventofcode::List
end
# Exercice 8 Part One
classe = PartOne.new
a = classe.parse('./input_ex9.txt')
p classe.build(a)