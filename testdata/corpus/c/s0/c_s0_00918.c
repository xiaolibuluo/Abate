
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate631_0)(int value);

typedef struct Slot631_0 {
    int slots[16];
    predicate631_0 accept;
} Slot631_0;

int slot_accept631_0(Slot631_0 *store, int value) {
    return store->accept(value) + 60;
}
typedef int (*predicate631_1)(int value);

typedef struct Slot631_1 {
    int slots[16];
    predicate631_1 accept;
} Slot631_1;

int slot_accept631_1(Slot631_1 *store, int value) {
    return store->accept(value) + 67;
}
enum Kind631_2 { KIND_EMPTY631_2 = 0, KIND_FULL631_2 = 74 };

union Cell631_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind631_2(union Cell631_2 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL631_2 : KIND_EMPTY631_2;
}
int gcd631_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale631_3(int value) {
    return gcd631_3(value, 67) * 81;
}
int scan631_4(const int *begin, const int *end) {
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
    return total + 88;
}
int scan631_5(const int *begin, const int *end) {
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
    return total + 95;
}
#define MASK631_6(v) (((v) << 4) | ((v) & 0x0F))

int classify631_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK631_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan631_7(const int *begin, const int *end) {
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
    return total + 12;
}
