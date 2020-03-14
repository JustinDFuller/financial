import React, { useState } from "react";
import * as service from "../service_pb";
import {
  CreateUser,
  CreateAccount,
  CreateContribution,
  CreateGoal
} from "./forms";
import { Chart } from "./Chart";

const STATE_USER = 3;
const STATE_ACCOUNTS = 30;
const STATE_CONTRIBUTIONS = 60;
const STATE_GOALS = 80;
const STATE_DONE = 100;

export function Form() {
  const [step, setStep] = useState(STATE_DONE);
  const [user, setUser] = useState();
  const [accounts, setAccounts] = useState([]);
  const [contributions, setContributions] = useState([]);
  const [goals, setGoals] = useState([]);

  function renderStep() {
    switch (step) {
      case STATE_USER:
        return <CreateUser onDone={handleCreateUserDone} />;
      case STATE_ACCOUNTS:
        return (
          <CreateAccount
            user={user}
            onSave={handleCreateAccountOnSave}
            onDone={handleCreateAccountOnDone}
          />
        );
      case STATE_CONTRIBUTIONS:
        return (
          <CreateContribution
            accounts={accounts}
            onSave={handleCreateContributionOnSave}
            onDone={handleCreateContributionOnDone}
          />
        );
      case STATE_GOALS:
        return (
          <CreateGoal
            accounts={accounts}
            user={user}
            onSave={handleCreateGoalOnSave}
            onDone={handleCreateGoalOnDone}
          />
        );
      case STATE_DONE:
        return <Chart user={user} />;
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
    setStep(STATE_CONTRIBUTIONS);
  }

  function handleCreateContributionOnSave(contribution) {
    setContributions([...contributions, contribution]);
  }

  function handleCreateContributionOnDone() {
    setStep(STATE_GOALS);
  }

  function handleCreateGoalOnSave(goal) {
    setGoals([...goals, goal]);
  }

  function handleCreateGoalOnDone() {
    setStep(STATE_DONE);
  }

  return (
    <div className="card">
      {step !== STATE_DONE && (
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
      )}
      <div className="card-body">{renderStep()}</div>
    </div>
  );
}
