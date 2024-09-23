import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import FindSniperGame from './FindSniperGame.jsx'

createRoot(document.getElementById('root')).render(
  <StrictMode>
        <FindSniperGame
			imageSrc="D:\sources\findthesniper.io\web\src\assets\find-the-pound-coin-v0-9f2f6z7q90qd1.webp"
			prompt="Click on the sniper!"
			/>

  </StrictMode>,
)
