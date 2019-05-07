var ws;

$(function(){
    connect();
})

function connect() {
    ws = new WebSocket("ws://192.168.22.117:8080/ws");
    ws.onopen = function(event){
        console.log(event);
    }
    ws.onmessage = function(event){
        var date = new Date();
        var msg = "<p>" + date.toLocaleString() + "</p><p>" + event.data + "</p>";
        $("#msgArea").append(msg);
    }
    ws.onclose = function(event){
        console.log("Disconnect from server");
    }
    ws.onerror = function(event) {
        console.log("error on websocket")
    }
}

function sendMsg() {
    var msg = $("#userMsg").val();
    ws.send(msg);
}
