const React = require("react");
const {
  AreaChart,
  Area,
  YAxis,
  CartesianGrid,
  Tooltip,
  Dot
} = require("recharts");
import { calculate } from "./api";

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

/* function limitToFifty(period, index, arr) {
  return (
    period.Goals ||
    index === 0 ||
    index % Math.round(arr.length / 50) === 0 ||
    index === arr.length - 1
  );
}*/

/* function withNetWorth(period) {
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
} */

/* the main page for the index route of this app */
class Chart extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  componentDidMount() {
    calculate().then(data => this.setState({ data }));
  }

  render() {
    const { data } = this.state;

    return (
      <div className="card" style={{ width: 600 }}>
        <div className="card-body">
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
        </div>
      </div>
    );
  }
}

export default Chart;
