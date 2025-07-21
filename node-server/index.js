const express = require('express');
const app = express();
const port = 3000;

app.get('/', (req, res) => {
  res.send('Hello from Node.js root!');
});

app.get('/api', (req, res) => {
  res.send('API handled by Node.js');
});

app.listen(port, () => {
  console.log(`Node.js server running at http://localhost:${port}`);
});
