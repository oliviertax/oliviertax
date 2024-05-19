
function uploadimg(theform){

    theform.submit();
  
   
  
     setStatus("Loading...", "showimg");
  
    return false;
  
  }
  
  function doneloading(rezultat) {
  
      rezultat = decodeURIComponent(rezultat.replace(/\+/g,  " "));
  
   
  
      document.getElementById('showimg').innerHTML = rezultat;
  
  }
  
   
  
  function setStatus(theStatus, theloc) {
  
    var tag = document.getElementById(theloc);
  
   
  
    if (tag) {
  
          tag.innerHTML = '<b>'+ theStatus + "</b>";
  
    }
  
  }