package lib;

import (
    "testing"
    "golang.org/x/exp/slices"
)

func TestCrearOpciones(t *testing.T) {
    esperado := []string{"/usuario/1", "/usuario/3215"};
    resultado := insertarOpciones(1, []string{"usuario", "${usuarioId}"}, []string{"1", "3215"})
    
    if !slices.Equal(resultado, esperado) {
       t.Fatalf("Se esperaba %v. Obtenido %v", esperado, resultado);
    }
    
    esperado = []string{"/usuario/1/detalles", "/usuario/3215/detalles"};
    resultado = insertarOpciones(1,  []string{"usuario", "${usuarioId}", "detalles"}, []string{"1", "3215"})
    
    if !slices.Equal(resultado, esperado) {
       t.Fatalf("Se esperaba %v. Obtenido %v", esperado, resultado);
    }

    
    esperado = []string{"/api/usuario/1/perfil/${perfilId}", "/api/usuario/32/perfil/${perfilId}", "/api/usuario/15/perfil/${perfilId}"};
    resultado = insertarOpciones(2, []string{"api", "usuario", "${usuarioId}", "perfil", "${perfilId}"}, []string{"1", "32", "15"})
    
    if !slices.Equal(resultado, esperado) {
       t.Fatalf("Se esperaba %v. Obtenido %v", esperado, resultado);
    }
}

func TestObtenerComponentesVariables(t *testing.T) {
    esperado := []int{1, 3};

    resultado := obtenerComponentesVariables("/usuario/${usuarioId}/tareas/${tareasId}", "${");

    if !slices.Equal(esperado, resultado) {
        t.Fatalf("Se esperaba %v. Obtenido %v", esperado, resultado);
    }
    
    esperado = []int{1, 3};

    resultado = obtenerComponentesVariables("/usuario/{{usuarioId}}/tareas/{{tareasId}}", "{{");

    if !slices.Equal(esperado, resultado) {
        t.Fatalf("Se esperaba %v. Obtenido %v", esperado, resultado);
    }
}

func TestCrearPathsAlternativos(t *testing.T) {
    esperado := []string{"/usuario/{{usuarioId}}/tareas/4/detalle", "/usuario/{{usuarioId}}/tareas/7/detalle","/usuario/{{usuarioId}}/tareas/14/detalle", "/usuario/{{usuarioId}}/tareas/21/detalle"};
    datos := []string{"/usuario/{{usuarioId}}/tareas/{{tareaId}}/detalle"};
    resultado := crearPathsAlternativos(3, datos, []string{"4", "7", "14", "21"});

    if !slices.Equal(esperado, resultado) {
        t.Fatalf("Se esperaba %v. Obtenido %v", esperado, resultado);
    }

}

func TestCrearVariantes(t *testing.T) {
    esperado := []string{"/usuario/1/tareas/1", "/usuario/1/tareas/3215","/usuario/3215/tareas/1", "/usuario/3215/tareas/3215",};
    resultado := crearListaPaths("/usuario/${usuarioId}/tareas/${tareasId}", "${", []string{"1", "3215"});

    if !slices.Equal(esperado, resultado) {
        t.Fatalf("Se esperaba %v. Obtenido %v", esperado, resultado);
    }
}


func TestCrearUrls(t *testing.T) {
    esperado := []string{"https://dominio.com/api/v1/usuario/1", "https://dominio.com/api/v1/usuario/2"};
    resultado := CrearUrls("https://dominio.com", "/api/v1/usuario/${usuarioId}", "${", []string{"1", "2"});
    
    if !slices.Equal(resultado, esperado) {
        t.Fatalf("Se esperaba %v. Obtenido %v", esperado, resultado);
    }
    
    
}
