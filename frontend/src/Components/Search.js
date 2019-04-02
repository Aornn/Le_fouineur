import React, { Component } from 'react';
import './Search.css';

class Search extends Component {
    constructor(props) {
        super(props);
        this.state = {
            query: '',
            data: [],
            time : 0
        };
        this.performRequest = this.performRequest.bind(this);
        this.dispData = this.dispData.bind(this);

    }
    performRequest(e) {
        var query = e.target.value
        this.setState({ query: query })
        if (query.length > 1) {
            fetch('http://localhost:8080/search?query=' + query)
                .then(response => response.json())
                .then(data => {
                    console.log(data)
                    if (data.Data == null) {
                        this.setState({ data: [], time : 0 })
                    }
                    else {
                        this.setState({ data: data.Data, time : data.Time })
                    }
                })
        }
        else {
            this.setState({ query: '', data: [], time:0 })

        }
    }

    dispData() {
        if (this.state.data.length > 0) {
            const elemdata = this.state.data.map((elem, i) => {
                let url = 'https://profile.intra.42.fr/users/'+elem.login
                return (
                <div id='user' key={i}>
                    <img src={elem.image_url} alt="stud_photo" />
                    <p>{elem.displayname}</p>
                    <p><a target="_blank" rel="noopener noreferrer" href={url}>{elem.login}</a></p>
                    <p>Level :{elem.level}</p>
                </div>
                )
            });
            return (
                <ul>
                    {elemdata}
                </ul>
            )
        }
        else {
            return (
                <p>Aucun resultat</p>
            )
        }
    }

    render() {
        return (
            <div id="container">
                <div id="inputStyle">
                    <input id="input" type="text" placeholder="Entrez votre requete" value={this.state.value} onChange={this.performRequest} />
                    <p>Temp de recherche : {this.state.time} ms</p>
                </div>
                {this.dispData()}
            </div>
        )
    }
}
export default Search;
