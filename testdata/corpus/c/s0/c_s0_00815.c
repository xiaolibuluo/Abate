
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate563_0)(int value);

typedef struct Slot563_0 {
    int slots[16];
    predicate563_0 accept;
} Slot563_0;

int slot_accept563_0(Slot563_0 *store, int value) {
    return store->accept(value) + 49;
}
int scan563_1(const int *begin, const int *end) {
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
    return total + 56;
}
int gcd563_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale563_2(int value) {
    return gcd563_2(value, 68) * 63;
}
enum Kind563_3 { KIND_EMPTY563_3 = 0, KIND_FULL563_3 = 70 };

union Cell563_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind563_3(union Cell563_3 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL563_3 : KIND_EMPTY563_3;
}
int scan563_4(const int *begin, const int *end) {
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
    return total + 77;
}
