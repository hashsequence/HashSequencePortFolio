import React, { Component } from 'react';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import { Jumbotron, Button } from 'reactstrap';
import btwist from './btwist.MOV'


const NameHeader = (props) => {
  return (
    <div>
      <Jumbotron>
        <h1 className="display-3">My name is Avery Wong</h1>
        <p className="lead">This is my Portfolio Website</p>
        <hr className="my-2" />
        <p>Future Side Projects Will be posted here, so I will try to keep this updated!</p>
        <p className="lead">
          <Button color="primary">Check Out My Projects</Button>
        </p>
        <img width="400"  src={require('./profile.jpg')} />
      </Jumbotron>
    </div>
  );
};

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
        <NameHeader />
        <video width="1200"  id="background-video" className='videoTag' autoPlay loop muted>
         <source src={btwist} type='video/mp4' />
       </video>
        </header>
      </div>
    );
  }
}

export default App;
