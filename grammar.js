/**
 * @file GMPL bindings to the Tree-sitter parsing library
 * @author Florian P. Walter <florian.p.walter@gmail.com>
 * @license MIT
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

export default grammar({
  name: "gmpl",

  rules: {
    // TODO: add the actual grammar rules
    source_file: $ => "hello"
  }
});
