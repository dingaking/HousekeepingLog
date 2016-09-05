var React = require('react');
var Axios = require('axios');

var LoginForm = React.createClass({

	getInitialState: function () {
		return {
			userid: '',
			password: ''
		};
	},

	useridChange: function (e) {
		console.log(e.target.value);
		this.setState({
			userid: e.target.value
		})
	},

	passwordChange: function (e) {
		this.setState({
			password: e.target.value
		})
	},

	requestLogin : function (e) {
		e.preventDefault(); 
		console.log(this.state.userid);
		Axios.post('/api/authC', {
			userid: this.state.userid,
			password: this.state.password
		}).then(response => {
			console.log(response.request.response);
        });
	},

	render: function () {
		return (
			<div className="loginmodal-container">
				<h1>Login to Your Account</h1><br />
				<form>
					<input 
						type="text" 
						name="user" 
						placeholder="Username" 
						onChange={this.useridChange} 
						value={this.state.userid} />
					<input 
						type="password" 
						name="pass" 
						placeholder="Password" 
						onChange={this.passwordChange} 
						value={this.state.password} />
					<input type="submit" name="login" className="login loginmodal-submit" value="Login" onClick={this.requestLogin} val={this.state.password} />
				</form>
					
				<div className="login-help">
					<a href="#">Register</a> - <a href="#">Forgot Password</a>
				</div>
			</div>);
	}
});

module.exports = LoginForm;
