import React, { useState, useEffect } from 'react'
import axios from 'axios'
import ChatTextField from './ChatTextField'
import ChatBotMessage from './ChatBotMessage'
import Card from '../UI/Card'

interface History {
  id: number
  created_at: Date
  user_query: string
  response: string
  session_id: string
}

interface allHistory {
  data: History[]
}

interface Session {
  session_id: string
}

interface ChatBotProps {
  session: string
  isKMP: boolean
  setSession: (sessionId: string) => void
  onNewSession: () => void
  setCurrentSession: (session: string) => void
}

const ChatBotLayout: React.FC<ChatBotProps> = ({ session, setSession, isKMP, onNewSession, setCurrentSession }) => {
  const [messages, setMessages] = useState<string[]>([])
  const backendUrl: string = import.meta.env.VITE_BACKEND_URL

  useEffect(() => {
    try {
      if (session !== '') {
        axios.get<allHistory>(`${backendUrl}/history/${session}`).then((response) => {
          const sortedHistory = response.data.data.sort((a, b) => a.id - b.id)
          const messagesFromHistory = sortedHistory.map((history) => [history.user_query, history.response]).flat() // create an array of [user_query, response], and flatten the array
          setMessages(messagesFromHistory)
        })
          .catch((error) => {
            console.log(error)
          })
      }
    } catch (error) {
      console.error(error)
    }
  }, [session])

  const handleTextSubmit = (text: string): void => {
    if (session === '') {
      axios.post<{ data: Session }>(`${backendUrl}/session`)
        .then((response) => {
          console.log(response.data.data.session_id)
          setSession(response.data.data.session_id)
          setCurrentSession(response.data.data.session_id)
          axios.post(`${backendUrl}/query`, { session_id: response.data.data.session_id, input: text, is_kmp: isKMP })
            .then((response) => {
              console.log(response.data.data.response)
              setMessages([...messages, text, response.data.data.response])
              console.log(messages)
            })
            .catch((error) => {
              console.error(error)
            })
          onNewSession()
        })
        .catch((error) => {
          console.error(error)
        })
    } else {
      axios.post(`${backendUrl}/query`, { session_id: session, input: text, is_kmp: false })
        .then((response) => {
          setMessages([...messages, text, response.data.data.response])
        })
        .catch((error) => {
          console.error(error)
        })
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
                {messages.length === 0 &&
                    <div className="flex flex-col m-auto text-white">
                        <label className="font-semibold text-xl">Welcome!<br />Start chatting with CARL, your personal chatbot!</label>
                        <div className="flex flex-row">
                            <div className="flex flex-col">
                                <label className="label font-semibold text-base justify-center">Features</label>
                                <Card text={'Ask questions from our database'} />
                                <Card text={'Calculator'} />
                                <Card text={'Date'} />
                                <Card text={'Add question to database'} />
                                <Card text={'Delete question from database'} />
                            </div>
                            <div className="flex flex-col">
                                <label className="label font-semibold text-base justify-center">Examples</label>
                                <Card text={'What is the capital city of Indonesia?'} />
                                <Card text={'2*5+10-3/2'} />
                                <Card text={'2-5-2023'} />
                                <Card text={'Tambahkan pertanyan xxx dengan jawaban yyy'} />
                                <Card text={'Hapus pertanyaan xxx'} />
                            </div>
                        </div>
                    </div>}
            </div>
            <div className="sticky bottom-0 left-0 w-full border-none md:border-transparent md:bg-vert-light-gradient bg-gray-800 md:bg-vert-dark-gradient pt-2">
                <ChatTextField onSubmit={handleTextSubmit} session={session} />
            </div>

        </div>
  )
}

export default ChatBotLayout
