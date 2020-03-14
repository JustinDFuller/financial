const React = require("react");
const {
  AreaChart,
  Area,
  YAxis,
  CartesianGrid,
  Tooltip,
  Dot
} = require("recharts");
import * as api from "./api";
import * as service from "../service_pb";

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
    "Net Worth": period.accountsList.reduce(function(balance, account) {
      if (account.type === service.Mode.DEBT) {
        balance -= account.balance;
      } else {
        balance += account.balance;
      }
      return balance;
    }, 0)
  };
}

/* the main page for the index route of this app */
export class Chart extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  async componentDidMount() {
    const calculate = new service.GetCalculateData()
      .setUserid(1)
      .setPeriods(500);
    const response = await api.getCalculate(calculate);
    const data = response.message.toObject();
    data.periodsList = data.periodsList.filter(limitToFifty).map(withNetWorth);
    this.setState({ data: data.periodsList });
  }

  render() {
    const { data } = this.state;

    console.log(data);

    if (data === undefined) {
      return null;
    }

    return (
      <div>
        <h5>Your forecast</h5>
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
    );
  }
}

export default Chart;
