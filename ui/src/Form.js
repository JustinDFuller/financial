import React, { useState } from "react";
import { User, Mode } from '../service_pb'
import { postUser } from './api'

export function CreateUser({ onDone }) {
  const [email, setEmail] = useState("")
  const [error, setError] = useState()

  async function handleClick() {
    const user = new User().setEmail(email)
    const response = await postUser(user)
    if (response.error !== undefined) {
      setError(response.error)
    } else {
      onDone(response)
    }
  }

  return (
    <form onSubmit={handleClick}>
      <h5 className="card-title">Sign Up</h5>
      {
        error !== undefined && <div className="alert alert-danger">{error.message}</div>
      }
      <input
        type="email"
        className="form-control"
        placeholder="What's your email?"
        value={email}
        onChange={e => setEmail(e.target.value)}
      />
      <button type="submit" className="btn btn-primary mt-3">Sign Up</button>
    </form>
  )
}

function CreateAccounts() {
  const [accounts, setAccounts] = useState([{}])

  function addAccount() {
    setAccounts([...accounts, {}])
  }

  function setAccountName(index, e) {
    const name = e.target.value;
    accounts[index].name = name;
    setAccounts(accounts);
  }

  function setAccountBalance(index, e) {
    const balance = e.target.value;
    accounts[index].balance = balance;
    setAccounts(accounts);
  }

  return (
      <form>
        <h5 className="card-title">Add Accounts</h5>
        {accounts.map((account, index) => (
          <div key={index}>
            {
              index > 0 &&   <div className="dropdown-divider my-4"></div>
            }
            <div className="form-group">
              <label>What should we call this account?</label>
              <input
                type="text"
                className="form-control mb-1"
                placeholder="Emergency Savings"
                value={account.name}
                onChange={e => setAccountName(index, e)}
              />
            </div>
            <div className="form-group">
              <label>What's the account's current balance?</label>
              <div className="input-group">
                 <div className="input-group-prepend">
                  <span className="input-group-text">$</span>
                </div> 
                <input
                  type="number"
                  className="form-control"
                  placeholder="5349.34"
                  value={account.balance}
                  onChange={e => setAccountBalance(index, e)}
                />
              </div>
            </div>
            <div className="form-group">
              <label>What type of account is it?</label>
              <select className="form-control">
                <option value={Mode.INVESTMENTS}>Investment</option>
                <option value={Mode.DEBT}>Debt</option>
              </select>
            </div>
          </div>
        ))}
        <button type="button" className="btn btn-primary" onClick={() => addAccount()}>
          Add Another Account
        </button>
      </form>
  );

}

const STATE_USER = 1
const STATE_ACCOUNTS = 25

export function Form() {
  const [step, setStep] = useState(STATE_USER)
  const [user, setUser] = useState()

  function renderStep() {
    switch(step) {
      case STATE_USER:
        return <CreateUser onDone={handleCreateUserDone} />
      case STATE_ACCOUNTS:
        return <CreateAccounts />
    }
  }

  function handleCreateUserDone(user) {
      setUser(user)
      setStep(STATE_ACCOUNTS)
  }

  return (
    <div className="card">
      <div className="progress mt-2 mx-2">
        <div className="progress-bar" role="progressbar" style={{ width: `${step}%` }} aria-valuenow={step} aria-valuemin="0" aria-valuemax="100"></div>
      </div>
      <div className="card-body">
      {
        renderStep()
      }
      </div>
    </div>
  )
}
