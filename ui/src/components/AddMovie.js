import React, { Component } from 'react'

export class AddMovie extends Component {
  state = {
    title: ''
  }

  onChange = (e) => {
    this.setState({
      [e.target.name]: e.target.value
    })
  }

  onSubmit = (e) => {
    e.preventDefault();
    this.props.addMovie(this.state.title)
    this.setState({title: ''})
  }
  
  render() {
    return (
      <form style={{display: 'flex'}}
            onSubmit={this.onSubmit}>
        <input type="text" name="title" placeholder="Add Movie Title..." 
               style={{flex:'10', padding: '5px'}} 
               value={this.state.title}
               onChange={this.onChange}/>
        <input type="submit" value="Submit" className="btn" 
               style={{flex:'1'}}/>
      </form>
    )
  }
}