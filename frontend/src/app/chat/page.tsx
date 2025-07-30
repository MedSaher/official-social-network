'use client';

import React, { useEffect, useState } from 'react';
import './chat.css';
import styles from '../../components/css/home.module.css';
import AllUsers from '@/components/user/fetchuser';


interface Message {
    id: number;
    senderId: number;
    receiverId: number;
    content: string;
}

// Replace with actual current user ID
const CURRENT_USER_ID = 5;
// Replace with actual receiver ID logic (could come from session or group metadata)
const RECEIVER_ID = 10;

export default function Chatpage({ params }: any) {
    const groupId = params.id;
    const [messages, setMessages] = useState<Message[]>([]);
    const [newMessage, setNewMessage] = useState('');
    const [sending, setSending] = useState(false);

    useEffect(() => {
        const fetchMessages = async () => {
            try {
                const res = await fetch(`/api/chats?id=${groupId}`);
                if (!res.ok) throw new Error('Failed to fetch messages');
                const data = await res.json();
                setMessages(data);
            } catch (err) {
                console.error('Error fetching chats:', err);
            }
        };

        fetchMessages();
    }, [groupId]);

    const sendMessage = async () => {
        if (!newMessage.trim()) return;

        const messagePayload = {
            senderId: CURRENT_USER_ID,
            receiverId: RECEIVER_ID,
            groupId,
            content: newMessage.trim(),
        };

        try {
            setSending(true);
            const res = await fetch('/api/chats', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(messagePayload),
            });

            if (!res.ok) throw new Error('Failed to send message');

            const createdMessage: Message = await res.json();
            setMessages((prev) => [...prev, createdMessage]);
            setNewMessage('');
        } catch (err) {
            console.error('Send error:', err);
        } finally {
            setSending(false);
        }
    };

    return (
        <div className= "chat-container">
            <div>
                <div className="chat">
                    <AllUsers />
                </div>
            </div>
            <div className="chat-wrapper">
                <div className="chat-page">
                    <div className="messages">
                        {messages.map((msg) => (
                            <div
                                key={msg.id}
                                className={`message ${msg.senderId === CURRENT_USER_ID ? 'sent' : 'received'
                                    }`}
                            >
                                {msg.content}
                            </div>
                        ))}
                    </div>

                    <div className="chat-input">
                        <input
                            type="text"
                            placeholder="Type a message..."
                            value={newMessage}
                            onChange={(e) => setNewMessage(e.target.value)}
                            onKeyDown={(e) => e.key === 'Enter' && sendMessage()}
                            disabled={sending}
                        />
                        <button onClick={sendMessage} disabled={sending}>
                            {sending ? 'Sending...' : 'Send'}
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
}