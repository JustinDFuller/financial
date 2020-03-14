import * as service from "../service_pb";

const baseURL = "http://localhost:8080";

const endpointCalculate = baseURL + "/svc/v1/calculate";
const endpointUser = baseURL + "/svc/v1/user";
const endpointAccount = baseURL + "/svc/v1/account";
const endpointContribution = baseURL + "/svc/v1/contribution"

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

export async function calculate() {
  const response = await fetch(endpointCalculate);
  const result = await tryDecode(response, service.GetCalculateResponse);
  console.log(result);
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
    body: new service.PostContributionRequest().setData(contribution).serializeBinary()
  })
  return tryDecode(response, service.PostContributionResponse);
}
