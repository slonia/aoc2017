input = File.readlines('input.txt')
->(input) do
  result = input.inject(0) do |sum, line|
    line = line.split(/\s/).map(&:to_i)
    sum += line.max - line.min
    sum
  end
  puts result
end

# part 2

->(input) do
  dividers = {}
  result = input.inject(0) do |sum, line|
    line = line.split(/\s/).map(&:to_i)
    max = nil
    min = nil
    line.each_with_index do |num, i|
      (i+1).upto(line.size - 1).each do |j|
        max_cand = [num, line[j]].max
        min_cand = [num, line[j]].min
        if max_cand % min_cand == 0
          max, min = max_cand, min_cand
          break
        end
      end
      break if !max.nil?
    end
    sum += max / min
    sum
  end
  puts result
end[input]
