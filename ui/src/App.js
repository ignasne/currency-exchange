import React from 'react'

import Currency from './components/Currency'
import Result from './components/Result'
import Header from './components/Header'
import Error from './components/Error'
import LadingIndicator from './components/LoadingIndicator'
import styled from 'styled-components'

const StyledApp = styled.div`
  margin: 0 auto;
  max-width: 600px;
  height: 100%;
  background-color: #ffffff;
  box-shadow: 0px 0px 11px 0px rgba(0, 0, 0, 0.2);
`

function App () {
  return (
    <StyledApp>
      <Header/>
      <Currency/>
      <LadingIndicator/>
      <Error/>
      <Result/>
    </StyledApp>
  )
}

export default App
