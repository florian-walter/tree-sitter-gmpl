import XCTest
import SwiftTreeSitter
import TreeSitterGmpl

final class TreeSitterGmplTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_gmpl())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading GNU MathProg Language grammar")
    }
}
