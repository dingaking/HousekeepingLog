var React = require('react');
var Auth = require('./Auth.react');
var Admin = require('./Admin.react');

var Application = React.createClass({

	getInitialState: function () {
		return {
			isAuthorized: false,
            accessToken: ''
		};
	},

    setIsAuthorized: function (isAuthorized, accessToken) {
        this.setState({
            isAuthorized: isAuthorized,
            accessToken: accessToken
        });
    },

	render: function () {
		if (this.state.isAuthorized) {
			return (<Admin updateAuthorized={this.setIsAuthorized} />);
		}
		return (<Auth updateAuthorized={this.setIsAuthorized} />);
	}

});

module.exports = Application;
