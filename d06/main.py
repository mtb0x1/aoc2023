import math

#///source of challenge is : https://adventofcode.com/2023/day/6

def solve_part1() -> int:
    result = 1
    for time, dis in races:
        b = -time
        c = dis
        x1 = (-b - math.sqrt(b ** 2 - 4 * c)) / 2
        x2 = (-b + math.sqrt(b ** 2 - 4 * c)) / 2
        start = x1 + 1 if x1.is_integer() else math.ceil(x1)
        end = x2 - 1 if x2.is_integer() else math.floor(x2)
        winning_ways = int(end - start + 1)
        result *= winning_ways
    return result

def solve_part2() -> int:
    time = int("".join([str(i) for i in times]))
    distance = int("".join([str(i) for i in distances]))
    wins = 0
    possible_holds = [i for i in range(time + 1)]
    possible_distances = [x * (time - x) for x in possible_holds]
    for p in possible_distances:
        if p > distance:
            wins += 1

    return wins

#main
with open('input.txt') as f:
    times = [int(n) for n in f.readline().split()[1:]]
    distances = [int(n) for n in f.readline().split()[1:]]
    races = list(zip(times, distances))

p1 = solve_part1()
p2 = solve_part2()
print("part 1 = ", p1, "\npart 2 = ", p2)