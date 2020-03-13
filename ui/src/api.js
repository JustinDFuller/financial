import {
  Error,
  UserResponse,
  GetCalculateResponse,
  PostUserRequest,
  PostAccountRequest,
  PostAccountResponse
} from "../service_pb";

const baseURL = "http://localhost:8080";

const endpointCalculate = baseURL + "/svc/v1/calculate";
const endpointUser = baseURL + "/svc/v1/user";
const endpointAccount = baseURL + "/svc/v1/account";

async function tryDecode(response, message) {
  const text = await response.arrayBuffer();
  const bytes = new Uint8Array(text);
  const result = {};

  if (!response.ok || response.status >= 400) {
    result.error = Error.deserializeBinary(bytes);
  } else {
    try {
      result.message = message.deserializeBinary(bytes);
    } catch (e) {
      result.error = Error.deserializeBinary(bytes);
    }
  }

  return result;
}

export async function calculate() {
  const response = await fetch(endpointCalculate);
  const result = await tryDecode(response, GetCalculateResponse);
  console.log(result);
}

export async function postUser(user) {
  const response = await fetch(endpointUser, {
    method: "POST",
    body: new PostUserRequest().setData(user).serializeBinary()
  });
  return tryDecode(response, UserResponse);
}

export async function postAccount(account) {
  const response = await fetch(endpointAccount, {
    method: "POST",
    body: new PostAccountRequest().setData(account).serializeBinary()
  });
  return tryDecode(response, PostAccountResponse);
}
