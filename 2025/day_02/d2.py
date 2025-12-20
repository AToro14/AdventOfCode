# Day 02

class Solution:
    def parse_data(self, file_name):
        input_file = open(file_name, 'r')
        lines = input_file.readlines()
        print(lines)
        lines = lines[0].split("\n")
        product_ranges = lines[0].split(",")
        print(product_ranges)

day02 = Solution()
day02.parse_data('d2_test_input.txt')

    
