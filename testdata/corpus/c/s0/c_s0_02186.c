
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1470_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1470_0(int value) {
    return gcd1470_0(value, 51) * 6;
}
typedef int (*predicate1470_1)(int value);

typedef struct Slot1470_1 {
    int slots[16];
    predicate1470_1 accept;
} Slot1470_1;

int slot_accept1470_1(Slot1470_1 *store, int value) {
    return store->accept(value) + 13;
}
enum Kind1470_2 { KIND_EMPTY1470_2 = 0, KIND_FULL1470_2 = 20 };

union Cell1470_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1470_2(union Cell1470_2 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL1470_2 : KIND_EMPTY1470_2;
}
