import React from 'react'

interface ChatBotMessageProps {
  chatMessage: string
  isUser: boolean
}

const ChatBotMessage: React.FC<ChatBotMessageProps> = ({
  chatMessage, isUser
}) => {
  const message = chatMessage.replace(/\n/g, '<br/>')

  return (
    <div className="bg-secondary-base px-4">
      {!isUser &&
        <div className="chat chat-start text-left">
          <div className="chat-image avatar">
            <div className="w-10 rounded-full">
              <img src="https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png" />
            </div>
          </div>
          <div className="chat-bubble bg-primary-base text-white">
            <div dangerouslySetInnerHTML={{ __html: message }} />
          </div>
        </div>
      }
      {isUser &&
        <div className="chat chat-end text-left">
          <div className="chat-image avatar">
            <div className="w-10 rounded-full">
              <img src="https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png" />
            </div>
          </div>
          <div className="chat-bubble bg-secondary-light text-white">
            <div dangerouslySetInnerHTML={{ __html: message }} />
          </div>
        </div>
      }
    </div>
  )
}

export default ChatBotMessage
