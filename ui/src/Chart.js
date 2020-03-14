import React, { useEffect, useState } from "react";
import { AreaChart, Area, YAxis, CartesianGrid, Tooltip, Dot } from "recharts";
import * as api from "./api";
import * as service from "../service_pb";

function formatMoney(t, name, props) {
  if (props && props.payload.goalsList.length !== 0) {
    return [props.payload.goalsList[0].name, "You met a goal"];
  }
  return `$${Math.round(t / 1000)}k`;
}

function shouldShowDot(dot) {
  if (dot.payload.goalsList.length !== 0) {
    return <Dot {...dot} r={dot.r * 2} />;
  }

  return null;
}

function limitToFifty(period, index, arr) {
  return (
    period.goalsList.length !== 0 ||
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
export function Chart({ user }) {
  const [data, setData] = useState();
  const [error, setError] = useState();

  async function fetchData() {
    const calculate = new service.GetCalculateData()
      .setUserid(user.getId())
      .setPeriods(500);
    const response = await api.getCalculate(calculate);
    setError(response.error);

    if (response.error === undefined) {
      const data = response.message.toObject();
      const periods = data.periodsList.filter(limitToFifty).map(withNetWorth);
      console.log(periods);
      setData(periods);
    }
  }

  useEffect(function() {
    if (data === undefined && error === undefined) {
      fetchData();
    }
  });

  if (error !== undefined) {
    return <div className="alert alert-danger">{error.getMessage()}</div>;
  }

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
