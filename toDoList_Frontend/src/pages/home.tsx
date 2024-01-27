import '../../app/globals.css';

import { Toaster } from "@/components/ui/toaster";

import { AddToDo } from '../components/Button/add_TodoList';
import { CardShowHome } from '../components/Card/card_ShowHome'

export const Home = () => {

  return (
    <>
    <Toaster />
    
    <CardShowHome />

    <AddToDo />

    </>
  );
};
