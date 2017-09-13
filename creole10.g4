///////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2007, 2008 Martin Junghans, Dirk Riehle, SAP Labs LLC.      //
//  Taken from: Martin Junghans, Dirk Riehle, Rama Gurram,                   //
//  Matthias Kaiser, Mario Lopes, and Umit Yalcinalp. An EBNF                //
//  grammar for Wiki Creole 1.0. In ACM SIGWEB Newsletter, Volume            //
//  2007, Issue Winter (Winter 2007), Article No. 4.                         //
//                                                                           //
//                                                                           //
// This file is part of 'Wiki Creole Markup Parser'.                         //
//                                                                           //
// The 'Wiki Creole Markup Parser' is free software: you can redistribute it //
// and/or modify it under the terms of the GNU General Public License as     //
// published by the Free Software Foundation, either version 3 of the        //
// License, or (at your option) any later version.                           //
//                                                                           //
// 'Wiki Creole Markup Parser' is distributed in the hope that it will be    //
// useful, but WITHOUT ANY WARRANTY; without even the implied warranty of    //
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General  //
// Public License for more details.                                          //
//                                                                           //
// Some packages include third-party code which is either non-GPL code or    //
// whose copyright and licence status is unclear. These packages are         //
// therefore not (or not clearly) GPL.                                       //
//                                                                           //
// You should have received a copy of the GNU General Public License along   //
// with the 'Wiki Creole Markup Parser'. If not, see                         //
// <http://www.gnu.org/licenses/>.                                           //
///////////////////////////////////////////////////////////////////////////////



grammar creole10;

start
	:	wikipage
	;


///////////////////////////////////////////////////////////////////////////////
////////////////////////   P A R S E R     R U L E S  /////////////////////////
///////////////////////////////////////////////////////////////////////////////

wikipage
	:	( whitespaces )?  paragraphs  EOF
	;

paragraphs
	:	( paragraph )*
	;
paragraph
	:	nowiki_block
	|	blanks  paragraph_separator
	|	( blanks )?
			(	heading
			|	{input.LA(1)==DASH && input.LA(2)==DASH && input.LA(3)==DASH && input.LA(4)==DASH}? horizontalrule
			|	list_unord
			|	list_ord
			|	table
			|	text_paragraph
			)  ( paragraph_separator )?
	;


///////////////////////   T E X T  P A R A G R A P H   ////////////////////////

text_paragraph
	:	(	text_line
		|   nowiki_inline  ( text_element )*  ( text_lineseparator )
		)+
	;
text_line
	:	text_firstelement  ( text_element )*  ( text_lineseparator )
	;
text_firstelement
	:	// predicate prevents using this rule if only ONE star
		{input.LA(1) != STAR || (input.LA(1) == STAR && input.LA(2) == STAR)}? text_formattedelement
	|	text_first_unformattedelement
	;	

text_formattedelement
	:	ital_markup  text_italcontent  ( ( NEWLINE )?  ital_markup )?
	|	bold_markup  text_boldcontent  ( ( NEWLINE )?  bold_markup )?
	;

text_boldcontent
	:	nachmarkup?  ( /*{input.LA(2) != STAR}? STAR?*/ /*onestar*/ text_boldcontentpart )*
	|	EOF
	;
text_element
	:	onestar /*{input.LA(2) != STAR}? STAR?*/  text_unformattedelement
	|	text_unformattedelement  onestar /*STAR?*/
	|	text_formattedelement
	;
nachmarkup
	:	/*{input.LA(2) != DASH && input.LA(2) != POUND && input.LA(2) != EQUAL && input.LA(2) != NEWLINE}? =>*/ ( NEWLINE )
	;
text_italcontent
	:	nachmarkup?  ( /*onestar*/ text_italcontentpart )*
	|	EOF
	;
text_boldcontentpart
	:	ital_markup  text_bolditalcontent  ( ital_markup )?
	|	text_formattedcontent
	;
text_italcontentpart
	:	bold_markup  text_bolditalcontent  ( bold_markup )?
	|	text_formattedcontent
	;
text_bolditalcontent
	:	nachmarkup?  ( text_formattedcontent )?
	|	EOF
	;
text_formattedcontent
	:	( onestar text_unformattedelement   ( text_lineseparator1 )? )+
	;
text_lineseparator1
	:	{input.LA(2) != DASH && input.LA(2) != POUND && input.LA(2) != EQUAL && input.LA(2) != NEWLINE}? text_lineseparator
	;
text_inlineelement
	:	text_first_inlineelement
	|	nowiki_inline
	;
text_first_inlineelement 
	:	link
	|	image
	|	extension
	;
text_first_unformattedelement
	:	text_first_unformatted
	|	text_first_inlineelement
	;
text_first_unformatted
	:	(	~(	//DASH			// horizontal rule
				POUND			// ordered list
			|	STAR			// unordered list
			|	EQUAL			// heading
			|	PIPE			// table
			|	ITAL
			//|	BOLD
			|	LINK_OPEN
			|	IMAGE_OPEN 
			|	NOWIKI_OPEN
			|	EXTENSION
			|	FORCED_LINEBREAK
			|	ESCAPE
			|	NEWLINE
			|	EOF )
		|	forced_linebreak 
		|	escaped )+
	;
text_unformattedelement
	:	text_unformatted
	|	text_inlineelement
	;
text_unformatted
	:	(	~(	ITAL
			//|	BOLD
			|	STAR
			|	LINK_OPEN
			|	IMAGE_OPEN
			|	NOWIKI_OPEN
			|	EXTENSION
			|	FORCED_LINEBREAK
			|	ESCAPE
			|	NEWLINE
			|	EOF )
		|	forced_linebreak 
		|	escaped )+
	;



//////////////////////////////   H E A D I N G   //////////////////////////////

heading
	:	heading_markup  heading_content  ( heading_markup )?  ( blanks )?  paragraph_separator
	;
heading_content
	:	heading_markup  heading_content  ( heading_markup )?
	|	( ~( EQUAL | ESCAPE | NEWLINE | EOF ) | escaped )+
	;



/////////////////////////////////   L I S T   /////////////////////////////////

list_ord
	:	( list_ordelem )+  ( end_of_list )?
	;	
list_ordelem
	:	list_ordelem_markup  list_elem
	;
list_unord
	:	( list_unordelem )+  ( end_of_list )?
	;
list_unordelem
	:	list_unordelem_markup  list_elem
	;
list_elem
	:	( list_elem_markup )*  list_elemcontent  list_elemseparator
	;
list_elem_markup
	:	list_ordelem_markup
	|	list_unordelem_markup
	;
list_elemcontent
	:	( list_elemcontentpart )*
	;
list_elemcontentpart
	:	text_unformattedelement onestar
	|	list_formatted_elem
	;
list_formatted_elem
	:	bold_markup  ( onestar list_boldcontentpart )*/*list_boldcontent?*/  ( bold_markup )?
	|	ital_markup  ( onestar list_italcontentpart )*  ( ital_markup )?
	;
onestar	:	({input.LA(2)!=STAR}? STAR?) | 
	;
//list_boldcontent
//	:	( {input.LA(2)!=STAR}? STAR? list_boldcontentpart )+
//	;
list_boldcontentpart
	:	ital_markup  list_bolditalcontent  ( ital_markup )?
	|	( text_unformattedelement )+
	;
//list_italcontent
//	:	( list_italcontentpart )+
//	;
list_italcontentpart
	:	bold_markup  list_bolditalcontent  ( bold_markup )?
	|	( text_unformattedelement )+
	;
list_bolditalcontent
	:	( text_unformattedelement )+
	;



////////////////////////////////   T A B L E   ////////////////////////////////

table
	:	( table_row )+
	;
table_row 
	:	/*table_cell_markup*/  ( table_cell )+  (table_cell_markup)? table_rowseparator
	;
table_cell
	:	{input.LA(2)==EQUAL}? table_headercell
//	|	table_emptyheadercell
	|	table_normalcell
//	|	table_emptynormalcell
	;
table_headercell
	:	table_headercell_markup  table_cellcontent  //( table_cell_markup )?
	;	
//table_emptyheadercell
//	:	table_headercell_markup  //( table_cell_markup )?
//	;
table_normalcell
	:	table_cell_markup table_cellcontent  //( table_cell_markup )?
	;
//table_emptynormalcell
//	:	table_cell_markup//( table_cell_markup )?
//	;



table_cellcontent
	:	( table_cellcontentpart )+ STAR?
	;
table_cellcontentpart
	:	table_formattedelement
	|	onestar/*{input.LA(2)!= STAR}? STAR?*/ table_unformattedelement
	;
table_formattedelement
	:	ital_markup  table_italcontent?  ( ital_markup )?
	|	bold_markup  table_boldcontent?  ( bold_markup )?
	;
table_boldcontent
	:	( onestar/*{input.LA(2)!= STAR}? STAR?*/ table_boldcontentpart )+
	|	EOF
	;
table_italcontent
	:	( table_italcontentpart )+
	|	EOF
	;
table_boldcontentpart
	:	table_formattedcontent
	|	ital_markup  table_bolditalcontent  ( ital_markup )?
	;
table_italcontentpart
	:	bold_markup  table_bolditalcontent  ( bold_markup )?
	|	table_formattedcontent
	;
table_bolditalcontent
	:	( table_formattedcontent )?
	|	EOF
	;
table_formattedcontent
	:	( table_unformattedelement )+
	;
table_inlineelement
	:	link
	|	image
	|	extension
	|	nowiki_inline
	;
table_unformattedelement
	:	table_unformatted
	|	table_inlineelement
	;
table_unformatted
	:	(	~(	PIPE
			|	ITAL
			|	STAR
			|	LINK_OPEN
			|	IMAGE_OPEN
			|	NOWIKI_OPEN
			|	EXTENSION
			|	FORCED_LINEBREAK
			|	ESCAPE
			|	NEWLINE
			|	EOF )
		|	forced_linebreak 
		|	escaped )+
	;

/*
table_cellcontent
	:	( table_cellcontentpart )+
	;
table_cellcontentpart
	:	table_unformatted_cellelement
	|	table_formatted_cellelement
	;
table_formatted_cellelement
	:	bold_markup  ( table_boldcellcontent )?  ( bold_markup )?
	|	ital_markup  ( table_italcellcontent )?  ( ital_markup )?
	;
table_boldcellcontent
	:	( {input.LA(2) != STAR}? STAR? table_boldcellcontentpart )+
	;
table_boldcellcontentpart
	:	ital_markup  ( table_bolditalcellcontent )?  ( ital_markup )?
	|	( table_unformatted_cellelement )+
	;
table_italcellcontent
	:	( table_italcellcontentpart )+
	;
table_italcellcontentpart
	:	bold_markup  ( table_bolditalcellcontent )?  ( bold_markup )?
	|	( table_unformatted_cellelement )+
	;
table_bolditalcellcontent
	:	( table_unformatted_cellelement )+
	;
table_unformatted_cellelement
	:	table_celltext
	|	text_inlineelement
	;
table_celltext
	:	(	~(	PIPE
			|	ITAL
			|	STAR
			//|	BOLD
			|	LINK_OPEN
			|	IMAGE_OPEN
			|	NOWIKI_OPEN
			|	EXTENSION
			|	FORCED_LINEBREAK
			|	ESCAPE
			|	NEWLINE
			|	EOF )
		|	forced_linebreak
		|	escaped )+
	;
*/


///////////////////////////////   N O W I K I   ///////////////////////////////

nowiki_block
	:	nowikiblock_open_markup  ( ~( NOWIKI_BLOCK_CLOSE | EOF ) )*  nowikiblock_close_markup  paragraph_separator
	;
nowikiblock_open_markup
	:	nowiki_open_markup  newline
	;
nowikiblock_close_markup
	:	NOWIKI_BLOCK_CLOSE
	;
//nowiki_block
//	:	nowiki_open_markup  NEWLINE  ( ~( NOWIKI_CLOSE | EOF ) )*  NEWLINE  nowiki_close_markup  paragraph_separator
//	;
nowiki_inline
	:	nowiki_open_markup  ( ~( NOWIKI_CLOSE | NEWLINE | EOF ) )*  nowiki_close_markup
	;



//////////////////////   H O R I Z O N T A L   R U L E   //////////////////////

horizontalrule
	:	horizontalrule_markup  ( blanks )?  paragraph_separator
	;



/////////////////////////////////   L I N K   /////////////////////////////////

link
	:	link_open_markup  link_address  ( link_description_markup  link_description )?  link_close_markup
	;
link_address
	:	link_interwiki_uri  ':'  link_interwiki_pagename
	|	link_uri
	;
link_interwiki_uri
	:	'C' '2'
/*	|	'D' 'o' 'k' 'u' 'W' 'i' 'k' 'i'
	|	'F' 'l' 'i' 'c' 'k' 'r'
	|	'G' 'o' 'o' 'g' 'l' 'e'
	|	'J' 'S' 'P' 'W' 'i' 'k' 'i'
	|	'M' 'e' 'a' 't' 'b' 'a' 'l' 'l'
	|	'M' 'e' 'd' 'i' 'a' 'W' 'i' 'k' 'i'
	|	'M' 'o' 'i' 'n' 'M' 'o' 'i' 'n'
	|	'O' 'd' 'd' 'm' 'u' 's' 'e'
	|	'O' 'h' 'a' 'n' 'a'
	|	'P' 'm' 'W' 'i' 'k' 'i'
	|	'P' 'u' 'k' 'i' 'W' 'i' 'k' 'i'
	|	'P' 'u' 'r' 'p' 'l' 'e' 'W' 'i' 'k' 'i'
	|	'R' 'a' 'd' 'e' 'o' 'x'
	|	'S' 'n' 'i' 'p' 'S' 'n' 'a' 'p'
	|	'T' 'i' 'd' 'd' 'l' 'y' 'W' 'i' 'k' 'i'
	|	'T' 'W' 'i' 'k' 'i'
	|	'U' 's' 'e' 'm' 'o' 'd'
	|	'W' 'i' 'k' 'i' 'p' 'e' 'd' 'i' 'a'
	|	'X' 'W' 'i' 'k' 'i'
*/	;
link_interwiki_pagename
	:	~( PIPE | LINK_CLOSE | NEWLINE | EOF )+
	;
link_description
	:	( onestar link_descriptionpart | image )+
	;
link_descriptionpart
	:	bold_markup  ( onestar link_bold_descriptionpart )+  bold_markup
	|	ital_markup  ( onestar link_ital_descriptionpart )+  ital_markup
	|	( onestar link_descriptiontext )+ onestar//link_descriptiontext
	;
link_bold_descriptionpart
	:	ital_markup  link_boldital_description  ital_markup
	|	link_descriptiontext
	;
link_ital_descriptionpart
	:	bold_markup  link_boldital_description  bold_markup
	|	link_descriptiontext
	;
link_boldital_description
	:	( onestar link_descriptiontext )+  onestar
	;
link_descriptiontext
	:	(	~(	LINK_CLOSE
			|	ITAL
			|	STAR
			//|	BOLD 
			|	LINK_OPEN
			|	IMAGE_OPEN
			|	NOWIKI_OPEN
			|	EXTENSION
			|	FORCED_LINEBREAK
			|	ESCAPE
			|	NEWLINE
			|	EOF )
		|	forced_linebreak
		|	escaped )+
	;
link_uri
	:	~( PIPE | LINK_CLOSE | NEWLINE | EOF )+
	;



////////////////////////////////   I M A G E   ////////////////////////////////

image
	:	image_open_markup  image_uri  ( image_alternative )?  image_close_markup
	;
image_uri
	:	~( PIPE | IMAGE_CLOSE | NEWLINE | EOF )+
	;
image_alternative
	:	image_alternative_markup  ( image_alternativepart )+
	;
image_alternativepart
	:	bold_markup  ( onestar  image_bold_alternativepart )+  bold_markup
	|	ital_markup  ( onestar  image_ital_alternativepart )+  ital_markup
	|	image_alternativetext
	;
image_bold_alternativepart
	:	ital_markup  link_boldital_description  ital_markup
	|	image_alternativetext
	;
image_ital_alternativepart
	:	bold_markup  link_boldital_description  bold_markup
	|	image_alternativetext
	;
image_boldital_alternative
	:	( onestar  image_alternativetext )+  onestar
	;
image_alternativetext
	:	(	~( 	IMAGE_CLOSE
			|	ITAL
			//|	BOLD
			|	STAR
			|	LINK_OPEN
			|	IMAGE_OPEN
			|	NOWIKI_OPEN
			|	EXTENSION
			|	FORCED_LINEBREAK
			|	NEWLINE
			|	EOF )
		|	forced_linebreak )+
	;


/////////////////////////////  E X T E N S I O N  /////////////////////////////

extension
	:	extension_markup  extension_handler  blanks  extension_statement  extension_markup
	;
extension_handler
	:	(~( EXTENSION  |  BLANKS  |  ESCAPE  |  NEWLINE  |  EOF ) | escaped )+
	;
extension_statement
	:	(~( EXTENSION  |  ESCAPE  |  EOF ) | escaped )*
	;





escaped
	:	ESCAPE STAR STAR
	|	ESCAPE .	// in parser rule . means arbitrary TOKEN, not character
	;
paragraph_separator
	:	( newline )+
	|	EOF
	;
whitespaces
	:	( blanks  |  newline )+
	;
blanks
	:	BLANKS
	;
text_lineseparator
	:	newline  ( blanks )?  //TODO blanks
	|	EOF
	;
newline
	:	NEWLINE
	;
bold_markup
	:	STAR STAR
	;
ital_markup
	:	ITAL
	;
heading_markup
	:	EQUAL
	;
list_ordelem_markup
	:	POUND
	;
list_unordelem_markup
	:	STAR
	;
list_elemseparator
	:	newline ( blanks )? //TODO blanks
	|	EOF
	;
end_of_list
	:	newline
	|	EOF
	;
table_cell_markup
	:	PIPE
	;
table_headercell_markup
	:	PIPE EQUAL
	;
table_rowseparator
	:	newline
	|	EOF
	;
nowiki_open_markup
	:	NOWIKI_OPEN
	;
nowiki_close_markup
	:	NOWIKI_CLOSE
	;
horizontalrule_markup
	:	DASH DASH DASH DASH
	;
link_open_markup
	:	LINK_OPEN
	;
link_close_markup
	:	LINK_CLOSE
	;
link_description_markup
	:	PIPE
	;
image_open_markup
	:	IMAGE_OPEN
	;
image_close_markup
	:	IMAGE_CLOSE
	;
image_alternative_markup
	:	PIPE
	;
extension_markup
	:	EXTENSION
	;
forced_linebreak
	:	FORCED_LINEBREAK
	;


///////////////////////////////////////////////////////////////////////////////
////////////////////////// S C A N E R    R U L E S ///////////////////////////
///////////////////////////////////////////////////////////////////////////////

ESCAPE				: '~';
NOWIKI_BLOCK_CLOSE	: NEWLINE '}}}';
NEWLINE				: ( CR )? LF | CR;
fragment CR			: '\r';
fragment LF			: '\n';

BLANKS				: ( SPACE | TABULATOR )+;
fragment SPACE		: ' ';
fragment TABULATOR	: '\t';

// if followed by a SLASH, then it's an URI indicator '://'
// prohibited to match '://' with COLON ITAL !
COLON_SLASH			: ':' '/';
ITAL				: '//';
//BOLD				: '**';
NOWIKI_OPEN			: '{{{';
NOWIKI_CLOSE		: '}}}';
LINK_OPEN			: '[[';
LINK_CLOSE			: ']]';
IMAGE_OPEN			: '{{';
IMAGE_CLOSE			: '}}';
FORCED_LINEBREAK	: '\\\\';
EQUAL				: '=';
PIPE				: '|';
POUND				: '#';
DASH				: '-';
STAR				: '*';
SLASH				: '/';
EXTENSION			: '@@';

// ATTENTION: NOT CREOLE !!!
// LINE_COMMENT		: '%' ( ~( '\n' | '\r' ) )* ( NEWLINE | EOF ) {$channel=HIDDEN;};

INSIGNIFICANT_CHAR	: .;


