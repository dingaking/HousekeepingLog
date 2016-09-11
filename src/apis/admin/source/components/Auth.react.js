var React = require('react');
var LoginForm = require('./LoginForm.react');
var FindPW = require('./FindPW.react');
var ChangePW = require('./ChangePW.react');

var Auth = React.createClass({

	getInitialState: function () {
		return {
			mode: 'LoginForm'
		};
	},

	setAuthState: function (mode) {
		this.setState({
			mode: mode
		});
	},

    setAccessToken: function (accessToken) {
        this.props.updateAuthorized(true, accessToken);
    },

	render: function () {
		if (this.state.mode == 'LoginForm') {
			return (<LoginForm updateState={this.setAuthState} setAccessToken={this.setAccessToken} />);
		} else if (this.state.mode == 'FindPW') {
			return (<FindPW updateState={this.setAuthState} setAccessToken={this.setAccessToken} />);
		} else if (this.state.mode == 'ChangePW') {
			return (<ChangePW updateState={this.setAuthState} setAccessToken={this.setAccessToken} />);
		}
        return (<LoginForm updateState={this.setAuthState} />);
	}
});

module.exports = Auth;
