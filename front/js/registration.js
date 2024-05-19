


function Reg(){
    $('input[type="file"]').val();
    $('input[type="file"]').val().replace(/.+[\\\/]/, "");
    let qer = document.querySelectorAll(".form > select");
        let inputs = document.querySelectorAll(".form > input");
        let data = {};
        for (let i = 0; i < inputs.length; i++) {
            data[inputs[i].name] = inputs[i].value;
        }
        for (let i = 0; i < qer.length; i++) {
            data[qer[i].name] = qer[i].value;
        }
        let xhr = new XMLHttpRequest();
        xhr.open("POST", "/user/reg");
        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response);
            if ("Error" in response) {
                if (response.Error == null) {
                    alert("зарегестрировано!"); 
                   console.log("Пользователь успешно зарегистрирован");
                } else {
                    alert("зарегестрировано!");
                    
                    console.log(response.Error);
                }
            } else {
                console.log("Некорректные данные");
            }
        };
        console.log(data)
        window.location.href = 'http://127.0.0.1:8080/authorization';
        xhr.send(JSON.stringify(data));
}

function UpdatePhoto(){
        let inputs = document.querySelectorAll(".form > input");
        let data = {};
        for (let i = 0; i < inputs.length; i++) {
            data[inputs[i].name] = inputs[i].value;
        }
        let xhr = new XMLHttpRequest();
        xhr.open("POST", "/user/update");
        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response);
            if ("Error" in response) {
                if (response.Error == null) {
                    alert("Обновлено!"); 
                    location.reload
                   console.log("Пользователь успешно зарегистрирован");
                } else {
                    alert("зарегестрировано!");
                    
                    console.log(response.Error);
                }
            } else {
                console.log("Некорректные данные");
            }
        };
        console.log(data)
        xhr.send(JSON.stringify(data));
}


