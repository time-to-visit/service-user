package objectValues

type Verify struct {
	Correo string `json:"correo" validate:"required"`
	Clave  string `json:"clave" validate:"required"`
}
