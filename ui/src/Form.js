import React from "react";

export default class Form extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      user: {
        name: ""
      },
      accounts: []
    };
  }

  setName(e) {
    this.setState({
      user: {
        name: e.target.value
      }
    });
  }

  addAccount() {
    this.setState(state => ({
      accounts: [...state.accounts, {}]
    }));
  }

  setAccountName(index, e) {
    const name = e.target.value;
    this.setState(state => {
      const { accounts } = state;
      accounts[index].name = name;
      this.setState({
        accounts
      });
    });
  }

  setAccountBalance(index, e) {
    const balance = e.target.value;
    this.setState(state => {
      const { accounts } = state;
      accounts[index].balance = balance;
      this.setState({
        accounts
      });
    });
  }

  render() {
    return (
      <div className="card" >
        <form className="card-body">
          <input
            type="text"
            className="form-control"
            placeholder="What's your name?"
            value={this.state.user.name}
            onChange={e => this.setName(e)}
          />
          {this.state.accounts.map((account, index) => (
            <div key={index}>
              <br />
              <input
                type="text"
                className="form-control"
                placeholder="What account is this?"
                value={account.name}
                onChange={e => this.setAccountName(index, e)}
              />
              <input
                type="number"
                className="form-control"
                placeholder="What's the current balance?"
                value={account.balance}
                onChange={e => this.setAccountBalance(index, e)}
              />
            </div>
          ))}
          <br />
          <button type="button" className="btn btn-primary" onClick={() => this.addAccount()}>
            Add Account
          </button>
        </form>
      </div>
    );
  }
}
