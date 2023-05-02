import './App.css'
import ChatBotLayout from './components/ChatBot/ChatBotLayout'
import SideBarLayout from './components/SideBar/SideBarLayout'

function App() {
    return (
    <div className='bg-secondary-dark w-screen h-full flex relative'>
        <SideBarLayout />
        <ChatBotLayout  />
    </div>
    )
}

export default App
