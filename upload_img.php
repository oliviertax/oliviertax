<?php

$savefolder = 'imgs'; // folder for upload

$max_size = 250;   // maxim size for image file, in KiloBytes

 

          // Allowed image types

$allowtype = array('bmp', 'gif', 'jpg', 'jpeg', 'gif', 'png');

 

$rezultat = '';

 

if (isset ($_FILES['myfile'])) {

    $type = end(explode(".", strtolower($_FILES['myfile']['name'])));

  if (in_array($type, $allowtype)) {

    // check its size

       if ($_FILES['myfile']['size']<=$max_size*1000) {

      if ($_FILES['myfile']['error'] == 0) {

        $thefile = $savefolder . "/" . $_FILES['myfile']['name'];

        // if the file can`t be uploaded, return a message

        if (!move_uploaded_file ($_FILES['myfile']['tmp_name'], $thefile)) {

          $rezultat = 'The file can`t be uploaded, try again';

        }

        else {

          $rezultat = '<img src="'.$thefile.'" />';

          echo 'The image was successfully loaded';

        }

      }

    }

       else { $rezultat = 'The file <b>'. $_FILES['myfile']['name']. '</b> exceeds the maximum permitted size <i>'. $max_size. 'KB</i>'; }

  }

  else { $rezultat = 'The file <b>'. $_FILES['myfile']['name']. '</b> has not an allowed extension'; }

}

 

$rezultat = urlencode($rezultat);

echo '<body onload="parent.doneloading(\''.$rezultat.'\')"></body>';

