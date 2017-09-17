// Generated from Creole10.g4 by ANTLR 4.7.

package parser // Creole10

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCreole10Listener is a complete listener for a parse tree produced by Creole10Parser.
type BaseCreole10Listener struct{}

var _ Creole10Listener = &BaseCreole10Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCreole10Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCreole10Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCreole10Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCreole10Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseCreole10Listener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseCreole10Listener) ExitStart(ctx *StartContext) {}

// EnterWikipage is called when production wikipage is entered.
func (s *BaseCreole10Listener) EnterWikipage(ctx *WikipageContext) {}

// ExitWikipage is called when production wikipage is exited.
func (s *BaseCreole10Listener) ExitWikipage(ctx *WikipageContext) {}

// EnterParagraphs is called when production paragraphs is entered.
func (s *BaseCreole10Listener) EnterParagraphs(ctx *ParagraphsContext) {}

// ExitParagraphs is called when production paragraphs is exited.
func (s *BaseCreole10Listener) ExitParagraphs(ctx *ParagraphsContext) {}

// EnterParagraph is called when production paragraph is entered.
func (s *BaseCreole10Listener) EnterParagraph(ctx *ParagraphContext) {}

// ExitParagraph is called when production paragraph is exited.
func (s *BaseCreole10Listener) ExitParagraph(ctx *ParagraphContext) {}

// EnterText_paragraph is called when production text_paragraph is entered.
func (s *BaseCreole10Listener) EnterText_paragraph(ctx *Text_paragraphContext) {}

// ExitText_paragraph is called when production text_paragraph is exited.
func (s *BaseCreole10Listener) ExitText_paragraph(ctx *Text_paragraphContext) {}

// EnterText_line is called when production text_line is entered.
func (s *BaseCreole10Listener) EnterText_line(ctx *Text_lineContext) {}

// ExitText_line is called when production text_line is exited.
func (s *BaseCreole10Listener) ExitText_line(ctx *Text_lineContext) {}

// EnterText_firstelement is called when production text_firstelement is entered.
func (s *BaseCreole10Listener) EnterText_firstelement(ctx *Text_firstelementContext) {}

// ExitText_firstelement is called when production text_firstelement is exited.
func (s *BaseCreole10Listener) ExitText_firstelement(ctx *Text_firstelementContext) {}

// EnterText_formattedelement is called when production text_formattedelement is entered.
func (s *BaseCreole10Listener) EnterText_formattedelement(ctx *Text_formattedelementContext) {}

// ExitText_formattedelement is called when production text_formattedelement is exited.
func (s *BaseCreole10Listener) ExitText_formattedelement(ctx *Text_formattedelementContext) {}

// EnterText_boldcontent is called when production text_boldcontent is entered.
func (s *BaseCreole10Listener) EnterText_boldcontent(ctx *Text_boldcontentContext) {}

// ExitText_boldcontent is called when production text_boldcontent is exited.
func (s *BaseCreole10Listener) ExitText_boldcontent(ctx *Text_boldcontentContext) {}

// EnterText_element is called when production text_element is entered.
func (s *BaseCreole10Listener) EnterText_element(ctx *Text_elementContext) {}

// ExitText_element is called when production text_element is exited.
func (s *BaseCreole10Listener) ExitText_element(ctx *Text_elementContext) {}

// EnterNachmarkup is called when production nachmarkup is entered.
func (s *BaseCreole10Listener) EnterNachmarkup(ctx *NachmarkupContext) {}

// ExitNachmarkup is called when production nachmarkup is exited.
func (s *BaseCreole10Listener) ExitNachmarkup(ctx *NachmarkupContext) {}

// EnterText_italcontent is called when production text_italcontent is entered.
func (s *BaseCreole10Listener) EnterText_italcontent(ctx *Text_italcontentContext) {}

// ExitText_italcontent is called when production text_italcontent is exited.
func (s *BaseCreole10Listener) ExitText_italcontent(ctx *Text_italcontentContext) {}

// EnterText_boldcontentpart is called when production text_boldcontentpart is entered.
func (s *BaseCreole10Listener) EnterText_boldcontentpart(ctx *Text_boldcontentpartContext) {}

// ExitText_boldcontentpart is called when production text_boldcontentpart is exited.
func (s *BaseCreole10Listener) ExitText_boldcontentpart(ctx *Text_boldcontentpartContext) {}

// EnterText_italcontentpart is called when production text_italcontentpart is entered.
func (s *BaseCreole10Listener) EnterText_italcontentpart(ctx *Text_italcontentpartContext) {}

// ExitText_italcontentpart is called when production text_italcontentpart is exited.
func (s *BaseCreole10Listener) ExitText_italcontentpart(ctx *Text_italcontentpartContext) {}

// EnterText_bolditalcontent is called when production text_bolditalcontent is entered.
func (s *BaseCreole10Listener) EnterText_bolditalcontent(ctx *Text_bolditalcontentContext) {}

// ExitText_bolditalcontent is called when production text_bolditalcontent is exited.
func (s *BaseCreole10Listener) ExitText_bolditalcontent(ctx *Text_bolditalcontentContext) {}

// EnterText_formattedcontent is called when production text_formattedcontent is entered.
func (s *BaseCreole10Listener) EnterText_formattedcontent(ctx *Text_formattedcontentContext) {}

// ExitText_formattedcontent is called when production text_formattedcontent is exited.
func (s *BaseCreole10Listener) ExitText_formattedcontent(ctx *Text_formattedcontentContext) {}

// EnterText_lineseparator1 is called when production text_lineseparator1 is entered.
func (s *BaseCreole10Listener) EnterText_lineseparator1(ctx *Text_lineseparator1Context) {}

// ExitText_lineseparator1 is called when production text_lineseparator1 is exited.
func (s *BaseCreole10Listener) ExitText_lineseparator1(ctx *Text_lineseparator1Context) {}

// EnterText_inlineelement is called when production text_inlineelement is entered.
func (s *BaseCreole10Listener) EnterText_inlineelement(ctx *Text_inlineelementContext) {}

// ExitText_inlineelement is called when production text_inlineelement is exited.
func (s *BaseCreole10Listener) ExitText_inlineelement(ctx *Text_inlineelementContext) {}

// EnterText_first_inlineelement is called when production text_first_inlineelement is entered.
func (s *BaseCreole10Listener) EnterText_first_inlineelement(ctx *Text_first_inlineelementContext) {}

// ExitText_first_inlineelement is called when production text_first_inlineelement is exited.
func (s *BaseCreole10Listener) ExitText_first_inlineelement(ctx *Text_first_inlineelementContext) {}

// EnterText_first_unformattedelement is called when production text_first_unformattedelement is entered.
func (s *BaseCreole10Listener) EnterText_first_unformattedelement(ctx *Text_first_unformattedelementContext) {
}

// ExitText_first_unformattedelement is called when production text_first_unformattedelement is exited.
func (s *BaseCreole10Listener) ExitText_first_unformattedelement(ctx *Text_first_unformattedelementContext) {
}

// EnterText_first_unformatted is called when production text_first_unformatted is entered.
func (s *BaseCreole10Listener) EnterText_first_unformatted(ctx *Text_first_unformattedContext) {}

// ExitText_first_unformatted is called when production text_first_unformatted is exited.
func (s *BaseCreole10Listener) ExitText_first_unformatted(ctx *Text_first_unformattedContext) {}

// EnterText_unformattedelement is called when production text_unformattedelement is entered.
func (s *BaseCreole10Listener) EnterText_unformattedelement(ctx *Text_unformattedelementContext) {}

// ExitText_unformattedelement is called when production text_unformattedelement is exited.
func (s *BaseCreole10Listener) ExitText_unformattedelement(ctx *Text_unformattedelementContext) {}

// EnterText_unformatted is called when production text_unformatted is entered.
func (s *BaseCreole10Listener) EnterText_unformatted(ctx *Text_unformattedContext) {}

// ExitText_unformatted is called when production text_unformatted is exited.
func (s *BaseCreole10Listener) ExitText_unformatted(ctx *Text_unformattedContext) {}

// EnterHeading is called when production heading is entered.
func (s *BaseCreole10Listener) EnterHeading(ctx *HeadingContext) {}

// ExitHeading is called when production heading is exited.
func (s *BaseCreole10Listener) ExitHeading(ctx *HeadingContext) {}

// EnterHeading_content is called when production heading_content is entered.
func (s *BaseCreole10Listener) EnterHeading_content(ctx *Heading_contentContext) {}

// ExitHeading_content is called when production heading_content is exited.
func (s *BaseCreole10Listener) ExitHeading_content(ctx *Heading_contentContext) {}

// EnterList_ord is called when production list_ord is entered.
func (s *BaseCreole10Listener) EnterList_ord(ctx *List_ordContext) {}

// ExitList_ord is called when production list_ord is exited.
func (s *BaseCreole10Listener) ExitList_ord(ctx *List_ordContext) {}

// EnterList_ordelem is called when production list_ordelem is entered.
func (s *BaseCreole10Listener) EnterList_ordelem(ctx *List_ordelemContext) {}

// ExitList_ordelem is called when production list_ordelem is exited.
func (s *BaseCreole10Listener) ExitList_ordelem(ctx *List_ordelemContext) {}

// EnterList_unord is called when production list_unord is entered.
func (s *BaseCreole10Listener) EnterList_unord(ctx *List_unordContext) {}

// ExitList_unord is called when production list_unord is exited.
func (s *BaseCreole10Listener) ExitList_unord(ctx *List_unordContext) {}

// EnterList_unordelem is called when production list_unordelem is entered.
func (s *BaseCreole10Listener) EnterList_unordelem(ctx *List_unordelemContext) {}

// ExitList_unordelem is called when production list_unordelem is exited.
func (s *BaseCreole10Listener) ExitList_unordelem(ctx *List_unordelemContext) {}

// EnterList_elem is called when production list_elem is entered.
func (s *BaseCreole10Listener) EnterList_elem(ctx *List_elemContext) {}

// ExitList_elem is called when production list_elem is exited.
func (s *BaseCreole10Listener) ExitList_elem(ctx *List_elemContext) {}

// EnterList_elem_markup is called when production list_elem_markup is entered.
func (s *BaseCreole10Listener) EnterList_elem_markup(ctx *List_elem_markupContext) {}

// ExitList_elem_markup is called when production list_elem_markup is exited.
func (s *BaseCreole10Listener) ExitList_elem_markup(ctx *List_elem_markupContext) {}

// EnterList_elemcontent is called when production list_elemcontent is entered.
func (s *BaseCreole10Listener) EnterList_elemcontent(ctx *List_elemcontentContext) {}

// ExitList_elemcontent is called when production list_elemcontent is exited.
func (s *BaseCreole10Listener) ExitList_elemcontent(ctx *List_elemcontentContext) {}

// EnterList_elemcontentpart is called when production list_elemcontentpart is entered.
func (s *BaseCreole10Listener) EnterList_elemcontentpart(ctx *List_elemcontentpartContext) {}

// ExitList_elemcontentpart is called when production list_elemcontentpart is exited.
func (s *BaseCreole10Listener) ExitList_elemcontentpart(ctx *List_elemcontentpartContext) {}

// EnterList_formatted_elem is called when production list_formatted_elem is entered.
func (s *BaseCreole10Listener) EnterList_formatted_elem(ctx *List_formatted_elemContext) {}

// ExitList_formatted_elem is called when production list_formatted_elem is exited.
func (s *BaseCreole10Listener) ExitList_formatted_elem(ctx *List_formatted_elemContext) {}

// EnterOnestar is called when production onestar is entered.
func (s *BaseCreole10Listener) EnterOnestar(ctx *OnestarContext) {}

// ExitOnestar is called when production onestar is exited.
func (s *BaseCreole10Listener) ExitOnestar(ctx *OnestarContext) {}

// EnterList_boldcontentpart is called when production list_boldcontentpart is entered.
func (s *BaseCreole10Listener) EnterList_boldcontentpart(ctx *List_boldcontentpartContext) {}

// ExitList_boldcontentpart is called when production list_boldcontentpart is exited.
func (s *BaseCreole10Listener) ExitList_boldcontentpart(ctx *List_boldcontentpartContext) {}

// EnterList_italcontentpart is called when production list_italcontentpart is entered.
func (s *BaseCreole10Listener) EnterList_italcontentpart(ctx *List_italcontentpartContext) {}

// ExitList_italcontentpart is called when production list_italcontentpart is exited.
func (s *BaseCreole10Listener) ExitList_italcontentpart(ctx *List_italcontentpartContext) {}

// EnterList_bolditalcontent is called when production list_bolditalcontent is entered.
func (s *BaseCreole10Listener) EnterList_bolditalcontent(ctx *List_bolditalcontentContext) {}

// ExitList_bolditalcontent is called when production list_bolditalcontent is exited.
func (s *BaseCreole10Listener) ExitList_bolditalcontent(ctx *List_bolditalcontentContext) {}

// EnterTable is called when production table is entered.
func (s *BaseCreole10Listener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BaseCreole10Listener) ExitTable(ctx *TableContext) {}

// EnterTable_row is called when production table_row is entered.
func (s *BaseCreole10Listener) EnterTable_row(ctx *Table_rowContext) {}

// ExitTable_row is called when production table_row is exited.
func (s *BaseCreole10Listener) ExitTable_row(ctx *Table_rowContext) {}

// EnterTable_cell is called when production table_cell is entered.
func (s *BaseCreole10Listener) EnterTable_cell(ctx *Table_cellContext) {}

// ExitTable_cell is called when production table_cell is exited.
func (s *BaseCreole10Listener) ExitTable_cell(ctx *Table_cellContext) {}

// EnterTable_headercell is called when production table_headercell is entered.
func (s *BaseCreole10Listener) EnterTable_headercell(ctx *Table_headercellContext) {}

// ExitTable_headercell is called when production table_headercell is exited.
func (s *BaseCreole10Listener) ExitTable_headercell(ctx *Table_headercellContext) {}

// EnterTable_normalcell is called when production table_normalcell is entered.
func (s *BaseCreole10Listener) EnterTable_normalcell(ctx *Table_normalcellContext) {}

// ExitTable_normalcell is called when production table_normalcell is exited.
func (s *BaseCreole10Listener) ExitTable_normalcell(ctx *Table_normalcellContext) {}

// EnterTable_cellcontent is called when production table_cellcontent is entered.
func (s *BaseCreole10Listener) EnterTable_cellcontent(ctx *Table_cellcontentContext) {}

// ExitTable_cellcontent is called when production table_cellcontent is exited.
func (s *BaseCreole10Listener) ExitTable_cellcontent(ctx *Table_cellcontentContext) {}

// EnterTable_cellcontentpart is called when production table_cellcontentpart is entered.
func (s *BaseCreole10Listener) EnterTable_cellcontentpart(ctx *Table_cellcontentpartContext) {}

// ExitTable_cellcontentpart is called when production table_cellcontentpart is exited.
func (s *BaseCreole10Listener) ExitTable_cellcontentpart(ctx *Table_cellcontentpartContext) {}

// EnterTable_formattedelement is called when production table_formattedelement is entered.
func (s *BaseCreole10Listener) EnterTable_formattedelement(ctx *Table_formattedelementContext) {}

// ExitTable_formattedelement is called when production table_formattedelement is exited.
func (s *BaseCreole10Listener) ExitTable_formattedelement(ctx *Table_formattedelementContext) {}

// EnterTable_boldcontent is called when production table_boldcontent is entered.
func (s *BaseCreole10Listener) EnterTable_boldcontent(ctx *Table_boldcontentContext) {}

// ExitTable_boldcontent is called when production table_boldcontent is exited.
func (s *BaseCreole10Listener) ExitTable_boldcontent(ctx *Table_boldcontentContext) {}

// EnterTable_italcontent is called when production table_italcontent is entered.
func (s *BaseCreole10Listener) EnterTable_italcontent(ctx *Table_italcontentContext) {}

// ExitTable_italcontent is called when production table_italcontent is exited.
func (s *BaseCreole10Listener) ExitTable_italcontent(ctx *Table_italcontentContext) {}

// EnterTable_boldcontentpart is called when production table_boldcontentpart is entered.
func (s *BaseCreole10Listener) EnterTable_boldcontentpart(ctx *Table_boldcontentpartContext) {}

// ExitTable_boldcontentpart is called when production table_boldcontentpart is exited.
func (s *BaseCreole10Listener) ExitTable_boldcontentpart(ctx *Table_boldcontentpartContext) {}

// EnterTable_italcontentpart is called when production table_italcontentpart is entered.
func (s *BaseCreole10Listener) EnterTable_italcontentpart(ctx *Table_italcontentpartContext) {}

// ExitTable_italcontentpart is called when production table_italcontentpart is exited.
func (s *BaseCreole10Listener) ExitTable_italcontentpart(ctx *Table_italcontentpartContext) {}

// EnterTable_bolditalcontent is called when production table_bolditalcontent is entered.
func (s *BaseCreole10Listener) EnterTable_bolditalcontent(ctx *Table_bolditalcontentContext) {}

// ExitTable_bolditalcontent is called when production table_bolditalcontent is exited.
func (s *BaseCreole10Listener) ExitTable_bolditalcontent(ctx *Table_bolditalcontentContext) {}

// EnterTable_formattedcontent is called when production table_formattedcontent is entered.
func (s *BaseCreole10Listener) EnterTable_formattedcontent(ctx *Table_formattedcontentContext) {}

// ExitTable_formattedcontent is called when production table_formattedcontent is exited.
func (s *BaseCreole10Listener) ExitTable_formattedcontent(ctx *Table_formattedcontentContext) {}

// EnterTable_inlineelement is called when production table_inlineelement is entered.
func (s *BaseCreole10Listener) EnterTable_inlineelement(ctx *Table_inlineelementContext) {}

// ExitTable_inlineelement is called when production table_inlineelement is exited.
func (s *BaseCreole10Listener) ExitTable_inlineelement(ctx *Table_inlineelementContext) {}

// EnterTable_unformattedelement is called when production table_unformattedelement is entered.
func (s *BaseCreole10Listener) EnterTable_unformattedelement(ctx *Table_unformattedelementContext) {}

// ExitTable_unformattedelement is called when production table_unformattedelement is exited.
func (s *BaseCreole10Listener) ExitTable_unformattedelement(ctx *Table_unformattedelementContext) {}

// EnterTable_unformatted is called when production table_unformatted is entered.
func (s *BaseCreole10Listener) EnterTable_unformatted(ctx *Table_unformattedContext) {}

// ExitTable_unformatted is called when production table_unformatted is exited.
func (s *BaseCreole10Listener) ExitTable_unformatted(ctx *Table_unformattedContext) {}

// EnterNowiki_block is called when production nowiki_block is entered.
func (s *BaseCreole10Listener) EnterNowiki_block(ctx *Nowiki_blockContext) {}

// ExitNowiki_block is called when production nowiki_block is exited.
func (s *BaseCreole10Listener) ExitNowiki_block(ctx *Nowiki_blockContext) {}

// EnterNowiki_block_content is called when production nowiki_block_content is entered.
func (s *BaseCreole10Listener) EnterNowiki_block_content(ctx *Nowiki_block_contentContext) {}

// ExitNowiki_block_content is called when production nowiki_block_content is exited.
func (s *BaseCreole10Listener) ExitNowiki_block_content(ctx *Nowiki_block_contentContext) {}

// EnterNowikiblock_open_markup is called when production nowikiblock_open_markup is entered.
func (s *BaseCreole10Listener) EnterNowikiblock_open_markup(ctx *Nowikiblock_open_markupContext) {}

// ExitNowikiblock_open_markup is called when production nowikiblock_open_markup is exited.
func (s *BaseCreole10Listener) ExitNowikiblock_open_markup(ctx *Nowikiblock_open_markupContext) {}

// EnterNowikiblock_close_markup is called when production nowikiblock_close_markup is entered.
func (s *BaseCreole10Listener) EnterNowikiblock_close_markup(ctx *Nowikiblock_close_markupContext) {}

// ExitNowikiblock_close_markup is called when production nowikiblock_close_markup is exited.
func (s *BaseCreole10Listener) ExitNowikiblock_close_markup(ctx *Nowikiblock_close_markupContext) {}

// EnterNowiki_inline is called when production nowiki_inline is entered.
func (s *BaseCreole10Listener) EnterNowiki_inline(ctx *Nowiki_inlineContext) {}

// ExitNowiki_inline is called when production nowiki_inline is exited.
func (s *BaseCreole10Listener) ExitNowiki_inline(ctx *Nowiki_inlineContext) {}

// EnterNowiki_inline_content is called when production nowiki_inline_content is entered.
func (s *BaseCreole10Listener) EnterNowiki_inline_content(ctx *Nowiki_inline_contentContext) {}

// ExitNowiki_inline_content is called when production nowiki_inline_content is exited.
func (s *BaseCreole10Listener) ExitNowiki_inline_content(ctx *Nowiki_inline_contentContext) {}

// EnterHorizontalrule is called when production horizontalrule is entered.
func (s *BaseCreole10Listener) EnterHorizontalrule(ctx *HorizontalruleContext) {}

// ExitHorizontalrule is called when production horizontalrule is exited.
func (s *BaseCreole10Listener) ExitHorizontalrule(ctx *HorizontalruleContext) {}

// EnterLink is called when production link is entered.
func (s *BaseCreole10Listener) EnterLink(ctx *LinkContext) {}

// ExitLink is called when production link is exited.
func (s *BaseCreole10Listener) ExitLink(ctx *LinkContext) {}

// EnterLink_address is called when production link_address is entered.
func (s *BaseCreole10Listener) EnterLink_address(ctx *Link_addressContext) {}

// ExitLink_address is called when production link_address is exited.
func (s *BaseCreole10Listener) ExitLink_address(ctx *Link_addressContext) {}

// EnterLink_interwiki_uri is called when production link_interwiki_uri is entered.
func (s *BaseCreole10Listener) EnterLink_interwiki_uri(ctx *Link_interwiki_uriContext) {}

// ExitLink_interwiki_uri is called when production link_interwiki_uri is exited.
func (s *BaseCreole10Listener) ExitLink_interwiki_uri(ctx *Link_interwiki_uriContext) {}

// EnterLink_interwiki_pagename is called when production link_interwiki_pagename is entered.
func (s *BaseCreole10Listener) EnterLink_interwiki_pagename(ctx *Link_interwiki_pagenameContext) {}

// ExitLink_interwiki_pagename is called when production link_interwiki_pagename is exited.
func (s *BaseCreole10Listener) ExitLink_interwiki_pagename(ctx *Link_interwiki_pagenameContext) {}

// EnterLink_description is called when production link_description is entered.
func (s *BaseCreole10Listener) EnterLink_description(ctx *Link_descriptionContext) {}

// ExitLink_description is called when production link_description is exited.
func (s *BaseCreole10Listener) ExitLink_description(ctx *Link_descriptionContext) {}

// EnterLink_descriptionpart is called when production link_descriptionpart is entered.
func (s *BaseCreole10Listener) EnterLink_descriptionpart(ctx *Link_descriptionpartContext) {}

// ExitLink_descriptionpart is called when production link_descriptionpart is exited.
func (s *BaseCreole10Listener) ExitLink_descriptionpart(ctx *Link_descriptionpartContext) {}

// EnterLink_bold_descriptionpart is called when production link_bold_descriptionpart is entered.
func (s *BaseCreole10Listener) EnterLink_bold_descriptionpart(ctx *Link_bold_descriptionpartContext) {}

// ExitLink_bold_descriptionpart is called when production link_bold_descriptionpart is exited.
func (s *BaseCreole10Listener) ExitLink_bold_descriptionpart(ctx *Link_bold_descriptionpartContext) {}

// EnterLink_ital_descriptionpart is called when production link_ital_descriptionpart is entered.
func (s *BaseCreole10Listener) EnterLink_ital_descriptionpart(ctx *Link_ital_descriptionpartContext) {}

// ExitLink_ital_descriptionpart is called when production link_ital_descriptionpart is exited.
func (s *BaseCreole10Listener) ExitLink_ital_descriptionpart(ctx *Link_ital_descriptionpartContext) {}

// EnterLink_boldital_description is called when production link_boldital_description is entered.
func (s *BaseCreole10Listener) EnterLink_boldital_description(ctx *Link_boldital_descriptionContext) {}

// ExitLink_boldital_description is called when production link_boldital_description is exited.
func (s *BaseCreole10Listener) ExitLink_boldital_description(ctx *Link_boldital_descriptionContext) {}

// EnterLink_descriptiontext is called when production link_descriptiontext is entered.
func (s *BaseCreole10Listener) EnterLink_descriptiontext(ctx *Link_descriptiontextContext) {}

// ExitLink_descriptiontext is called when production link_descriptiontext is exited.
func (s *BaseCreole10Listener) ExitLink_descriptiontext(ctx *Link_descriptiontextContext) {}

// EnterLink_uri is called when production link_uri is entered.
func (s *BaseCreole10Listener) EnterLink_uri(ctx *Link_uriContext) {}

// ExitLink_uri is called when production link_uri is exited.
func (s *BaseCreole10Listener) ExitLink_uri(ctx *Link_uriContext) {}

// EnterImage is called when production image is entered.
func (s *BaseCreole10Listener) EnterImage(ctx *ImageContext) {}

// ExitImage is called when production image is exited.
func (s *BaseCreole10Listener) ExitImage(ctx *ImageContext) {}

// EnterImage_uri is called when production image_uri is entered.
func (s *BaseCreole10Listener) EnterImage_uri(ctx *Image_uriContext) {}

// ExitImage_uri is called when production image_uri is exited.
func (s *BaseCreole10Listener) ExitImage_uri(ctx *Image_uriContext) {}

// EnterImage_alternative is called when production image_alternative is entered.
func (s *BaseCreole10Listener) EnterImage_alternative(ctx *Image_alternativeContext) {}

// ExitImage_alternative is called when production image_alternative is exited.
func (s *BaseCreole10Listener) ExitImage_alternative(ctx *Image_alternativeContext) {}

// EnterImage_alternativepart is called when production image_alternativepart is entered.
func (s *BaseCreole10Listener) EnterImage_alternativepart(ctx *Image_alternativepartContext) {}

// ExitImage_alternativepart is called when production image_alternativepart is exited.
func (s *BaseCreole10Listener) ExitImage_alternativepart(ctx *Image_alternativepartContext) {}

// EnterImage_bold_alternativepart is called when production image_bold_alternativepart is entered.
func (s *BaseCreole10Listener) EnterImage_bold_alternativepart(ctx *Image_bold_alternativepartContext) {
}

// ExitImage_bold_alternativepart is called when production image_bold_alternativepart is exited.
func (s *BaseCreole10Listener) ExitImage_bold_alternativepart(ctx *Image_bold_alternativepartContext) {
}

// EnterImage_ital_alternativepart is called when production image_ital_alternativepart is entered.
func (s *BaseCreole10Listener) EnterImage_ital_alternativepart(ctx *Image_ital_alternativepartContext) {
}

// ExitImage_ital_alternativepart is called when production image_ital_alternativepart is exited.
func (s *BaseCreole10Listener) ExitImage_ital_alternativepart(ctx *Image_ital_alternativepartContext) {
}

// EnterImage_boldital_alternative is called when production image_boldital_alternative is entered.
func (s *BaseCreole10Listener) EnterImage_boldital_alternative(ctx *Image_boldital_alternativeContext) {
}

// ExitImage_boldital_alternative is called when production image_boldital_alternative is exited.
func (s *BaseCreole10Listener) ExitImage_boldital_alternative(ctx *Image_boldital_alternativeContext) {
}

// EnterImage_alternativetext is called when production image_alternativetext is entered.
func (s *BaseCreole10Listener) EnterImage_alternativetext(ctx *Image_alternativetextContext) {}

// ExitImage_alternativetext is called when production image_alternativetext is exited.
func (s *BaseCreole10Listener) ExitImage_alternativetext(ctx *Image_alternativetextContext) {}

// EnterExtension is called when production extension is entered.
func (s *BaseCreole10Listener) EnterExtension(ctx *ExtensionContext) {}

// ExitExtension is called when production extension is exited.
func (s *BaseCreole10Listener) ExitExtension(ctx *ExtensionContext) {}

// EnterExtension_handler is called when production extension_handler is entered.
func (s *BaseCreole10Listener) EnterExtension_handler(ctx *Extension_handlerContext) {}

// ExitExtension_handler is called when production extension_handler is exited.
func (s *BaseCreole10Listener) ExitExtension_handler(ctx *Extension_handlerContext) {}

// EnterExtension_statement is called when production extension_statement is entered.
func (s *BaseCreole10Listener) EnterExtension_statement(ctx *Extension_statementContext) {}

// ExitExtension_statement is called when production extension_statement is exited.
func (s *BaseCreole10Listener) ExitExtension_statement(ctx *Extension_statementContext) {}

// EnterEscaped is called when production escaped is entered.
func (s *BaseCreole10Listener) EnterEscaped(ctx *EscapedContext) {}

// ExitEscaped is called when production escaped is exited.
func (s *BaseCreole10Listener) ExitEscaped(ctx *EscapedContext) {}

// EnterParagraph_separator is called when production paragraph_separator is entered.
func (s *BaseCreole10Listener) EnterParagraph_separator(ctx *Paragraph_separatorContext) {}

// ExitParagraph_separator is called when production paragraph_separator is exited.
func (s *BaseCreole10Listener) ExitParagraph_separator(ctx *Paragraph_separatorContext) {}

// EnterWhitespaces is called when production whitespaces is entered.
func (s *BaseCreole10Listener) EnterWhitespaces(ctx *WhitespacesContext) {}

// ExitWhitespaces is called when production whitespaces is exited.
func (s *BaseCreole10Listener) ExitWhitespaces(ctx *WhitespacesContext) {}

// EnterBlanks is called when production blanks is entered.
func (s *BaseCreole10Listener) EnterBlanks(ctx *BlanksContext) {}

// ExitBlanks is called when production blanks is exited.
func (s *BaseCreole10Listener) ExitBlanks(ctx *BlanksContext) {}

// EnterText_lineseparator is called when production text_lineseparator is entered.
func (s *BaseCreole10Listener) EnterText_lineseparator(ctx *Text_lineseparatorContext) {}

// ExitText_lineseparator is called when production text_lineseparator is exited.
func (s *BaseCreole10Listener) ExitText_lineseparator(ctx *Text_lineseparatorContext) {}

// EnterNewline is called when production newline is entered.
func (s *BaseCreole10Listener) EnterNewline(ctx *NewlineContext) {}

// ExitNewline is called when production newline is exited.
func (s *BaseCreole10Listener) ExitNewline(ctx *NewlineContext) {}

// EnterBold_markup is called when production bold_markup is entered.
func (s *BaseCreole10Listener) EnterBold_markup(ctx *Bold_markupContext) {}

// ExitBold_markup is called when production bold_markup is exited.
func (s *BaseCreole10Listener) ExitBold_markup(ctx *Bold_markupContext) {}

// EnterItal_markup is called when production ital_markup is entered.
func (s *BaseCreole10Listener) EnterItal_markup(ctx *Ital_markupContext) {}

// ExitItal_markup is called when production ital_markup is exited.
func (s *BaseCreole10Listener) ExitItal_markup(ctx *Ital_markupContext) {}

// EnterHeading_markup is called when production heading_markup is entered.
func (s *BaseCreole10Listener) EnterHeading_markup(ctx *Heading_markupContext) {}

// ExitHeading_markup is called when production heading_markup is exited.
func (s *BaseCreole10Listener) ExitHeading_markup(ctx *Heading_markupContext) {}

// EnterList_ordelem_markup is called when production list_ordelem_markup is entered.
func (s *BaseCreole10Listener) EnterList_ordelem_markup(ctx *List_ordelem_markupContext) {}

// ExitList_ordelem_markup is called when production list_ordelem_markup is exited.
func (s *BaseCreole10Listener) ExitList_ordelem_markup(ctx *List_ordelem_markupContext) {}

// EnterList_unordelem_markup is called when production list_unordelem_markup is entered.
func (s *BaseCreole10Listener) EnterList_unordelem_markup(ctx *List_unordelem_markupContext) {}

// ExitList_unordelem_markup is called when production list_unordelem_markup is exited.
func (s *BaseCreole10Listener) ExitList_unordelem_markup(ctx *List_unordelem_markupContext) {}

// EnterList_elemseparator is called when production list_elemseparator is entered.
func (s *BaseCreole10Listener) EnterList_elemseparator(ctx *List_elemseparatorContext) {}

// ExitList_elemseparator is called when production list_elemseparator is exited.
func (s *BaseCreole10Listener) ExitList_elemseparator(ctx *List_elemseparatorContext) {}

// EnterEnd_of_list is called when production end_of_list is entered.
func (s *BaseCreole10Listener) EnterEnd_of_list(ctx *End_of_listContext) {}

// ExitEnd_of_list is called when production end_of_list is exited.
func (s *BaseCreole10Listener) ExitEnd_of_list(ctx *End_of_listContext) {}

// EnterTable_cell_markup is called when production table_cell_markup is entered.
func (s *BaseCreole10Listener) EnterTable_cell_markup(ctx *Table_cell_markupContext) {}

// ExitTable_cell_markup is called when production table_cell_markup is exited.
func (s *BaseCreole10Listener) ExitTable_cell_markup(ctx *Table_cell_markupContext) {}

// EnterTable_headercell_markup is called when production table_headercell_markup is entered.
func (s *BaseCreole10Listener) EnterTable_headercell_markup(ctx *Table_headercell_markupContext) {}

// ExitTable_headercell_markup is called when production table_headercell_markup is exited.
func (s *BaseCreole10Listener) ExitTable_headercell_markup(ctx *Table_headercell_markupContext) {}

// EnterTable_rowseparator is called when production table_rowseparator is entered.
func (s *BaseCreole10Listener) EnterTable_rowseparator(ctx *Table_rowseparatorContext) {}

// ExitTable_rowseparator is called when production table_rowseparator is exited.
func (s *BaseCreole10Listener) ExitTable_rowseparator(ctx *Table_rowseparatorContext) {}

// EnterNowiki_open_markup is called when production nowiki_open_markup is entered.
func (s *BaseCreole10Listener) EnterNowiki_open_markup(ctx *Nowiki_open_markupContext) {}

// ExitNowiki_open_markup is called when production nowiki_open_markup is exited.
func (s *BaseCreole10Listener) ExitNowiki_open_markup(ctx *Nowiki_open_markupContext) {}

// EnterNowiki_close_markup is called when production nowiki_close_markup is entered.
func (s *BaseCreole10Listener) EnterNowiki_close_markup(ctx *Nowiki_close_markupContext) {}

// ExitNowiki_close_markup is called when production nowiki_close_markup is exited.
func (s *BaseCreole10Listener) ExitNowiki_close_markup(ctx *Nowiki_close_markupContext) {}

// EnterHorizontalrule_markup is called when production horizontalrule_markup is entered.
func (s *BaseCreole10Listener) EnterHorizontalrule_markup(ctx *Horizontalrule_markupContext) {}

// ExitHorizontalrule_markup is called when production horizontalrule_markup is exited.
func (s *BaseCreole10Listener) ExitHorizontalrule_markup(ctx *Horizontalrule_markupContext) {}

// EnterLink_open_markup is called when production link_open_markup is entered.
func (s *BaseCreole10Listener) EnterLink_open_markup(ctx *Link_open_markupContext) {}

// ExitLink_open_markup is called when production link_open_markup is exited.
func (s *BaseCreole10Listener) ExitLink_open_markup(ctx *Link_open_markupContext) {}

// EnterLink_close_markup is called when production link_close_markup is entered.
func (s *BaseCreole10Listener) EnterLink_close_markup(ctx *Link_close_markupContext) {}

// ExitLink_close_markup is called when production link_close_markup is exited.
func (s *BaseCreole10Listener) ExitLink_close_markup(ctx *Link_close_markupContext) {}

// EnterLink_description_markup is called when production link_description_markup is entered.
func (s *BaseCreole10Listener) EnterLink_description_markup(ctx *Link_description_markupContext) {}

// ExitLink_description_markup is called when production link_description_markup is exited.
func (s *BaseCreole10Listener) ExitLink_description_markup(ctx *Link_description_markupContext) {}

// EnterImage_open_markup is called when production image_open_markup is entered.
func (s *BaseCreole10Listener) EnterImage_open_markup(ctx *Image_open_markupContext) {}

// ExitImage_open_markup is called when production image_open_markup is exited.
func (s *BaseCreole10Listener) ExitImage_open_markup(ctx *Image_open_markupContext) {}

// EnterImage_close_markup is called when production image_close_markup is entered.
func (s *BaseCreole10Listener) EnterImage_close_markup(ctx *Image_close_markupContext) {}

// ExitImage_close_markup is called when production image_close_markup is exited.
func (s *BaseCreole10Listener) ExitImage_close_markup(ctx *Image_close_markupContext) {}

// EnterImage_alternative_markup is called when production image_alternative_markup is entered.
func (s *BaseCreole10Listener) EnterImage_alternative_markup(ctx *Image_alternative_markupContext) {}

// ExitImage_alternative_markup is called when production image_alternative_markup is exited.
func (s *BaseCreole10Listener) ExitImage_alternative_markup(ctx *Image_alternative_markupContext) {}

// EnterExtension_markup is called when production extension_markup is entered.
func (s *BaseCreole10Listener) EnterExtension_markup(ctx *Extension_markupContext) {}

// ExitExtension_markup is called when production extension_markup is exited.
func (s *BaseCreole10Listener) ExitExtension_markup(ctx *Extension_markupContext) {}

// EnterForced_linebreak is called when production forced_linebreak is entered.
func (s *BaseCreole10Listener) EnterForced_linebreak(ctx *Forced_linebreakContext) {}

// ExitForced_linebreak is called when production forced_linebreak is exited.
func (s *BaseCreole10Listener) ExitForced_linebreak(ctx *Forced_linebreakContext) {}
