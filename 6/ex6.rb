require '../libs/common'

module Adventofcode
  module FileParser
    include Adventofcode::StreamFile

    def parse(file)
      parsedArray = readFullContent(file).split("\n").map do |coords|
        coords.split(",").map(&:to_i)
      end
    end
  end
  module Exercice6
    private

    def countMax(countArray)
      countArray.max_by {|k, v| v}
    end

    def computeDistance(x1, x2, y1, y2)
      (x1 - x2).abs + (y1 - y2).abs
    end

    def computeBorders(x, minMaxX, y, minMaxY)
      (minX, maxX) = minMaxX
      (minY, maxY) = minMaxY
      ((y == minY || y == maxY) || (x == minX || x == maxX))
    end

    def deleteInfinite(arrayCount, pointToExclude)
      arrayCount.select {|val, key| !pointToExclude.include?(val)}
    end
  end
end

class PartOne
  include Adventofcode::FileParser
  include Adventofcode::Exercice6

  def distance(arrayCoords)
    (minX, maxX) = arrayCoords.map(&:first).minmax
    (minY, maxY) = arrayCoords.map(&:last).minmax
    pointCount = Hash.new
    pointToExclude = Array.new
    (minX..maxX).each do |x|
      (minY..maxY).each do |y|
        dist = Hash.new
        arrayCoords.map.with_index do |(valX, valY), index|
          dist[index] = computeDistance(valX, x, valY, y)
        end
        minDist = dist.values.min
        selected = dist.select do |index, distance|
          distance == minDist
        end
        if selected.length == 1
          selectedKey = selected.keys.first
          pointToExclude.push(selectedKey) if computeBorders(x, [minX, maxX], y, [minY, maxY]) && !pointToExclude.include?(selectedKey)
          pointCount[selectedKey] = 0 unless pointCount.has_key?(selectedKey)
          pointCount[selectedKey] += 1
        end
      end
    end
    countMax(deleteInfinite(pointCount, pointToExclude))
  end
end

class PartTwo
  include Adventofcode::FileParser
  include Adventofcode::Exercice6

  def distance(arrayCoords)
    (minX, maxX) = arrayCoords.map(&:first).minmax
    (minY, maxY) = arrayCoords.map(&:last).minmax
    count = 0
    (minX..maxX).each do |x|
      (minY..maxY).each do |y|
        sum = arrayCoords.sum do |(valX, valY)|
          computeDistance(valX, x, valY, y)
        end
        count += 1 if sum < 10_000
      end
    end
    count
  end

end

# Exercice 6 Part One
#classe = PartOne.new
#a = classe.parse('./input_ex6.txt')
#p classe.distance(a)

# Exercice 6 Part Two
classe = PartTwo.new
a = classe.parse('./input_ex6.txt')
p classe.distance(a)