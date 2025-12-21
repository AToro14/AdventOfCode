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

    def validate_range(self, nested_list):
        for s_list in nested_list:
            if not (self.count_digits(s_list[0]) % 2 == 0 or self.count_digits(s_list[1]) % 2 == 0):
                print(nested_list.index(s_list))
                nested_list.pop(nested_list.index(s_list))
        return nested_list
    
    def find_invalid_id(self, nested_list):
        # start with start_range and take the first half of the digits
        # repeat that first half to create an invalid id
        # check if less than end_range
        # if so, add to invalid id list
        # inc first half of digits by 1 and repeat
        invalid_id_lst = []
        for s_list in nested_list: # [999, 1010]
            print("Working on", s_list)
            digits = self.count_digits(s_list[0]) # 3
            if digits % 2 != 0: 
                digits += 1 # 4
            half_digits = digits // 2 # 2
            repeat_chunk = s_list[0] // 10 ** half_digits # 999 // 100 = 9
            invalid_id = repeat_chunk * 10 ** half_digits + repeat_chunk
            print("Calculated start as", repeat_chunk, invalid_id) 
            print("s_list[1]", s_list[1])
            while  invalid_id <= s_list[1]:
                if self.count_digits(invalid_id) % 2 == 0 and invalid_id >= s_list[0]:           
                    invalid_id_lst.append(invalid_id)
                    print(invalid_id_lst)
                repeat_chunk += 1
                invalid_id = repeat_chunk * 10 ** half_digits + repeat_chunk
                print("Next invalid_id to check", invalid_id)
        return invalid_id_lst
    
def main():
    day02 = Solution()
    working_lst = day02.parse_data('d2_input.txt')
    number = 12345
    print(number, "has", day02.count_digits(number), "digits")
    working_lst = day02.format_list(working_lst) 
    print(working_lst)
    day02.validate_range(working_lst)
    print("Valid lists only", working_lst)
    invalid_ids = day02.find_invalid_id(working_lst)
    print(invalid_ids)
    print(sum(invalid_ids))

if __name__ == '__main__':
    main()
