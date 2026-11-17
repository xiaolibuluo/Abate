
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2189_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2189_0(int value) {
    return gcd2189_0(value, 21) * 41;
}
int scan2189_1(const int *begin, const int *end) {
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
    return total + 48;
}
typedef int (*predicate2189_2)(int value);

typedef struct Slot2189_2 {
    int slots[16];
    predicate2189_2 accept;
} Slot2189_2;

int slot_accept2189_2(Slot2189_2 *store, int value) {
    return store->accept(value) + 55;
}
enum Kind2189_3 { KIND_EMPTY2189_3 = 0, KIND_FULL2189_3 = 62 };

union Cell2189_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2189_3(union Cell2189_3 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL2189_3 : KIND_EMPTY2189_3;
}
typedef int (*predicate2189_4)(int value);

typedef struct Slot2189_4 {
    int slots[16];
    predicate2189_4 accept;
} Slot2189_4;

int slot_accept2189_4(Slot2189_4 *store, int value) {
    return store->accept(value) + 69;
}
int scan2189_5(const int *begin, const int *end) {
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
    return total + 76;
}
int scan2189_6(const int *begin, const int *end) {
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
    return total + 83;
}
int scan2189_7(const int *begin, const int *end) {
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
    return total + 90;
}
