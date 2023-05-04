import React, { useState } from 'react'
import './App.css'
import ChatBotLayout from './components/ChatBot/ChatBotLayout'
import SideBarLayout from './components/SideBar/SideBarLayout'

function App (): JSX.Element {
  const [selectedSession, setSelectedSession] = useState<string>('')
  const [isKMP, setIsKMP] = useState<boolean>(false)

  return (
        <div className='flex-row items-stretch bg-secondary-dark w-full h-full gap-2 flex relative p-2'>
            <SideBarLayout isKMP={isKMP} setIsKMP={setIsKMP} session={selectedSession} onClickHistory={(id) => { setSelectedSession(id) }} />
            <ChatBotLayout session={selectedSession} setSession={setSelectedSession} isKMP={isKMP}/>
        </div>
  )
}

export default App
