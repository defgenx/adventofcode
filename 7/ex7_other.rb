require '../libs/common'
require "set"

module Adventofcode
    module FileParser
        include Adventofcode::StreamFile
        def parse(file)
            tree = Hash.new { |h,k| h[k] = Set.new }
            readFullContent(file).split("\n").map do |line|
                matching = /^Step (\w).+step (\w).+$/.match(line.chomp)
                tree[matching[2]].add(matching[1])
            end
            finalString = Set.new
            remaining = tree.map { |k,v| [k, v.to_a] }.flatten.uniq.sort.to_set

            until remaining.empty?
                step = remaining.find {|step|
                    # p step, tree[step], finalString
                    # p "--------------"
                    tree[step].subset?(finalString)
                }
                # p "Add #{step}"
                remaining.delete(step)
                finalString.add(step)
            end
            finalString
        end
    end
end

class PartOne
    include Adventofcode::FileParser
end

class PartTwo
    include Adventofcode::FileParser
end

# Exercice 7 Part One
classe = PartOne.new
a = classe.parse('./input_ex7.txt')
p a.to_a.join