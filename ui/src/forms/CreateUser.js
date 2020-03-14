import React, { useState } from "react";
import * as service from "../../service_pb";
import * as api from ".././api";

export function CreateUser({ onDone }) {
  const [email, setEmail] = useState("");
  const [error, setError] = useState();

  async function handleSubmit(e) {
    e.preventDefault();
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
