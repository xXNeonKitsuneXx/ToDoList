import { Button } from "@/components/ui/button";
import Swal from 'sweetalert2'

export const DeleteToDo = () => {

  return (
    <Button className="ml-auto bg-red-600 hover:bg-orange-400"
      onClick={() => {
        Swal.fire({
            title: "Good job!",
            html: `<strong class="text-purple-600">(Name of Work MOCKKK)</strong> is done!`,
            icon: "success",
            confirmButtonText: "Close",
            confirmButtonColor: "#9333ea",
          });
      }}
    >
      Done
    </Button>
  );
};
