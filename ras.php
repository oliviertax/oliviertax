<?php
$uploaddir = '/assets/raspisanie/';
$uploadfile = $uploaddir . basename($_FILES['raspis']['name']);

echo '<pre>';
if (move_uploaded_file($_FILES['raspis']['tmp_name'], $uploadfile)) {
    echo "Файл корректен и был успешно загружен.\n";
} else {
    echo "Возможная атака с помощью файловой загрузки!\n";
}

echo 'Некоторая отладочная информация:';
print_r($_FILES);

print "</pre>";
