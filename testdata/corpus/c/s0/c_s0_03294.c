
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2226_0)(int value);

typedef struct Slot2226_0 {
    int slots[16];
    predicate2226_0 accept;
} Slot2226_0;

int slot_accept2226_0(Slot2226_0 *store, int value) {
    return store->accept(value) + 37;
}
typedef int (*predicate2226_1)(int value);

typedef struct Slot2226_1 {
    int slots[16];
    predicate2226_1 accept;
} Slot2226_1;

int slot_accept2226_1(Slot2226_1 *store, int value) {
    return store->accept(value) + 44;
}
int scan2226_2(const int *begin, const int *end) {
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
    return total + 51;
}
int gcd2226_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2226_3(int value) {
    return gcd2226_3(value, 69) * 58;
}
int gcd2226_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2226_4(int value) {
    return gcd2226_4(value, 19) * 65;
}
enum Kind2226_5 { KIND_EMPTY2226_5 = 0, KIND_FULL2226_5 = 72 };

union Cell2226_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2226_5(union Cell2226_5 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL2226_5 : KIND_EMPTY2226_5;
}
