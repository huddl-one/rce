function dailyTemperatures(temperatures) {
    let ans = new Array(temperatures.length).fill(0);
    let stack = [];

    for (let i = 0; i < temperatures.length; i++) {
        while (stack.length !== 0 && temperatures[i] > temperatures[stack[stack.length - 1]]) {
            let idx = stack.pop();
            ans[idx] = i - idx;
        }
        stack.push(i);
    }
    return ans;
}

// Inputs and corresponding expected outputs
const inputs = [
    [73, 74, 75, 71, 69, 72, 76, 73],
    [30, 40, 50, 60],
    // Additional inputs
];

const expectedOutputs = [
    [1, 1, 4, 2, 1, 1, 0, 0],
    [1, 1, 1, 0],
    // Corresponding expected outputs
];

// Evaluation function
function runTests(inputs, expectedOutputs, userFunction) {
    return inputs.map((input, index) => {
        const output = userFunction(input);
        const expected = expectedOutputs[index];
        const passed = JSON.stringify(output) === JSON.stringify(expected);

        return {
            input:input,
            expected: expected,
            output: output,
            passed: passed
        };
    });
}

// Run the tests and output the results
const results = runTests(inputs, expectedOutputs, dailyTemperatures);
console.log(JSON.stringify(results, null, 2));
