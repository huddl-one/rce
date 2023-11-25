import json
import concurrent.futures
import traceback

# Helper function with the actual implementation
def max_subarray_sum_helper(nums):
    # Implement the logic here (Kadane’s algorithm, for example)
    max_sum = nums[0]
    current_sum = nums[0]

    for num in nums[1:]:
        current_sum = max(num, current_sum + num)
        max_sum = max(max_sum, current_sum)
    return max_sum

# Primary function that calls the helper
def max_subarray_sum(nums):
    return max_subarray_sum_helper(nums)

# Test Cases
inputs = [
    [-2, 1, -3, 4, -1, 2, 1, -5, 4],
    [1],
    [0, -1, 2, -3],
    [-1, -2, -3, -4],
    [1, 2, 3, 4, 5],
    # Additional test cases
]

expected_outputs = [
    6,  # Explanation: [4,-1,2,1] has the largest sum = 6.
    1,
    2,
    -1,
    15,
    # Corresponding expected outputs
]

# Evaluation Function
def evaluate_test_case(user_function, input_case, expected_output):
    try:
        output = user_function(input_case)
        return {
            'passed': output == expected_output,
            'input': input_case,
            'expected': expected_output,
            'output': output,
            'error': None
        }
    except Exception as e:
        return {
            'passed': False,
            'input': input_case,
            'expected': expected_output,
            'output': None,
            'error': str(traceback.format_exc())
        }

# Evaluate the primary function
def evaluate_user_function(user_func, inputs, expected_outputs):
    with concurrent.futures.ThreadPoolExecutor() as executor:
        futures = [executor.submit(evaluate_test_case, user_func, inp, exp) for inp, exp in zip(inputs, expected_outputs)]
        return [future.result() for future in futures]

# Running the Evaluation
results = evaluate_user_function(max_subarray_sum, inputs, expected_outputs)

# Output Results
print(json.dumps(results, indent=4))
