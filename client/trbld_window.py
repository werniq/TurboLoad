import tkinter as tk

class TRBLDApp:
    def __init__(self, master):
        self.master = master
        master.title("TRBLD")
        master.geometry("600x400+50+50")

        self.message = tk.Label(master, text="TRBLD")
        self.message.pack()

if __name__ == "__main__":
    root = tk.Tk()
    app = TRBLDApp(root)
    root.mainloop()
