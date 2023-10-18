package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func PasswordHash(password string) ([]byte, []byte, error) {
	// Senha a ser hasheada
	//password = "senha_a_ser_hasheada"

	// Gere um "salt" aleatório
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		fmt.Println("Erro ao gerar o salt:", err)
		return nil, nil, err
	}

	// Parâmetros do algoritmo Scrypt
	N := 16384
	r := 8
	p := 1
	keyLen := 32

	// Crie o hash da senha
	hashedPassword, err := scrypt.Key([]byte(password), salt, N, r, p, keyLen)
	if err != nil {
		fmt.Println("Erro ao gerar o hash da senha:", err)
		return nil, nil, err
	}

	fmt.Printf("Senha original: %s\n", password)
	fmt.Printf("Hash da senha (hexadecimal): %s\n", hex.EncodeToString(hashedPassword))

	return hashedPassword, salt, nil
}

func ValidatePasswordHash(passwordLogin string, passwordDB []byte, salt []byte) (bool, error) {
	// Parâmetros do algoritmo Scrypt
	N := 16384
	r := 8
	p := 1
	keyLen := 32

	comparisonPassword, err := scrypt.Key([]byte(passwordLogin), salt, N, r, p, keyLen)
	if err != nil {
		fmt.Println("Erro ao gerar o hash para verificação:", err)
		return false, err
	}

	println("Senha salva no banco: ", passwordDB)
	println("Senha a ser comaparada: ", comparisonPassword)
	if hex.EncodeToString(comparisonPassword) == hex.EncodeToString(passwordDB) {
		fmt.Println("A senha é válida.")
		return true, nil
	} else {
		fmt.Println("A senha está incorreta.")
		return false, nil
	}
}
