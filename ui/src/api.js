import * as service from "../service_pb";

const API_URL = process.env.REACT_APP_API_URL;

if (API_URL === undefined) {
  throw new Error("API_URL is required.");
}

const endpointCalculate = API_URL + "/svc/v1/calculate";
const endpointUser = API_URL + "/svc/v1/user";
const endpointAccount = API_URL + "/svc/v1/account";
const endpointContribution = API_URL + "/svc/v1/contribution";
const endpointGoal = API_URL + "/svc/v1/goal";

async function tryDecode(response, message) {
  const text = await response.arrayBuffer();
  const bytes = new Uint8Array(text);
  const result = {};

  if (!response.ok || response.status >= 400) {
    result.error = service.Error.deserializeBinary(bytes);
  } else {
    try {
      result.message = message.deserializeBinary(bytes);
    } catch (e) {
      result.error = service.Error.deserializeBinary(bytes);
    }
  }

  return result;
}

export async function getCalculate(calculate) {
  const response = await fetch(endpointCalculate, {
    method: "POST",
    body: new service.GetCalculateRequest().setData(calculate).serializeBinary()
  });
  const result = await tryDecode(response, service.GetCalculateResponse);
  return result;
}

export async function postUser(user) {
  const response = await fetch(endpointUser, {
    method: "POST",
    body: new service.PostUserRequest().setData(user).serializeBinary()
  });
  return tryDecode(response, service.UserResponse);
}

export async function postAccount(account) {
  const response = await fetch(endpointAccount, {
    method: "POST",
    body: new service.PostAccountRequest().setData(account).serializeBinary()
  });
  return tryDecode(response, service.PostAccountResponse);
}

export async function postContribution(contribution) {
  const response = await fetch(endpointContribution, {
    method: "POST",
    body: new service.PostContributionRequest()
      .setData(contribution)
      .serializeBinary()
  });
  return tryDecode(response, service.PostContributionResponse);
}

export async function postGoal(goal) {
  const response = await fetch(endpointGoal, {
    method: "POST",
    body: new service.PostGoalRequest().setData(goal).serializeBinary()
  });
  return tryDecode(response, service.PostGoalResponse);
}
