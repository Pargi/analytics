# GET /prod/health

+ Request

    + Headers

            x-api-key: 8CyeEFIruH9U5CVRyXQFT7GcvfCKH9Wk7Cjzg4Zb
            x-app-version: 3.0.1



+ Response 200 (application/json)

    + Headers

            X-Cache: Miss from cloudfront
            x-amzn-RequestId: 009a13c0-246e-11e8-bb2b-7bad8ef0e7ee
            X-Amz-Cf-Id: 6vNi7gK212bn-URO3ZccJwvYJXx8XwlezXkAQS70vVhdJUAdkAirgA==
            X-Amzn-Trace-Id: sampled=0;root=1-5aa3e95e-3711fc7d723c6f246ca044b4

    + Body

            {"status":"OK"}


# POST /prod/events

+ Request (application/json; charset=utf-8)

    + Headers

            x-api-key: 8CyeEFIruH9U5CVRyXQFT7GcvfCKH9Wk7Cjzg4Zb
            x-app-version: 3.0.1

    + Body

            {
                "start": "2018-03-10T14:00:09Z",
                "end": "2018-03-10T14:30:09Z",
                "zone": "KESKLINN",
                "device": "6b84620f-29bc-491f-b6aa-71d8aecbb278",
                "latitude": 59.437,
                "longitude": 24.7536
            }

+ Response 200 (application/json)

    + Headers

            X-Cache: Miss from cloudfront
            x-amzn-RequestId: 77127c91-246f-11e8-aefc-0391bca4f916
            X-Amz-Cf-Id: EzYxaIUu94vbdh_nXM9DmJ3ki6iGi6j8FZQrCL6bkpgj2cQnOUf32Q==
            X-Amzn-Trace-Id: sampled=0;root=1-5aa3ebd3-0628792fc43fe8bbf4d4a957

    + Body

            {"guid":"1d6e40a1-8b0f-4b28-9480-bc5738932efa","device":"587c5ad1-484f-46db-a2b0-67564a6250a4","appVersion":"3.0.1","start":"2018-03-10T13:59:39Z","end":"2018-03-10T14:29:39Z","zone":"KESKLINN","latitude":59.437,"longitude":24.7536}


