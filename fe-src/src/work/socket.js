
function websocket() {

    let port = "8000"
    let socket = new WebSocket("ws://127.0.0.1:" + port + "/ws");
    console.log("Attempting Connection...");

    socket.onopen = () => {
        console.log("Successfully Connected");
        socket.send("Hi From the Client!")
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
        socket.send("Client Closed!")
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };

    socket.onmessage = (event) => {
        console.log("------from server------");
        console.log(event.data);
        console.log(event);
        console.log("-----------------------");
    };

}

export default websocket;
