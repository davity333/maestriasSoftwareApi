package repositories

import entities "software/src/Domain/Entities"

type AreaStaticRepository struct {
    data []*entities.Maestria
}

func NewAreaStaticRepository() *AreaStaticRepository {
    return &AreaStaticRepository{
        data: []*entities.Maestria{
            {
                IdMaestria:         1,
                Nombre_maestria:   "Maestría en Ingeniería de Software",
                Description:       "Enfocada en el diseño, construcción, mantenimiento y calidad del software. Es la maestría más completa para quienes quieren ser desarrolladores profesionales.",
                Areas:       []string{
					"Frontend","Backend","Full-Stack","DevOps","QA / Testing","Arquitectura de software",
				},
                Salario:           30000,
                Cantidad_escuelas: 7,
                Escuelas:          []string{
					"UNAM (México)","UNAM (México)", "Tec de Monterrey (México)", "Universidad de Guadalajara","Universidad Politécnica de Madrid (España)",
					"Universidad de Buenos Aires (Argentina)", "Universidad de Barcelona (España)",
				},
                Imagen_miniatura:  "https://www.aunnait.com/wp-content/uploads/2024/02/desarrollo-de-software.jpg",
                Imagen_principal:  "https://unives.mx/wp-content/uploads/2021/02/rear-view-of-programmer-working-all-night-long-min.jpg",
            },

			{
				IdMaestria: 2,
				Nombre_maestria: "Maestría en Ciencias de la Computación",
				Description: "Maestría más teórica y científica. Profundiza en algoritmos, sistemas, IA y fundamentos matemáticos.",
				Areas: []string{
					"Backend avanzado", "Inteligencia Artificial","Sistemas distribuidos",
					"Investigación","Arquitectura de software",
				},
				Salario: 14500,
				Cantidad_escuelas: 6,
				Escuelas: []string{
					"UNAM (México)","CINVESTAV (México)","Tec de Monterrey","Universidad de Chile",
					"Universidad de Buenos Aires","Universidad Politécnica de Cataluña (España)",
				},
				Imagen_miniatura: "https://www.unidiversidad.com.ar/cache/nota-1-opcion1_1000_1100.jpg",
				Imagen_principal: "https://blogs.ucontinental.edu.pe/wp-content/uploads/2025/02/ciencia-de-la-computacion-sera-la-carrera-que-dominara-el-mundo-universidad-continental-1-696x392.jpg",
			},

			{
				IdMaestria: 3,
				Nombre_maestria: "Maestría en Desarrollo de Software",
				Description: "Más práctica que Ingeniería de Software. Enfocada en programación avanzada, APIs, cloud y desarrollo moderno.",
				Areas: []string{
					"Backend","Frontend","Full‑Stack","DevOps","Cloud",
				},
				Salario: 17000,
				Cantidad_escuelas: 4,
				Escuelas: []string{
					"Universidad de Guadalajara","Universidad Autónoma de Nuevo León",
					"Universidad Europea (España)","Universidad Internacional de Valencia (España)",
				},
				Imagen_miniatura: "https://unab.edu.co/wp-content/uploads/2022/03/3-mae-desarrollo-software.jpg",
				Imagen_principal: "https://www.grupocibernos.com/hs-fs/hubfs/desarrollo%20de%20software%20a%20medida.jpg?width=1280&height=720&name=desarrollo%20de%20software%20a%20medida.jpg",
			},

			{
				IdMaestria: 5,
				Nombre_maestria: "Maestría en Inteligencia Artificial",
				Description: "Enfocada en machine learning, deep learning, visión computacional y NLP.",
				Areas: []string{
					"Machine Learning","IA","Data Science","Automatización","Robótica,",
				},
				Salario: 90000,
				Cantidad_escuelas: 6,
				Escuelas: []string{
					"UNAM","Tec de Monterrey","IPN – CIC","Universidad de Barcelona",
					"Universidad Politécnica de Madrid","Universidad de Buenos Aires",
				},
				Imagen_miniatura: "https://apd-images.imgix.net/uploads/sites/2/2021/01/ramas_ia_1.jpg?auto=compress%2Cformat&crop=edges&fit=crop&ixlib=php-1.1.0&w=900&s=0655022506ec76e0f2e64d5735c1465f",
				Imagen_principal: "https://kschool.com/wp-content/uploads/2025/02/educacio-n-e-inteligencia-artificial.webp",
			},

			{
				IdMaestria: 6,
				Nombre_maestria: "Maestría en Ciberseguridad",
				Description: "Enfocada en seguridad ofensiva y defensiva, criptografía y análisis forense.",
				Areas: []string{
					"Ciberseguridad","Pentesting","Seguridad ofensiva","Seguridad defensiva","Auditoría de sistemas",
				},
				Salario: 45000,
				Cantidad_escuelas: 5,
				Escuelas: []string{
					"IPN – ESIME","UNAM","Tec de Monterrey","Universidad Europea (España)","Universidad de Barcelona",
				},
				Imagen_miniatura: "https://img.gerencia.cl/wp-content/uploads/2024/05/07165915/Iplacex-abre-la-carrera-de-Ingenieria-en-Ciberseguridad.jpg",
				Imagen_principal: "https://media.licdn.com/dms/image/v2/D4E12AQH_vPixihcnxQ/article-cover_image-shrink_720_1280/article-cover_image-shrink_720_1280/0/1730382495275?e=2147483647&v=beta&t=Yl_2KWa3fXvXCrfiGRWUfm8O0lPKqhxNJmPKScDQv1A",
			},
			{
				IdMaestria: 7,
				Nombre_maestria: "Maestría en Desarrollo de Videojuegos",
				Description: "Enfocada en motores gráficos, animación, física y programación de juegos.",
				Areas: []string{
					"Game Developer","Animación","Diseño de videojuegos","Programación gráfica",
				},
				Salario: 25000,
				Cantidad_escuelas: 4,
				Escuelas: []string{
					"Universidad de Guadalajara","Universidad Europea (España)","Universidad de Barcelona","Universidad de Buenos Aires",
				},
				Imagen_miniatura: "https://micarrerauniversitaria.com/wp-content/uploads/2018/04/desarrolladordevideojuegos-1.jpg",
				Imagen_principal: "https://online.anahuac.mx/images/d_desarrollo_videojuegos_anahuac_online_thumb_1200w.jpg",
			},
			{
				IdMaestria: 8,
				Nombre_maestria: "Maestría en Sistemas Embebidos / IoT",
				Description: "Enfocada en hardware + software, microcontroladores, electrónica y sistemas en tiempo real.",
				Areas: []string{
					"IoT","Sistemas embebidos","Robótica","Automatización industrial",
				},
				Salario: 32000,
				Cantidad_escuelas: 5,
				Escuelas: []string{
					"IPN – ESIME","UNAM","Tec de Monterrey","Universidad Politécnica de Cataluña","Universidad de Buenos Aires",
				},
				Imagen_miniatura: "https://img.innovaciondigital360.com/wp-content/uploads/2023/08/23221330/computadora-cuantica-microsoft-1.jpg",
				Imagen_principal: "https://elciudadanoweb.com/wp-content/uploads/2025/02/Embedded_Systems_Engineer_9d6f1a.jpg",
			},
        },
    }
}

func (r *AreaStaticRepository) GetAllAreas() ([]*entities.Maestria, error) {
    return r.data, nil
}

func (r *AreaStaticRepository) GetById(id int) (*entities.Maestria, error) {
    for _, m := range r.data {
        if m.IdMaestria == id {
            return m, nil
        }
    }
    return nil, nil
}
