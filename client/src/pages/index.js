import Navbar from "../navbar/components/navbar";
import TopCardComponent from '../top-card/components/top-card-component';
import { TaskList } from '../task-list/components/task-list-component';

function Home() {
  return (
    <>
      <Navbar name="Name" />
      <TopCardComponent />
      <main>
        <TaskList />
      </main>

      <footer></footer>
    </>
  );
}

export default Home;
