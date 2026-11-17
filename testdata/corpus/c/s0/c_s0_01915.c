
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1289_0)(int value);

typedef struct Slot1289_0 {
    int slots[16];
    predicate1289_0 accept;
} Slot1289_0;

int slot_accept1289_0(Slot1289_0 *store, int value) {
    return store->accept(value) + 78;
}
int scan1289_1(const int *begin, const int *end) {
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
    return total + 85;
}
enum Kind1289_2 { KIND_EMPTY1289_2 = 0, KIND_FULL1289_2 = 92 };

union Cell1289_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1289_2(union Cell1289_2 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL1289_2 : KIND_EMPTY1289_2;
}
int gcd1289_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1289_3(int value) {
    return gcd1289_3(value, 37) * 99;
}
