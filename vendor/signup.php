<?php

    require_once 'connect.php';

        $path = '/front/uploads' . time() . $_FILES['Avatar']['error'];
        if (!move_uploaded_file($_FILES['Avatar']['tmp_name'], '../' . $path)) {
            $_SESSION['message'] = 'Ошибка при загрузке сообщения';
            header('Location: ../profile.php');
        }

        $pass = md5($pass);

        mysqli_query($connect, `UPDATE "User" SET "Avatar"=$1, WHERE "Login"=$2 `,$path,);





