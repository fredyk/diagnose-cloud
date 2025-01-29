package tests

/*
	Uno de nuestros stakeholders nos solicita que desarrollemos una solución para consultar y almacenar datos de diagnóstico de pacientes en nuestros sistemas, esta solución aparte de funcionar internamente, necesita tener la capacidad para integrarse con otras aplicaciones. Ambas partes atesoran datos sensibles por lo que será necesario implementar una forma de autenticar las peticiones.

	Para lograr este objetivo, haremos lo siguiente:

	Un endpoint que requiera de un usuario y contraseña para generar un token de autenticación
	Un endpoint protegido que permita consultar diagnósticos y filtrarlos por: nombre del paciente y/o fecha
	Un endpoint protegido que permita almacenar diagnósticos para un paciente en concreto
	Un endpoint externo y protegido que permita consultar diagnósticos y filtrarlos por: nombre del paciente y/o fecha
	Un endpoint externo y protegido que permita almacenar diagnósticos para un paciente en concreto
	La diferencia entre el endpoint interno y el externo existe debido a que el sistema con el cual nos estamos integrando quizás cuente con una cantidad menor o mayor de datos, o usé una nomenclatura distinta, por lo que debemos tener esto en consideración a la hora de enviar/recibir información entre ambos sistemas.

	Sobre la estructura de la base de datos, necesitamos que existan pacientes, sobre los cuales guardaremos los siguientes datos:

	nombre
	dni
	email
	teléfono (opcional)
	dirección (opcional)
	Estos pacientes tendrán una relación con sus datos de diagnóstico, sobre los cuáles guardaremos los siguientes datos:

	paciente
	diagnóstico
	prescripción (opcional)
	fecha
*/

import (
	"testing"

	"github.com/fredyk/westack-go/client/v2/wstfuncs"
	"github.com/stretchr/testify/assert"

	wst "github.com/fredyk/westack-go/v2/common"
)

func TestLoginInvalidCredentials(t *testing.T) {

	t.Parallel()

	// Test login with invalid credentials
	resp, err := wstfuncs.InvokeApiFullResponse("POST", "/accounts/login", wst.M{
		"username": "invalid",
		"password": "invalid",
	}, wst.M{
		"Content-Type": "application/json",
	})
	assert.Nil(t, err)
	assert.Equal(t, 401, resp.StatusCode)

}

func TestLoginValidCredentials(t *testing.T) {

	t.Parallel()

	t.Error("Test not implemented")

}
