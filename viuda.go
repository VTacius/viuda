package main

import (
    "fmt"
    "github.com/alecthomas/kong"
    "github.com/vtacius/viuda/lib"
)

var Opciones struct {
    Dominio string `help:"Dominio base" required:""`
    Prefijo string `help:"Prefijo de las componentes variables del path" default:"${"`
    Fichero string `help:"Fichero con el contenido deseado" type:"existingfile"`
    Opciones []string `help:"Opciones para sustituir en los componentes variables del path" default:"12,13"`
    Metodos []string `help:"Metodos a aplicar en cada URL" default:"GET,POST,PUT,PATCH,DELETE"`
    Json bool `help:"Configura la petición como tipo JSON" default:"false"`
}


func main() {
    _ = kong.Parse(&Opciones);
    dominio := Opciones.Dominio;
    prefijo := Opciones.Prefijo;
    fichero := Opciones.Fichero;
    paths := lib.LeerArchivo(fichero);
    opciones := Opciones.Opciones;
    metodos := Opciones.Metodos;
    isJson := Opciones.Json;

    cabeceras := map[string]string {
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:122.0) Gecko/20100101 Firefox/122.0",
    }

    if isJson {
        cabeceras["Content-Type"] = "application/json";
        cabeceras["Accept"] = "application/json";
    }
    
    urls := []string{};
    
    for _, path := range(paths) {
        urls = append(urls, lib.CrearUrls(dominio, path, prefijo, opciones)...);
    }
    
    for _, url := range(urls) {
        for _, metodo := range(metodos){
            fmt.Println(lib.Peticion(metodo, url, cabeceras));
            fmt.Println("");
        } 
    }
}
