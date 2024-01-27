import { Button } from "@/components/ui/button";
import { useToast } from "@/components/ui/use-toast";

import { FaPlus } from "react-icons/fa";

export const AddToDo = () => {
  const { toast } = useToast();

  return (
    <Button className="bg-purple-600 hover:bg-orange-400 mt-8 rounded-full flex items-center justify-center ml-auto" 
      onClick={() => {
        toast({
          variant: "destructive",
          title: "Done adding",
          description: "",
        });
      }}
    >
      <FaPlus className="text-lg" />
    </Button>
  );
};
