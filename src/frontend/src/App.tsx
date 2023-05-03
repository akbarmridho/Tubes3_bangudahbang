import { useEffect, useState } from 'react';
import axios from 'axios';
import './App.css'
import ChatBotLayout from './components/ChatBot/ChatBotLayout'
import SideBarLayout from './components/SideBar/SideBarLayout'

interface allHistory {
    data: History[]
}

interface History {
    id: number;
    created_at: Date;
    user_query: string;
    response: string;
    session_id: string;
}

function App() {
    const [histories, setHistories] = useState<History[]>([]);
    const [selectedSession, setSelectedSession] = useState<string>('');
    const [isKMP, setIsKMP] = useState<boolean>(false);

    useEffect(() => {
        // Fetch all histories from the server
        axios.get<allHistory>("http://localhost:8080/history")
            .then((response) => {
                setHistories(response.data.data);
            })
            .catch((error) => {
                console.log(error);
            });
    }, []);

    return (
        <div className='bg-secondary-dark w-screen h-full min-h-screen gap-2 flex relative p-2'>
            <SideBarLayout isKMP={isKMP} setIsKMP={setIsKMP} history={histories} onClickHistory={(id) => { setSelectedSession(id) }} />
            <ChatBotLayout session={selectedSession} setSession={setSelectedSession} />
        </div>
    )
}

export default App
