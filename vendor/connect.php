<?php
$connect_data = "host=localhost port=5432 dbname=board user=postgres password=123";
$db_connect = pg_connect($connect_data);
if (!$db_connect) {
    die("Ошибка подключения: " );
}
echo "Подключение к БД прошло успешно.";
