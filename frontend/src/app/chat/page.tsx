'use client';

import { useEffect, useRef } from 'react';

// You can export sendMessage to use it elsewhere in your app, or keep it inside the component as needed
let sendMessageExternal: ((msg: any) => void) | null = null;

export function sendMessage(msg: any) {
  if (sendMessageExternal) {
    sendMessageExternal(msg);
  } else {
    console.warn("SharedWorker port not ready yet");
  }
}

export default function Home() {
  const workerPort = useRef<MessagePort | null>(null);

  useEffect(() => {
    // Create SharedWorker instance
    const worker = new SharedWorker('/shared-worker.js');
    workerPort.current = worker.port;
    worker.port.start();

    // Save sendMessage function so it can be called outside (optional)
    sendMessageExternal = (message: any) => {
      // Add auth token if needed, here is an example, adjust as you want
      const token = getCookieValue('session_token');
      message.token = token;
      worker.port.postMessage(message);
    };

    // Handle incoming messages from shared worker
    worker.port.onmessage = (event) => {
      const msg = event.data;
      console.log("From shared worker:", msg);

      // Your big switch-case logic here, or call another handler function
      switch (msg.type) {
        case 'message':
          // update UI or state accordingly
          // you can also use React state/hooks for UI updates here
          console.log('New message:', msg);
          break;
        case 'start_typing':
          // show typing indicator
          break;
        // Add all other cases from your switch here...
        default:
          console.warn('Unknown message type', msg);
      }
    };

    worker.port.onmessageerror = (err) => {
      console.error("Worker port message error:", err);
    };



    // Cleanup on unmount
    return () => {
      worker.port.close();
      sendMessageExternal = null;
    };
  }, []);

  return (
    <div>
      SharedWorker Connected
      {/* Your UI components here */}
    </div>
  );
}

// Helper function to read cookies
function getCookieValue(key: string): string | null {
  const cookies = document.cookie.split("; ");
  for (let i = 0; i < cookies.length; i++) {
    const [k, v] = cookies[i].split("=");
    if (k === key) {
      return decodeURIComponent(v);
    }
  }
  return null;
}
