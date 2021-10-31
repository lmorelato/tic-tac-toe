import "./StartGame.scss";
import Button from "react-bootstrap/Button";
import Collapse from "react-bootstrap/Collapse";
import Form from "react-bootstrap/Form";
import { useState } from "react";

const StartGame = (props: any) => {
  const [startGameMode, setStartGameMode] = useState(false);
  const [joinGameMode, setJoinGameMode] = useState(false);
  const [validated, setValidated] = useState(false);

  const startGameHandler = () => {
    if (startGameMode) return;
    const timeout = !startGameMode && !joinGameMode ? 0 : 450;
    setStartGameMode(false);
    setJoinGameMode(false);

    window.setTimeout(() => {
      setStartGameMode(true);
      setJoinGameMode(false);
    }, timeout);
  };

  const joinGameHandler = () => {
    if (joinGameMode) return;
    const timeout = !startGameMode && !joinGameMode ? 100 : 600;
    setStartGameMode(false);
    setJoinGameMode(false);

    window.setTimeout(() => {
      setStartGameMode(false);
      setJoinGameMode(true);
    }, timeout);
  };

  const submitHandler = (event: any) => {
    setValidated(true);

    const form = event.currentTarget;
    if (form.checkValidity() === false) {
      event.preventDefault();
      event.stopPropagation();
      return;
    }

    const sessionId = joinGameMode ? form["sessionId"].value: "";
    props.onGameLaunched(form["playerName"].value, sessionId);
  };


  return (
    <div className="d-grid gap-2 w-75">
      <Button size="lg" variant="success" onClick={startGameHandler}>
        Start a New Game
      </Button>
      <Button size="lg" variant="secondary" onClick={joinGameHandler}>
        Join a Game
      </Button>
      <Collapse in={startGameMode || joinGameMode}>
        <div className="my-3">
          <Form id="gameForm" noValidate validated={validated} onSubmit={submitHandler}>
            <Form.Group className="mb-3" controlId="playerName">
              <Form.Label>Player Name</Form.Label>
              <Form.Control
                required
                type="text"
                placeholder="Enter your name"
              />
              <Form.Control.Feedback type="invalid">
                Please provide a player name.
              </Form.Control.Feedback>
            </Form.Group>

            {joinGameMode && (
              <Form.Group className="mb-3" controlId="sessionId">
                <Form.Label>Invitation Code</Form.Label>
                <Form.Control
                  required
                  type="text"
                  placeholder="Enter the invitation code"
                />
                <Form.Control.Feedback type="invalid">
                  Please provide a invitation code.
                </Form.Control.Feedback>
              </Form.Group>
            )}

            <Button variant="primary" type="submit">
              Let's Play!
            </Button>
          </Form>
        </div>
      </Collapse>
    </div>
  );
};

export default StartGame;
