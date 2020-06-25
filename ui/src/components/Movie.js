import React, { Component } from 'react';
import PropTypes from 'prop-types';

class Movie extends Component {

  getStyle = () => ({
    textDecoration: this.props.movie.watched ? 'line-through' : 'none',
  })

  getLoadingStyle = () => ({
    display: this.props.movie.loading ? 'inline' : 'none',
  })


  render() {
    const { movie: { title, id, watched } } = this.props;
    return (
      
      <div key={title} style={rowStyle}>
        <p>
          <span style={this.getLoadingStyle()}>loading...</span>
          <input type='checkbox' onChange={this.props.toggleWatched.bind(this, id, watched)} checked={watched} /> {' '}
          <span style={this.getStyle()}>{title}</span>
          <button style={btnStyle} onClick={this.props.deleteMovie.bind(this, id)}>x</button>
        </p>
      </div>
    );
  }
}

Movie.propTypes = {
  movie: PropTypes.object.isRequired
}

const rowStyle = {
  backgroundColor: '#f4f4f4f4',
  padding: '10px',
  borderBottom: '1px #ccc dotted',
}

const btnStyle = {
  background: '#ff0000',
  color: '#fff',
  border: 'none',
  padding: '5px 9px',
  borderRadius: '50%',
  cursor: 'pointer',
  float: 'right',
}

export default Movie;