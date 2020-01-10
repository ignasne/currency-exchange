import React, { Component } from 'react'
import styled from 'styled-components'

const StyledHeader = styled.h1`
  margin: 0 auto;
  color: #43b02a;
  letter-spacing: 0px;
  text-transform: none;
  font-size: 22px;
  padding: 20px 0;
  text-align: center;
  max-width: 600px;
`

class Header extends Component {
  render () {
    return (
      <header>
        <StyledHeader>
          Currency Exchange
        </StyledHeader>
      </header>
    )
  }
}

export default Header
