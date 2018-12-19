RubyVM::InstructionSequence.compile_option = {
  :tailcall_optimization => true,
  :trace_instruction => false
}


RubyVM::InstructionSequence.new(<<-EOF).eval
    # LOCAL LIBS
    require '../libs/common'

    # BUNDLER LIBS
    require 'bundler/setup'
    require 'pattern-match'
    using PatternMatch

    module Exercice5
        def react(polymer, tmpArray)
            test = match([polymer, tmpArray]) do
                with(_[a & String, b & Array], guard { !a.empty? && b.empty?} ) do
                    firstChar = a[0]
                    stringToArray = a
                    stringToArray[0]= ''
                    [stringToArray, b.insert(0, firstChar)]
                end
                with(_[a & String, b & Array], guard { !a.empty? && !b.empty? && (a[0].ord - b.first.ord).abs == 32 } ) do
                    stringToArray = a
                    stringToArray[0]= ''
                    [stringToArray, b.drop(1)]
                end
                with(_[a & String, b & Array], guard { !a.empty? && !b.empty?} ) do
                    firstChar = a[0]
                    stringToArray = a
                    stringToArray[0]= ''
                    [stringToArray, b.insert(0, firstChar)]
                end
                with(_[a & String, b & Array], guard { a.empty? } ) do
                    b.length
                end
                with(_) do
                    puts "No matching pattern found ! \nStopping program..."
                    exit
                end
            end
            return test unless test.kind_of?(Array)
            react(test[0], test[1])
        end

        def reactPerf(polymer, tmpArray)
            if !polymer.empty? && tmpArray.empty?
                firstChar = polymer[0]
                stringToArray = polymer
                stringToArray[0]= ''
                returnVal = [stringToArray, tmpArray.insert(0, firstChar)]
            elsif !polymer.empty? && !tmpArray.empty? && (polymer[0].ord - tmpArray.first.ord).abs == 32
                stringToArray = polymer
                stringToArray[0]= ''
                returnVal = [stringToArray, tmpArray.drop(1)]
            elsif !polymer.empty? && !tmpArray.empty?
                firstChar = polymer[0]
                stringToArray = polymer
                stringToArray[0]= ''
                returnVal = [stringToArray, tmpArray.insert(0, firstChar)]
            elsif polymer.empty?
                returnVal = tmpArray.length
            else
                puts "No matching pattern found ! \nStopping program..."
                exit
            end
            return returnVal unless returnVal.kind_of?(Array)
            reactPerf(returnVal.first, returnVal.last)
        end
    end

    class PartOne
        include Adventofcode::StreamFile
        include Exercice5
    end

    class PartTwo
        include Adventofcode::StreamFile
        include Exercice5
    end
EOF
# Exercice 5 Part One
#classe = PartOne.new
#fileString = classe.readFullContent('./input_ex5.txt')
#puts classe.react(fileString, Array.new)

# Exercice 5 Part Two
classe = PartTwo.new
fileString = classe.readFullContent('./input_ex5.txt')
tmpArray =  ("a".."z").map do |l|
    tmpStr = fileString.clone
    tmpStr.delete!(l)
    tmpStr.delete!(l.upcase)
    classe.reactPerf(tmpStr, Array.new)
end
p tmpArray.min
