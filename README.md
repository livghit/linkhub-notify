<div align="center">
    <kbd>
    <img src="https://static.vecteezy.com/system/resources/previews/007/511/386/original/the-cute-octopus-is-smiling-clipart-trendy-doodle-style-vector.jpg" alt="Logo" height="200" width="200"/>
    </kbd>

# LinkHub - Notify

![Go Badge](https://img.shields.io/badge/Go-%2300ADD8?style=for-the-badge&logo=go&logoColor=white)
![Docker Badge](https://img.shields.io/badge/Docker-%234682B4?style=for-the-badge&logo=docker&logoColor=white)

# Notification Service written in Go

<div>

    This is my sideproject exploring the Golang wolrd and using Websockets to create a Notification system .

</div>

In order to start the Server run the followind inside the terminal

```bash
make build && make run
```

# How to connect to the ws via JS

```js
var input = document.querySelector("#input");
var button = document.querySelector("#send");

function WsClient(url) {
  this.ws = new WebSocket(url);
  this.eventListener = {};

  this.on = (event, cb) => (this.eventListener[event] = cb);

  this.emit = (name, data) => {
    let event = {
      event: name,
      data: data,
    };
    let rawData = JSON.stringify(event);
    this.ws.send(rawData);
  };

  this.ws.onmessage = (response) => {
    try {
      let data = JSON.parse(response.data);
      if (data) {
        let cb = this.eventListener[data.event];
        if (cb) cb(data.data);
      }
    } catch (e) {
      console.log(e);
    }
  };
}

var ws = new WsClient(
  // "ws://" + window.location.origin.replace("http://", "") + "/ws",
  "ws://localhost:3000/ws",
);

ws.on("response", (data) => {
  console.log("response: ", data);
  let msg = document.createElement("p");
  msg.innerText = data;
  document.body.appendChild(msg);
});

button.onclick = () => {
  ws.emit("message", input.value);
};
```

This snippet connects to the ws and emites a message via the input and takes the Server response and renders it

</div>
