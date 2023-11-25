import json
import concurrent.futures
import traceback

# User-defined function
def daily_temperatures(temperatures):
    ans = [0] * len(temperatures)
    stack = []

    for i, t in enumerate(temperatures):
        while stack and t > stack[-1][0]:
            tempt, tempi = stack.pop()
            ans[tempi] = i - tempi
        stack.append([t, i])
    return ans

# Test Cases
inputs = [
    [73, 74, 75, 71, 69, 72, 76, 73],
    [30, 40, 50, 60],
    [30, 60, 90],
    [90, 80, 70, 60],
]

expected_outputs = [
    [1, 1, 4, 2, 1, 1, 0, 0],
    [1, 1, 1, 0],
    [1, 1, 0],
    [0, 0, 0, 0],
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

# Evaluate the user function
def evaluate_user_function(user_func, inputs, expected_outputs):
    with concurrent.futures.ThreadPoolExecutor() as executor:
        futures = [executor.submit(evaluate_test_case, user_func, inp, exp) for inp, exp in zip(inputs, expected_outputs)]
        return [future.result() for future in futures]

# Running the Evaluation
results = evaluate_user_function(daily_temperatures, inputs, expected_outputs)

# Output Results
print(json.dumps(results, indent=4))
