import React, { Component } from "react";
import ApolloClient, { gql } from "apollo-boost";
import { Habits } from "./habits/Habits";

export class App extends Component {
  client = new ApolloClient({
    uri: `${process.env.REACT_APP_API_URL}/query`
  });
  state = {
    habits: []
  };

  componentDidMount() {
    this.requestProgrammers();
  }

  requestProgrammers() {
    this.client
      .query({
        query: gql`
                {
                    habits() {
                        name,
                        details{
                            startDate,
                            values,
                        }
                    }
                }
            `
      })
      .then(result =>
        this.setState({
          habits: result.data.habits
        })
      );
  }

  render() {
    return (
      <div className="container collection">
        <Habits habits={this.state.habits} />
      </div>
    );
  }
}
