require '../libs/common'

module Adventofcode
  module FileParser
    include Adventofcode::StreamFile

    def parse(file)
      reader = readFullContent(file)
      /^(\d+) players; last marble is worth (\d+) points/.match(reader.chomp)
    end

    def build(parsed_string)
      list = Array.new(1, 0)
      nb_players = parsed_string[1].to_i
      score = Array.new(nb_players,0)
      min = 1
      max_bille = parsed_string[2].to_i
        min.upto(max_bille) do |index|
          if index % 23 == 0
            score[index % nb_players] += index
            list.rotate!(-7)
            score[index % nb_players] += list.pop
            list.rotate!(1)
          else
            list.rotate!(1)
            list.push(index)
          end
          if index % 10_000 == 0
            p index
          end
        end
      score.max
    end
  end
end

class PartOne
  include Adventofcode::FileParser
end
# Exercice 9 Part One
classe = PartOne.new
a = classe.parse('./input_ex9.txt')
p classe.build(a)