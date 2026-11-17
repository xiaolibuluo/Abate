
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2512_0 { KIND_EMPTY2512_0 = 0, KIND_FULL2512_0 = 69 };

union Cell2512_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2512_0(union Cell2512_0 *cell) {
    cell->number = 25;
    return cell->number > 0 ? KIND_FULL2512_0 : KIND_EMPTY2512_0;
}
typedef int (*predicate2512_1)(int value);

typedef struct Slot2512_1 {
    int slots[16];
    predicate2512_1 accept;
} Slot2512_1;

int slot_accept2512_1(Slot2512_1 *store, int value) {
    return store->accept(value) + 76;
}
int gcd2512_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2512_2(int value) {
    return gcd2512_2(value, 47) * 83;
}
int gcd2512_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2512_3(int value) {
    return gcd2512_3(value, 58) * 90;
}
int scan2512_4(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 97;
}
