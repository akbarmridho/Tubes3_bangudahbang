import { useState } from "react";

interface ChatTextFieldProps {
    onSubmit: (text: string) => void;
    session: string
}

const ChatTextField: React.FC<ChatTextFieldProps> = ({ onSubmit, session }) => {
    const [text, setText] = useState("");

    const handleSubmit = () => {
        const regex = /\S+/;
        if (regex.test(text)) {
            onSubmit(text);
            setText("");
        }
    };

    const handleKeyDown = (event: React.KeyboardEvent<HTMLTextAreaElement>) => {
        if (event.key === "Enter" && !event.shiftKey) {
          event.preventDefault();
          handleSubmit();
        }
      };


    return (
        <form onSubmit={handleSubmit} className="stretch mx-2 flex flex-row gap-3 last:mb-2 md:mx-4 md:last:mb-6 lg:mx-auto lg:max-w-2xl xl:max-w-3xl">

            <div className="flex flex-col w-full py-2 flex-grow md:py-3 md:pl-4 relative border border-black/10 border-gray-900/50 text-white bg-secondary-light rounded-md shadow-[0_0_15px_rgba(0,0,0,0.10)] ">
                <textarea
                    value={text}
                    onChange={(e) => setText(e.target.value)}
                    className="textarea textarea-bordered m-0 w-full resize-none min-h-3rem h-auto overflow-hidden max-h-200 focus:border-none focus:outline-none border-0 bg-transparent p-0 pr-7 pl-2 md:pl-0"
                    placeholder="Send your question here"
                    onKeyDown={handleKeyDown}
                    rows={1}
                ></textarea>
                <button
                    className="absolute p-1 rounded-md text-gray-500 bottom-1.5 md:bottom-2.5 hover:bg-gray-100 enabled:hover:text-gray-400 hover:bg-gray-900 disabled:hover:bg-transparent right-1 md:right-2 disabled:opacity-40"

                >
                    <svg
                        stroke="currentColor"
                        fill="none"
                        strokeWidth="2"
                        viewBox="0 0 24 24"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        className="h-4 w-4 mr-1"
                        height="1em"
                        width="1em"
                        xmlns="http://www.w3.org/2000/svg"
                    >
                        <line x1="22" y1="2" x2="11" y2="13"></line>
                        <polygon points="22 2 15 22 11 13 2 9 22 2"></polygon>
                    </svg>
                </button>
            </div>

        </form>
    );
};

export default ChatTextField;
