print("Hello world!")

combo_lock = [x for x in range(0, 100)] 
position = 50
password = 0

print(combo_lock)
print('At position ', combo_lock[position])

input_file = open('d1_input.txt', 'r')
lines = input_file.readlines()

def find_direction(rotation):
    print(rotation[0])
    return rotation[0]

def find_clicks(rotation):
    print(rotation[1:])
    return int(rotation[1:])

def track_zero(position, password):
    if position == 0:
        password += 1
    print('track_zero() password', password)
    return password

def turn_dial(direction, clicks, position):
    if direction == 'L':
        position = (position - clicks) % 100
    elif direction == 'R':
        position = (position + clicks) % 100
    else:
        print(direction, 'is not a valid direction')
    print('turn_dial() position', position)
    return position

# Main loop
count = 0
print('Current position', position)
print('Current password', password)
for line in lines:
    count += 1
    line = line.strip()
    #print(line)
    #find_direction(line)
    #find_clicks(line)
    position = turn_dial(find_direction(line), find_clicks(line), position)
    #position += find_steps(line)
    password = track_zero(position, password)
    #print('Current position', position)
    print('Current password', password)
    
input_file.close()
print('Total lines', count)
