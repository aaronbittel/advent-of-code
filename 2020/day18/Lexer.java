package day18;

import java.util.ArrayList;
import java.util.List;

public class Lexer {
    private final String source;
    private int index;

    public Lexer(String source) {
        this.source = source;
    }

    public List<Token> tokenize() {
        List<Token> tokens = new ArrayList<>();
        while(!isEof()) {
            skipSpaces();

            char c = current();
            Token token = switch (c) {
                case '+' -> new Token(TokenType.PLUS, String.valueOf(advance()));
                case '*' -> new Token(TokenType.STAR, String.valueOf(advance()));
                case '(' -> new Token(TokenType.OPEN_PAREN, String.valueOf(advance()));
                case ')' -> new Token(TokenType.CLOSE_PAREN, String.valueOf(advance()));
                default -> {
                    if (!Character.isDigit(c)) {
                        throw new IllegalArgumentException("Illegal char: " + c);
                    }
                    yield parseNumber();
                }
            };
            tokens.add(token);
        }
        return tokens;
    }

    private Token parseNumber() {
        int start = index;
        while (Character.isDigit(current())) {
            advance();
        }
        int end = index;
        return new Token(TokenType.NUMBER, source.substring(start, end));
    }

    private char current() {
        if (isEof()) return 0;
        return source.charAt(index);
    }

    private char advance() {
        if (isEof()) return 0;
        char c = current();
        index++;
        return c;
    }

    private void skipSpaces() {
        while(current() == ' ') {
            advance();
        }
    }

    private boolean isEof() {
        return index >= source.length();
    }
}
