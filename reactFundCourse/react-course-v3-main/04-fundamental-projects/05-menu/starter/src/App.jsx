import menu from "./data";
import SingleMenu from "./SingleMenu";
import Menu from "./Menu";
import Categories from "./Categories";
import { useState } from "react";
const allCategories = ['all', ...new Set(menu.map((menu) => menu.category))];
const App = () => {
  const filterCatgories = () => {

  }
  const [categories, setCategories] = useState(allCategories)
  console.log(allCategories)
  return <div>
    <h2>Menu Starter</h2>
    <Categories categories={categories} />
    <Menu menu={menu} />
  </div>;

};
export default App;
