require '../libs/common'

module Adventofcode
    module FileParser
        include Adventofcode::StreamFile
        def parse(file)
            tree = Hash.new
            readFullContent(file).split("\n").map do |line|
                matching = /^Step (\w).+step (\w).+$/.match(line.chomp)
                tree[matching[1]] = Hash.new  if tree[matching[1]].nil?
                tree[matching[2]] = Hash.new  if tree[matching[2]].nil?
                tree[matching[1]][:children] = Array.new  if tree[matching[1]][:children].nil?
                tree[matching[1]][:parents] = Array.new  if tree[matching[1]][:parents].nil?
                tree[matching[2]][:children] = Array.new  if tree[matching[2]][:children].nil?
                tree[matching[2]][:parents] = Array.new  if tree[matching[2]][:parents].nil?
                tree[matching[2]][:parents].push(matching[1])
                tree[matching[1]][:children].push(matching[2])
                tree[matching[1]][:parents].sort!
                tree[matching[1]][:children].sort!
                tree[matching[2]][:parents].sort!
                tree[matching[2]][:children].sort!
            end
            tree = sortHashByKey(tree)
            foundRoots = findRoots(tree)
            aggrChar = ''
            foundRoots.each do |char|
                aggrChar += findNext(char, tree)
            end
            p aggrChar
        end

        private

        def search?(tree, (valToFind, keyVal))
            return true if tree[valToFind][:parents].length > 1
            false
        end

        def findNext(root, tree)
            aggrChar = root
            tree[root][:children].each do |currentCharChild, _|
                if search?(tree, [currentCharChild, root])
                    tree[currentCharChild][:parents].delete(tree[currentCharChild][:parents].first)
                    next
                end
                aggrChar += findNext(currentCharChild, tree)
            end
            aggrChar
        end

        def sortHashByKey(hash)
            hash.sort_by {|k, v| k}.to_h
        end

        def findRoots(tree)
            tree.map{|key, subHash| key if subHash[:parents].empty?}.delete_if{|i| i.nil?}.sort!
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