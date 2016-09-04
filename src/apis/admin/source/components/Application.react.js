var React = require('react');
var Auth = require('./Auth.react');
var AdminMenu = require('./Admin.react');

var Application = React.createClass({

	getInitialState: function () {
		return {
			isAuthorized: false
		};
	},

	render: function () {
		if (this.state.isAuthorized) {
			return (<Admin />);
		}
		return (<Auth />);
	}

});

module.exports = Application;
