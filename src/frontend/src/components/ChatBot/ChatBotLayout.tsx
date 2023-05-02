import React, { useState } from "react";
import ChatTextField from "./ChatTextField";
import ChatBotMessage from "./ChatBotMessage";

const ChatBotLayout: React.FC = () => {
    const [messages, setMessages] = useState<string[]>([]);

    const handleTextSubmit = (text: string) => {
        setMessages([...messages, text]);
    };

    return (
        <div className="flex flex-col h-full max-w-full min-h-screen flex-1 bg-secondary-base rounded m-2 ml-0">
            <div className="items-stretch flex-grow bg-gray-100 flex w-full flex-col px-4 py-8">
                {messages &&
                    messages.map((message, index) => (
                        <ChatBotMessage
                            key={index}
                            chatMessage={message}
                            isUser={index % 2 === 0}
                        />
                    ))}
                {!messages && <h2 className="text-white">Start chatting with chat rawr</h2>}
            </div>
            <div className="sticky bottom-0 left-0 w-full border-none md:border-transparent md:bg-vert-light-gradient bg-gray-800 md:bg-vert-dark-gradient pt-2">
                <ChatTextField onSubmit={handleTextSubmit} />
            </div>

        </div>
    );
};

export default ChatBotLayout;
