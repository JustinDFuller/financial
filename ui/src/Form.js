import React, { useState } from "react";

export function CreateUser({ onDone }) {
  const [email, setEmail] = useState("")

  return (
    <form>
      <input
        type="text"
        className="form-control"
        placeholder="What's your email?"
        value={email}
        onChange={e => setEmail(e.target.value)}
      />
      <button className="btn btn-primary mt-2" onClick={onDone}>Sign Up</button>
    </form>
  )
}

function CreateAccounts() {
  const [accounts, setAccounts] = useState([])

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
        {accounts.map((account, index) => (
          <div key={index}>
            <br />
            <input
              type="text"
              className="form-control mb-1"
              placeholder="What account is this?"
              value={account.name}
              onChange={e => setAccountName(index, e)}
            />
            <input
              type="number"
              className="form-control"
              placeholder="What's the current balance?"
              value={account.balance}
              onChange={e => setAccountBalance(index, e)}
            />
          </div>
        ))}
        <br />
        <button type="button" className="btn btn-primary" onClick={() => addAccount()}>
          Add Account
        </button>
      </form>
  );

}

const STATE_USER = "USER"
const STATE_ACCOUNTS = "ACCOUNTS"

export function Form() {
  const [step, setStep] = useState(STATE_USER)

  function renderStep() {
    switch(step) {
      case STATE_USER:
        return <CreateUser onDone={updateStep(STATE_ACCOUNTS)} />
      case STATE_ACCOUNTS:
        return <CreateAccounts />
    }
  }

  function updateStep(step) {
    return function() {
      setStep(step)
    }
  }

  return (
    <div className="card">
      <div className="card-body">
      {
        renderStep()
      }
      </div>
    </div>
  )
}
