# CV-Engine: Generador de Currículum Premium

Este proyecto es un motor de generación de currículums profesionales basado en Markdown, Go e Inteligencia Artificial. Permite mantener la información modular y generar versiones en PDF con diseños premium de forma automatizada.

## Estructura del Proyecto

```text
.
├── cv_data.json         # Fuente de verdad estructurada (JSON)
├── jos/                 # Archivos fuente en Markdown (Modulares)
├── skills/              # Protocolos de IA (Skills)
│   └── CV_Transformer.md # Instrucciones para transformar MD a JSON
├── references/          # CVs originales y otros documentos de referencia (Ignorado por Git)
├── src/                 # Generador en Go

│   ├── main.go          # Lógica de renderizado y PDF
│   └── templates/       # Plantillas HTML/CSS
│       ├── default/     # Diseño clásico
│       └── modern/      # Diseño premium de dos columnas
└── curriculum_*.pdf      # Archivos PDF generados
```

## Flujo de Trabajo

### 1. Actualizar Datos (Markdown)
Los datos están organizados por secciones en la carpeta `jos/`. Puedes editar cualquiera de estos archivos para añadir experiencia, proyectos o habilidades.

### 2. Transformar a JSON (The Skill)
Para que el generador de Go procese los datos, deben estar consolidados en `cv_data.json`.
- Usa el skill definido en `skills/CV_Transformer.md`.
- Puedes pedirle a una IA: *"Usa el skill CV_Transformer para actualizar el cv_data.json con los cambios de la carpeta jos/"*.

### 3. Generar el PDF
El generador utiliza Go y Chromedp (para renderizado de alta fidelidad).

```bash
# Ir al directorio del código
cd src

# Generar con la plantilla moderna (Por defecto)
go run main.go -template modern

# Generar con la plantilla clásica
go run main.go -template default
```

## Requisitos Técnicos
- **Go** 1.20 o superior.
- **Google Chrome** o **Edge** (instalado en el sistema, necesario para Chromedp).
- Dependencias de Go (se instalan automáticamente con `go mod tidy`).

## Personalización de Plantillas
Las plantillas están en `src/templates/`. Cada plantilla consta de un archivo `template.html` que utiliza la sintaxis de **Go Templates** para acceder a los campos del `cv_data.json`.

---