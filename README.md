# amazonparser



## POST /api/sendurls

**http request**

```json
{
 "urls": ["https://www.amazon.co.uk/gp/product/1472949293", "https://www.amazon.co.uk/gp/product/0241003008"]
}

```

**http response**

```
Your request is processing, requestID=3ab9d743
```

## GET /api/requests/{id}

**http request**

```
http://localhost:7755/api/requests/3ab9d743

```

**http response**

```json
{
    "id": "3ab9d743",
    "products": [
        {
            "url": "https://www.amazon.co.uk/gp/product/1472949293",
            "meta": {
                "title": "Lose Weight for Good: Full-flavour cooking for a low-calorie diet",
                "price": "7.99",
                "image": "https://images-na.ssl-images-amazon.com/images/I/51x7IABfcgL._SX258_BO1,204,203,200_.jpg"
            }
        },
        {
            "url": "https://www.amazon.co.uk/gp/product/0241003008",
            "meta": {
                "title": "The Very Hungry Caterpillar [Board Book]",
                "price": "3.49",
                "image": "https://images-na.ssl-images-amazon.com/images/I/51nXr9QCsIL._SY354_BO1,204,203,200_.jpg"
            }
        }
    ]
}
```

## GET /api/requests

**http request**

```
http://localhost:7755/api/requests

```

**http response**

```json
[
    {
        "id": "3ab9d743",
        "products": [
            {
                "url": "https://www.amazon.co.uk/gp/product/1472949293",
                "meta": {
                    "title": "Lose Weight for Good: Full-flavour cooking for a low-calorie diet",
                    "price": "7.99",
                    "image": "https://images-na.ssl-images-amazon.com/images/I/51x7IABfcgL._SX258_BO1,204,203,200_.jpg"
                }
            },
            {
                "url": "https://www.amazon.co.uk/gp/product/0241003008",
                "meta": {
                    "title": "The Very Hungry Caterpillar [Board Book]",
                    "price": "3.49",
                    "image": "https://images-na.ssl-images-amazon.com/images/I/51nXr9QCsIL._SY354_BO1,204,203,200_.jpg"
                }
            }
        ]
    },
    {
        "id": "92352fb7",
        "products": [
            {
                "url": "https://www.amazon.co.uk/gp/product/1472949293",
                "meta": {
                    "title": "Lose Weight for Good: Full-flavour cooking for a low-calorie diet",
                    "price": "7.99",
                    "image": "https://images-na.ssl-images-amazon.com/images/I/51x7IABfcgL._SX258_BO1,204,203,200_.jpg"
                }
            }
        ]
    }
]
```