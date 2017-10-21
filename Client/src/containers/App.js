import React, { Component } from 'react';
import axios from 'axios';
import Slider from 'react-rangeslider';

class App extends Component {
  constructor(props, context) {
    super(props, context)
    this.state = {
      volume: 100,
      url: 'http://192.168.1.9:8080'
    }
  }

  handleOnChange = (value) => {
    console.log(value);
    this.setState({
      volume: value
    })
    const self = this;
    axios.get(self.state.url + '/volume/' + self.state.volume).then(function(response) {
      console.log(response.data.data)
    }).catch(function(response){
    });
  }

  handleChangeFinish = (val) => {
    console.log(val);
  }

  handleMute = () => {
    const self = this;
    axios.get(self.state.url + '/mute').then(function(response) {

    });
  }

  componentDidMount() {
    const self = this;
    axios.get(self.state.url).then(function(response){
      self.setState({volume: response.data.data.volume});
    });
  }

  render() {
    let { volume } = this.state
    return (
      <div>
      <h3 style={{textAlign: 'center'}}>{this.state.volume}</h3>
      <Slider
        value={volume}
        orientation="vertical"
        step={1}
        onChange={this.handleOnChange}
        onChangeComplete={this.handleChangeFinish}
      />
      <center><button onClick={this.handleMute}>Mute</button></center>
      </div>
    )
  }
}

export default App;