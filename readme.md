# Box_upload

box_upload is a single purpose program to upload things to box, utilising JWT with RSA keys.

# Args
 - '-folderid=<id>'
 - '-path=<item to upload>'


# Installation
 Generate a RSA key pair

 ```sh
$ openssl genrsa -des3 -out private.pem 2048
$ openssl rsa -in private.pem -outform PEM -pubout -out public.pem
```

Configure config.json and place in same folder as binary. It should contain the app information from box.com

 ```json
{
  "boxAppSettings": {
    "clientID": "",
    "clientSecret": "",
    "appAuth": {
      "publicKeyID": "",
      "privateKey": "/etc/ssl/certs/box_private.pem",
      "passphrase": ""
    }
  },
  "enterpriseID": ""
}
```
