package entities

type Maestria struct {
    IdMaestria      int    `json:"idMaestria"`
    Nombre_maestria string `json:"nombre_maestria"`
    Description string `json:"description"`
    Areas []string `json:"areas"`
    Salario     int`json:"salario"`
    Cantidad_escuelas int `json:"cantidad_escuelas"`
    Escuelas    []string `json:"escuelas"`
	Imagen_miniatura string `json:"imagen_miniatura"`
    Imagen_principal string `json:"imagen_principal"`
}
