var express = require('express');
var router = express.Router();
const fs = require('fs');

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Express' });
});

router.post('/', (req, res, next) => {
  const bdy = {
    id: req.id,
    headers: req.headers,
    body: req.body
  };
  fs.writeFile(`./tmp/${ req.id }.json`, JSON.stringify(bdy), (err) => {
    if (err) throw err;
  });
  res.send(bdy);
});

module.exports = router;
