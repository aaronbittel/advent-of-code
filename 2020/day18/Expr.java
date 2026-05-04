package day18;

public sealed interface Expr permits NumberExpr, AddExpr, MultExpr {
    long eval();
}

record NumberExpr(long value) implements Expr {
    @Override
    public long eval() {
        return value;
    }

    @Override
    public String toString() {
        return String.valueOf(value);
    }
}

record AddExpr(Expr left, Expr right) implements Expr {
    @Override
    public long eval() {
        return left.eval() + right.eval();
    }

    @Override
    public String toString() {
        return String.format("%s + %s", left, right);
    }
}

record MultExpr(Expr left, Expr right) implements Expr {
    @Override
    public long eval() {
        return left.eval() * right.eval();
    }

    @Override
    public String toString() {
        return String.format("%s * %s", left, right);
    }
}

