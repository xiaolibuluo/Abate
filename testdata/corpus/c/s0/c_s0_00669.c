
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate460_0)(int value);

typedef struct Slot460_0 {
    int slots[16];
    predicate460_0 accept;
} Slot460_0;

int slot_accept460_0(Slot460_0 *store, int value) {
    return store->accept(value) + 68;
}
int scan460_1(const int *begin, const int *end) {
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
    return total + 75;
}
int gcd460_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale460_2(int value) {
    return gcd460_2(value, 18) * 82;
}
enum Kind460_3 { KIND_EMPTY460_3 = 0, KIND_FULL460_3 = 89 };

union Cell460_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind460_3(union Cell460_3 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL460_3 : KIND_EMPTY460_3;
}
