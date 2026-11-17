
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind959_0 { KIND_EMPTY959_0 = 0, KIND_FULL959_0 = 56 };

union Cell959_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind959_0(union Cell959_0 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL959_0 : KIND_EMPTY959_0;
}
int scan959_1(const int *begin, const int *end) {
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
#if defined(__STDC_VERSION__)
int variant959_2(int value, ...) {
    return value + 70;
}
#else
int variant959_2(int value) {
    return value + 70;
}
#endif
typedef int (*predicate959_3)(int value);

typedef struct Slot959_3 {
    int slots[16];
    predicate959_3 accept;
} Slot959_3;

int slot_accept959_3(Slot959_3 *store, int value) {
    return store->accept(value) + 77;
}
int gcd959_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale959_4(int value) {
    return gcd959_4(value, 56) * 84;
}
enum Kind959_5 { KIND_EMPTY959_5 = 0, KIND_FULL959_5 = 91 };

union Cell959_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind959_5(union Cell959_5 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL959_5 : KIND_EMPTY959_5;
}
