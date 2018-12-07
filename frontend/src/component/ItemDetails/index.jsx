import React, { Component } from "react";
import { Link } from "react-router-dom";
import uniqid from "uniqid";
import axios from "axios";
import "./index.css";

class ItemDetails extends Component {
  constructor(props) {
    super(props);
    this.state = {
      item_name: this.props.location.state.item_name,
      item_description: this.props.location.state.item_description,
      reviews: [],
      reviewStatus: false
    };
  }
  render() {
    console.log("reviews", this.state.reviews);
    let review_list;
    if (!this.state.reviewStatus) {
      review_list = "No reviews found.";
    } else {
      review_list = this.state.reviews.map(review => {
        console.log("review:", review);
        return (
          <div key={uniqid()} className="user-review">
            {console.log(review.ReviewerName)}
            <div className="reviewer-name">{review.ReviewerName}</div>
            <div className="reviewer-rating">Rating: {review.Rating}/5</div>
            <div className="reviewer-summary">{review.ReviewSummary}</div>
            <div className="reviewer-description">{review.Review}</div>
          </div>
        );
      });
    }
    console.log(review_list);
    return (
      <React.Fragment>
        <div className="container">
          <div id="title">
            <h2>Item Details</h2>
          </div>
          <hr />
          <div className="row">
            <span className="col-lg-10 item-name">{this.state.item_name}</span>
            <span className="col-lg-2">
              <Link
                to="/inventory"
                className="btn btn-block btn-login rounded-0"
              >
                Go Back
              </Link>
            </span>
          </div>
          <div>Description:</div>
          <div className="description">{this.state.item_description}</div>
          <div className="row">
            <div className="ratings col-lg-3">
              <div className="heading">Avg. Ratings</div>
              <div>{"item.ratings"} / 5</div>
            </div>
            <div className="all-reviews col-lg-9">
              <div className="heading">Customer Reviews</div>
              {review_list}
            </div>
          </div>
        </div>
      </React.Fragment>
    );
  }

  componentWillMount() {
    let REVIEW_HOST_ELB = "ELB1-1383213972.us-west-2.elb.amazonaws.com";
    let PORT = 3000;
    let item_name = this.props.location.state.item_name;
    axios
      .get(`http://${REVIEW_HOST_ELB}:${PORT}/getReviews/${item_name}`)
      .then(response => {
        if (response.status === 204) {
          this.setState({ reviewStatus: false });
        }
        console.log("status:", response.status);
        console.log("Status Code GET Reviews:", response.data[0].Reviews);
        this.setState({
          reviews: response.data[0].Reviews,
          reviewStatus: true
        });
      })
      .catch(err => {
        console.log(err);
      });
  }
}

export default ItemDetails;
