package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!global`:         `global`,
	`!set`:            `set`,
	`!event`:          `event`,
	`(`:               `brace-quote`,
	`echo`:            `out`,
	`!and`:            `and`,
	`!or`:             `or`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!if`:             `if`,
	`!catch`:          `catch`,
	`set`:             `set`,
	`swivel-datatype`: `swivel-datatype`,
	`brace-quote`:     `brace-quote`,
	`tout`:            `tout`,
	`g`:               `g`,
	`rx`:              `rx`,
	`try`:             `try`,
	`global`:          `global`,
	`err`:             `err`,
	`if`:              `if`,
	`append`:          `append`,
	`getfile`:         `getfile`,
	`out`:             `out`,
	`>>`:              `>>`,
	`event`:           `event`,
	`murex-docs`:      `murex-docs`,
	`or`:              `or`,
	`read`:            `read`,
	`and`:             `and`,
	`catch`:           `catch`,
	`trypipe`:         `trypipe`,
	`swivel-table`:    `swivel-table`,
	`pt`:              `pt`,
	`ttyfd`:           `ttyfd`,
	`f`:               `f`,
	`alter`:           `alter`,
	`prepend`:         `prepend`,
	`post`:            `post`,
	`>`:               `>`,
	`tread`:           `tread`,
	`get`:             `get`,
	`export`:          `export`,
}
