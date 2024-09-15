let currentOperation = '';
let firstOperand = '';
let secondOperand = '';
let operator = null;


const display = document.getElementById('display');

// Initialize event listeners for calculator buttons
document.querySelectorAll('.calculator button').forEach(button => {
    button.addEventListener('click', () => handleButtonClick(button.textContent));
});




// Handle button click logic
function handleButtonClick(value) {

    var isNumberOrDecimal = !isNaN(value) || value === '.';
    if (isNumberOrDecimal) {
        updateOperands(value);
        return
    }
    
    var isOperator = ['+', '-', '×', '÷', '%', '^'].includes(value);
    if (isOperator) {
        operator = value;
        return
    } 
    
    var isReset = value === 'C';
    if (isReset) {
        resetCalculator();
        return
    }
    
    var isEqual = value === '=';
    if (isEqual) {
        performCalculation();
    }
}





// Update operands based on the current operator
function updateOperands(value) {
    if (operator === null) {
        firstOperand += value;
        updateDisplay(firstOperand);
    } else {
        secondOperand += value;
        updateDisplay(secondOperand);
    }
}






// Reset calculator to initial state
function resetCalculator() {
    firstOperand = '';
    secondOperand = '';
    operator = null;
    updateDisplay('0');
}


// Update the calculator display
function updateDisplay(value) {
    display.textContent = value;
}






// Perform the calculation by making an API request
function performCalculation() {
    if (!firstOperand || !operator || !secondOperand) return;

    const operand1 = parseFloat(firstOperand);
    const operand2 = parseFloat(secondOperand);
    const operation = mapOperatorToEndpoint(operator);

    if (!operation) return;

    const apiUrl = `http://localhost:8080/${operation}?operand1=${operand1}&operand2=${operand2}`;

    // Make API request and update display
    fetchApiAndUpdate(apiUrl);
}







// Mapping for selected operator to corresponding API endpoint
function mapOperatorToEndpoint(operator) {
    const operatorMap = {
        '+': 'add',
        '-': 'subtract',
        '×': 'multiply',
        '÷': 'divide',
        '%': 'modulo',
        '^': 'power',
    };
    return operatorMap[operator];
}



// Make API request and update display with result
function fetchApiAndUpdate(apiUrl) {
    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            updateDisplay(data.result);
            firstOperand = data.result.toString();
            secondOperand = '';
            operator = null;
        })
        .catch(() => {
            display.textContent = "Error occurred";
        });
}
