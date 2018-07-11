package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`(`:               `brace-quote`,
	`!and`:            `and`,
	`!if`:             `if`,
	`!global`:         `global`,
	`!set`:            `set`,
	`!event`:          `event`,
	`echo`:            `out`,
	`!or`:             `or`,
	`!catch`:          `catch`,
	`!export`:         `export`,
	`unset`:           `export`,
	`ttyfd`:           `ttyfd`,
	`g`:               `g`,
	`tread`:           `tread`,
	`alter`:           `alter`,
	`>`:               `>`,
	`trypipe`:         `trypipe`,
	`out`:             `out`,
	`f`:               `f`,
	`if`:              `if`,
	`brace-quote`:     `brace-quote`,
	`global`:          `global`,
	`set`:             `set`,
	`append`:          `append`,
	`swivel-datatype`: `swivel-datatype`,
	`err`:             `err`,
	`and`:             `and`,
	`event`:           `event`,
	`murex-docs`:      `murex-docs`,
	`getfile`:         `getfile`,
	`pt`:              `pt`,
	`rx`:              `rx`,
	`read`:            `read`,
	`export`:          `export`,
	`prepend`:         `prepend`,
	`tout`:            `tout`,
	`>>`:              `>>`,
	`or`:              `or`,
	`catch`:           `catch`,
	`try`:             `try`,
	`get`:             `get`,
	`post`:            `post`,
	`swivel-table`:    `swivel-table`,
}
