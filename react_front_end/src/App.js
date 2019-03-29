import React, { Component } from 'react';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import {Button, ButtonToolbar } from 'reactstrap';
import btwist from './btwist.MOV';
import axios from 'axios';

class NameHeader extends Component {

  constructor(props) {
      super(props);
      this.state = {Name : "", Summary : ""}
    }

    componentWillMount() {
     axios.get(`http://localhost:8080/GetUserData`)
       .then(res => {
         const user = res.data;
         this.setState({Name: user.Name.String, Summary: user.Summary.String});
       })
   }

    render() {
      return (
        <div id="Header">
            <h1 className="display">{this.state.Name}</h1>
            <p className="lead">{"Hey guys this my personal website!"}</p>
            <p className="lead"> {this.state.Summary}</p>
            <hr className="my-2" />
            <p></p>
            <img width="400"  src={require('./profile.jpg')} />
            <p className="lead">
            </p>
        </div>
      );
    }

};


class ExperienceList extends Component {
  constructor(props) {
      super(props);
      this.state = {exps: []};
    }

  componentWillMount() {
   axios.get(`http://localhost:8080/GetUserExperienceData`)
     .then(res => {
       const exps = res.data;
       this.setState({ exps });
     })
 }

 render() {
   const bulletPoints = this.state.exps.map((exp)=>
   <div className='Resume' key={exp.toString()}>
   {exp.Title.String}
   <p> {exp.Bulletpoints.String} </p>
   </div>);
   return (
     <ul>{bulletPoints}</ul>
   );
  }
}

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
        <ButtonToolbar>
          <Button size="large" variant="primary" href="#profile">Profile</Button>
          <Button size="large" variant="primary" href="#experience">Work Experience</Button>
          <Button size="large" variant="primary" href="#flip-video">Me Flipping</Button>
        </ButtonToolbar>
        <div id="profile">
          <NameHeader />
        </div>
        <div id="experience">
          <ExperienceList />
        </div>
        <video width="1200"  id="flip-video" className='videoTag' autoPlay loop muted>
         <source src={btwist} type='video/mp4' />
       </video>
        </header>
      </div>
    );
  }
}



export default App;
