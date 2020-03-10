const React = require("react");
const {
  AreaChart,
  Area,
  YAxis,
  CartesianGrid,
  Tooltip,
  Dot
} = require("recharts");
import { Error, GetCalculateResponse } from "../service_pb";

function formatMoney(t, name, props) {
  if (props && props.payload.Goals) {
    return [props.payload.Goals[0].Name, "You met a goal"];
  }
  return `$${Math.round(t / 1000)}k`;
}

function shouldShowDot(dot) {
  if (dot.payload.Goals !== null) {
    return <Dot {...dot} r={dot.r * 2} />;
  }

  return null;
}

function limitToFifty(period, index, arr) {
  return (
    period.Goals ||
    index === 0 ||
    index % Math.round(arr.length / 50) === 0 ||
    index === arr.length - 1
  );
}

function withNetWorth(period) {
  return {
    ...period,
    "Net Worth": period.Accounts.reduce(function(balance, account) {
      if (account.Type === "Debt") {
        balance -= account.Balance;
      } else {
        balance += account.Balance;
      }
      return balance;
    }, 0)
  };
}

async function tryDecode(response, message) {
  const text = await response.arrayBuffer()
  const bytes = new Uint8Array(text);
  const result = {}
  
  if (!response.ok || response.status >= 400) {
    result.error = Error.deserializeBinary(bytes)
  } else {
    try {
    result.message = message.deserializeBinary(bytes);
    } catch (e) {
      result.error = Error.deserializeBinary(bytes)
    }
  }
  
  return result
}

/* the main page for the index route of this app */
class Chart extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  componentDidMount() {
    fetch("http://localhost:8080/svc/v1/calculate")
      .then(async function(response) {
        const result = await tryDecode(response, GetCalculateResponse)
        console.log(result.error.toObject())
        /* return getCalculateResponse.Periods.map(withNetWorth).filter(
          limitToFifty
        ); */
      })
      .then(data => this.setState({ data }));
  }

  render() {
    const { data } = this.state;

    return (
      <AreaChart
        width={600}
        height={300}
        data={data}
        margin={{ top: 25, right: 30, left: -35, bottom: 5 }}
      >
        <defs>
          <linearGradient id="netWorthGradient" x1="0" y1="0" x2="0" y2="1">
            <stop offset="5%" stopColor="#82ca9d" stopOpacity={0.8} />
            <stop offset="95%" stopColor="#82ca9d" stopOpacity={0.3} />
          </linearGradient>
        </defs>
        <YAxis
          axisLine={false}
          tickLine={false}
          tickFormatter={formatMoney}
          style={{ left: 10, top: 50, position: "relative" }}
        />
        <Tooltip
          formatter={formatMoney}
          separator=" — "
          labelFormatter={() => undefined}
        />
        <CartesianGrid vertical={false} />
        <Area
          type="basis"
          dataKey="Net Worth"
          stroke="#82ca9d"
          fillOpacity={1}
          fill="url(#netWorthGradient)"
          dot={shouldShowDot}
        />
      </AreaChart>
    );
  }
}

export default Chart;
