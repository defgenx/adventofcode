module Adventofcode
    module StreamFile
        def readContentStream(file)
            f = File.open(file, 'r')
            f.each_line do |line|
                    yield line
            end
        ensure
            puts "Close file"
            f.close
        end

        def readFullContent(file)
            f = File.open(file, 'r')
            f.read
        ensure
            puts "Close file"
            f.close
        end
    end
end