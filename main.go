package main

import (
	"embed"
)

const demo_link = "https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/58b12a05-e5a3-4b3c-adc1-14a27f45df6c/dgqq4oy-98680950-a442-438e-88a3-c86cf1156d74.jpg/v1/fit/w_828,h_466,q_70,strp/tamamo_no_mae_by_yuchenghong_dgqq4oy-414w-2x.jpg?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7ImhlaWdodCI6Ijw9NzIxIiwicGF0aCI6IlwvZlwvNThiMTJhMDUtZTVhMy00YjNjLWFkYzEtMTRhMjdmNDVkZjZjXC9kZ3FxNG95LTk4NjgwOTUwLWE0NDItNDM4ZS04OGEzLWM4NmNmMTE1NmQ3NC5qcGciLCJ3aWR0aCI6Ijw9MTI4MCJ9XV0sImF1ZCI6WyJ1cm46c2VydmljZTppbWFnZS5vcGVyYXRpb25zIl19.bepnG8bs-pOcc05YZcND_RYgMcNecPmwXhRZIbKTD4Q"
const video_link = "https://www.youtube.com/watch?v=Aavt_JRhDZA&t=818s"

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp("image/png", demo_link)
	app.Run()
}

func MediaProcess(*Media) {
	//unimplemented
}

func InputHandler() (*Media, error) {
  //unimplemented
  return nil, nil
}
