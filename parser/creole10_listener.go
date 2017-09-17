// Generated from Creole10.g4 by ANTLR 4.7.

package parser // Creole10

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Creole10Listener is a complete listener for a parse tree produced by Creole10Parser.
type Creole10Listener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterWikipage is called when entering the wikipage production.
	EnterWikipage(c *WikipageContext)

	// EnterParagraphs is called when entering the paragraphs production.
	EnterParagraphs(c *ParagraphsContext)

	// EnterParagraph is called when entering the paragraph production.
	EnterParagraph(c *ParagraphContext)

	// EnterText_paragraph is called when entering the text_paragraph production.
	EnterText_paragraph(c *Text_paragraphContext)

	// EnterText_line is called when entering the text_line production.
	EnterText_line(c *Text_lineContext)

	// EnterText_firstelement is called when entering the text_firstelement production.
	EnterText_firstelement(c *Text_firstelementContext)

	// EnterText_formattedelement is called when entering the text_formattedelement production.
	EnterText_formattedelement(c *Text_formattedelementContext)

	// EnterText_boldcontent is called when entering the text_boldcontent production.
	EnterText_boldcontent(c *Text_boldcontentContext)

	// EnterText_element is called when entering the text_element production.
	EnterText_element(c *Text_elementContext)

	// EnterNachmarkup is called when entering the nachmarkup production.
	EnterNachmarkup(c *NachmarkupContext)

	// EnterText_italcontent is called when entering the text_italcontent production.
	EnterText_italcontent(c *Text_italcontentContext)

	// EnterText_boldcontentpart is called when entering the text_boldcontentpart production.
	EnterText_boldcontentpart(c *Text_boldcontentpartContext)

	// EnterText_italcontentpart is called when entering the text_italcontentpart production.
	EnterText_italcontentpart(c *Text_italcontentpartContext)

	// EnterText_bolditalcontent is called when entering the text_bolditalcontent production.
	EnterText_bolditalcontent(c *Text_bolditalcontentContext)

	// EnterText_formattedcontent is called when entering the text_formattedcontent production.
	EnterText_formattedcontent(c *Text_formattedcontentContext)

	// EnterText_lineseparator1 is called when entering the text_lineseparator1 production.
	EnterText_lineseparator1(c *Text_lineseparator1Context)

	// EnterText_inlineelement is called when entering the text_inlineelement production.
	EnterText_inlineelement(c *Text_inlineelementContext)

	// EnterText_first_inlineelement is called when entering the text_first_inlineelement production.
	EnterText_first_inlineelement(c *Text_first_inlineelementContext)

	// EnterText_first_unformattedelement is called when entering the text_first_unformattedelement production.
	EnterText_first_unformattedelement(c *Text_first_unformattedelementContext)

	// EnterText_first_unformatted is called when entering the text_first_unformatted production.
	EnterText_first_unformatted(c *Text_first_unformattedContext)

	// EnterText_unformattedelement is called when entering the text_unformattedelement production.
	EnterText_unformattedelement(c *Text_unformattedelementContext)

	// EnterText_unformatted is called when entering the text_unformatted production.
	EnterText_unformatted(c *Text_unformattedContext)

	// EnterHeading is called when entering the heading production.
	EnterHeading(c *HeadingContext)

	// EnterHeading_content is called when entering the heading_content production.
	EnterHeading_content(c *Heading_contentContext)

	// EnterList_ord is called when entering the list_ord production.
	EnterList_ord(c *List_ordContext)

	// EnterList_ordelem is called when entering the list_ordelem production.
	EnterList_ordelem(c *List_ordelemContext)

	// EnterList_unord is called when entering the list_unord production.
	EnterList_unord(c *List_unordContext)

	// EnterList_unordelem is called when entering the list_unordelem production.
	EnterList_unordelem(c *List_unordelemContext)

	// EnterList_elem is called when entering the list_elem production.
	EnterList_elem(c *List_elemContext)

	// EnterList_elem_markup is called when entering the list_elem_markup production.
	EnterList_elem_markup(c *List_elem_markupContext)

	// EnterList_elemcontent is called when entering the list_elemcontent production.
	EnterList_elemcontent(c *List_elemcontentContext)

	// EnterList_elemcontentpart is called when entering the list_elemcontentpart production.
	EnterList_elemcontentpart(c *List_elemcontentpartContext)

	// EnterList_formatted_elem is called when entering the list_formatted_elem production.
	EnterList_formatted_elem(c *List_formatted_elemContext)

	// EnterOnestar is called when entering the onestar production.
	EnterOnestar(c *OnestarContext)

	// EnterList_boldcontentpart is called when entering the list_boldcontentpart production.
	EnterList_boldcontentpart(c *List_boldcontentpartContext)

	// EnterList_italcontentpart is called when entering the list_italcontentpart production.
	EnterList_italcontentpart(c *List_italcontentpartContext)

	// EnterList_bolditalcontent is called when entering the list_bolditalcontent production.
	EnterList_bolditalcontent(c *List_bolditalcontentContext)

	// EnterTable is called when entering the table production.
	EnterTable(c *TableContext)

	// EnterTable_row is called when entering the table_row production.
	EnterTable_row(c *Table_rowContext)

	// EnterTable_cell is called when entering the table_cell production.
	EnterTable_cell(c *Table_cellContext)

	// EnterTable_headercell is called when entering the table_headercell production.
	EnterTable_headercell(c *Table_headercellContext)

	// EnterTable_normalcell is called when entering the table_normalcell production.
	EnterTable_normalcell(c *Table_normalcellContext)

	// EnterTable_cellcontent is called when entering the table_cellcontent production.
	EnterTable_cellcontent(c *Table_cellcontentContext)

	// EnterTable_cellcontentpart is called when entering the table_cellcontentpart production.
	EnterTable_cellcontentpart(c *Table_cellcontentpartContext)

	// EnterTable_formattedelement is called when entering the table_formattedelement production.
	EnterTable_formattedelement(c *Table_formattedelementContext)

	// EnterTable_boldcontent is called when entering the table_boldcontent production.
	EnterTable_boldcontent(c *Table_boldcontentContext)

	// EnterTable_italcontent is called when entering the table_italcontent production.
	EnterTable_italcontent(c *Table_italcontentContext)

	// EnterTable_boldcontentpart is called when entering the table_boldcontentpart production.
	EnterTable_boldcontentpart(c *Table_boldcontentpartContext)

	// EnterTable_italcontentpart is called when entering the table_italcontentpart production.
	EnterTable_italcontentpart(c *Table_italcontentpartContext)

	// EnterTable_bolditalcontent is called when entering the table_bolditalcontent production.
	EnterTable_bolditalcontent(c *Table_bolditalcontentContext)

	// EnterTable_formattedcontent is called when entering the table_formattedcontent production.
	EnterTable_formattedcontent(c *Table_formattedcontentContext)

	// EnterTable_inlineelement is called when entering the table_inlineelement production.
	EnterTable_inlineelement(c *Table_inlineelementContext)

	// EnterTable_unformattedelement is called when entering the table_unformattedelement production.
	EnterTable_unformattedelement(c *Table_unformattedelementContext)

	// EnterTable_unformatted is called when entering the table_unformatted production.
	EnterTable_unformatted(c *Table_unformattedContext)

	// EnterNowiki_block is called when entering the nowiki_block production.
	EnterNowiki_block(c *Nowiki_blockContext)

	// EnterNowiki_block_content is called when entering the nowiki_block_content production.
	EnterNowiki_block_content(c *Nowiki_block_contentContext)

	// EnterNowikiblock_open_markup is called when entering the nowikiblock_open_markup production.
	EnterNowikiblock_open_markup(c *Nowikiblock_open_markupContext)

	// EnterNowikiblock_close_markup is called when entering the nowikiblock_close_markup production.
	EnterNowikiblock_close_markup(c *Nowikiblock_close_markupContext)

	// EnterNowiki_inline is called when entering the nowiki_inline production.
	EnterNowiki_inline(c *Nowiki_inlineContext)

	// EnterNowiki_inline_content is called when entering the nowiki_inline_content production.
	EnterNowiki_inline_content(c *Nowiki_inline_contentContext)

	// EnterHorizontalrule is called when entering the horizontalrule production.
	EnterHorizontalrule(c *HorizontalruleContext)

	// EnterLink is called when entering the link production.
	EnterLink(c *LinkContext)

	// EnterLink_address is called when entering the link_address production.
	EnterLink_address(c *Link_addressContext)

	// EnterLink_interwiki_uri is called when entering the link_interwiki_uri production.
	EnterLink_interwiki_uri(c *Link_interwiki_uriContext)

	// EnterLink_interwiki_pagename is called when entering the link_interwiki_pagename production.
	EnterLink_interwiki_pagename(c *Link_interwiki_pagenameContext)

	// EnterLink_description is called when entering the link_description production.
	EnterLink_description(c *Link_descriptionContext)

	// EnterLink_descriptionpart is called when entering the link_descriptionpart production.
	EnterLink_descriptionpart(c *Link_descriptionpartContext)

	// EnterLink_bold_descriptionpart is called when entering the link_bold_descriptionpart production.
	EnterLink_bold_descriptionpart(c *Link_bold_descriptionpartContext)

	// EnterLink_ital_descriptionpart is called when entering the link_ital_descriptionpart production.
	EnterLink_ital_descriptionpart(c *Link_ital_descriptionpartContext)

	// EnterLink_boldital_description is called when entering the link_boldital_description production.
	EnterLink_boldital_description(c *Link_boldital_descriptionContext)

	// EnterLink_descriptiontext is called when entering the link_descriptiontext production.
	EnterLink_descriptiontext(c *Link_descriptiontextContext)

	// EnterLink_uri is called when entering the link_uri production.
	EnterLink_uri(c *Link_uriContext)

	// EnterImage is called when entering the image production.
	EnterImage(c *ImageContext)

	// EnterImage_uri is called when entering the image_uri production.
	EnterImage_uri(c *Image_uriContext)

	// EnterImage_alternative is called when entering the image_alternative production.
	EnterImage_alternative(c *Image_alternativeContext)

	// EnterImage_alternativepart is called when entering the image_alternativepart production.
	EnterImage_alternativepart(c *Image_alternativepartContext)

	// EnterImage_bold_alternativepart is called when entering the image_bold_alternativepart production.
	EnterImage_bold_alternativepart(c *Image_bold_alternativepartContext)

	// EnterImage_ital_alternativepart is called when entering the image_ital_alternativepart production.
	EnterImage_ital_alternativepart(c *Image_ital_alternativepartContext)

	// EnterImage_boldital_alternative is called when entering the image_boldital_alternative production.
	EnterImage_boldital_alternative(c *Image_boldital_alternativeContext)

	// EnterImage_alternativetext is called when entering the image_alternativetext production.
	EnterImage_alternativetext(c *Image_alternativetextContext)

	// EnterExtension is called when entering the extension production.
	EnterExtension(c *ExtensionContext)

	// EnterExtension_handler is called when entering the extension_handler production.
	EnterExtension_handler(c *Extension_handlerContext)

	// EnterExtension_statement is called when entering the extension_statement production.
	EnterExtension_statement(c *Extension_statementContext)

	// EnterEscaped is called when entering the escaped production.
	EnterEscaped(c *EscapedContext)

	// EnterParagraph_separator is called when entering the paragraph_separator production.
	EnterParagraph_separator(c *Paragraph_separatorContext)

	// EnterWhitespaces is called when entering the whitespaces production.
	EnterWhitespaces(c *WhitespacesContext)

	// EnterBlanks is called when entering the blanks production.
	EnterBlanks(c *BlanksContext)

	// EnterText_lineseparator is called when entering the text_lineseparator production.
	EnterText_lineseparator(c *Text_lineseparatorContext)

	// EnterNewline is called when entering the newline production.
	EnterNewline(c *NewlineContext)

	// EnterBold_markup is called when entering the bold_markup production.
	EnterBold_markup(c *Bold_markupContext)

	// EnterItal_markup is called when entering the ital_markup production.
	EnterItal_markup(c *Ital_markupContext)

	// EnterHeading_markup is called when entering the heading_markup production.
	EnterHeading_markup(c *Heading_markupContext)

	// EnterList_ordelem_markup is called when entering the list_ordelem_markup production.
	EnterList_ordelem_markup(c *List_ordelem_markupContext)

	// EnterList_unordelem_markup is called when entering the list_unordelem_markup production.
	EnterList_unordelem_markup(c *List_unordelem_markupContext)

	// EnterList_elemseparator is called when entering the list_elemseparator production.
	EnterList_elemseparator(c *List_elemseparatorContext)

	// EnterEnd_of_list is called when entering the end_of_list production.
	EnterEnd_of_list(c *End_of_listContext)

	// EnterTable_cell_markup is called when entering the table_cell_markup production.
	EnterTable_cell_markup(c *Table_cell_markupContext)

	// EnterTable_headercell_markup is called when entering the table_headercell_markup production.
	EnterTable_headercell_markup(c *Table_headercell_markupContext)

	// EnterTable_rowseparator is called when entering the table_rowseparator production.
	EnterTable_rowseparator(c *Table_rowseparatorContext)

	// EnterNowiki_open_markup is called when entering the nowiki_open_markup production.
	EnterNowiki_open_markup(c *Nowiki_open_markupContext)

	// EnterNowiki_close_markup is called when entering the nowiki_close_markup production.
	EnterNowiki_close_markup(c *Nowiki_close_markupContext)

	// EnterHorizontalrule_markup is called when entering the horizontalrule_markup production.
	EnterHorizontalrule_markup(c *Horizontalrule_markupContext)

	// EnterLink_open_markup is called when entering the link_open_markup production.
	EnterLink_open_markup(c *Link_open_markupContext)

	// EnterLink_close_markup is called when entering the link_close_markup production.
	EnterLink_close_markup(c *Link_close_markupContext)

	// EnterLink_description_markup is called when entering the link_description_markup production.
	EnterLink_description_markup(c *Link_description_markupContext)

	// EnterImage_open_markup is called when entering the image_open_markup production.
	EnterImage_open_markup(c *Image_open_markupContext)

	// EnterImage_close_markup is called when entering the image_close_markup production.
	EnterImage_close_markup(c *Image_close_markupContext)

	// EnterImage_alternative_markup is called when entering the image_alternative_markup production.
	EnterImage_alternative_markup(c *Image_alternative_markupContext)

	// EnterExtension_markup is called when entering the extension_markup production.
	EnterExtension_markup(c *Extension_markupContext)

	// EnterForced_linebreak is called when entering the forced_linebreak production.
	EnterForced_linebreak(c *Forced_linebreakContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitWikipage is called when exiting the wikipage production.
	ExitWikipage(c *WikipageContext)

	// ExitParagraphs is called when exiting the paragraphs production.
	ExitParagraphs(c *ParagraphsContext)

	// ExitParagraph is called when exiting the paragraph production.
	ExitParagraph(c *ParagraphContext)

	// ExitText_paragraph is called when exiting the text_paragraph production.
	ExitText_paragraph(c *Text_paragraphContext)

	// ExitText_line is called when exiting the text_line production.
	ExitText_line(c *Text_lineContext)

	// ExitText_firstelement is called when exiting the text_firstelement production.
	ExitText_firstelement(c *Text_firstelementContext)

	// ExitText_formattedelement is called when exiting the text_formattedelement production.
	ExitText_formattedelement(c *Text_formattedelementContext)

	// ExitText_boldcontent is called when exiting the text_boldcontent production.
	ExitText_boldcontent(c *Text_boldcontentContext)

	// ExitText_element is called when exiting the text_element production.
	ExitText_element(c *Text_elementContext)

	// ExitNachmarkup is called when exiting the nachmarkup production.
	ExitNachmarkup(c *NachmarkupContext)

	// ExitText_italcontent is called when exiting the text_italcontent production.
	ExitText_italcontent(c *Text_italcontentContext)

	// ExitText_boldcontentpart is called when exiting the text_boldcontentpart production.
	ExitText_boldcontentpart(c *Text_boldcontentpartContext)

	// ExitText_italcontentpart is called when exiting the text_italcontentpart production.
	ExitText_italcontentpart(c *Text_italcontentpartContext)

	// ExitText_bolditalcontent is called when exiting the text_bolditalcontent production.
	ExitText_bolditalcontent(c *Text_bolditalcontentContext)

	// ExitText_formattedcontent is called when exiting the text_formattedcontent production.
	ExitText_formattedcontent(c *Text_formattedcontentContext)

	// ExitText_lineseparator1 is called when exiting the text_lineseparator1 production.
	ExitText_lineseparator1(c *Text_lineseparator1Context)

	// ExitText_inlineelement is called when exiting the text_inlineelement production.
	ExitText_inlineelement(c *Text_inlineelementContext)

	// ExitText_first_inlineelement is called when exiting the text_first_inlineelement production.
	ExitText_first_inlineelement(c *Text_first_inlineelementContext)

	// ExitText_first_unformattedelement is called when exiting the text_first_unformattedelement production.
	ExitText_first_unformattedelement(c *Text_first_unformattedelementContext)

	// ExitText_first_unformatted is called when exiting the text_first_unformatted production.
	ExitText_first_unformatted(c *Text_first_unformattedContext)

	// ExitText_unformattedelement is called when exiting the text_unformattedelement production.
	ExitText_unformattedelement(c *Text_unformattedelementContext)

	// ExitText_unformatted is called when exiting the text_unformatted production.
	ExitText_unformatted(c *Text_unformattedContext)

	// ExitHeading is called when exiting the heading production.
	ExitHeading(c *HeadingContext)

	// ExitHeading_content is called when exiting the heading_content production.
	ExitHeading_content(c *Heading_contentContext)

	// ExitList_ord is called when exiting the list_ord production.
	ExitList_ord(c *List_ordContext)

	// ExitList_ordelem is called when exiting the list_ordelem production.
	ExitList_ordelem(c *List_ordelemContext)

	// ExitList_unord is called when exiting the list_unord production.
	ExitList_unord(c *List_unordContext)

	// ExitList_unordelem is called when exiting the list_unordelem production.
	ExitList_unordelem(c *List_unordelemContext)

	// ExitList_elem is called when exiting the list_elem production.
	ExitList_elem(c *List_elemContext)

	// ExitList_elem_markup is called when exiting the list_elem_markup production.
	ExitList_elem_markup(c *List_elem_markupContext)

	// ExitList_elemcontent is called when exiting the list_elemcontent production.
	ExitList_elemcontent(c *List_elemcontentContext)

	// ExitList_elemcontentpart is called when exiting the list_elemcontentpart production.
	ExitList_elemcontentpart(c *List_elemcontentpartContext)

	// ExitList_formatted_elem is called when exiting the list_formatted_elem production.
	ExitList_formatted_elem(c *List_formatted_elemContext)

	// ExitOnestar is called when exiting the onestar production.
	ExitOnestar(c *OnestarContext)

	// ExitList_boldcontentpart is called when exiting the list_boldcontentpart production.
	ExitList_boldcontentpart(c *List_boldcontentpartContext)

	// ExitList_italcontentpart is called when exiting the list_italcontentpart production.
	ExitList_italcontentpart(c *List_italcontentpartContext)

	// ExitList_bolditalcontent is called when exiting the list_bolditalcontent production.
	ExitList_bolditalcontent(c *List_bolditalcontentContext)

	// ExitTable is called when exiting the table production.
	ExitTable(c *TableContext)

	// ExitTable_row is called when exiting the table_row production.
	ExitTable_row(c *Table_rowContext)

	// ExitTable_cell is called when exiting the table_cell production.
	ExitTable_cell(c *Table_cellContext)

	// ExitTable_headercell is called when exiting the table_headercell production.
	ExitTable_headercell(c *Table_headercellContext)

	// ExitTable_normalcell is called when exiting the table_normalcell production.
	ExitTable_normalcell(c *Table_normalcellContext)

	// ExitTable_cellcontent is called when exiting the table_cellcontent production.
	ExitTable_cellcontent(c *Table_cellcontentContext)

	// ExitTable_cellcontentpart is called when exiting the table_cellcontentpart production.
	ExitTable_cellcontentpart(c *Table_cellcontentpartContext)

	// ExitTable_formattedelement is called when exiting the table_formattedelement production.
	ExitTable_formattedelement(c *Table_formattedelementContext)

	// ExitTable_boldcontent is called when exiting the table_boldcontent production.
	ExitTable_boldcontent(c *Table_boldcontentContext)

	// ExitTable_italcontent is called when exiting the table_italcontent production.
	ExitTable_italcontent(c *Table_italcontentContext)

	// ExitTable_boldcontentpart is called when exiting the table_boldcontentpart production.
	ExitTable_boldcontentpart(c *Table_boldcontentpartContext)

	// ExitTable_italcontentpart is called when exiting the table_italcontentpart production.
	ExitTable_italcontentpart(c *Table_italcontentpartContext)

	// ExitTable_bolditalcontent is called when exiting the table_bolditalcontent production.
	ExitTable_bolditalcontent(c *Table_bolditalcontentContext)

	// ExitTable_formattedcontent is called when exiting the table_formattedcontent production.
	ExitTable_formattedcontent(c *Table_formattedcontentContext)

	// ExitTable_inlineelement is called when exiting the table_inlineelement production.
	ExitTable_inlineelement(c *Table_inlineelementContext)

	// ExitTable_unformattedelement is called when exiting the table_unformattedelement production.
	ExitTable_unformattedelement(c *Table_unformattedelementContext)

	// ExitTable_unformatted is called when exiting the table_unformatted production.
	ExitTable_unformatted(c *Table_unformattedContext)

	// ExitNowiki_block is called when exiting the nowiki_block production.
	ExitNowiki_block(c *Nowiki_blockContext)

	// ExitNowiki_block_content is called when exiting the nowiki_block_content production.
	ExitNowiki_block_content(c *Nowiki_block_contentContext)

	// ExitNowikiblock_open_markup is called when exiting the nowikiblock_open_markup production.
	ExitNowikiblock_open_markup(c *Nowikiblock_open_markupContext)

	// ExitNowikiblock_close_markup is called when exiting the nowikiblock_close_markup production.
	ExitNowikiblock_close_markup(c *Nowikiblock_close_markupContext)

	// ExitNowiki_inline is called when exiting the nowiki_inline production.
	ExitNowiki_inline(c *Nowiki_inlineContext)

	// ExitNowiki_inline_content is called when exiting the nowiki_inline_content production.
	ExitNowiki_inline_content(c *Nowiki_inline_contentContext)

	// ExitHorizontalrule is called when exiting the horizontalrule production.
	ExitHorizontalrule(c *HorizontalruleContext)

	// ExitLink is called when exiting the link production.
	ExitLink(c *LinkContext)

	// ExitLink_address is called when exiting the link_address production.
	ExitLink_address(c *Link_addressContext)

	// ExitLink_interwiki_uri is called when exiting the link_interwiki_uri production.
	ExitLink_interwiki_uri(c *Link_interwiki_uriContext)

	// ExitLink_interwiki_pagename is called when exiting the link_interwiki_pagename production.
	ExitLink_interwiki_pagename(c *Link_interwiki_pagenameContext)

	// ExitLink_description is called when exiting the link_description production.
	ExitLink_description(c *Link_descriptionContext)

	// ExitLink_descriptionpart is called when exiting the link_descriptionpart production.
	ExitLink_descriptionpart(c *Link_descriptionpartContext)

	// ExitLink_bold_descriptionpart is called when exiting the link_bold_descriptionpart production.
	ExitLink_bold_descriptionpart(c *Link_bold_descriptionpartContext)

	// ExitLink_ital_descriptionpart is called when exiting the link_ital_descriptionpart production.
	ExitLink_ital_descriptionpart(c *Link_ital_descriptionpartContext)

	// ExitLink_boldital_description is called when exiting the link_boldital_description production.
	ExitLink_boldital_description(c *Link_boldital_descriptionContext)

	// ExitLink_descriptiontext is called when exiting the link_descriptiontext production.
	ExitLink_descriptiontext(c *Link_descriptiontextContext)

	// ExitLink_uri is called when exiting the link_uri production.
	ExitLink_uri(c *Link_uriContext)

	// ExitImage is called when exiting the image production.
	ExitImage(c *ImageContext)

	// ExitImage_uri is called when exiting the image_uri production.
	ExitImage_uri(c *Image_uriContext)

	// ExitImage_alternative is called when exiting the image_alternative production.
	ExitImage_alternative(c *Image_alternativeContext)

	// ExitImage_alternativepart is called when exiting the image_alternativepart production.
	ExitImage_alternativepart(c *Image_alternativepartContext)

	// ExitImage_bold_alternativepart is called when exiting the image_bold_alternativepart production.
	ExitImage_bold_alternativepart(c *Image_bold_alternativepartContext)

	// ExitImage_ital_alternativepart is called when exiting the image_ital_alternativepart production.
	ExitImage_ital_alternativepart(c *Image_ital_alternativepartContext)

	// ExitImage_boldital_alternative is called when exiting the image_boldital_alternative production.
	ExitImage_boldital_alternative(c *Image_boldital_alternativeContext)

	// ExitImage_alternativetext is called when exiting the image_alternativetext production.
	ExitImage_alternativetext(c *Image_alternativetextContext)

	// ExitExtension is called when exiting the extension production.
	ExitExtension(c *ExtensionContext)

	// ExitExtension_handler is called when exiting the extension_handler production.
	ExitExtension_handler(c *Extension_handlerContext)

	// ExitExtension_statement is called when exiting the extension_statement production.
	ExitExtension_statement(c *Extension_statementContext)

	// ExitEscaped is called when exiting the escaped production.
	ExitEscaped(c *EscapedContext)

	// ExitParagraph_separator is called when exiting the paragraph_separator production.
	ExitParagraph_separator(c *Paragraph_separatorContext)

	// ExitWhitespaces is called when exiting the whitespaces production.
	ExitWhitespaces(c *WhitespacesContext)

	// ExitBlanks is called when exiting the blanks production.
	ExitBlanks(c *BlanksContext)

	// ExitText_lineseparator is called when exiting the text_lineseparator production.
	ExitText_lineseparator(c *Text_lineseparatorContext)

	// ExitNewline is called when exiting the newline production.
	ExitNewline(c *NewlineContext)

	// ExitBold_markup is called when exiting the bold_markup production.
	ExitBold_markup(c *Bold_markupContext)

	// ExitItal_markup is called when exiting the ital_markup production.
	ExitItal_markup(c *Ital_markupContext)

	// ExitHeading_markup is called when exiting the heading_markup production.
	ExitHeading_markup(c *Heading_markupContext)

	// ExitList_ordelem_markup is called when exiting the list_ordelem_markup production.
	ExitList_ordelem_markup(c *List_ordelem_markupContext)

	// ExitList_unordelem_markup is called when exiting the list_unordelem_markup production.
	ExitList_unordelem_markup(c *List_unordelem_markupContext)

	// ExitList_elemseparator is called when exiting the list_elemseparator production.
	ExitList_elemseparator(c *List_elemseparatorContext)

	// ExitEnd_of_list is called when exiting the end_of_list production.
	ExitEnd_of_list(c *End_of_listContext)

	// ExitTable_cell_markup is called when exiting the table_cell_markup production.
	ExitTable_cell_markup(c *Table_cell_markupContext)

	// ExitTable_headercell_markup is called when exiting the table_headercell_markup production.
	ExitTable_headercell_markup(c *Table_headercell_markupContext)

	// ExitTable_rowseparator is called when exiting the table_rowseparator production.
	ExitTable_rowseparator(c *Table_rowseparatorContext)

	// ExitNowiki_open_markup is called when exiting the nowiki_open_markup production.
	ExitNowiki_open_markup(c *Nowiki_open_markupContext)

	// ExitNowiki_close_markup is called when exiting the nowiki_close_markup production.
	ExitNowiki_close_markup(c *Nowiki_close_markupContext)

	// ExitHorizontalrule_markup is called when exiting the horizontalrule_markup production.
	ExitHorizontalrule_markup(c *Horizontalrule_markupContext)

	// ExitLink_open_markup is called when exiting the link_open_markup production.
	ExitLink_open_markup(c *Link_open_markupContext)

	// ExitLink_close_markup is called when exiting the link_close_markup production.
	ExitLink_close_markup(c *Link_close_markupContext)

	// ExitLink_description_markup is called when exiting the link_description_markup production.
	ExitLink_description_markup(c *Link_description_markupContext)

	// ExitImage_open_markup is called when exiting the image_open_markup production.
	ExitImage_open_markup(c *Image_open_markupContext)

	// ExitImage_close_markup is called when exiting the image_close_markup production.
	ExitImage_close_markup(c *Image_close_markupContext)

	// ExitImage_alternative_markup is called when exiting the image_alternative_markup production.
	ExitImage_alternative_markup(c *Image_alternative_markupContext)

	// ExitExtension_markup is called when exiting the extension_markup production.
	ExitExtension_markup(c *Extension_markupContext)

	// ExitForced_linebreak is called when exiting the forced_linebreak production.
	ExitForced_linebreak(c *Forced_linebreakContext)
}
