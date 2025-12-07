#combo_lock = [x for x in range(0, 100)] 
last_position = 50
current_position = last_position
password = 0

#print(combo_lock)
#print('At position ', combo_lock[current_position])

input_file = open('d1_input.txt', 'r')
lines = input_file.readlines()

def find_direction(rotation):
    print(rotation[0])
    return rotation[0]

def find_clicks(rotation):
    print(rotation[1:])
    return int(rotation[1:])

def track_landed_zero(position, password, clicks):
    if position == 0:
        password += 1
    print('track_zero() password', password)
    return password

def track_passed_zeroes(direction, clicks, position, password, current_position):
    print('track_passed_zeroes() starting with', password, 'zeroes')
    if direction == 'L':
        if position < (clicks % 100) and not position == 0:
            password += 1
            print('added 1 more going L for total of', password)
    elif direction == 'R':
        if (100 - position) < (clicks % 100) and not position == 0:
            password += 1
            print('added 1 more going R for total of', password)
    if clicks >= 100:
        password += (clicks // 100)

    print('adding more for total of', password)
    return password

def turn_dial(direction, clicks, position, password):
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
print('Current position', current_position)
print('Last position', last_position)
print('Current password', password)
for line in lines:
    count += 1
    line = line.strip()
    clicks = find_clicks(line)
    direction = find_direction(line)
    
    print('last_position', last_position)
    print(clicks, 'clicks to the', direction)
    current_position = turn_dial(direction, clicks, last_position, password)
    print('current_position', current_position)
    password = track_passed_zeroes(direction, clicks, last_position, password, current_position)
    password = track_landed_zero(current_position, password, clicks)
    last_position = current_position
    print('Current password', password)
    
input_file.close()
print('Total lines', count)
