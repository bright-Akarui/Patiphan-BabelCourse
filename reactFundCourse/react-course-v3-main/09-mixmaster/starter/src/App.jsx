import {createBrowserRouter, RouterProvider} from 'react-router-dom';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import {About,HomeLayout,Cocktail,Error,Landing,Newsletter} from './pages'
import { Loading as LandingLoader } from './pages/Landing';
import { Loading as CocktailLoader } from './pages/Cocktail';
import { action as NewsletterAction } from './pages/Newsletter';
const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60 * 5,
    },
  },
});
const router = createBrowserRouter([
  {
    path: '/',
    element: <HomeLayout />,
    errorElement: <Error />,
    children: [
      {
        index: 'true',
        element: <Landing />,
        errorElement: <Error />,
        loader: LandingLoader ,
      },
      {
        path: 'cocktail/:id',
        element: <Cocktail />,
        errorElement: <Error />,
        loader: CocktailLoader ,
      },
      {
        path: 'newsletter',
        element: <Newsletter />,
        action: NewsletterAction,
      },
      {
        path: 'about',
        element: <About />,
      }
    ],
  },
]);

const App = () => {
  return<QueryClientProvider client={queryClient}>
    <RouterProvider router={router} />;
  </QueryClientProvider>
};
export default App;
