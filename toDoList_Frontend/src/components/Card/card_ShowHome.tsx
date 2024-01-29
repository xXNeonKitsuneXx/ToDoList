import { DeleteToDo } from "../Button/delete_TodoList"; 
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";


function updateDiglog() {
    
  }

export function CardShowHome() {
  return (
    <Card className="w-[350px] shadow-xl hover:bg-gray-50" onClick={updateDiglog}>
      <CardHeader className="pb-0">
        <CardTitle>Name of Work MOCKKK</CardTitle>
        <CardDescription><span>Date Start</span><span> - </span><span>Date End</span></CardDescription>
      </CardHeader>
      <CardContent className="pb-2">
        Content mock
      </CardContent>
      <CardFooter className="flex">
        <DeleteToDo />
      </CardFooter>
    </Card>
  );
}
