import React, { useState, useEffect } from 'react'
import './App.css'
import ChatBotLayout from './components/ChatBot/ChatBotLayout'
import SideBarLayout from './components/SideBar/SideBarLayout'
import axios from 'axios'

interface allHistory {
  data: History[]
}

interface History {
  id: number
  created_at: Date
  user_query: string
  response: string
  session_id: string
}

function App (): JSX.Element {
  const [histories, setHistories] = useState<History[]>([])
  const backendUrl: string = import.meta.env.VITE_BACKEND_URL
  const [selectedSession, setSelectedSession] = useState<string>('')
  const [currentSession, setCurrentSession] = useState<string>('')
  const [isKMP, setIsKMP] = useState<boolean>(false)

  useEffect(() => {
    // Fetch all histories from the server
    axios.get<allHistory>(`${backendUrl}/history`)
      .then((response) => {
        setHistories(response.data.data)
      })
      .catch((error) => {
        console.log(error)
      })
  }, [])

  const handleNewSession = (session: string): void => {
    setCurrentSession(currentSession)
  }

  return (
        <div className='flex-row items-stretch bg-secondary-dark w-full h-full gap-2 flex relative p-2'>
            <SideBarLayout isKMP={isKMP} setIsKMP={setIsKMP} histories={histories} currentSession={currentSession} session={selectedSession} onClickHistory={(id) => { setSelectedSession(id) }} />
            <ChatBotLayout session={selectedSession} setSession={setSelectedSession} setCurrentSession={handleNewSession} isKMP={isKMP} onNewSession={() => {
              axios.get<allHistory>(`${backendUrl}/history`)
                .then((response) => {
                  setHistories(response.data.data)
                })
                .catch((error) => {
                  console.log(error)
                })
            }} />
        </div>
  )
}

export default App
