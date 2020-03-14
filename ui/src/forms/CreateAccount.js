import React, { useState } from "react";
import * as service from "../../service_pb";
import * as api from ".././api";

export function CreateAccount({ user, onSave, onDone }) {
  const [name, setName] = useState("");
  const [balance, setBalance] = useState("");
  const [interestRate, setInterestRate] = useState("");
  const [mode, setMode] = useState("");
  const [addInterestEveryNPeriods, setAddInterestEveryNPeriods] = useState("");
  const [error, setError] = useState();
  const [created, setCreated] = useState();

  async function handleSubmit(e) {
    e.preventDefault();
    const account = new service.Account()
      .setName(name)
      .setMode(mode)
      .setBalance(Number(balance))
      .setInterestrate(Number(interestRate) / 100)
      .setAddinteresteverynperiods(Number(addInterestEveryNPeriods))
      .setUserid(user.getId());
    const response = await api.postAccount(account);
    setError(response.error);
    if (response.error === undefined) {
      account.setId(response.message.getId());

      // Save Progress
      onSave(account);
      setCreated(name);

      // Reset
      setName("");
      setMode("");
      setBalance("");
      setInterestRate("");
      setAddInterestEveryNPeriods("");
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      {created !== undefined && (
        <div className="alert alert-success">
          Successfully Created {created}
        </div>
      )}
      {error !== undefined && (
        <div className="alert alert-danger">{error.getMessage()}</div>
      )}
      <h5 className="card-title">Add Accounts</h5>
      <div className="form-group">
        <label>What should we call this account?</label>
        <input
          type="text"
          className="form-control mb-1"
          placeholder="Emergency Savings"
          value={name}
          onChange={e => setName(e.target.value)}
          required
        />
      </div>
      <div className="form-group">
        <label>What type of account is it?</label>
        <select
          onChange={e => setMode(Number(e.target.value))}
          value={mode}
          required
          className="form-control"
        >
          <option value="" disabled>
            Select an account type
          </option>
          <option value={service.Mode.INVESTMENTS}>Investment</option>
          <option value={service.Mode.DEBT}>Debt</option>
        </select>
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
            onChange={e => setBalance(e.target.value)}
            required
          />
        </div>
      </div>
      <div className="form-group">
        <label>What's this account's interest rate?</label>
        <div className="input-group">
          <input
            type="number"
            value={interestRate}
            onChange={e => setInterestRate(e.target.value)}
            placeholder="4.5"
            className="form-control"
            required
          />
          <div className="input-group-append">
            <span className="input-group-text">%</span>
          </div>
        </div>
      </div>
      <div className="form-group">
        <label>How often is interest calculated?</label>
        <div className="input-group">
          <div className="input-group-prepend">
            <span className="input-group-text">Every</span>
          </div>
          <input
            type="number"
            className="form-control"
            placeholder="12"
            value={addInterestEveryNPeriods}
            onChange={e => setAddInterestEveryNPeriods(e.target.value)}
            required
          />
          <div className="input-group-append">
            <span className="input-group-text">Months</span>
          </div>
        </div>
        {mode === service.Mode.INVESTMENTS && (
          <small id="emailHelp" className="form-text text-muted">
            Interest is usually calculated every one or twelve months.
          </small>
        )}
        {mode === service.Mode.DEBT && (
          <small id="emailHelp" className="form-text text-muted">
            Interest is usually calculated every one month.
          </small>
        )}
      </div>
      <button type="submit" className="btn btn-primary">
        Save Account
      </button>
      {created && (
        <button type="button" className="btn btn-link" onClick={onDone}>
          All done
        </button>
      )}
    </form>
  );
}
