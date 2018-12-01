import React, { Component } from "react";
import uniqid from "uniqid";
import "./index.css";

class ItemDetails extends Component {
  constructor() {
    super();
    this.state = {
      inventory: [
        {
          id: 1,
          name: "Cereals",
          description:
            "Breakfast cereal is a food product made from processed cereal grains and often eaten for breakfast, primarily in Western societies. It is most often mixed with milk, yogurt, or fruit. Some companies promote their products for the health benefits from eating oat-based and high-fiber cereals.",
          ratings: 4.5
        }
      ],
      all_reviews: [
        {
          id: 1,
          reviews: [
            {
              reviewer_name: "Shivam Waghela",
              reviewer_rating: 4.9,
              reviewer_summary: "Highly recommended",
              reviewer_description:
                "Must have for everyone who wants to start their day with a healthy breakfast."
            }
          ]
        }
      ]
    };
  }
  render() {
    let item = this.state.inventory.find(item => {
      return item.id === 1;
    });
    let item_reviews = this.state.all_reviews.find(item => {
      return item.id === 1;
    });
    console.log(item_reviews);
    let review_list = item_reviews.reviews.map(review => {
      return (
        <div key={uniqid()} className="user-review">
          <div className="reviewer-name">{review.reviewer_name}</div>
          <div className="reviewer-rating">
            Rating: {review.reviewer_rating}/5
          </div>
          <div className="reviewer-summary">{review.reviewer_summary}</div>
          <div className="reviewer-description">
            {review.reviewer_description}
          </div>
        </div>
      );
    });
    console.log(review_list);
    return (
      <React.Fragment>
        <div className="container">
          <div id="title">
            <h2>Item Details</h2>
          </div>
          <hr />
          <div className="item-name">{item.name}</div>
          <div>Description:</div>
          <div className="description">{item.description}</div>
          <div className="row">
            <div className="ratings col-lg-3">
              <div className="heading">Avg. Ratings</div>
              <div>{item.ratings} / 5</div>
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
}

export default ItemDetails;