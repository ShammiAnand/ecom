# ecom

## UN-AUTHENTICATED:

- `POST /api/v1/register`
- `POST /api/v1/login`
- `GET /api/v1/products`

## AUTHENTICATED:

- `GET /api/v1/orders`
- `POST /api/v1/cart/checkout`
- `POST /api/v1/cart/cancel`

### everything server related is written using `net/http` library alone
