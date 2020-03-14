import React, { useState } from "react";
import * as service from "../service_pb";
import { CreateUser, CreateAccount, CreateContribution } from './forms'

const STATE_USER = 3;
const STATE_ACCOUNTS = 30;
const STATE_CONTRIBUTIONS = 60;
const STATE_GOALS = 80

export function Form() {
  const [step, setStep] = useState(STATE_CONTRIBUTIONS);
  const [user, setUser] = useState(new service.User().setId(1));
  const [accounts, setAccounts] = useState([
    new service.Account().setName("Credit Card").setId(1),
    new service.Account().setName("Mortgage").setId(2),
    new service.Account().setName("Investments").setId(3),
    new service.Account().setName("Emergency Savings").setId(4),
  ]);
  const [contributions, setContributions] = useState([])

  function renderStep() {
    switch (step) {
      case STATE_USER:
        return <CreateUser onDone={handleCreateUserDone} />;
      case STATE_ACCOUNTS:
        return (
          <CreateAccount user={user} onSave={handleCreateAccountOnSave} onDone={handleCreateAccountOnDone} />
        );
      case STATE_CONTRIBUTIONS:
        return <CreateContribution accounts={accounts} onSave={handleCreateContributionOnSave} onDone={handleCreateContributionOnDone} />;
      default:
        break;
    }
  }

  function handleCreateUserDone(user) {
    setUser(user);
    setStep(STATE_ACCOUNTS);
  }

  function handleCreateAccountOnSave(account, nextStep) {
    setAccounts([...accounts, account]);
  }

  function handleCreateAccountOnDone() {
    setStep(STATE_CONTRIBUTIONS)
  }

  function handleCreateContributionOnSave(contribution) {
    setContributions([...contributions, contribution])
  }

  function handleCreateContributionOnDone() {
    setStep(STATE_GOALS)
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
