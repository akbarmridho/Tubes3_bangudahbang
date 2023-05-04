import React, { useState, useEffect } from "react";
import axios from 'axios'
import ChatTextField from "./ChatTextField";
import ChatBotMessage from "./ChatBotMessage";

interface History {
    id: number;
    created_at: Date;
    user_query: string;
    response: string;
    session_id: string;
}

interface allHistory {
    data: History[]
}

interface Session {
    session_id: string;
}

interface ChatBotProps {
    session: string
    isKMP: boolean
    setSession: (sessionId: string) => void;
}

const ChatBotLayout: React.FC<ChatBotProps> = ({ session, setSession, isKMP }) => {
    const [messages, setMessages] = useState<string[]>([]);
    const backendUrl = import.meta.env.VITE_BACKEND_URL

    useEffect(() => {
        if (session) {
            axios.get<allHistory>(`${backendUrl}/history/${session}`).then((response) => {
                console.log(response.data)
                const sortedHistory = response.data.data.sort((a, b) => a.id - b.id);
                const messagesFromHistory = sortedHistory.map((history) => [history.user_query, history.response,]).flat(); // create an array of [user_query, response], and flatten the array
                setMessages(messagesFromHistory);
            });
        }
    }, [session]);

    const handleTextSubmit = (text: string) => {
        if (!session) {
            axios.post<{ data: Session }>(`${backendUrl}/session`)
                .then((response) => {
                    console.log(response.data.data.session_id)
                    setSession(response.data.data.session_id);
                    console.log(session)
                    axios.post(`${backendUrl}/query`, { session_id: response.data.data.session_id, input: text, is_kmp: isKMP })
                        .then((response) => {
                            console.log(response.data.data.response)
                            setMessages([...messages, text, response.data.data.response]);
                            console.log(messages)
                        })
                        .catch((error) => {
                            console.error(error);
                        });
                })
                .catch((error) => {
                    console.error(error);
                });
        } else {
            axios.post(`${backendUrl}/query`, { session_id: session, input: text, is_kmp: false })
                .then((response) => {
                    setMessages([...messages, text, response.data.data.response]);
                })
                .catch((error) => {
                    console.error(error);
                });
        }
    }

    return (
        <div className="flex flex-col h-full max-w-full flex-1 bg-secondary-base rounded">
            <div className="items-stretch flex-1 flex-grow bg-gray-100 flex w-full flex-col px-4 py-8 overflow-y-scroll">
                {messages.length > 0 &&
                    messages.map((message, index) => (
                        <ChatBotMessage
                            key={index}
                            chatMessage={message}
                            isUser={index % 2 === 0}
                        />
                    ))}
                {messages.length === 0 && <label className="text-white m-auto font-semibold text-xl">Welcome!<br/>Start chatting with CARL, your personal chatbot!</label>}
            </div>
            <div className="sticky bottom-0 left-0 w-full border-none md:border-transparent md:bg-vert-light-gradient bg-gray-800 md:bg-vert-dark-gradient pt-2">
                <ChatTextField onSubmit={handleTextSubmit} session={session} />
            </div>

        </div>
    );
}

export default ChatBotLayout;
