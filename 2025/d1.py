print("Hello world!")

combo_lock = [x for x in range(0, 100)] 
position = 0
password = 0

print(combo_lock)
print('At position ', combo_lock[position])

input_file = open('d1_input.txt', 'r')
lines = input_file.readlines()

def find_direction(rotation):
    print(rotation[0])
    return rotation[0]

def find_steps(rotation):
    print(rotation[1:])
    return int(rotation[1:])

def track_zero(position):
    if position == 0:
        password += 1
    return

# Main loop
for line in lines[0:4]:
    line = line.strip()
    print(line)
    find_direction(line)
    find_steps(line)
    position += find_steps(line)
    print('Current position', position)

input_file.close()
