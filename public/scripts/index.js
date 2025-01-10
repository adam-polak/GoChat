const obj = { id: 0 };

var loc = window.location;
var uri = "ws:"

if(loc.protocol === "https:") {
  uri = "wss:";
}

uri += "//" + loc.host;

uri += loc.pathname + "ws"

ws = new WebSocket(uri)

ws.onopen = function() {
  console.log("Connected")
}

ws.onmessage = function(event) {
  var out = document.getElementById("output");
  out.innerHTML += `${obj.id++}: ` + event.data + "<br>";
}

setInterval(function() {
  ws.send("Hello server!");
}, 1000);
