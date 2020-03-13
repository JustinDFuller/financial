import React, { useState } from "react";
import * as service from "../service_pb";
import * as api from "./api";

export function CreateUser({ onDone }) {
  const [email, setEmail] = useState("");
  const [error, setError] = useState();

  async function handleSubmit(e) {
    e.stopPropagation();
    const user = new service.User().setEmail(email);
    const response = await api.postUser(user);
    if (response.error !== undefined) {
      setError(response.error);
    } else {
      user.setId(response.message.getId());
      onDone(user);
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <h5 className="card-title">Sign Up</h5>
      {error !== undefined && (
        <div className="alert alert-danger">{error.getMessage()}</div>
      )}
      <input
        type="email"
        className="form-control"
        placeholder="What's your email?"
        value={email}
        onChange={e => setEmail(e.target.value)}
      />
      <button type="submit" className="btn btn-primary mt-3">
        Sign Up
      </button>
    </form>
  );
}

function CreateAccounts({ user }) {
  const [name, setName] = useState();
  const [balance, setBalance] = useState();
  const [mode, setMode] = useState(service.Mode.INVESTMENTS);
  const [error, setError] = useState();

  function setAccountName(e) {
    setName(e.target.value);
  }

  function setAccountBalance(e) {
    setBalance(e.target.value);
  }

  function setAccountMode(e) {
    setMode(e.target.value);
  }

  async function handleSubmit() {
    const account = new service.Account()
      .setName(name)
      .setBalance(balance)
      .setMode(mode)
      .setUserid(user.getId());
    const response = await api.postAccount(account);
    if (response.error !== undefined) {
      setError(response.error);
    } else {
      account.setId(response.message.getId());
      console.log(account.toObject());
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <h5 className="card-title">Add Accounts</h5>
      {error !== undefined && (
        <div className="alert alert-danger">{error.getMessage()}</div>
      )}
      <div className="form-group">
        <label>What should we call this account?</label>
        <input
          type="text"
          className="form-control mb-1"
          placeholder="Emergency Savings"
          value={name}
          onChange={e => setAccountName(e)}
          required
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
            value={balance}
            onChange={e => setAccountBalance(e)}
            required
          />
        </div>
      </div>
      <div className="form-group">
        <label>What type of account is it?</label>
        <select
          onChange={setAccountMode}
          value={mode}
          required
          className="form-control"
        >
          <option value={service.Mode.INVESTMENTS}>Investment</option>
          <option value={service.Mode.DEBT}>Debt</option>
        </select>
      </div>
      <button type="submit" className="btn btn-primary">
        Add Another Account
      </button>
    </form>
  );
}

const STATE_USER = 1;
const STATE_ACCOUNTS = 25;

export function Form() {
  const [step, setStep] = useState(STATE_USER);
  const [user, setUser] = useState();
  const [accounts, setAccounts] = useState([]);

  function renderStep() {
    switch (step) {
      case STATE_USER:
        return <CreateUser onDone={handleCreateUserDone} />;
      case STATE_ACCOUNTS:
        return (
          <CreateAccounts user={user} onDone={handleCreateAccountOnDone} />
        );
    }
  }

  function handleCreateUserDone(user) {
    setUser(user);
    setStep(STATE_ACCOUNTS);
  }

  function handleCreateAccountOnDone(account) {
    setAccounts([...accounts, account]);
  }

  return (
    <div className="card">
      <div className="progress mt-2 mx-2">
        <div
          className="progress-bar"
          role="progressbar"
          style={{ width: `${step}%` }}
          aria-valuenow={step}
          aria-valuemin="0"
          aria-valuemax="100"
        ></div>
      </div>
      <div className="card-body">{renderStep()}</div>
    </div>
  );
}
