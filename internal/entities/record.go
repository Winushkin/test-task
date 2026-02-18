// Package entities contains app's entities
package entities

type Record struct {
	Number       string `info:"Номер записи в файле" db:"num"`
	Mqtt         string `info:"mqtt" db:"mqtt"`
	InvID        string `info:"Инвентарный id" db:"inv_id"`
	UnitGUID     string `info:"ID Записи" db:"unit_guid"`
	MessageID    string `info:"ID сообщения" db:"message_id"`
	MessageText  string `info:"Текст сообщения" db:"message_text"`
	Context      string `info:"Среда" db:"context"`
	MessageClass string `info:"Класс сообщения" db:"message_class"`
	MessageLevel string `info:"Уровень сообщения" db:"message_level"`
	Area         string `info:"Зона переменных" db:"area"`
	VarAddress   string `info:"Адрес переменной в контроллере" db:"var_addr"`
	Block        string `info:"Начало блока" db:"block_sign"`
	MessageType  string `info:"Тип" db:"message_type"`
	BitNumber    string `info:"Номер бита в регистре" db:"bit_number"`
	InvertBit    string `info:"Инвертированный бит" db:"invert_bit"`
	FileID       int    `info:"Имя tsv файла" db:"file_id"`
}
