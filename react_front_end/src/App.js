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

    componentDidMount() {
     axios.get(`http://localhost:8080/GetUserData`)
     .then(res => {
       const user = res.data;
       this.setState({user});
     }).catch(err => {
       console.log('Show error notification!')
       return Promise.reject(err)
     })
 }

  render() {
          const userProfile = this.state.user.map((e) =>
            <div className="Header" key={e.toString()}>
            <h1 className="display">{e.Name.String}</h1>
            <p className="lead">{"Hey guys this my personal website!"}</p>
            <p className="lead"> {e.Summary.String}</p>
            <hr className="my-2" />
            <img alt="profilePic" width="400"  src={require('./profile.jpg')} />
            </div>);
          return userProfile;
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
   <h3> {exp.Title.String} at {exp.Company.String}</h3>
   <h4> {exp.Startdate.String.replace("T00:00:00Z","")} to {exp.Enddate.Valid === true ? exp.Enddate.String.replace("T00:00:00Z","") : "Current"}</h4>
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
