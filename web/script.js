let currentOperation = '';
let firstOperand = '';
let secondOperand = '';
let operator = null;


const display = document.getElementById('display');

// Initialize event listeners for buttons
document.getElementById('historyButton').addEventListener('click', toggleHistoryVisibility);

document.querySelectorAll('.calculator button').forEach(button => {
    button.addEventListener('click', () => handleButtonClick(button.textContent));
});






const operatorEndpointMap = {
    '+': 'add',
    '-': 'subtract',
    '×': 'multiply',
    '÷': 'divide',
    '%': 'modulo',
    '^': 'power',
};


const endpointOperatorMap = {
    'Add': '+',
    'Subtract': '-',
    'Multiply': '*',
    'Divide': '÷',
    'Modulo': '%',
    'Power': '^'
};








// Handle button click logic based on current display and the button clicked.
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





// Update either the first or second operand based on if the operator has been selected.
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






// Perform the calculation by making an API request to the google cloud backend server.
function performCalculation() {
    var notReadyForCalculation = !(firstOperand && operator && secondOperand);
    if (notReadyForCalculation) return;

    const operand1 = parseFloat(firstOperand);
    const operand2 = parseFloat(secondOperand);
    const operation = mapOperatorToEndpoint(operator);

    if (!operation) return;

    //const apiUrl = `https://overengineered-calculator-360186502614.europe-west1.run.app/${operation}?operand1=${operand1}&operand2=${operand2}`;
    const apiUrl = `http://localhost:8080/${operation}?operand1=${operand1}&operand2=${operand2}`; // Local testing

    // Make API request and update display
    fetchApiAndUpdate(apiUrl);
}







// Mapping for selected operator to corresponding API endpoint
function mapOperatorToEndpoint(operator) {
    return operatorEndpointMap[operator];
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







// Method to toggle the visibility of the history list and fetch history
function toggleHistoryVisibility() {
    const historyList = document.getElementById('historyList');
    
    var historyHidden = historyList.classList.contains('hidden');
    if (historyHidden) {
        fetchAndDisplayHistory();
    } else {
        historyList.classList.add('hidden');
    }
}







// Method to fetch history from the API and update the history list
function fetchAndDisplayHistory() {
    
    //const apiUrl = 'https://overengineered-calculator-360186502614.europe-west1.run.app/history';
    const apiUrl = 'http://localhost:8080/history'; // Local testing

    fetch(apiUrl)
        .then(response => response.json())
        .then(historyData => updateHistoryList(historyData))
        .catch(() => console.error('Error fetching history'));
}




// Method to update history list with fetched data
function updateHistoryList(historyData) {

    // Delete existing items
    const historyList = document.getElementById('historyList');
    historyList.innerHTML = ''; 

    historyData.forEach(entry => {
        const listItem = createHistoryListItem(entry);
        historyList.appendChild(listItem);
    });

    historyList.classList.remove('hidden');
}






// Mapping for API endpoint to selected operator
function mapEndpointToOperator(endpointName) {
    return endpointOperatorMap[endpointName];
}



// Create a single list item for the history list.
// Elements are sorted by timestamp in descending order on the backend.
function createHistoryListItem(entry) {
    const listItem = document.createElement('li');

    // Format the timestamp as DD/MM/YY HH:MM
    const timestamp = new Date(entry.Timestamp);
    const day = timestamp.getDate();
    const month = timestamp.getMonth() + 1; // Months are 0-based in JS...
    const year = timestamp.getFullYear().toString().slice(-2); // Last two digits of the year
    const hours = timestamp.getHours();
    const minutes = timestamp.getMinutes().toString().padStart(2, '0'); // Make minutes always two digits

    const formattedDate = `${day}/${month}/${year} ${hours}.${minutes}`;

    const operatorSymbol = mapEndpointToOperator(entry.Operation);


    listItem.textContent = `${formattedDate}: ${entry.Operand1} ${operatorSymbol} ${entry.Operand2} = ${entry.Result}`;

    return listItem;
}