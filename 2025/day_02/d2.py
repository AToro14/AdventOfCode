# Day 02

class Solution:
    def parse_data(self, file_name):
        input_file = open(file_name, 'r')
        lines = input_file.readlines()
        print(lines)
        lines = lines[0].split("\n")
        ranges = lines[0].split(",")
        product_ranges = []
        for item in ranges:
            product_ranges.append((item.split("-")))
        print(product_ranges)
        return product_ranges

    def format_list(self, str_list):
        for item in str_list:
            # print('string typ', item)
            for i in range(len(item)):
                item[i] = int(item[i])
            # print('int-ified', item)
        # print(str_list)
        return str_list

    def count_digits(self, end_range):
        str_end_range = str(end_range)
        return len(str_end_range)

def main():
    day02 = Solution()
    working_lst = day02.parse_data('d2_test_input.txt')
    number = 12345
    print(number, "has", day02.count_digits(number), "digits")
    working_lst = day02.format_list(working_lst) 
    print(working_lst)

if __name__ == '__main__':
    main()
