import React, { Component } from "react";
import "./index.css";

class ItemDetails extends Component {
  render() {
    return (
      <React.Fragment>
        <div className="container">
          <div id="title">
            <h2>Item Details</h2>
          </div>
          <hr />
          <div className="item-name">Cereals</div>
          <div>Description:</div>
          <div className="description">
            Breakfast cereal is a food product made from processed cereal grains
            and often eaten for breakfast, primarily in Western societies. It is
            most often mixed with milk, yogurt, or fruit. Some companies promote
            their products for the health benefits from eating oat-based and
            high-fiber cereals.
          </div>
          <div className="row">
            <div className="ratings col-lg-3">
              <div className="heading">Avg. Ratings</div>
              <div>4.5 / 5</div>
            </div>
            <div className="all-reviews col-lg-9">
              <div className="heading">Customer Reviews</div>
              <div className="user-review">
                <div className="user-name">Shivam Waghela</div>
                <div className="user-review-rating">Rating: 5/5</div>
                <div className="user-review-summary">Highly recommended</div>
                <div className="user-review-details">
                  Must have for everyone who wants to start their day with a
                  healthy breakfast.
                </div>
              </div>
              <div className="user-review">
                <div className="user-name">Shivam Waghela</div>
                <div className="user-review-rating">Rating: 5/5</div>
                <div className="user-review-summary">Highly recommended</div>
                <div className="user-review-details">
                  Must have for everyone who wants to start their day with a
                  healthy breakfast.
                </div>
              </div>
              <div className="user-review">
                <div className="user-name">Shivam Waghela</div>
                <div className="user-review-rating">Rating: 5/5</div>
                <div className="user-review-summary">Highly recommended</div>
                <div className="user-review-details">
                  Must have for everyone who wants to start their day with a
                  healthy breakfast.
                </div>
              </div>
            </div>
          </div>
        </div>
      </React.Fragment>
    );
  }
}

export default ItemDetails;
