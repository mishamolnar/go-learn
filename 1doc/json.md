# Questions json

### 1. What is the difference between: `json.Marshal/json.Unmarshal` 
### and  `json.NewDecoder(resp.Body).Decode(&result)`?

The difference is that Decoder that allows entities to be decoded in a stream, unlike json.Marshal


