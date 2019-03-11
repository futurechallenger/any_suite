# Config
```json
{
  "dependencies": {},
  "permissions": [],
  "fetchWhiteList": [],
  "plugins": {
    "call": {
      "enterTrigger": {
        "onEnter": "onCallEnter" // When a user click call tab, this function will be called
      },
      "conditionalTriggers":[
        {
          "condition":{
            "key": "startCall"
          },
          "onHandle": "onCallStart"
        },
        // Or set condition like this, if there are no additional conditions
        {
          "condition": "callError",
          "onHandle": "onCallError"
        }, 
        {
          "condition": {
            "key": "callEnded"
          },
          "onHandle": "onCallEnded"
        }
      ]
      "voip":{},
      "ringout":{},
    },
    "account ": {},
    "callHistory": {},
    "messages": {}
  }
}

```

1. When a user logined in, first check the user's permissions are OK with his configurations. He may not have to permission to access what he configured.

2. 