require "./common"

class PartOne
    include Adventofcode::StreamFile
    attr_reader :storage, :mapStorage

    def initialize
        @storage = Array.new
        @mapStorage = Hash.new
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
            tmpArray = Hash.new
            tmpArray[rectangle[3].to_i..rectangle[5].to_i] = rectangle[1].to_i
            p tmpArray
            exit
            @mapStorage[rectangle[2].to_i..rectangle[4].to_i] = tmpArray

        end
    end
end

# Exercice 3 Part One
classe = PartOne.new
classe.parse('./input_ex3.txt')
classe.builder
p classe.mapStorage