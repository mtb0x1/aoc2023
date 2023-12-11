def calculate_weights(line, p1: true)
  all_dots = line.all? { |c| c == "." }
  p1 ? (all_dots ? 2 : 1) : (all_dots ? 1_000_000 : 1)
end

def parse()
  File.readlines('input.txt').map(&:chomp).map(&:chars)
end

def solve(data, p1: true)
  rows = data.map { |line| calculate_weights(line, p1: p1) }
  cols = data.transpose.map { |line| calculate_weights(line, p1: p1) }
  distance_weight = ->(from, to) {
    ([from[0], to[0]].min...[from[0], to[0]].max).sum { |y| rows[y] } +
    ([from[1], to[1]].min...[from[1], to[1]].max).sum { |y| cols[y] }
  }

  xy = []
  data.each.with_index { |row, y| row.each.with_index { |v, x| xy << [y, x] if v == "#" } }
  xy.combination(2).sum { |from, to| distance_weight[from, to] }
end

def solve_part1(data)
  solve(data, p1: true)
end

def solve_part2(data)
  solve(data, p1: false)
end

# Main
data = parse()
p1 = solve_part1(data)
p2 = solve_part2(data)
puts "part 1 = %d\npart 2 = %d" % [p1, p2]