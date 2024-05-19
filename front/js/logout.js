function Logout() {
          let data = {};
          let xhr = new XMLHttpRequest();
          xhr.open("POST", "/logout");
          xhr.onload = function (e) {
              location.reload();
              if ("Error" in response) {
                  if (response.Error == null) {
                      alert("авторизован!"); 
                      console.log("Успешно авторизован");
                      window.location.href = 'http://127.0.0.1:8080/';
                  } else {
                      console.log(response.Error);
                  }
              } else {
                  console.log("Некорректные данные");
              }
          };
          xhr.send(JSON.stringify(data));

      }   