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
    const backendUrl = import.meta.env.VITE_BACKNED_URL

    useEffect(() => {
        // Fetch all histories from the server
        axios.get<allHistory>(`${backendUrl}/history`)
            .then((response) => {
                setHistories(response.data.data);
            })
            .catch((error) => {
                console.log(error);
            });
    }, []);

    return (
        <div className='flex-row items-stretch bg-secondary-dark w-full h-full gap-2 flex relative p-2'>
            <SideBarLayout isKMP={isKMP} setIsKMP={setIsKMP} history={histories} session={selectedSession} onClickHistory={(id) => { setSelectedSession(id) }} />
            <ChatBotLayout session={selectedSession} setSession={setSelectedSession} isKMP={isKMP}/>
        </div>
    )
}

export default App
