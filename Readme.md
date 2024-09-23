# Coupon Management API

## Overview
This API is designed to manage and apply discount coupons for an e-commerce platform. The project supports various coupon types such as **cart-wise**, **product-wise**, and **Buy X, Get Y (BxGy)** coupons. The system is built with flexibility in mind to allow easy addition of new coupon types in the future.

This README documents the cases that were considered, the ones implemented, limitations, assumptions, and suggestions for future improvements.

## Implemented Cases

### 1. **Cart-wise Coupons**
- **Description**: Applies a discount to the entire cart if the cart total exceeds a defined threshold.
- **Example**:
    - Condition: Cart total > 100
    - Discount: 10% off the entire cart
- **Status**: Fully implemented and tested.

### 2. **Product-wise Coupons**
- **Description**: Applies a discount to specific products in the cart.
- **Example**:
    - Condition: Product A is in the cart.
    - Discount: 20% off on Product A.
- **Status**: Fully implemented and tested.

### 3. **Buy X, Get Y (BxGy) Coupons**
- **Description**: Buy a certain quantity of products from a list and get other products for free.
- **Example**:
    - Buy 3 of Product X or Product Y, get 1 of Product Z free.
- **Status**: Implemented with a repetition limit.

## Unimplemented/Partially Implemented Cases

### 1. **Coupon Expiry**
- **Description**: Implementing expiration dates for coupons.
- **Status**: Not implemented due to time constraints, but can be added in the future by introducing an `expiry_date` field to the coupon schema.

### 2. **Advanced Validation for Multiple Coupons**
- **Description**: Ensuring that multiple coupons don’t conflict or overlap.
- **Status**: Partially implemented; currently, multiple coupons can apply to the cart but conflict resolution (like determining priority between different coupon types) is not fully handled.

### 3. **More Complex BxGy Scenarios**
- **Description**: Handling scenarios where different quantities or combinations of buy/get products are allowed (e.g., buy 3 get 2 free).
- **Status**: Partially implemented; currently handles basic "buy X, get Y free" rules, but more complex repetition and quantity combinations are not fully supported.

## Assumptions

1. **Coupon Validity**: It is assumed that all coupons provided in the system are valid and not expired unless an expiration feature is added in the future.
2. **Cart Structure**: It is assumed that the cart will have a clear and consistent structure, including product IDs, quantities, and prices.
3. **Single Coupon Application**: For simplicity, each coupon is applied independently without considering interactions or conflicts between multiple coupons (except BxGy, which accounts for a repetition limit).

## Limitations

1. **Complex Coupon Interaction**: There is no logic to handle complex interactions between multiple coupons (e.g., when two product-wise coupons apply to the same product).
2. **Performance Optimization**: While the project currently works well for in-memory data, optimization for large datasets (such as those involving persistent databases) would require further indexing and query optimizations.
3. **BxGy Flexibility**: The "Buy X, Get Y" functionality supports simple repetition limits but does not yet allow for more advanced coupon rules (e.g., progressive discounts like "buy 3, get 2 free").
4. **No Expiry Management**: Currently, there is no support for coupon expiry or scheduled deactivation.

## API Endpoints

### Coupon Management
- `POST /coupons`: Create a new coupon.
- `GET /coupons`: Retrieve all coupons.
- `GET /coupons/{id}`: Retrieve a specific coupon by its ID.
- `PUT /coupons/{id}`: Update a specific coupon by its ID.
- `DELETE /coupons/{id}`: Delete a specific coupon by its ID.

### Coupon Application
- `POST /applicable-coupons`: Fetch all applicable coupons for a given cart and calculate the total discount.
- `POST /apply-coupon/{id}`: Apply a specific coupon to the cart and return the updated cart with discounted prices for each item.

## Error Handling

- The API returns appropriate error messages for cases such as:
    - **Invalid Input**: If the input data for coupons or carts is missing or malformed.
    - **Coupon Not Found**: If a coupon is requested by ID but doesn’t exist.
    - **Conditions Not Met**: If a coupon is applied but the conditions (e.g., cart total threshold) are not met.

## Suggestions for Future Improvement

1. **Database Integration**: The current in-memory storage can be replaced with a persistent database (e.g., MySQL, MongoDB) for better scalability and data persistence.
2. **Coupon Expiry**: Add support for coupon expiration to automatically invalidate coupons based on their expiry date.
3. **Coupon Conflict Resolution**: Introduce more sophisticated rules for resolving conflicts when multiple coupons are applied to a single product or cart.
4. **Bulk Coupon Management**: Add features to manage coupons in bulk (e.g., bulk creation, updating, and deletion).

## How to Run

1. **Install Dependencies**:
    - Make sure Go is installed on your system.
    - Run `go mod tidy` to install all required dependencies.

2. **Run the Server**:
   ```bash
   go run main.go
