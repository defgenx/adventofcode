require '../libs/common'
require 'date'

module Adventofcode
  module FileParser
    def parse(file)
      array = Array.new
      parsedArray = Array.new
      readContentStream(file) do |line|
        next if line == 0
        array.push(line)
      end
      array.sort.each do |line|
        matching = /^\[(\d+-\d+-\d+ \d+:\d+)\] (Guard #(\d+) begins shift|wakes (up)|falls (asleep))?$/.match(line.chomp)
        parsedArray.push(matching)
      end
      parsedArray
    end
  end

  module DataStructure
    private

    def sumMaxSleepByGuard(hashDataStructure)
      newHash = Hash.new
      hashDataStructure.each do |guardId, subHash|
        newHash[guardId] = subHash.inject(0) {|sum, tuple| sum += tuple[1]}
      end
      newHash
    end

    def largestHashKeyVal(hash)
      hash.max_by {|k, v| v}
    end
  end
end

class PartOne
  include Adventofcode::StreamFile
  include Adventofcode::FileParser
  include Adventofcode::DataStructure

  def initialize
    @structure = Hash.new {Array.new}
  end


  def guardMaxMinuteSleep(dataStructure)
    hashDataStructure = timeForGuard(dataStructure)
    maxByGuardHash = sumMaxSleepByGuard(hashDataStructure)
    tupleResult = largestHashKeyVal(maxByGuardHash)
    tupleResult[0].to_i * (largestHashKeyVal(hashDataStructure[tupleResult[0]])[0])
  end

  private

  def timeForGuard(dataStructure)
    parsedDateAsleep = nil
    parsedDateWakeup = nil
    currentGuard = nil
    dataStructure.each do |matchData|
      if matchData[3] != nil
        @structure[matchData[3]] = Hash.new until @structure.include?(matchData[3])
        currentGuard = matchData[3]
      end
      if matchData[5] != nil
        parsedDateAsleep = Date._parse(matchData[1])
      end
      if matchData[4] != nil
        parsedDateWakeup = Date._parse(matchData[1])
        (parsedDateAsleep[:min]..parsedDateWakeup[:min]).each do |val|
          if @structure[currentGuard][val] == nil
            @structure[currentGuard][val] = 1
          else
            @structure[currentGuard][val] += 1
          end
        end

      end
    end
    @structure
  end
end

class PartTwo
  include Adventofcode::StreamFile
  include Adventofcode::FileParser
  include Adventofcode::DataStructure

  def initialize
    @structure = Hash.new {Array.new}
  end

  def allGuardsMaxMinuteSleep(dataStructure)
    guardMax = Array.new(2, 0)
    currentGuard = nil
    hashDataStructure = timeForGuard(dataStructure)
    hashDataStructure.each do |guardId, subHash|
      tmpGuardMax = largestHashKeyVal(subHash)
      if tmpGuardMax != nil && tmpGuardMax[1] > guardMax[1]
        guardMax = tmpGuardMax
        currentGuard = guardId.to_i
      end
    end
    currentGuard * guardMax[0]
  end

  private

  def timeForGuard(dataStructure)
    parsedDateAsleep = nil
    parsedDateWakeup = nil
    currentGuard = nil
    dataStructure.each do |matchData|
      if matchData[3] != nil
        @structure[matchData[3]] = Hash.new until @structure.include?(matchData[3])
        currentGuard = matchData[3]
      end
      if matchData[5] != nil
        parsedDateAsleep = Date._parse(matchData[1])
      end
      if matchData[4] != nil
        parsedDateWakeup = Date._parse(matchData[1])
        (parsedDateAsleep[:min]..parsedDateWakeup[:min]).each do |val|
          if @structure[currentGuard][val] == nil
            @structure[currentGuard][val] = 1
          else
            @structure[currentGuard][val] += 1
          end
        end

      end
    end
    @structure
  end
end

# Exercice 4 Part One
classe = PartOne.new
dataStructure = classe.parse('./input_ex4.txt')
p classe.guardMaxMinuteSleep(dataStructure)

# Exercice 4 Part Two
classe = PartTwo.new
dataStructure = classe.parse('./input_ex4.txt')
p classe.allGuardsMaxMinuteSleep(dataStructure)