import React, {Component} from 'react';
import axios from 'axios';
import 'rc-slider/assets/index.css';
import Slider from 'rc-slider';
import './App.css';

class App extends Component {
    constructor(props, context) {
        super(props, context);
        this.state = {
            volume: 100,
            mute: false,
            url: 'http://192.168.1.9:8080'
        }
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
        });
    };

    handleChangeFinish = (val) => {
        const self = this;
        axios.get(self.state.url + '/volume/' + self.state.volume).then(function (response) {
        }).catch(function (response) {
        });
    };

    handleMute = () => {
        const self = this;

        axios.get(self.state.url + '/mute').then(function (response) {
            self.setState({mute: !self.state.mute});
        });

    };

    componentDidMount() {
        const self = this;
        axios.get(self.state.url).then(function (response) {
            self.setState({
                volume: response.data.data.volume,
                mute: response.data.data.muted
            });
        });
    }

    render() {
        const style = {width: 100 + '%', height: 400, display: 'flex', justifyContent: 'center'};
        const centerStyle = {display: 'flex', justifyContent: 'center'};
        return (
            <div>

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
                </div>
            </div>
        )
    }
}

export default App;