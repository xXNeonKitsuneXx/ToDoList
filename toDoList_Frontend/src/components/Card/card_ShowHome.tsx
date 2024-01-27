import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

export function CardShowHome() {
  return (
    <Card className="w-[350px] shadow-xl">
      <CardHeader className="pb-0">
        <CardTitle>Name of Work MOCKKK</CardTitle>
        <CardDescription>Date MOCK ?</CardDescription>
      </CardHeader>
      <CardContent className="pb-2">
        Content mock
      </CardContent>
      <CardFooter className="flex">
        <Button className="ml-auto bg-red-600 hover:bg-orange-400">Done</Button>
      </CardFooter>
    </Card>
  );
}
