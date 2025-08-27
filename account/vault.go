package account

import (
	"encoding/json"
	"errors"
	"strings"
	"test/test_app_4/encrypter"
	"test/test_app_4/output"
	"time"

	"github.com/fatih/color"
)

type ByteWriter interface {
	Write([]byte)
}
type ByteReader interface {
	Read() ([]byte, error)
}
type DB interface {
	ByteWriter
	ByteReader
}
type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDB struct {
	Vault
	db DB
	enc encrypter.Encrypter
}

func NewVault(db DB, enc encrypter.Encrypter) (*VaultWithDB, error) {

	file, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
			enc: enc,
		}, nil
	}
	decFile := enc.Decrup(file)
	var vault Vault
	err = json.Unmarshal(decFile, &vault)
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
			enc: enc,
		}, err

	}
	return &VaultWithDB{
		Vault: vault,
		db:    db,
		enc: enc,
	}, nil
}

func (vault *VaultWithDB) FindAccount(input string, checker func(Account, string)bool) ([]Account, error) {
	var accounts []Account
	for _, acc := range vault.Accounts {
		if checker(acc, input) {
			accounts = append(accounts, acc)
		}
	}
	return accounts, nil
}

func (vault *VaultWithDB) AddAccount(acc *Account) {
	vault.Accounts = append(vault.Accounts, *acc)
	err := vault.save()
	if err != nil {
		output.PrintError(err)
	} else{
		color.HiGreen("Успешное добавление")
	}
}

func (vault *VaultWithDB) DeleteAccount(loginURL string) error {
	var accounts []Account
	flag := false
	for _, acc := range vault.Accounts {
		if !strings.Contains(acc.Url, loginURL) {
			accounts = append(accounts, acc)
			continue
		}
		flag = true
	}
	if !flag {
		return errors.New("nothing deleted")
	}
	vault.Accounts = accounts
	err := vault.save()
	if err != nil {
		output.PrintError(err)
	}
	return nil
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *VaultWithDB) save() error {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		return err
	}
	encData := vault.enc.Encryp(data)
	vault.db.Write(encData)
	return nil
}
