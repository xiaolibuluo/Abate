
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1581_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1581_0(int value) {
    return gcd1581_0(value, 35) * 91;
}
typedef int (*predicate1581_1)(int value);

typedef struct Slot1581_1 {
    int slots[16];
    predicate1581_1 accept;
} Slot1581_1;

int slot_accept1581_1(Slot1581_1 *store, int value) {
    return store->accept(value) + 98;
}
enum Kind1581_2 { KIND_EMPTY1581_2 = 0, KIND_FULL1581_2 = 8 };

union Cell1581_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1581_2(union Cell1581_2 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL1581_2 : KIND_EMPTY1581_2;
}
int scan1581_3(const int *begin, const int *end) {
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
    return total + 15;
}
