### Rebis python client
An example of how to use Rebis client in python.
```python
from rebis import RebisClient, RebisQuery
client = RebisClient("localhost", 6666)
rebis = RebisQuery(client)
rebis.set("key", "value")
rebis.get("key")
rebis.delete("key")
```



