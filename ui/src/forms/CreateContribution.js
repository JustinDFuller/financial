import React, { useState } from "react";
import * as service from "../../service_pb";
import * as api from ".././api";

export function CreateContribution({ accounts, onSave, onDone }) {
  const [selectedAccount, setSelectedAccount] = useState("");
  const [amount, setAmount] = useState("");
  const [alreadyCreated, setAlreadyCreated] = useState([]);
  const [error, setError] = useState();
  const [created, setCreated] = useState();

  async function handleSubmit(e) {
    e.preventDefault();
    const contribution = new service.Contribution()
      .setAccountid(selectedAccount)
      .setAmount(Number(amount));

    const response = await api.postContribution(contribution);
    setError(response.error);
    if (response.error === undefined) {
      contribution.setId(response.message.getId());

      // Save contribution
      onSave(contribution);
      setCreated(
        accounts.find(acc => acc.getId() === selectedAccount).getName()
      );
      setAlreadyCreated([...alreadyCreated, selectedAccount]);

      // Reset form
      setSelectedAccount("");
      setAmount("");
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      {created !== undefined && (
        <div className="alert alert-success">
          Successfully created contribution for account "{created}"
        </div>
      )}
      {error !== undefined && (
        <div className="alert alert-danger">{error.getMessage()}</div>
      )}
      <h5 className="card-title">Add Contributions</h5>
      <div className="form-group">
        <label>Which account is this contribution for?</label>
        <select
          onChange={e => setSelectedAccount(Number(e.target.value))}
          value={selectedAccount}
          required
          className="form-control"
        >
          <option value="" disabled>
            Select an account
          </option>
          {accounts
            .filter(acc => !alreadyCreated.includes(acc.getId()))
            .map(function(account) {
              return (
                <option key={account.getId()} value={account.getId()}>
                  {account.getName()}
                </option>
              );
            })}
        </select>
      </div>
      <div className="form-group">
        <label>How much will you contribute each paycheck?</label>
        <div className="input-group">
          <div className="input-group-prepend">
            <span className="input-group-text">$</span>
          </div>
          <input
            type="number"
            className="form-control"
            placeholder="100"
            value={amount}
            onChange={e => setAmount(e.target.value)}
            required
          />
        </div>
      </div>
      <button type="submit" className="btn btn-primary">
        Save Contribution
      </button>
      {created && (
        <button type="button" className="btn btn-link" onClick={onDone}>
          All done
        </button>
      )}
    </form>
  );
}
