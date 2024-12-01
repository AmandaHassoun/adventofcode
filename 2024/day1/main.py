filename = "input.txt"

left = []
right = []
with open(filename) as file:
    for line in file:
        nums = line.rstrip().split()
        left.append(int(nums[0]))
        right.append(int(nums[1]))

left.sort()
right.sort()

total_distance = 0
for i in range(len(left)):
    diff = abs(left[i] - right[i])
    total_distance += diff

print(f"Part 1: Total distance is {total_distance}")

similarity_score = 0

for num_left in left:
    for num_right in right:
        if num_left == num_right:
            similarity_score += num_left

print(f"Part 2: Similarity score is {similarity_score}")

