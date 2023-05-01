import React from "react";
import AlgorithmChooser from "./AlgorithmChooser";
import ChatHistory from "./ChatHistory";

interface SideBarLayoutProps {
    // history: Object[]
}

const SideBarLayout: React.FC<SideBarLayoutProps> = ({
    // history,
}) => {

    return (
        <div className="flex flex-col h-full bg-secondary-base w-60 py-8">
            <AlgorithmChooser />
            <div className="divider"></div>
            <label className="label justify-start ml-4">Chat History</label>
            {/* array of chat history */}
            <ChatHistory message="halooooo" />
            <ChatHistory message="hehehheheh" />
        </div>
    );
};

export default SideBarLayout;