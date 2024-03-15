import "./App.css";
import Main from "./Main";

import { BrowserRouter} from "react-router-dom";
import "react-notifications/lib/notifications.css";
import { NotificationContainer } from "react-notifications";

function App() {
  return (
    <BrowserRouter>
      <div className="App" data-testid="app-head">
        <Main/>
        <NotificationContainer/>
      </div>
      </BrowserRouter>
  );
}

export default App;
