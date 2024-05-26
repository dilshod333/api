window.onload = sendApiRequest;

async function sendApiRequest(){
    let response = await fetch(`https://opentdb.com/api.php?amount=10&category=18&type=multiple`);
    let data = await response.json();
    useApiData(data);
}

function useApiData(data){
    document.querySelector("#category").innerHTML = `Category: ${data.results[0].category}`;
    document.querySelector("#difficulty").innerHTML = `Difficulty: ${data.results[0].difficulty}`;
    const questionsContainer = document.querySelector("#questions");

    // Loop through the results to display questions and answers
    for(let i = 0; i < data.results.length; i++) {
        // Display the question
        questionsContainer.innerHTML += `<p>Question ${i + 1}: ${data.results[i].question}</p>`;

        // Display the answers
        const answersContainer = document.createElement("div");
        answersContainer.classList.add("answers-container");

        let answers = [data.results[i].correct_answer, ...data.results[i].incorrect_answers];
        answers = shuffleArray(answers); // Shuffle the answers to randomize their order

        for(let j = 0; j < answers.length; j++) {
            const answerButton = document.createElement("button");
            answerButton.innerHTML = answers[j];
            answerButton.classList.add("answer-button");
            answersContainer.appendChild(answerButton);
        }

        questionsContainer.appendChild(answersContainer);
    }
}

// Function to shuffle an array (Fisher-Yates shuffle algorithm)
function shuffleArray(array) {
    for (let i = array.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [array[i], array[j]] = [array[j], array[i]];
    }
    return array;
}
