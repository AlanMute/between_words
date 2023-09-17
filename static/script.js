document.addEventListener("DOMContentLoaded", function () {
    const form = document.querySelector("form");
    const resultElement = document.getElementById("result");

    form.addEventListener("submit", function (e) {
        e.preventDefault(); // Предотвращаем отправку формы по умолчанию

        const input1 = document.querySelector("input[name='input1']").value;
        const input2 = document.querySelector("input[name='input2']").value;

        // Отправляем данные на сервер
        fetch("/combine", {
            method: "POST",
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
            body: `input1=${encodeURIComponent(input1)}&input2=${encodeURIComponent(input2)}`,
        })
        .then((response) => response.json())
        .then((data) => {
                resultElement.textContent = data.result;        
        })
    });
});