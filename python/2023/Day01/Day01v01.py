#  Day 01 - Trebuchet?!

# iterate over each line in file
filename = "testData.txt"
def main():
    with open("testData.txt") as inFile:
        for line in inFile:
            currStr = line
            isNum(currStr, 0)
            #print("\n".rstrip("\n"))
            file1.write("\n")
            #print(line.rstrip("\n"))

# check if each character is a number
def isNum(inStr, direction):
    if direction == 0:
        for element in inStr:
            if element.isnumeric():
                #print (element, end = '')
                file1.write(element)
                isNumRev(inStr)
                return

# if number then add to a new string

# iterate over same string backwards
def isNumRev(inStr):
    for element in inStr[ : :-1]:
        if element.isnumeric():
            #print(element, end = '')
            file1.write(element)
            return
# if number then add to same string

# convert strings into ints

# sum column of ints for answer
def sumLines():
    sum = 0
    with open("testOutput.txt") as inFile:
        for line in inFile:
            num = int(line)
            sum += num
    print("Sum = ", sum)

# find all digits as words
def findWordNum():
    with open("test2.txt") as inFile:
        for line in inFile:
            if "two" in inFile.read():
                print("two found")

#12381577 function calls
file1 = open("testOutput.txt", "a")
main()
file1.close()
sumLines()
file2 = open("test2.txt", "a")
findWordNum()
