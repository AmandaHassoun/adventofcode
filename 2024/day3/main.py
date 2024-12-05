import re

filename = "input.txt"
pattern = r'mul\((\d+),(\d+)\)'


def multiply(matches: list) -> int:
    total = 0
    for pairs in matches:
        total += int(pairs[0]) * int(pairs[1])

    return total


total_sum = 0

with open(filename) as file:
    for line in file:
        command = line.rstrip()
        command_matches = re.findall(pattern, command)
        total_sum += multiply(command_matches)
file.close()

print(f"Part1: Multiplication result is {total_sum}")