function getgps() {
    if (!navigator.geolocation) {
        document.getElementById("txt").innerHTML = "お使いのブラウザは位置情報の取得をサポートしていません。";
    } else {
        navigator.geolocation.watchPosition((position) => {
            var lat = position.coords.latitude;
            var lng = position.coords.longitude;
            var accu = position.coords.accuracy;
            var txt = document.getElementById("txt");
            txt.innerHTML = "緯度, 経度: " + lat + ", " + lng + "<br>" + "精度: " + accu;

        }, (error) => {
            switch (error.code) {
                case error.PERMISSION_DENIED:
                    txt.innerHTML = "ユーザーが位置情報の取得を拒否しました。";
                    break;
                case error.POSITION_UNAVAILABLE:
                    txt.innerHTML = "位置情報が利用できません。";
                    break;
                case error.TIMEOUT:
                    txt.innerHTML = "位置情報の取得がタイムアウトしました。";
                    break;
                case error.UNKNOWN_ERROR:
                    txt.innerHTML = "不明なエラーが発生しました。";
                    break;
            }
        }, {
            enableHighAccuracy: true
        });

    }
}


var num = 0;
var watch_id;

function test() {
    watch_id = navigator.geolocation.watchPosition(test2, function (e) { alert(e.message); }, { "enableHighAccuracy": true, "timeout": 20000, "maximumAge": 100 });
}

function clearview() {

    navigator.geolocation.clearWatch(watch_id);
    document.getElementById('position_view').innerHTML = "";
}

function test2(position) {

    var geo_text = "緯度:" + position.coords.latitude + "\n";
    geo_text += "経度:" + position.coords.longitude + "\n";
    geo_text += "高度:" + position.coords.altitude + "\n";
    geo_text += "位置精度:" + position.coords.accuracy + "\n";
    geo_text += "高度精度:" + position.coords.altitudeAccuracy + "\n";
    geo_text += "移動方向:" + position.coords.heading + "\n";
    geo_text += "速度:" + position.coords.speed + "\n";

    var date = new Date(position.timestamp);

    geo_text += "取得時刻:" + date.toLocaleString() + "\n";
    geo_text += "取得回数:" + (++num) + "\n";

    url = "http://localhost:9000/savegps?lat=" + position.coords.latitude + "&lng=" + position.coords.longitude;
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);
    xhr.send(null);

    document.getElementById('position_view').innerHTML = geo_text;

}