import React from 'react'
import { Provider } from 'react-redux'

import getStore from './store'
import CurrencySection from './components/CurrencySection'
import ResultSection from './components/ResultSection'

const store = getStore()

function App () {
  return (
    <Provider store={store}>
      <div className="container">
        <header>
          <h1>Currency Exchange</h1>
        </header>
        <CurrencySection/>
        <ResultSection/>
      </div>
    </Provider>
  )
}

export default App
