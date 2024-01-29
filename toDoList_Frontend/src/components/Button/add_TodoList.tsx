import { FaPlus } from "react-icons/fa";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  DialogClose
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";

import { zodResolver } from "@hookform/resolvers/zod";
import { format } from "date-fns";
import { CalendarIcon } from "lucide-react";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
// import { toast } from "@/components/ui/use-toast";

import Swal from "sweetalert2";

// import * as dayjs from 'dayjs'



const FormSchema = z.object({
  name: z.string({
    required_error: "Please enter the name of task.",
  }),
  deadline: z.date({
    required_error: "Please add your deadline.",
  }),
  description: z.string({
    required_error: "Please enter the description of task",
  }),
});

export const AddToDo = () => {
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
  });

  // data: z.infer<typeof FormSchema>
  function onSubmit() {
    const formData = form.getValues();
    const deadlineData = formData.deadline.toISOString().substring(0,10);


    // toast({
    //   title: "You submitted the following values:",
    //   description: (
    //     <pre className="mt-2 w-[340px] rounded-md bg-slate-950 p-4">
    //       <code className="text-white">{JSON.stringify(data, null, 2)}</code>
    //     </pre>
    //   ),
    // });
    Swal.fire({
      title: "Adding Task complete!",
      html: `<strong class="text-purple-600">${formData.name}</strong> is added <strong>&</strong> deadline is <strong class="text-red-600">${deadlineData}</strong>`,
      icon: "success",
      confirmButtonText: "Close",
      confirmButtonColor: "#9333ea",
    });
  }
  

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button className="mt-8 rounded-full flex items-center justify-center ml-auto">
          <FaPlus className="text-lg" />
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px] ">
        <DialogHeader>
          <DialogTitle className="text-xl">Add Your Task</DialogTitle>
        </DialogHeader>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem className="flex flex-col">
                  <FormLabel>Name of Deadline</FormLabel>
                  <Input {...field} placeholder="Enter the name" />
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="deadline"
              render={({ field }) => (
                <FormItem className="flex flex-col">
                  <FormLabel>Deadline</FormLabel>
                  <Popover>
                    <PopoverTrigger asChild>
                      <FormControl>
                        <Button
                          variant={"outline"}
                          className={cn(
                            "w-[240px] pl-3 text-left font-normal",
                            !field.value && "text-muted-foreground"
                          )}
                        >
                          {field.value ? (
                            format(field.value, "PPP")
                          ) : (
                            <span>Pick a date</span>
                          )}
                          <CalendarIcon className="ml-auto h-4 w-4 opacity-50" />
                        </Button>
                      </FormControl>
                    </PopoverTrigger>
                    <PopoverContent className="w-auto p-0" align="start">
                      <Calendar
                        mode="single"
                        selected={field.value}
                        onSelect={field.onChange}
                        disabled={(date) =>
                          date < new Date((new Date().setHours(0, 0, 0, 0)))
                        }
                        initialFocus
                      />
                    </PopoverContent>
                  </Popover>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="description"
              render={({ field }) => (
                <FormItem className="flex flex-col">
                  <FormLabel>Description</FormLabel>
                  <Input {...field} placeholder="Enter the description" />
                  <FormMessage />
                </FormItem>
              )}
            />
            <DialogClose asChild>
            <Button className="ml-auto" type="submit" disabled={!form.formState.isValid}>
              Submit
            </Button>
          </DialogClose>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
};
