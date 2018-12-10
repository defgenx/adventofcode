require "./common"

module Adventofcode
    module FileParser
        def parse(file)
            array = Array.new
            readContentStream(file) do |line|
                next if line == 0
                matching = /^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$/.match(line.chomp)
                array.push(matching)
            end
            array
        end
        def parseToHash(file)
            array = Array.new
            readContentStream(file) do |line|
                next if line == 0
                matching = /^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$/.match(line.chomp)
                rect = Hash.new
                rect.store(:id, matching[1].to_i)
                rect.store(:x, matching[2].to_i)
                rect.store(:width, matching[4].to_i)
                rect.store(:y, matching[3].to_i)
                rect.store(:height, matching[5].to_i)
                array.push(rect)
            end
            array
        end
    end
end

class PartOne
    include Adventofcode::StreamFile
    include Adventofcode::FileParser
    attr_reader :storage, :mapStorage, :overlapCount

    public

    def initialize
        @storage = Array.new
        @mapStorage = Array.new
        @overlapCount = 0
    end

    def builder
        @storage.each do |rectangle|
            (rectangle[2].to_i..(rectangle[2].to_i + rectangle[4].to_i - 1)).each do |parentPos|
                @mapStorage[parentPos] = Array.new if @mapStorage[parentPos] == nil
                (rectangle[3].to_i..(rectangle[3].to_i + rectangle[5].to_i - 1)).each do |val|
                    next if @mapStorage[parentPos][val]  == "X"

                    if @mapStorage[parentPos][val]  != "X" && @mapStorage[parentPos][val] != nil
                         @overlapCount += 1
                    end
                    if @mapStorage[parentPos][val]  == nil
                        @mapStorage[parentPos][val] = rectangle[1].to_i
                    else
                        @mapStorage[parentPos][val] = "X"
                    end

                end
            end
        end
    end
end

class PartTwo
    attr_reader :storage, :mapStorage, :finalVal

    public

    include Adventofcode::StreamFile
    include Adventofcode::FileParser

    def initialize
        @storage = Array.new
        @mapStorage = Array.new
        @finalVal = Array.new
    end

    def build(parsedStructure)
    @storage = parsedStructure
        @storage.each do |rectangle|
            @finalVal.push(rectangle[1].to_i) until @finalVal.include? rectangle[1].to_i
            (rectangle[2].to_i..(rectangle[2].to_i + rectangle[4].to_i - 1)).each do |parentPos|
                @mapStorage[parentPos] = Array.new if @mapStorage[parentPos] == nil
                (rectangle[3].to_i..(rectangle[3].to_i + rectangle[5].to_i - 1)).each do |val|
                    tmp = Array.new
                    if @mapStorage[parentPos][val]  == nil
                        tmp.push(rectangle[1].to_i)
                        @mapStorage[parentPos][val] = tmp
                    else
                        tmp.push(rectangle[1].to_i)
                        tmp = tmp + @mapStorage[parentPos][val]
                        tmp = tmp.uniq
                        @mapStorage[parentPos][val] = tmp
                    end
                    @finalVal = @finalVal.reject { |numberCurrent| @mapStorage[parentPos][val].size > 1 && @mapStorage[parentPos][val].include?(numberCurrent) }
                end
            end
        end
        @finalVal.join('').to_i
    end

    def resolve(parsedStructure)
        notFound = false
        firstelmt = parsedStructure.first
        parsedStructure.delete(firstelmt)

        if !parsedStructure.any?{|elmt| overlap(firstelmt, elmt)}
            return firstelmt[:id]
        else
            parsedStructure.push(firstelmt)
            resolve(parsedStructure)
        end
    end

    private

    def overlap (rectA, rectB)
         !(rectA[:x] + rectA[:width] <= rectB[:x] || rectA[:x] >= rectB[:x] + rectB[:width] || rectA[:y] + rectA[:height] <= rectB[:y] || rectA[:y] >= rectB[:y] + rectB[:height])
    end
end

# Exercice 3 Part One
#classe = PartOne.new
#classe.parse('./input_ex3.txt')
#classe.builder
#p classe.overlapCount

# Exercice 3 Part Two
classe = PartTwo.new
#struct = classe.parse('./input_ex3.txt')
#p classe.build struct
struct = classe.parseToHash('./input_ex3.txt')
p classe.resolve struct