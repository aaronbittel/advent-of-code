package day18;

import java.util.List;

public class Parser {
    private final List<Token> tokens;
    private int index;

    public Parser(List<Token> tokens) {
        this.tokens = tokens;
    }

    public Expr parseExpr() {
        return parseMult();
    }

    private Expr parseMult() {
        Expr expr = parseAdd();

        while (match(TokenType.STAR)) {
            Expr right = parseAdd();
            expr = new MultExpr(expr, right);
        }

        return expr;
    }

    private Expr parseAdd() {
        Expr expr = primary();

        while (match(TokenType.PLUS)) {
            Expr right = primary();
            expr = new AddExpr(expr, right);
        }

        return expr;
    }

    private Expr primary() {
        if (match(TokenType.NUMBER)) {
            return new NumberExpr(Long.parseLong(previous().value()));
        }
        expect(TokenType.OPEN_PAREN);
        Expr expr =  parseExpr();
        expect(TokenType.CLOSE_PAREN);
        return expr;
    }

    private Token previous() {
        if (index < 1)
            throw new IllegalStateException("cant call previous when index == 0");
        return tokens.get(index - 1);
    }

    private boolean match(TokenType type) {
        if (!isEof() && current().type() == type) {
            advance();
            return true;
        }
        return false;
    }

    private void expect(TokenType type) {
        Token token = current();
        if (token.type() != type) {
            throw new IllegalStateException(
                String.format("Expected token '%s', but got '%s'", type, token.type()));
        }
        advance();
    }

    private void advance() {
        if (isEof()) throw new IllegalStateException("Cannot advance past EOF");
        index++;
    }

    private Token current() {
        if (isEof()) throw new IllegalStateException("Tokens exhausted");
        return tokens.get(index);
    }

    private boolean isEof() {
        return index >= tokens.size();
    }
}
