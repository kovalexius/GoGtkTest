package main
 
import (
    //"gtk"
    "github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-gtk/glib"
)
 
func main() {
    gtk.Init(nil)
    window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
    window.SetTitle("gtk-go example")
    window.Connect("destroy", func(ctx *glib.CallbackContext) {
        println(ctx.Data().(string))
        gtk.MainQuit()
    }, "data for callback function")
	
	
	
    vbox := gtk.NewVBox(false, 1)
 
    frame := gtk.NewFrame("Demo")
    framebox := gtk.NewVBox(false, 1)
    frame.Add(framebox)
 
    entry := gtk.NewEntry()
    entry.SetText("<Name>")
    framebox.Add(entry)
 
    buttons := gtk.NewHBox(false, 1)
 
    button := gtk.NewButtonWithLabel("Hello me")
    button.Clicked(func() {
        print("button clicked: ", button.GetLabel(), "\n")
        messagedialog := gtk.NewMessageDialog(
            button.GetTopLevelAsWindow(),
            gtk.DIALOG_MODAL,
            gtk.MESSAGE_INFO,
            gtk.BUTTONS_OK,
            "Hello, " + entry.GetText())
        messagedialog.Response(func() {
 
             
            messagedialog.Destroy()
        })
        messagedialog.Run()
    })
    buttons.Add(button)
 
    framebox.Add(buttons)
 
    menubar := gtk.NewMenuBar()
    vbox.PackStart(menubar, false, false, 0)
 
 
    cascademenu := gtk.NewMenuItemWithMnemonic("_File")
    menubar.Append(cascademenu)
    submenu := gtk.NewMenu()
    cascademenu.SetSubmenu(submenu)
 
    menuitem := gtk.NewMenuItemWithMnemonic("E_xit")
    menuitem.Connect("activate", func() {
        gtk.MainQuit()
    })
    submenu.Append(menuitem)
 
     
    cascademenu = gtk.NewMenuItemWithMnemonic("_Help")
    menubar.Append(cascademenu)
    submenu = gtk.NewMenu()
    cascademenu.SetSubmenu(submenu)
 
    menuitem = gtk.NewMenuItemWithMnemonic("_About")
    menuitem.Connect("activate", func() {
        dialog := gtk.NewAboutDialog()
        dialog.SetName("go-gtk library")
        dialog.SetProgramName("helloworlder 1.0")
         
        dialog.Run()
        dialog.Destroy()
    })
    submenu.Append(menuitem)
     
 
    vbox.Add(frame)
 
    window.Add(vbox)
    window.SetSizeRequest(400, 100)
    window.ShowAll()
    gtk.Main()
}
