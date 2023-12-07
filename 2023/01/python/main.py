import sys

file = open(sys.argv[1])
lines = file.readlines()
sum = 0
for r in lines:
    chars = [x for x in r if x.isdigit()]
    sum += int(chars[0] + chars[-1])

print(sum)