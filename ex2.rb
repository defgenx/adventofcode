require "./common"

class PartOne
    include Adventofcode::StreamFile
    attr_reader :storage

    def initialize
        @storage = Hash.new
    end

    def counter(file)
        counter = 1
        readContentStream(file) do |line|
            tmpHash = Hash.new
            next if line == 0
            line.chomp.split('').each { |char|
                if tmpHash.has_key?(char)
                    tmpHash.store(char, tmpHash[char] + 1)
                else
                    tmpHash.store(char, 1)
                end

                break if tmpHash.has_value?(2) && tmpHash.has_value?(3)
            }
            @storage.store(counter, tmpHash)
            counter = counter + 1
        end
    end

    def checksum(intCompare)
        counter = 0
        @storage.each do |lineId, subHash|
            subHash.each do |charVal, numberVal|
                if numberVal % intCompare == 0
                    counter = counter + 1
                    break
                end
            end
        end
        counter
    end
end

class PartTwo
    include Adventofcode::StreamFile
    attr_reader :storage

    def initialize
        @storage = Array.new
    end

    def parse(file)
        readContentStream(file) do |line|
            next if line == 0
            @storage.push(line.chomp.split(''))
        end
    end

    def intersect
        foundFlag = false
        0.upto(@storage.length - 1) do |indexLine|
            diff = ''
            0.upto(@storage.length - 1) do |indexLineComp|
                0.upto(@storage[indexLineComp].length - 1) do |i|
                    if @storage[indexLine][i] != @storage[indexLineComp][i]
                        break if foundFlag
                        diff = i
                        foundFlag = true
                    end
                    if foundFlag && @storage[indexLine].length - 1 == i
                    copy = Array.new
                    copy = @storage
                    copy[indexLine].delete_at(diff)
                    return copy[indexLine].join('')
                    end
                end
                foundFlag = false
            end
        end
    end
end

# Exercice 2 Part One
#classe = PartOne.new
#classe.counter('./input_ex2.txt')
#countDouble = classe.checksum(2)
#countTriple = classe.checksum(3)
#p countDouble * countTriple

# Exercice 2 Part Two
classe = PartTwo.new
classe.parse('./input_ex2.txt')
p classe.intersect