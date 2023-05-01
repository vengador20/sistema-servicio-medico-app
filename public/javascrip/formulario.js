function iniciarMap(){
    var coord = {lat:19.502654 ,lng: -96.8827172};
    var map = new google.maps.Map(document.getElementById('map'),{
      zoom: 10,
      center: coord
    });
    var marker = new google.maps.Marker({
      position: coord,
      map: map
    });
}