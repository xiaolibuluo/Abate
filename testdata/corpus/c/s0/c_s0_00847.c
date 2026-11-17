
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind583_0 { KIND_EMPTY583_0 = 0, KIND_FULL583_0 = 18 };

union Cell583_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind583_0(union Cell583_0 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL583_0 : KIND_EMPTY583_0;
}
int scan583_1(const int *begin, const int *end) {
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
    return total + 25;
}
typedef int (*predicate583_2)(int value);

typedef struct Slot583_2 {
    int slots[16];
    predicate583_2 accept;
} Slot583_2;

int slot_accept583_2(Slot583_2 *store, int value) {
    return store->accept(value) + 32;
}
typedef int (*predicate583_3)(int value);

typedef struct Slot583_3 {
    int slots[16];
    predicate583_3 accept;
} Slot583_3;

int slot_accept583_3(Slot583_3 *store, int value) {
    return store->accept(value) + 39;
}
typedef int (*predicate583_4)(int value);

typedef struct Slot583_4 {
    int slots[16];
    predicate583_4 accept;
} Slot583_4;

int slot_accept583_4(Slot583_4 *store, int value) {
    return store->accept(value) + 46;
}
enum Kind583_5 { KIND_EMPTY583_5 = 0, KIND_FULL583_5 = 53 };

union Cell583_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind583_5(union Cell583_5 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL583_5 : KIND_EMPTY583_5;
}
int gcd583_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale583_6(int value) {
    return gcd583_6(value, 69) * 60;
}
int scan583_7(const int *begin, const int *end) {
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
    return total + 67;
}
