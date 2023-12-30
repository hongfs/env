const express = require('express');
const bodyParser = require('body-parser');

const app = express();

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
app.use(bodyParser.raw());

const port = 9000

app.get('/*', (req, res) => {
  res.send('Hello World!')
})

app.post('/*', (req, res) => {
  console.log(JSON.stringify(req.body));

  res.send('Hello World!')
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})