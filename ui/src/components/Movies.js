import React, {Component} from 'react';
import PropTypes from 'prop-types';
import Movie from './Movie.js';

class Movies extends Component {  
  render() {
    const { error, isLoaded, retrying, movies, toggleWatched, deleteMovie } = this.props;
    if (error || movies === null) {
      return (
        <React.Fragment>
          <div>Error: {error}</div>
          <div>{ retrying ? "Retrying..." : ""}</div>
        </React.Fragment>
      );
    } else if (!isLoaded) {
      return <div>Loading...</div>;
    } else {
      return (
        <div>
          {movies.map(movie => (
            <Movie key={movie.id} movie={movie} 
                   toggleWatched={toggleWatched}
                   deleteMovie={deleteMovie} />
          ))}
        </div>
      );
    }
  }
}

Movies.propTypes = {
  movies: PropTypes.array.isRequired
}

export default Movies;