import React, { Component } from 'react';
import { BrowserRouter as Router, Route} from 'react-router-dom';
import Header from './components/layout/Header.js';
import Healthz from './components/pages/Healthz.js';
import About from './components/pages/About.js';
import Movies from './components/Movies.js';
import { AddMovie } from './components/AddMovie';
import './App.css';
import Axios from 'axios';
import config from './config'

const apiHost = config.apiHost

class App extends Component {
  state = {
    movies: {
      data: [],
      isLoaded: false,
      retrying: false,
      error: ""
    }
  }

  componentDidMount() {
    this.getMovies()
  }

  getMovies = () => {
    Axios.get(`http://${apiHost}/movies`)
    .then(r => {
      const newState = this.state;
      newState.movies = {
        data: r.data,
        isLoaded: true,
        retrying: false,
        error: "",
      }
      this.setState(newState)
    })
    .catch(err => {
      console.log(err)
      const newState = this.state;
      newState.movies.isLoaded = false;
      newState.movies.error = "Could not retrieve movies.";
      newState.movies.retrying = true;
      this.setState(newState)
      setTimeout(() => {this.getMovies()}, 2000)
    })
  }

  findMovieByID = (id) => {
    const movies = this.state.movies.data
    for (let i = 0; i < movies.length; i++) {
      if ( movies[i].id === id ) {
        return movies[i]
      }
    }
    return null
  }

  movieIsLoading = (id) => {
    const newState = this.state
    newState.movies.data = this.state.movies.data.map(movie => {
      if (movie.id === id) {
        movie.loading = true
      }
      return movie;
    });
    this.setState(newState)
  }

  toggleWatched = (id, watched) => {
    if (id === null || id === "") {
      alert("id is required")
      return
    } else if (watched === null || watched === "") {
      alert("watched is required")
      return
    }
    
    Axios({
      method: "patch",
      url: `http://${apiHost}/movies/${id}/watched/${!watched}`
    }).then(r => {
      const newState = this.state;
      newState.movies.data = this.state.movies.data.map(movie => {
        if (movie.id === id) {
          movie.watched = !movie.watched
          movie.loading = false
        }
        return movie;
      });
      this.setState(newState)
    }).catch(err => {
      alert(err)
      // TODO: Handle 500 with a retry. Add loading icon for each movie.
    })

  }

  deleteMovie = (id) => {
    if (id === null || id === "") {
      alert("id is required")
      return
    }

    Axios({
      method: 'delete',
      url: `http://${apiHost}/movies/${id}`
    }).then(r => {
      const newState = this.state
      newState.movies.data = this.state.movies.data.filter(movie => movie.id !== id)
      this.setState(newState);
    }).catch(err => {
      alert(err)
      // TODO: Handle errors
    })

  }

  addMovie = (title) => {
    if (title === null || title === "") {
      alert("title is required")
      return
    }
    
    Axios({
      method: 'POST',
      url: `http://${apiHost}/movies?title=${title}`
    }).then( r => {
      const newState = this.state
      newState.movies.data.unshift(r.data.movie)
      this.setState(newState)
    }).catch( err => {
      alert(err)
      // TODO: Handle errors
    })
  }

  render() {
    return (
      <Router>
        <div className="App">
          <div className="container">
            <Header />
            <Route exact path="/" render={ props => (
              <React.Fragment>
                <AddMovie addMovie={this.addMovie} />
                <Movies movies={this.state.movies.data} 
                        error={this.state.movies.error}
                        retrying={this.state.movies.retrying}
                        isLoaded={this.state.movies.isLoaded}
                        toggleWatched={this.toggleWatched}
                        deleteMovie={this.deleteMovie}/>
              </React.Fragment>
            )}/>
            <Route path="/health" component={Healthz}/>
            <Route path="/about" component={About}/>
          </div>
        </div>
      </Router>
    );
  }
}

export default App;
