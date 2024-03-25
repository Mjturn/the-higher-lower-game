const score = 0
const score_element = document.getElementById("score")
score_element.innerHTML += score.toString()

function generate_random_number() {
    const MAX_NUMBER = 100
    const random_number = Math.floor(Math.random() * (MAX_NUMBER + 1))
    const random_number_element = document.getElementById("random-number")
    random_number_element.innerText = random_number.toString()
}

generate_random_number()
