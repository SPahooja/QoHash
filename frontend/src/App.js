import React from 'react';
import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';
import Form from 'react-bootstrap/Form';
import FormControl from 'react-bootstrap/FormControl';
import Button from 'react-bootstrap/Button';
import Table from 'react-bootstrap/Table';
import Card from 'react-bootstrap/Card'
import axios from "axios";
import qs from 'qs'
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';

class App extends React.Component{
  readData() {
    const self = this;
    axios    
      .get("http://localhost:8000/")
      .then(function (response) {
        console.log(response.data);
        if (response.data !== null) {
          self.setState({ products: response.data });
        }
      })
      .catch(function (error) {
        console.log(error);
      });
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
    this.state = { products: [] , addr: "\\"};

    this.readData = this.readData.bind(this);
  }

  handleChange = (event) =>{
    this.setState({addr: event.target.value})
  }

  handleSubmit = () =>{
    axios({
      method: 'post',
      url: 'http://localhost:8000/',
      data: qs.stringify({
        addr: this.state.addr
      }),
      headers: {
        'content-type': 'application/x-www-form-urlencoded;charset=utf-8'
      }
    })
    this.readData()
  }
  render(){
  return (
    <div className="App" bg="dark" variant="dark">
      <Navbar bg="dark" variant="dark" expand="lg">
        <Navbar.Brand>Sahil Pahooja</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="mr-auto"></Nav>
          <Form onSubmit={this.handleSubmit} inline>
            <FormControl onSubmit={this.handleSubmit} onChange={this.handleChange} type="text" placeholder="Search" className="mr-sm-2" />
            <Button variant="outline-success" onClick={this.handleSubmit} >Search</Button>
          </Form>
        </Navbar.Collapse>
      </Navbar>
      <Card className="bg-dark text-white">
        <Card.Header as="h1">Files</Card.Header>
  <Card.Header as="h3">Found: {this.state.products.length}</Card.Header>
      </Card>
      <Table striped bordered hover variant="dark" variant="dark">
        <thead><tr><th>Name</th><th>Size (bytes)</th><th>ModTime</th>{" "}</tr></thead>
        <tbody>{this.getProducts()}</tbody>{" "}
      </Table>
    </div>
  );
}
}

export default App;
