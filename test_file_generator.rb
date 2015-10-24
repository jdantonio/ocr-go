#!/usr/bin/env ruby

require 'optparse'

DEFAULT_ENTRY_COUNT = 500
DEFAULT_FILE_NAME = 'data.txt'

DEFAULT_OPTIONS = {
  file: DEFAULT_FILE_NAME,
  count: DEFAULT_ENTRY_COUNT
}

LCD = [
  # 0    1    2    3    4    5    6    7    8
  # 1    2    3    4    5    6    7    8    9
  #' '  '_'  ' '  '|'  '_'  '|'  '|'  '_'  '|'
  #--------------------------------------------
  [' ', '_', ' ', '|', ' ', '|', '|', '_', '|'], # 0
  [' ', ' ', ' ', ' ', ' ', '|', ' ', ' ', '|'], # 1
  [' ', '_', ' ', ' ', '_', '|', '|', '_', ' '], # 2
  [' ', '_', ' ', ' ', '_', '|', ' ', '_', '|'], # 3
  [' ', ' ', ' ', '|', '_', '|', ' ', ' ', '|'], # 4
  [' ', '_', ' ', '|', '_', ' ', ' ', '_', '|'], # 5
  [' ', '_', ' ', '|', '_', ' ', '|', '_', '|'], # 6
  [' ', '_', ' ', ' ', ' ', '|', ' ', ' ', '|'], # 7
  [' ', '_', ' ', '|', '_', '|', '|', '_', '|'], # 8
  [' ', '_', ' ', '|', '_', '|', ' ', '_', '|'], # 9
]

def parse_options(default_options)
  options = {}

  OptionParser.new do |opts|
    opts.banner = "Usage: #{File.basename(__FILE__)} [options]"

    opts.on("-c", "--count NUM", Integer,
            "Number of entries in the file (default: #{DEFAULT_ENTRY_COUNT})") do |c|
      options[:count] = c
    end

    opts.on("-f", "--file FILE",
            "Name of the output file (default: #{DEFAULT_FILE_NAME})") do |f|
      options[:file] = f
    end
  end.parse!

  default_options.merge(options)
end

def print_number(number, file)
  number = number.to_s.rjust(9, '0')

  3.times do |row|
    number.length.times do |digit|
      3.times do |column|
        file.write(LCD[number[digit].to_i][(row * 3) + column])
      end
    end
    file.write("\n")
  end
end

if $0 == __FILE__

  options = parse_options(DEFAULT_OPTIONS)

  File.open(options[:file], 'w') do |file|

    options[:count].times do
      number = rand(1_000_000_000)
      print_number(number, file)
      file.write("\n")
    end
  end
end
