// https://jsfiddle.net/rqkvs0cw/2/

const {AreaChart, Area, XAxis, YAxis, CartesianGrid, Tooltip, Legend, Dot} = Recharts;

function formatMoney (t) {
	return `$${Math.round(t / 1000)}k`
}

function shouldShowDot (dot) {
  if (dot.payload.Goals !== null) {
  	return <Dot {...dot} r={dot.r * 2} />
  }
  
  return null
}

fetch('http://127.0.0.01:8080/svc/v1/user/calculate')
	.then(res => res.json())
	.then(function (res) {
  	const data = res.map(function (period) {
      return {
      	Goals: period.Goals,
        "Net Worth": period.Accounts.reduce(function (balance, account) {
          if (account.Type === "Debt") {
            balance -= account.Balance
           } else {
            balance += account.Balance
           }
          return balance
        }, 0)
      }
    })
    /**
     * Limit to fifty, preserve start & end.
     */
    .filter(function (period, index, arr) {
      return period.Goals || index === 0 || (index % Math.round(arr.length / 50) === 0) || index === (arr.length - 1)
    });

    const SimpleLineChart = React.createClass({
      render () {
        return (
          <AreaChart width={600} height={300} data={data}
                margin={{top: 5, right: 30, left: 20, bottom: 5}}>
            <defs>
              <linearGradient id="netWorthGradient" x1="0" y1="0" x2="0" y2="1">
                <stop offset="5%" stopColor="#82ca9d" stopOpacity={0.8}/>
                <stop offset="95%" stopColor="#82ca9d" stopOpacity={0.3}/>
              </linearGradient>
            </defs>
           <YAxis axisLine={false} tickLine={false} tickFormatter={formatMoney} style={{ left: 10, top: 50, position: 'relative' }} />
           <Tooltip formatter={formatMoney} />
           <CartesianGrid vertical={false} />
           <Area type="basis" dataKey="Net Worth" stroke="#82ca9d" fillOpacity={1} fill="url(#netWorthGradient)" dot={shouldShowDot} />
          </AreaChart>
        );
      }
  })

    ReactDOM.render(
      <SimpleLineChart />,
      document.getElementById('container')
    );
  })
  
