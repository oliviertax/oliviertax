function Zayavka(){
    let button = document.querySelector(".form > button");
            let inputs = document.querySelectorAll(".form > input");
            let data = {};
            for (let i = 0; i < inputs.length; i++) {
                data[inputs[i].name] = inputs[i].value;
            }
            let xhr = new XMLHttpRequest();
            xhr.open("POST", "/user/zayavka");
            xhr.onload = function (e) {
                let response = JSON.parse(e.currentTarget.response);
                if ("Error" in response) {
                    if (response.Error == null) {
                        alert("отправлено!"); 
                       console.log("Заявка отправлена!");
                    } else {
                        console.log(response.Error);
                    }
                } else {
                    console.log("Некорректные данные");
                }
            };
            console.log(data)
            xhr.send(JSON.stringify(data));
    }
    function Delete() {
        let inputs = document.querySelectorAll(".form > input");
        let data = {};
        for (let i = 0; i < inputs.length; i++) {
            data[inputs[i].name] = inputs[i].value;
        }
            let xhr = new XMLHttpRequest();
            xhr.open("POST", "/user/delete");
            xhr.onload = function (e) {
                let response = JSON.parse(e.currentTarget.response);
                if ("Error" in response) {
                    if (response.Error == null) {
                        console.log("Успешно удалено");
                        location.reload();
                    } else {
                        console.log(response.Error);
                    }
                } else {
                    console.log("Некорректные данные");
                }
            };
            xhr.send(JSON.stringify(data));
            console.log(data)
        } 

        function Edit() {
            let inputs = document.querySelectorAll(".form > input");
            let data = {};
            for (let i = 0; i < inputs.length; i++) {
                data[inputs[i].name] = inputs[i].value;
            }
                let xhr = new XMLHttpRequest();
                xhr.open("POST", "/user/editsss");
                xhr.onload = function (e) {
                    let response = JSON.parse(e.currentTarget.response);
                    if ("Error" in response) {
                        if (response.Error == null) {
                            alert("Изменено!");
                            console.log("Успешно удалено");
                        } else {
                            console.log(response.Error);
                        }
                    } else {
                        console.log("Некорректные данные");
                    }
                };
                alert("Изменено!");
                xhr.send(JSON.stringify(data));
                window.location.href = 'http://127.0.0.1:8080/teacher'; 
                console.log(data)
            } 

            function Scroll() {
                
				var hiddenElement = document.getElementById("box");
  hiddenElement.scrollIntoView({ block: "center", behavior: "smooth" });



			}