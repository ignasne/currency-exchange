import React from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import styled from 'styled-components'

const StyledErrorContainer = styled.div`
  width: 100%;
  text-align: center;
  padding: 8px 0;
`

const StyledError = styled.span`
  border: 1px solid red;
  padding: 8px;
  font-size: 12px;
`

function Error (props) {
  if (props.message.length === 0) {
    return (null)
  }

  return (
    <StyledErrorContainer>
      <StyledError>{props.message}</StyledError>
    </StyledErrorContainer>
  )
}

Error.propTypes = {
  message: PropTypes.string.isRequired
}

const stateToProps = function (state) {
  return {
    message: state.error.message
  }
}

export default connect(stateToProps)(Error)