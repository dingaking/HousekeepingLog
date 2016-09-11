var React = require('react');
var Axios = require('axios');
var Cookie = require('react-cookie');

var ChangePW = React.createClass({

    getInitialState: function () {
        return {
            oldPassword: '',
            newPassword: '',
            newPasswordRepeat: ''
        };
    },

    oldPassword: function (e) {
        this.setState({
            oldPassword: e.target.value
        });
    },

    newPassword: function (e) {
        this.setState({
            newPassword: e.target.value
        });
    },

    newPasswordRepeat: function (e) {
        this.setState({
            newPasswordRepeat: e.target.value
        });
    },

    requestChangePassword: function (e) {
        e.preventDefault();
        if (this.state.newPassword != this.state.newPasswordRepeat) {
            alert("password not match");
            return;
        }
        Axios.post('/api/admin/userU', {
            action: "1",
            userid: Cookie.load('userid'),
            old_password: this.state.oldPassword,
            new_password: this.state.newPassword
        }).then(response => {
            var result = JSON.parse(response.request.response);
            if (result.result == "success") {
                Cookie.save('access_token', result.result.access_token, {path: '/'});
                this.props.setAccessToken(result.result.access_token);
            }
        });
    },

    render: function () {
        return (
            <div className="changepw-container">
                <h1>Change User Password</h1> 
                <form>
                    <input
                        type="text"
                        name="old_password"
                        placeholder="old passowrd"
                        onChange={this.oldPassword}
                        value={this.state.oldPassword} />
                    <input
                        type="text"
                        name="new_password"
                        placeholder="new passowrd"
                        onChange={this.newPassword}
                        value={this.state.newPassword} />
                    <input
                        type="text"
                        name="new_password_repeat"
                        placeholder="new passowrd repeat"
                        onChange={this.newPasswordRepeat}
                        value={this.state.newPasswordRepeat} />
                    <input
                        type="submit"
                        name="change_password"
                        onClick={this.requestChangePassword}
                        value="Change" />
                </form>
            </div>
        );
    }
});

module.exports = ChangePW;
