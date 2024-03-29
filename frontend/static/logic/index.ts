let score = 0
const score_element = document.getElementById("score")
score_element.innerHTML += score.toString()

function generate_random_numbers() {
    const MAX_NUMBER = 100
    const random_number_left = Math.floor(Math.random() * (MAX_NUMBER + 1))
    const random_number_right = Math.floor(Math.random()* (MAX_NUMBER + 1))
    const random_number_left_element = document.getElementById("random-number-left")
    random_number_left_element.innerText = random_number_left.toString()

    determine_result(random_number_left, random_number_right)
}

let result = ""

function determine_result(random_number_left, random_number_right) {
    if(random_number_left > random_number_right) {
        result = "higher"
    } else if(random_number_left < random_number_right) {
        result = "lower"
    } else {
        result = "equal"
    }
}
    
const higher_button = document.getElementById("higher-button")
const lower_button = document.getElementById("lower-button")

higher_button.addEventListener("click", () => handle_button_click("higher"))
lower_button.addEventListener("click", () => handle_button_click("lower"))

function handle_button_click(higher_or_lower) {
    if(higher_or_lower == "higher" && result == "higher") {
        score++
    } else if(higher_or_lower == "higher" && result == "lower") {
        score = 0
    } else if(higher_or_lower == "lower" && result == "lower") {
        score++
    } else if(higher_or_lower == "lower" && result == "higher") {
        score = 0
    } else if(result == "equal") {
        score++
    }

    score_element.innerText = "Score: " + score
    generate_random_numbers()
}

generate_random_numbers()
