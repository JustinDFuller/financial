import { Error, GetCalculateResponse } from "../service_pb";

const baseURL = "http://localhost:8080"

const endpointCalculate = baseURL + "/svc/v1/calculate"

async function tryDecode(response, message) {
  const text = await response.arrayBuffer()
  const bytes = new Uint8Array(text);
  const result = {}
  
  if (!response.ok || response.status >= 400) {
    result.error = Error.deserializeBinary(bytes).toObject();
  } else {
    try {
    result.message = message.deserializeBinary(bytes).toObject();
    } catch (e) {
      result.error = Error.deserializeBinary(bytes).toObject();
    }
  }
  
  return result
}

export async function calculate() {
  const response = await fetch(endpointCalculate)
  const result = await tryDecode(response, GetCalculateResponse)
  console.log(result)
}

