import React, {Component} from 'react';
import axios from 'axios';
import 'rc-slider/assets/index.css';
import AlertContainer from 'react-alert';
import Slider from 'rc-slider';
import './App.css';

class App extends React.Component {
    constructor(props, context) {
        super(props, context);
        this.state = {
            volume: 100,
            mute: false,
            url: 'http://192.168.1.2:8080',
        };

        this.alertOptions = {
            offset: 14,
            position: 'top left',
            theme: 'dark',
            time: 5000,
            transition: 'scale'
        };
    }

    handleOnChange = (value) => {
        this.setState({
            volume: value
        });

        const self = this;
        axios.get(self.state.url + '/volume/' + self.state.volume).then(function (response) {
            self.setState({
                mute: response.data.data.muted
            });
        }).catch(function (response) {
            self.msg.removeAll();
            self.msg.error("Couldn't connect to server.");
        });
    };

    handleChangeFinish = (val) => {
        const self = this;
        axios.get(self.state.url + '/volume/' + self.state.volume).then(function (response) {
        }).catch(function (response) {
            self.msg.removeAll();
            self.msg.error("Couldn't connect to server.");
        });
    };

    handleMute = () => {
        const self = this;
        axios.get(self.state.url + '/mute').then(function (response) {
            self.setState({mute: !self.state.mute});
        }).catch(function (response) {
            self.msg.removeAll();
            self.msg.error("Couldn't connect to server.");
        });
    };

    handlePlayPause = () => {
        const self = this;
        axios.get(self.state.url + '/playpause').then(function (Response) {
        }).catch(function (Response) {
            self.msg.removeAll();
            self.msg.errer("Couldn't connct to server");
        });
    };

    handleNextTrack = () => {
        const self = this;
        axios.get(self.state.url + '/next').then(function (Response) {
        }).catch(function (Response) {
            self.msg.removeAll();
            self.msg.errer("Couldn't connct to server");
        });
    };

    handlePrevTrack = () => {
        const self = this;
        axios.get(self.state.url + '/prev').then(function (Response) {
        }).catch(function (Response) {
            self.msg.removeAll();
            self.msg.errer("Couldn't connct to server");
        });
    };

    handleStopTrack = () => {
        const self = this;
        axios.get(self.state.url + '/stop').then(function (Response) {
        }).catch(function (Response) {
            self.msg.removeAll();
            self.msg.errer("Couldn't connct to server");
        });
    };

    componentDidMount = () => {
        const self = this;
        axios.get(self.state.url).then(function (response) {
            self.setState({
                volume: response.data.data.volume,
                mute: response.data.data.muted
            });
        }).catch(function (response) {
            self.msg.error("Couldn't connect so server.")
        });
    };


    render() {
        const style = {width: 100 + '%', height: 400, display: 'flex', justifyContent: 'center'};
        const centerStyle = {display: 'flex', justifyContent: 'center'};
        return (
            <div>
                <AlertContainer ref={a => this.msg = a} {...this.alertOptions}/>
                <h3 style={{textAlign: 'center'}}>{this.state.volume}</h3>
                <div style={style}>
                    <Slider
                        vertical
                        min={0}
                        max={100}
                        onChange={this.handleOnChange}
                        onAfterChange={this.handleChangeFinish}
                        defaultValue={50}
                        value={this.state.volume}
                    />
                </div>
                <div style={centerStyle}>
                    <button className="mute-btn"
                            onClick={this.handleMute}>{this.state.mute ? "Unmute" : "Mute" }</button>
                    <button className="mute-btn"
                            onClick={this.handleStopTrack}>Stop</button>
                </div>
                <div style={centerStyle}>
                    <button className="mute-btn"
                            onClick={this.handlePrevTrack}>Prev</button>
                    <button className="mute-btn"
                            onClick={this.handlePlayPause}>Play/Pause</button>
                    <button className="mute-btn"
                            onClick={this.handleNextTrack}>Next</button>
                </div>
            </div>
        )
    }
}

export default App;