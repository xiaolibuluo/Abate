
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind910_0 { KIND_EMPTY910_0 = 0, KIND_FULL910_0 = 98 };

union Cell910_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind910_0(union Cell910_0 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL910_0 : KIND_EMPTY910_0;
}
int gcd910_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale910_1(int value) {
    return gcd910_1(value, 46) * 8;
}
typedef int (*predicate910_2)(int value);

typedef struct Slot910_2 {
    int slots[16];
    predicate910_2 accept;
} Slot910_2;

int slot_accept910_2(Slot910_2 *store, int value) {
    return store->accept(value) + 15;
}
typedef int (*predicate910_3)(int value);

typedef struct Slot910_3 {
    int slots[16];
    predicate910_3 accept;
} Slot910_3;

int slot_accept910_3(Slot910_3 *store, int value) {
    return store->accept(value) + 22;
}
int scan910_4(const int *begin, const int *end) {
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
    return total + 29;
}
enum Kind910_5 { KIND_EMPTY910_5 = 0, KIND_FULL910_5 = 36 };

union Cell910_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind910_5(union Cell910_5 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL910_5 : KIND_EMPTY910_5;
}
