import React from 'react';
import { Provider } from 'react-redux';
import { hot } from 'react-hot-loader/root';
import store from './redux/index';
import IndexLayout from './components/layout/index'
import { HashRouter as Router, Switch, Route, Redirect } from 'react-router-dom'

class App extends React.Component {
	render() {
		return (
      <React.StrictMode>
        <Provider store={store}>
          <Router>
                <Switch>
                    <Route render={ () => <IndexLayout /> } />
                </Switch>
            </Router>
        </Provider>
      </React.StrictMode>
		);
	}
}

export default hot(App);
