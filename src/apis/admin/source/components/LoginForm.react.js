var React = require('react');
var Axios = require('axios');
var Cookie = require('react-cookie');

var LoginForm = React.createClass({

	getInitialState: function () {
		return {
			userid: '',
			password: ''
		};
	},

	useridChange: function (e) {
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
		Axios.post('/api/admin/userR', {
			userid: this.state.userid,
			password: this.state.password
		}).then(response => {
            var result = JSON.parse(response.request.response);
            if (result.result == "success" && result.access_token == "") {
                this.props.updateState("ChangePW");
                Cookie.save('userid', this.state.userid, { path: '/' });
                Cookie.save('access_token', this.state.userid, { path: '/' });
            } else if (result.result == "success" && result.access_token != "") {
                this.props.setAccessToken(result.access_token);
                Cookie.save('userid', this.state.userid, { path: '/' });
                Cookie.save('access_token', result.access_token, { path: '/' });
            }
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
					<input type="submit" 
						name="login" 
						className="login loginmodal-submit" 
						value="Login" 
						onClick={this.requestLogin} />
				</form>
					
				<div className="login-help">
					<a href="#">Register</a> - <a href="#">Forgot Password</a>
				</div>
			</div>);
	}
});

module.exports = LoginForm;
