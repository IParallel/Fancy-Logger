package fancy_logger

type Colors string

/*
Terminal colors

you can add decoration example : Colors_Black + Deco_Bold
*/
const (
	Colors_Reset                   = Colors("\u001b[0m")
	Colors_Black                   = Colors("\u001b[30m")
	Colors_Red                     = Colors("\u001b[31m")
	Colors_Green                   = Colors("\u001b[32m")
	Colors_Yellow                  = Colors("\u001b[33m")
	Colors_Blue                    = Colors("\u001b[34m")
	Colors_Magenta                 = Colors("\u001b[35m")
	Colors_Cyan                    = Colors("\u001b[36m")
	Colors_White                   = Colors("\u001b[37m")
	Colors_Bold                    = Colors("\u001b[1m")
	Colors_Underline               = Colors("\u001b[4m")
	Colors_Reversed                = Colors("\u001b[7m")
	Colors_BrightBlack             = Colors("\u001b[30;1m")
	Colors_BrightRed               = Colors("\u001b[31;1m")
	Colors_BrightGreen             = Colors("\u001b[32;1m")
	Colors_BrightYellow            = Colors("\u001b[33;1m")
	Colors_BrightBlue              = Colors("\u001b[34;1m")
	Colors_BrightMagenta           = Colors("\u001b[35;1m")
	Colors_BrightCyan              = Colors("\u001b[36;1m")
	Colors_BrightWhite             = Colors("\u001b[37;1m")
	Colors_BackgroundBlack         = Colors("\u001b[40m")
	Colors_BackgroundRed           = Colors("\u001b[41m")
	Colors_BackgroundGreen         = Colors("\u001b[42m")
	Colors_BackgroundYellow        = Colors("\u001b[43m")
	Colors_BackgroundBlue          = Colors("\u001b[44m")
	Colors_BackgroundMagenta       = Colors("\u001b[45m")
	Colors_BackgroundCyan          = Colors("\u001b[46m")
	Colors_BackgroundWhite         = Colors("\u001b[47m")
	Colors_BackgroundBrightBlack   = Colors("\u001b[40;1m")
	Colors_BackgroundBrightRed     = Colors("\u001b[41;1m")
	Colors_BackgroundBrightGreen   = Colors("\u001b[42;1m")
	Colors_BackgroundBrightYellow  = Colors("\u001b[43;1m")
	Colors_BackgroundBrightBlue    = Colors("\u001b[44;1m")
	Colors_BackgroundBrightMagenta = Colors("\u001b[45;1m")
	Colors_BackgroundBrightCyan    = Colors("\u001b[46;1m")
	Colors_BackgroundBrightWhite   = Colors("\u001b[47;1m")
)

// This are decorations you can combine them with colors or use alone
const (
	Deco_Bold      = Colors("\u001b[1m")
	Deco_Underline = Colors("\u001b[4m")
	Deco_Reversed  = Colors("\u001b[7m")
)

//future
// const (
// 	Cursor_Up    = Colors("\u001b[{n}A")
// 	Cursor_Down  = Colors("\u001b[{n}B")
// 	Cursor_Right = Colors("\u001b[{n}C")
// 	Cursor_Left  = Colors("\u001b[{n}D")
// )
