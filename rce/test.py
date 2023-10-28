import unittest
# from app import two_sum

def two_sum(nums, target):
    seen = {}
    for i, num in enumerate(nums):
        complement = target - num
        if complement in seen:
            return [seen[complement], i]
        seen[num] = i
    return None

class TestTwoSum(unittest.TestCase):
    def run_test(self, nums, target, expected):
        result = two_sum(nums, target)
        self.assertEqual(result, expected)

    def test_case1(self):
        nums = [2, 7, 11, 15]
        target = 9
        result = two_sum(nums, target)
        self.assertEqual(result, [0, 0], "Test Case 1 failed. Expected [0, 1], got {}".format(result))


    def test_cases(self):
        test_cases = [
            ([2, 7, 11, 15], 9, [0, 1]),
            ([3, 2, 4], 6, [1, 2]),
            ([3, 3], 6, [0, 1]),
            ([1, 2, 3, 4], 10, None)
        ]

        for nums, target, expected in test_cases:
            with self.subTest(nums=nums, target=target, expected=expected):
                self.run_test(nums, target, expected)

if __name__ == '__main__':
    unittest.main()




# import unittest
# from parameterized import parameterized
# # from app import two_sum

# class TestTwoSum(unittest.TestCase):

#     @parameterized.expand([
#         ([2, 7, 11, 15], 9, [0, 1]),
#         ([3, 2, 4], 6, [1, 2]),
#         ([3, 3], 6, [0, 1]),
#         ([1, 2, 3, 4], 10, None)
#     ])
#     def test_two_sum(self, nums, target, expected):
#         result = two_sum(nums, target)
#         self.assertEqual(result, expected)

# if __name__ == '__main__':
#     unittest.main()




