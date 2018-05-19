# amazonparser



## POST /api/sendurls

**http request**

```json
{
 "urls": ["https://www.amazon.co.uk/gp/product/1472949293", "https://www.amazon.co.uk/gp/product/1509836071", "https://www.amazon.co.uk/gp/product/0241003008"]
}

```

**http response**

```
Your request is processing, requestID=f600f1c5
```

## GET /api/requests/{id}

**http request**

```
http://localhost:7755/api/requests/f600f1c5

```

**http response**

```json
{
    "id": "f600f1c5",
    "products": [
        {
            "url": "https://www.amazon.co.uk/gp/product/1472949293",
            "meta": {
                "title": "Lose Weight for Good: Full-flavour cooking for a low-calorie diet",
                "price": "£7.99£4.80",
                "image": "https://images-na.ssl-images-amazon.com/images/I/51x7IABfcgL._SX258_BO1,204,203,200_.jpg"
            }
        },
        {
            "url": "https://www.amazon.co.uk/gp/product/1509836071",
            "meta": {
                "title": "The Fat-Loss Plan: 100 Quick and Easy Recipes with Workouts",
                "price": "£8.49",
                "image": "https://images-na.ssl-images-amazon.com/images/I/51IsTylYiPL._SX382_BO1,204,203,200_.jpg"
            }
        },
        {
            "url": "https://www.amazon.co.uk/gp/product/0241003008",
            "meta": {
                "title": "The Very Hungry Caterpillar [Board Book]",
                "price": "£3.49£0.93",
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
        "id": "c6e84799",
        "products": [
            {
                "url": "https://www.amazon.co.uk/gp/product/1509836071",
                "meta": {
                    "title": "The Fat-Loss Plan: 100 Quick and Easy Recipes with Workouts",
                    "price": "£8.49",
                    "image": "https://images-na.ssl-images-amazon.com/images/I/51IsTylYiPL._SX382_BO1,204,203,200_.jpg"
                }
            }
        ]
    },
    {
        "id": "f600f1c5",
        "products": [
            {
                "url": "https://www.amazon.co.uk/gp/product/1472949293",
                "meta": {
                    "title": "Lose Weight for Good: Full-flavour cooking for a low-calorie diet",
                    "price": "£7.99£4.80",
                    "image": "https://images-na.ssl-images-amazon.com/images/I/51x7IABfcgL._SX258_BO1,204,203,200_.jpg"
                }
            },
            {
                "url": "https://www.amazon.co.uk/gp/product/1509836071",
                "meta": {
                    "title": "The Fat-Loss Plan: 100 Quick and Easy Recipes with Workouts",
                    "price": "£8.49",
                    "image": "https://images-na.ssl-images-amazon.com/images/I/51IsTylYiPL._SX382_BO1,204,203,200_.jpg"
                }
            },
            {
                "url": "https://www.amazon.co.uk/gp/product/0241003008",
                "meta": {
                    "title": "The Very Hungry Caterpillar [Board Book]",
                    "price": "£3.49£0.93",
                    "image": "https://images-na.ssl-images-amazon.com/images/I/51nXr9QCsIL._SY354_BO1,204,203,200_.jpg"
                }
            }
        ]
    }
]
```