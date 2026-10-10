import { useState } from 'react'
import heroImg from './assets/hero.png'
import reactLogo from './assets/react.svg'
import viteLogo from './assets/vite.svg'
import './App.css'
import { errAsync } from './lib/errors'
import { ping } from './ping/ping'

function App() {
  const [count, setCount] = useState(0)

  const [data, setData] = useState<string>("waiting")

  const click = async () => {
		const [res, err] = await errAsync(ping())
		if (err) {
			setData(`error: ${err.message}`)
			return
		}
		console.log(res);
		if (res && res.msg) {
			setData(res.msg)
		}
	}

  return (
    <>
      <section id="center">
        <div>
          <h1>Test Connect RPC</h1>
        </div>
        <button
          type="button"
          className="counter"
          onClick={click}
        >
          Load
        </button>
        <div>
          {data}
        </div>
      </section>

      <div className="ticks"></div>
      <section id="spacer"></section>
    </>
  )
}

export default App
