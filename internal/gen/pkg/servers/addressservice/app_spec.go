// Code generated by sysl DO NOT EDIT.
package addressservice

const AppSpec = `{
 "apps": {
  "addressdirectory": {
   "name": {
    "part": [
     "addressdirectory"
    ]
   },
   "longName": "addressdirectory",
   "attrs": {
    "package": {
     "s": "addressdirectory"
    }
   },
   "endpoints": {
    "GET /address": {
     "name": "GET /address",
     "attrs": {
      "patterns": {
       "a": {
        "elt": [
         {
          "s": "rest"
         }
        ]
       }
      }
     },
     "stmt": [
      {
       "ret": {
        "payload": "ok <: Address"
       },
       "sourceContext": {
        "file": "specs/addressdirectory.sysl",
        "start": {
         "line": 5,
         "col": 16
        },
        "end": {
         "line": 5,
         "col": 22
        }
       }
      }
     ],
     "restParams": {
      "method": "GET",
      "path": "/address"
     },
     "sourceContext": {
      "file": "specs/addressdirectory.sysl",
      "start": {
       "line": 4,
       "col": 12
      },
      "end": {
       "line": 7,
       "col": 4
      }
     }
    }
   },
   "types": {
    "Address": {
     "tuple": {
      "attrDefs": {
       "country": {
        "primitive": "STRING",
        "sourceContext": {
         "file": "specs/addressdirectory.sysl",
         "start": {
          "line": 9,
          "col": 19
         },
         "end": {
          "line": 9,
          "col": 19
         }
        }
       },
       "pinCode": {
        "primitive": "STRING",
        "sourceContext": {
         "file": "specs/addressdirectory.sysl",
         "start": {
          "line": 10,
          "col": 19
         },
         "end": {
          "line": 10,
          "col": 19
         }
        }
       },
       "street": {
        "primitive": "STRING",
        "sourceContext": {
         "file": "specs/addressdirectory.sysl",
         "start": {
          "line": 8,
          "col": 19
         },
         "end": {
          "line": 8,
          "col": 19
         }
        }
       }
      }
     },
     "sourceContext": {
      "file": "specs/addressdirectory.sysl",
      "start": {
       "line": 7,
       "col": 4
      },
      "end": {
       "line": 11
      }
     }
    }
   },
   "sourceContext": {
    "file": "specs/addressdirectory.sysl",
    "start": {
     "line": 1,
     "col": 1
    },
    "end": {
     "line": 1,
     "col": 17
    }
   }
  },
  "addressservice": {
   "name": {
    "part": [
     "addressservice"
    ]
   },
   "longName": "addressservice",
   "attrs": {
    "package": {
     "s": "addressservice"
    }
   },
   "endpoints": {
    "GET /address-check": {
     "name": "GET /address-check",
     "docstring": "Get address details",
     "attrs": {
      "patterns": {
       "a": {
        "elt": [
         {
          "s": "rest"
         }
        ]
       }
      }
     },
     "stmt": [
      {
       "call": {
        "target": {
         "part": [
          "addressdirectory"
         ]
        },
        "endpoint": "GET /address"
       },
       "sourceContext": {
        "file": "specs/addressservice.sysl",
        "start": {
         "line": 7,
         "col": 12
        },
        "end": {
         "line": 7,
         "col": 32
        }
       }
      },
      {
       "ret": {
        "payload": "ok <: Address"
       },
       "sourceContext": {
        "file": "specs/addressservice.sysl",
        "start": {
         "line": 8,
         "col": 12
        },
        "end": {
         "line": 8,
         "col": 18
        }
       }
      }
     ],
     "restParams": {
      "method": "GET",
      "path": "/address-check"
     },
     "sourceContext": {
      "file": "specs/addressservice.sysl",
      "start": {
       "line": 5,
       "col": 8
      },
      "end": {
       "line": 10,
       "col": 4
      }
     }
    }
   },
   "types": {
    "Address": {
     "tuple": {
      "attrDefs": {
       "country": {
        "primitive": "STRING",
        "sourceContext": {
         "file": "specs/addressservice.sysl",
         "start": {
          "line": 12,
          "col": 19
         },
         "end": {
          "line": 12,
          "col": 19
         }
        }
       },
       "pinCode": {
        "primitive": "STRING",
        "sourceContext": {
         "file": "specs/addressservice.sysl",
         "start": {
          "line": 13,
          "col": 19
         },
         "end": {
          "line": 13,
          "col": 19
         }
        }
       },
       "street": {
        "primitive": "STRING",
        "sourceContext": {
         "file": "specs/addressservice.sysl",
         "start": {
          "line": 11,
          "col": 18
         },
         "end": {
          "line": 11,
          "col": 18
         }
        }
       }
      }
     },
     "sourceContext": {
      "file": "specs/addressservice.sysl",
      "start": {
       "line": 10,
       "col": 4
      },
      "end": {
       "line": 14
      }
     }
    }
   },
   "sourceContext": {
    "file": "specs/addressservice.sysl",
    "start": {
     "line": 2,
     "col": 1
    },
    "end": {
     "line": 2,
     "col": 15
    }
   }
  }
 }
}`
