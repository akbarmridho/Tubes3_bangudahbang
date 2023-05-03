import React from "react";
import AlgorithmChooser from "./AlgorithmChooser";
import ChatHistory from "./ChatHistory";

interface History {
    id: number;
    created_at: Date;
    user_query: string;
    response: string;
    session_id: string;
}

interface SideBarLayoutProps {
    history: History[]
    onClickHistory: (id: string) => void
    isKMP: boolean
    setIsKMP: (newVal: boolean) => void
}

const SideBarLayout: React.FC<SideBarLayoutProps> = ({
    history, onClickHistory, isKMP, setIsKMP
}) => {
    const handleOnClickHistory = (id: string) => {
        onClickHistory(id);
    }

    console.log(history)

    return (
        <div className="flex flex-col h-auto bg-secondary-base w-60 py-8 rounded">
            <AlgorithmChooser isKMP={isKMP} setIsKMP={setIsKMP} />
            <div className="divider"></div>
            <label className="label justify-start ml-4">Chat History</label>
            {history.length > 0 && history.map((item) =>
                <ChatHistory key={item.session_id} message={item.user_query} onClick={() => { handleOnClickHistory(item.session_id) }} />)}
            {history.length === 0 &&
            <label className="label">You don't have any history yet.</label>}

        </div>
    );
};

export default SideBarLayout;