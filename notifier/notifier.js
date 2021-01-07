const express = require('express');
const path = require('path');
const bodyParser = require('body-parser');
const notifier = require('node-notifier');
const app = express();
const port = process.env.PORT || 9000;

app.use(bodyParser.json());

app.get("/health", (req, res) => res.status(200).send({message: `It's alive 🎉`}));
app.post("/notify", (req, res) => {
    notify(req.body, reply => res.send(reply));
});
app.listen(port, () => console.log(`Server is up and running on port ${port} 🎉`));

const notify = ({title, message}, cb) => {
    notifier.notify(
        {
            title: title || "Unknown title 😕",
            message: message || "Unknown message 😕",
            icon: path.join(__dirname, "./golang_icon.png"),
            sound: true,
            wait: true,
            reply: true,
            closeLabel: "Completed?",
            timeout: 15
        },
        (err, response, reply) => {
            cb(reply)
        }
    )
}