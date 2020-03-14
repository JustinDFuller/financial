import React, { useState } from "react";
import * as service from "../../service_pb";
import * as api from ".././api";

export function CreateGoal({ user, accounts, onSave, onDone }) {
  const [selectedAccounts, setSelectedAccounts] = useState([]);
  const [balance, setBalance] = useState("");
  const [name, setName] = useState("");
  const [error, setError] = useState();
  const [created, setCreated] = useState();

  async function handleSubmit(e) {
    e.preventDefault();
    const goal = new service.Goal()
      .setUserid(user.getId())
      .setName(name)
      .setBalance(Number(balance))
      .setAccountidsList(selectedAccounts);

    const response = await api.postGoal(goal);
    setError(response.error);

    if (response.error === undefined) {
      goal.setId(response.message.getId());

      // Save progress
      onSave(goal);
      setCreated(name);

      // Reset
      setSelectedAccounts([]);
      setBalance("");
      setName("");
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      {created !== undefined && (
        <div className="alert alert-success">
          Successfully created goal "{created}"
        </div>
      )}
      {error !== undefined && (
        <div className="alert alert-danger">{error.getMessage()}</div>
      )}
      <h5 className="card-title">Add Goals</h5>
      <div className="form-group">
        <label>What should we call this goal?</label>
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
        <label>Which accounts does this goal apply to?</label>
        {accounts.map(function(account) {
          return (
            <div key={account.getId()} className="form-check">
              <input
                className="form-check-input"
                type="checkbox"
                checked={selectedAccounts.includes(account.getId())}
                id={account.getId()}
                onClick={e =>
                  setSelectedAccounts(accs =>
                    accs.includes(account.getId())
                      ? accs.filter(id => id !== account.getId())
                      : [...accs, account.getId()]
                  )
                }
              />
              <label className="form-check-label" htmlFor={account.getId()}>
                {account.getName()}
              </label>
            </div>
          );
        })}
      </div>
      <div className="form-group">
        <label>At what balance will this goal be met?</label>
        <div className="input-group">
          <div className="input-group-prepend">
            <span className="input-group-text">$</span>
          </div>
          <input
            type="number"
            className="form-control"
            placeholder="100"
            value={balance}
            onChange={e => setBalance(e.target.value)}
            required
          />
        </div>
        <small id="emailHelp" className="form-text text-muted">
          Goals are met when the balance of all accounts add up to the balance
          you enter here.
        </small>
      </div>
      <button type="submit" className="btn btn-primary">
        Save Goal
      </button>
      {created && (
        <button type="button" className="btn btn-link" onClick={onDone}>
          All done
        </button>
      )}
    </form>
  );
}
