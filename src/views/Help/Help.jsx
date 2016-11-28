import React from 'react'
import AuthenticatedView from '../AuthenticatedView'

class Help extends React.Component {
  render () {
    return (
      <div>
        <h2>Help</h2>
        <div><Link to="/">Home</Link></div>
        <p>Help page stuffs go here</p>
      </div>
    )
  }
}

export default AuthenticatedView(Help)
