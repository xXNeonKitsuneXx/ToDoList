import { Button } from "@/components/ui/button"
import { useToast } from "@/components/ui/use-toast"
 
export const ToastDemo = () => {
    const { toast } = useToast()
   
    return (
      <Button className="hover:bg-[#0000CD] mt-8" 
        onClick={() => {
          toast({
            title: "Scheduled: Catch up",
            description: "Friday, February 10, 2023 at 5:57 PM",
          })
        }}
      >
        Show Toast
      </Button>
    )
  }