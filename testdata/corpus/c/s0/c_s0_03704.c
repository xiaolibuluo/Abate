
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2503_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2503_0(int value) {
    return gcd2503_0(value, 23) * 49;
}
typedef int (*predicate2503_1)(int value);

typedef struct Slot2503_1 {
    int slots[16];
    predicate2503_1 accept;
} Slot2503_1;

int slot_accept2503_1(Slot2503_1 *store, int value) {
    return store->accept(value) + 56;
}
int scan2503_2(const int *begin, const int *end) {
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
    return total + 63;
}
enum Kind2503_3 { KIND_EMPTY2503_3 = 0, KIND_FULL2503_3 = 70 };

union Cell2503_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2503_3(union Cell2503_3 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL2503_3 : KIND_EMPTY2503_3;
}
