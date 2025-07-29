const ports = [];
let ws = null;
let reconnectTimeout = null;

function connectWebSocket() {
  if (ws) {
    console.log("[Worker] Closing existing WebSocket before reconnecting");
    ws.onopen = null;
    ws.onclose = null;
    ws.onerror = null;
    ws.onmessage = null;
    ws.close();
  }

  console.log("[Worker] Creating new WebSocket connection...");
 ws = new WebSocket("ws://localhost:8080/api/ws");

  ws.onopen = () => {
    console.log("[Worker] WebSocket connected");
    broadcast({ type: "status", status: "connected" });
  };

  ws.onmessage = (event) => {
    let data;
    try {
      data = JSON.parse(event.data);
    } catch (e) {
      console.error("[Worker] Invalid JSON received:", event.data);
      return;
    }
    console.log("[Worker] WebSocket message received:", data);
    broadcast(data);
  };

  ws.onclose = (event) => {
    console.warn(`[Worker] WebSocket closed (code: ${event.code}). Reconnecting in 3 seconds...`);
    broadcast({ type: "status", status: "disconnected" });
    reconnectTimeout = setTimeout(connectWebSocket, 3000);
  };

  ws.onerror = (err) => {
    console.error("[Worker] WebSocket error:", err);
    if (ws) ws.close(); // Trigger reconnect
  };
}

function broadcast(message, excludePort = null) {
  console.log("[Worker] Broadcasting message to ports:", message);
  ports.forEach((port) => {
    if (port !== excludePort) {
      try {
        port.postMessage(message);
      } catch (e) {
        console.error("[Worker] Broadcast error:", e);
      }
    }
  });
}

onconnect = (e) => {
  const port = e.ports[0];
  ports.push(port);
  port.start();

  // Send initial status to new port
  port.postMessage({
    type: "status",
    status: ws && ws.readyState === WebSocket.OPEN ? "connected" : "disconnected",
  });

  port.onmessage = (event) => {
    const { type, ...data } = event.data;

    switch (type) {
      case "read":
        console.log("[Worker] Read message received:", event.data);
        broadcast(event.data);
        break;

      case "login":
        // Trigger connection to backend WebSocket server
        connectWebSocket();
        break;

      case "logout":
        if (ws) ws.close();
        break;

      case "sent_message":
        broadcast(event.data);
        break;

      case "message":
      case "start_typing":
      case "stop_typing":
      case "typing":
        if (ws && ws.readyState === WebSocket.OPEN) {
          ws.send(JSON.stringify({ type, ...data }));
        } else {
          console.warn("[Worker] Cannot send message, WebSocket not connected");
        }
        break;

      default:
        console.warn("[Worker] Unknown message type:", type);
    }
  };

  port.onmessageerror = (err) => {
    console.error("[Worker] Port message error:", err);
  };

  port.onclose = () => {
    console.log("[Worker] Port closed");
    const idx = ports.indexOf(port);
    if (idx !== -1) ports.splice(idx, 1);
    console.log("[Worker] Remaining ports:", ports.length);
  };
};

console.log("[Worker] SharedWorker script loaded");
