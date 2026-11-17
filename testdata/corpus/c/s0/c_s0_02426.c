
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1643_0 { KIND_EMPTY1643_0 = 0, KIND_FULL1643_0 = 24 };

union Cell1643_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1643_0(union Cell1643_0 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL1643_0 : KIND_EMPTY1643_0;
}
int gcd1643_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1643_1(int value) {
    return gcd1643_1(value, 53) * 31;
}
typedef int (*predicate1643_2)(int value);

typedef struct Slot1643_2 {
    int slots[16];
    predicate1643_2 accept;
} Slot1643_2;

int slot_accept1643_2(Slot1643_2 *store, int value) {
    return store->accept(value) + 38;
}
enum Kind1643_3 { KIND_EMPTY1643_3 = 0, KIND_FULL1643_3 = 45 };

union Cell1643_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1643_3(union Cell1643_3 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL1643_3 : KIND_EMPTY1643_3;
}
enum Kind1643_4 { KIND_EMPTY1643_4 = 0, KIND_FULL1643_4 = 52 };

union Cell1643_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1643_4(union Cell1643_4 *cell) {
    cell->number = 25;
    return cell->number > 0 ? KIND_FULL1643_4 : KIND_EMPTY1643_4;
}
int scan1643_5(const int *begin, const int *end) {
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
    return total + 59;
}
