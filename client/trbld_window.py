import tkinter as tk
from tkinter import ttk
from dotenv import load_dotenv

from retrieve_data import (get_file_statistics,
                            get_all_files_statistics,
                            get_concurrent_requests,
                            get_total_downloads)


class TRBLDApp:
    def __init__(self, master):
        self.master = master
        master.title("TRBLD")
        master.geometry("1000x400+50+50")

        # Padding frame
        self.padding_frame = tk.Frame(master, bg='#092C46', padx=5, pady=5)
        self.padding_frame.pack(fill="both", expand=True)

        # First row
        self.first_row_frame = tk.Frame(self.padding_frame, bg="#092C46")
        self.first_row_frame.pack(fill="both", expand=True)

        self.first_left_frame = tk.Frame(self.first_row_frame, bg="#092C46", bd=1, relief="solid")
        self.first_left_frame.pack(side="left", fill="both", expand=True, padx=(0, 5), pady=2)

        self.first_right_frame = tk.Frame(self.first_row_frame, bg="#092C46", bd=1, relief="solid")
        self.first_right_frame.pack(side="left", fill="both", expand=True, padx=(0, 5), pady=2)

        # Second row
        self.second_row_frame = tk.Frame(self.padding_frame, bg="#092C46")
        self.second_row_frame.pack(fill="both", expand=True)

        self.second_frame = ttk.Treeview(self.second_row_frame, style="Custom.Treeview")
        self.second_frame["columns"] = ("1", "2", "3")
        self.second_frame.heading("#0", text="ID")
        self.second_frame.heading("1", text="Filename")
        self.second_frame.heading("2", text="Sized of the file")
        self.second_frame.heading("3", text="Downloads count")
        self.second_frame.pack(fill="both", expand=True, padx=0, pady=5)

        style = ttk.Style()
        style.theme_use("default")
        style.configure("Custom.Treeview", background="#092C46", fieldbackground="#092C46", foreground="white")


    def insert_data_to_first_left_frame(self, data):
        # Clear existing widgets
        for widget in self.first_left_frame.winfo_children():
            widget.destroy()

        # Insert new data
        label = tk.Label(self.first_left_frame, text=data)
        label.pack()

    def insert_data_to_first_right_frame(self, data):
        # Clear existing widgets
        for widget in self.first_right_frame.winfo_children():
            widget.destroy()

        # Insert new data
        label = tk.Label(self.first_right_frame, text=data)
        label.pack()

    def insert_data_to_second_frame(self, data):
        # Clear existing data
        self.second_frame.delete(*self.second_frame.get_children())

        # Insert new data
        for item in data:
            self.second_frame.insert("", "end", text=item[0], values=(item[1], item[2], item[3]))


if __name__ == "__main__":
    load_dotenv()

    root = tk.Tk()
    app = TRBLDApp(root)

    # Example data
    data_for_first_left_frame = "Files download statistics"
    data_for_first_right_frame = "Files statistics"
    data_for_second_frame = get_all_files_statistics()

    # Insert data into frames
    app.insert_data_to_first_left_frame(data_for_first_left_frame)
    app.insert_data_to_first_right_frame(data_for_first_right_frame)
    app.insert_data_to_second_frame(data_for_second_frame)

    root.mainloop()
