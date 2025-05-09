# Day XX Test for Setting Up CI Pipeline

def solve_part1(a, b):
    return (a ** b) - (b ** a)

def solve_part2(a, b):
    return (a + b) // (b - a)

def solve_part3(a, b):
    return (a * b) + (a - b)

arg_1 = 2
arg_2 = 5   

print(solve_part1(arg_1, arg_2))
print(solve_part2(arg_1, arg_2))
