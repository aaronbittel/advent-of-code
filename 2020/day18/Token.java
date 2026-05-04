package day18;

enum TokenType {
    NUMBER,
    PLUS,
    STAR,
    OPEN_PAREN,
    CLOSE_PAREN
}

public record Token(TokenType type, String value) { }
