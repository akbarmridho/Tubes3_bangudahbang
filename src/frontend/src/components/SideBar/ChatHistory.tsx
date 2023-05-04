import React from 'react'
import { ReactSVG } from 'react-svg'
import ChatIcon from '../../assets/chat.svg'
import clsx from 'clsx'

interface ChatHistoryProps {
  message: string
  isActive: boolean
  onClick: React.MouseEventHandler
}

const ChatHistory: React.FC<ChatHistoryProps> = ({
  message,
  onClick,
  isActive
}) => {
  const truncatedMessage =
    message.length > 200 ? `${message.substring(0, 200)}...` : message

  return (
    <div
      className={clsx('w-full p-2 bg-secondary-base rounded-lg cursor-pointer transition-colors hover:bg-secondary-hover', isActive && 'bg-secondary-light')}
      onClick={onClick}
    >
      <div className="flex items-center">
        <ReactSVG src={ChatIcon} style={{ transform: 'scale(0.4)' }} className="mr-1" />
        <p className="text-sm font-medium text-gray-700 truncate">{truncatedMessage}</p>
      </div>
    </div>
  )
}

export default ChatHistory
