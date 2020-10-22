import React from "react";
import Table from 'react-bootstrap/Table';
import axios from "axios";
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';


class FileList extends React.Component {
  readData() {
    const self = this;
    axios
      .get("http://localhost:8000/")
      .then(function (response) {
        console.log(response.data);

        self.setState({ products: response.data });
      })
      .catch(function (error) {
        console.log(error);
      });
  }

  getAlert(){
    alert('Got')
  }
  getProducts() {
    let table = [];

    for (let i = 0; i < this.state.products.length; i++) {
      table.push(
        <tr key={i}>
          <td> {this.state.products[i].name} </td>{" "}
          <td> {this.state.products[i].size} </td>{" "}
          <td> {this.state.products[i].modtime} </td>{" "}
        </tr>
      );
    }

    return table;
  }

  constructor(props) {
    super(props);
    this.readData();
    this.state = { products: [] };

    this.readData = this.readData.bind(this);
  }

  render() {
    return (
      <div>
      <style>{"table{border:1px solid black;}"}</style>
        <h1 style={{ marginBottom: "40px" }}> Inventory </h1>

        <Table>
          <thead>
            <tr>
              <th>Name </th> <th>Size </th> <th>ModTime </th>{" "}
            </tr>{" "}
          </thead>{" "}
          <tbody> {this.getProducts()} </tbody>{" "}
        </Table>{" "}
      </div>
    );
  }
}

export default FileList;
