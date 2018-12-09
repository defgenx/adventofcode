require "./common"

class PartOne
    include Adventofcode::StreamFile
    attr_reader :storage, :mapStorage, :overlapCount

    public

    def initialize
        @storage = Array.new
        @mapStorage = Array.new { Array.new }
        @overlapCount = 0
    end

    def parse(file)
        readContentStream(file) do |line|
            next if line == 0
            matching = /^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$/.match(line.chomp)
            @storage.push(matching)
        end
    end

    def builder
        @storage.each do |rectangle|
            tmpUnordered = [rectangle[2].to_i, rectangle[4].to_i]
            tmpMinF = tmpUnordered.min
            tmpMaxF = tmpUnordered.max
            (tmpMinF..tmpMaxF).each do |parentPos|
                @mapStorage[parentPos] = Array.new if @mapStorage[parentPos] == nil

                tmpUnordered = [rectangle[3].to_i, rectangle[5].to_i]
                tmpMin = tmpUnordered.min
                tmpMax = tmpUnordered.max
                (tmpMin..tmpMax).each do |val|

                    if @mapStorage[parentPos][val]  != nil

                         @mapStorage[parentPos][val] = "X"
                         @overlapCount += 1
                     else
                        @mapStorage[parentPos][val] = rectangle[1].to_i
                    end
                end
            end
        end
    end
end

# Exercice 3 Part One
classe = PartOne.new
classe.parse('./input_ex3.txt')
classe.builder
p classe.overlapCount