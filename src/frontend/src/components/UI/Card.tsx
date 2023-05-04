import React from 'react'

interface CardProps {
  text: string
}

const Card: React.FC<CardProps> = ({
  text
}) => {
  return (
    <div className="card w-72 bg-secondary-light h-20 p-1 flex items-center justify-center shadow-xl m-1">
    <div className="card-body text-center">
      <p className="break-words">{text}</p>
    </div>
  </div>
  )
}

export default Card
