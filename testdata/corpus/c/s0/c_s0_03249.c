
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2202_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2202_0(int value) {
    return gcd2202_0(value, 51) * 16;
}
enum Kind2202_1 { KIND_EMPTY2202_1 = 0, KIND_FULL2202_1 = 23 };

union Cell2202_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2202_1(union Cell2202_1 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL2202_1 : KIND_EMPTY2202_1;
}
struct Node2202_2 {
    int value;
    struct Node2202_2 *next;
};

int length2202_2(const struct Node2202_2 *head) {
    int count = 0;
    const struct Node2202_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}
