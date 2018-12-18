require '../libs/common'


class PartOne
    include Adventofcode::StreamFile
    def computeFinalVal(file)
        value = 0
         readContentStream(file) do |line|
              value += line.to_i
         end
        value
    end
end

class PartTwo
    #attr_accessor :value, :storeHash
    include Adventofcode::StreamFile

    def initialize
        @value = 0
        @storeHash = Hash.new
    end

    def computeFinalVal(file)
        readContentStream(file) do |line|
            if @storeHash.has_key?(@value)
                @storeHash.store(@value, @storeHash[@value] + 1)
            else
                @storeHash.store(@value, 1)
            end

            if @storeHash[@value] == 2
                return @value
            end
            @value = @value + line.to_i
        end
        puts "Double not found.. Retry."
        computeFinalVal(file)
    end
end

#p PartOne.new.computeFinalVal('./input_ex1.txt')
p PartTwo.new.computeFinalVal('./input_ex1.txt')