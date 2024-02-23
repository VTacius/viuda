package lib;

import (
    "fmt"
    "strings"
)

func insertarOpciones(indice int, componentes []string, opciones[]string) []string {
    urls := []string{}
    for _, opcion := range(opciones) {
        url := "";
        for i, j := range(componentes){
            if indice == i {
                url = fmt.Sprintf("%s/%v", url, opcion);
            } else {
                url = fmt.Sprintf("%s/%s", url, j);
            }
       } 
       urls = append(urls, url);
    }

    return urls;
}

func obtenerComponentesVariables(url string, prefijo string) (resultado []int) {
    componentes := strings.Split(strings.TrimLeft(url, "/"), "/");
    for indice, componente := range(componentes) {
        if strings.HasPrefix(componente, prefijo) {
            resultado = append(resultado, indice);
        }
    }

    return
}

func crearPathsAlternativos(indice int, urls []string, opciones []string)(resultado []string){
    for _, url := range(urls) {
        componentes := strings.Split(strings.TrimLeft(url, "/"), "/");
        opciones := insertarOpciones(indice, componentes, opciones);

        resultado = append(resultado, opciones...);
    }

    return
}

func crearListaPaths(path string, prefijo string, opciones []string) []string {
    variables := obtenerComponentesVariables(path, prefijo);
    paths := []string{path}
    for _, indice := range(variables) {
        paths = crearPathsAlternativos(indice, paths, opciones);
    }
    return paths;
}

func CrearUrls(baseurl string, path string, prefijo string, opciones []string) (urls []string) {
    for _, path := range(crearListaPaths(path, prefijo, opciones)) {
        baseurl = strings.TrimRight(baseurl, "/");
        path = strings.TrimLeft(path, "/");
        url := fmt.Sprintf("%s/%s", baseurl, path);
        urls = append(urls, url) 
    }
    
    return
}
