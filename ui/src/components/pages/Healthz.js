import React, {Component} from 'react';
import config from '../../config'
import Axios from 'axios';

const apiHost = config.apiHost

class Healthz extends Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            apiHealth: {
                status: ""
            }
        }
    }

    componentDidMount() {
        console.log("EHY")
        Axios.get(`http://${apiHost}/healthz`)
        .then( res => {
            console.log(res)
            const newState = this.state
            newState.isLoaded = true
            newState.apiHealth = res.data
            this.setState(newState)
        })
        .catch( err => this.setState({
            isLoaded: true,
            err
        }))
    }

    render() {
        const { error, isLoaded, apiHealth } = this.state;
        if (error) {
            return <div>Error: {error.message}</div>;
        } else if (!isLoaded) {
            return <div>Loading...</div>
        } else {
            return (
            <div>{apiHost} is {apiHealth.status}</div>
            )
        }
    }
}

export default Healthz;