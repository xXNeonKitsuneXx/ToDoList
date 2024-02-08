import { FaPlus } from "react-icons/fa";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  DialogClose,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";

import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";

import { zodResolver } from "@hookform/resolvers/zod";
import { format } from "date-fns";
import { CalendarIcon } from "lucide-react";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { cn } from "@/lib/utils";

import { Calendar } from "@/components/ui/calendar";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import Swal from "sweetalert2";

const formSchema = z.object({
  name: z
    .string()
    .min(3, { message: "3 to 20 character" })
    .max(20, { message: "3 to 20 character" }),
  deadline: z.date(),
  description: z
    .string()
    .min(5, { message: "5 to 30 character" })
    .max(30, { message: "5 to 30 character" }), //Didnt tell error from start
});

type FormSchema = z.infer<typeof formSchema>;

export const AddToDo = () => {
  const form = useForm<FormSchema>({
    defaultValues: {
      name: "",
      deadline: new Date(new Date().setHours(23, 59, 0, 0)),
      description: "",
    },
    resolver: zodResolver(formSchema), //Set mai a
  });

  const handleClick = (data: FormSchema) => {
    Swal.fire({
      title: "Adding Task complete!",
      html: `<strong class="text-purple-600">${data.name}${data.description}</strong> is added <strong>&</strong> deadline is <strong class="text-red-600">${data.deadline}}</strong>`,
      icon: "success",
      confirmButtonText: "Close",
      confirmButtonColor: "#9333ea",
    });
    form.reset();
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button className="mt-8 rounded-full flex items-center justify-center ml-auto">
          <FaPlus className="text-lg" />
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Add Task</DialogTitle>
          <DialogDescription>Make a to do list for yourself.</DialogDescription>
        </DialogHeader>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(handleClick)} className="space-y-8">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input placeholder="Enter name..." {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            {/* //---------------------------------------------- */}
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
                          date < new Date(new Date().setHours(23, 59, 0, 0)) ||
                          date < new Date("1900-01-01") // How to set today date and how I can deadline at 23:59 of that picked date
                        } // time zone too
                        initialFocus
                      />
                    </PopoverContent>
                  </Popover>
                  <FormMessage />
                </FormItem>
              )}
            />
            {/* //-------------------------------------------- */}
            <FormField
              control={form.control}
              name="description"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Description</FormLabel>
                  <FormControl>
                    <Input placeholder="Enter description..." {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            {/* ////////////////////////////////////////////////////////////// */}
            <DialogFooter>
              <DialogClose asChild>
                <Button type="submit" disabled={!form.formState.isValid}>Submit</Button>
              </DialogClose>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
};
