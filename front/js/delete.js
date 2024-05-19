    function Delete() {
        let inputs = document.querySelectorAll(".form > input");
        let data = {};
        for (let i = 0; i < inputs.length; i++) {
            data[inputs[i].name] = inputs[i].value;
        }
            let xhr = new XMLHttpRequest();
            xhr.open("POST", "/user/delete1");
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

        
        function Relocate() {       
            location.reload();
            window.location.href = 'http://127.0.0.1:8080/reg';
        }   



    function Change() {
        let inputs = document.querySelectorAll(".form1 > input");
        let data = {};
        for (let i = 0; i < inputs.length; i++) {
            data[inputs[i].name] = inputs[i].value;
        }
            let xhr = new XMLHttpRequest();
            xhr.open("POST", "/user/change");
            xhr.onload = function (e) {
                let response = JSON.parse(e.currentTarget.response);
                if ("Error" in response) {
                    if (response.Error == null) {
                        console.log("Успешно изменено");
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

        function A1() {
            let data = {};
            data.Log = 'A1'
          let xhr = new XMLHttpRequest();
                xhr.open("POST", "/user/teacher");
                    xhr.send(JSON.stringify(data));
                    window.location.href = 'http://127.0.0.1:8080/teacher';
                    console.log(data)              
            }  
     