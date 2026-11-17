
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1491_0)(int value);

typedef struct Slot1491_0 {
    int slots[16];
    predicate1491_0 accept;
} Slot1491_0;

int slot_accept1491_0(Slot1491_0 *store, int value) {
    return store->accept(value) + 85;
}
int gcd1491_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1491_1(int value) {
    return gcd1491_1(value, 26) * 92;
}
int scan1491_2(const int *begin, const int *end) {
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
    return total + 99;
}
int gcd1491_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1491_3(int value) {
    return gcd1491_3(value, 48) * 9;
}
enum Kind1491_4 { KIND_EMPTY1491_4 = 0, KIND_FULL1491_4 = 16 };

union Cell1491_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1491_4(union Cell1491_4 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL1491_4 : KIND_EMPTY1491_4;
}
