// https://jsfiddle.net/alidingling/xqjtetw0/
//
const {LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend} = Recharts;

fetch('http://127.0.0.01:8080/svc/v1/user/calculate')
	.then(res => res.json())
	.then(function (res) {
  	const data = res.map(function (period) {
      return {
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
    .filter(function (val, index, arr) {
      return index === 0 || (index % Math.round(arr.length / 50) === 0) || index === (arr.length - 1)
    });

    const SimpleLineChart = React.createClass({
      render () {
        return (
          <LineChart width={600} height={300} data={data}
                margin={{top: 5, right: 30, left: 20, bottom: 5}}>
           <XAxis dataKey="name" />
           <YAxis />
           <Tooltip/>
           <Legend />
           <Line type="monotone" dataKey="Net Worth" stroke="green" />
          </LineChart>
        );
      }
  })

    ReactDOM.render(
      <SimpleLineChart />,
      document.getElementById('container')
    );
  })
  




