require '../libs/common'

module Adventofcode
  module FileParser
    include Adventofcode::StreamFile

    def parse(file)
      readFullContent(file).split(" ").map(&:to_i)
    end
  end

  module TreeNode
    class Node
      attr_accessor :children, :metadata

      def initialize(nb_child, nb_meta)
        @children = Array.new {|a, k| a[k] = self.new}
        @header = {nb_child: nb_child, nb_meta: nb_meta}
        @metadata = Array.new
      end

      def add_child(node)
        @children.push(node)
      end

      def add_metadata(metadata)
        @metadata.push(metadata)
      end

      def get_nb_child
        @header[:nb_child]
      end

      def get_nb_meta
        @header[:nb_meta]
      end
    end
  end

  module ExecuteEx
    include Adventofcode::TreeNode
    def execute((nb_child, nb_meta, *rest))
      t = Node.new(nb_child, nb_meta)
      t.get_nb_child.times do |_|
        (child, rest) = execute(rest)
        t.add_child(child)
      end
      t.get_nb_meta.times {|meta| t.add_metadata(rest[meta])}
      [t, rest[t.get_nb_meta..rest.size]]
    end
  end
end

class PartOne
  include Adventofcode::FileParser
  include Adventofcode::ExecuteEx

  def sum_meta_tree(node)
    sum = 0
    node.metadata.each {|val| sum += val}
    node.children.each {|child| sum += sum_meta_tree(child)}
    sum
  end

end

class PartTwo
  include Adventofcode::FileParser
  include Adventofcode::ExecuteEx

  def sum_meta_tree(node)
    sum = 0
    if node.get_nb_child == 0
      node.metadata.each {|val| sum += val}
    else
      node.metadata.each do |val|
        next if node.children[val-1].nil?
        sum += sum_meta_tree(node.children[val-1])
      end
    end
    sum
  end
end

# Exercice 8 Part One
# classe = PartOne.new
# a = classe.parse('./input_ex8.txt').freeze
# tmp_a = a.clone
# (t, _) = classe.execute(tmp_a)
# p classe.sum_meta_tree(t)

# Exercice 8 Part Two
classe = PartTwo.new
a = classe.parse('./input_ex8.txt').freeze
tmp_a = a.clone
(t, _) = classe.execute(tmp_a)
p classe.sum_meta_tree(t)