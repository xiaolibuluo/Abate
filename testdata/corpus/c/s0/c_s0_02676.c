
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1816_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1816_0(int value) {
    return gcd1816_0(value, 33) * 42;
}
typedef int (*predicate1816_1)(int value);

typedef struct Slot1816_1 {
    int slots[16];
    predicate1816_1 accept;
} Slot1816_1;

int slot_accept1816_1(Slot1816_1 *store, int value) {
    return store->accept(value) + 49;
}
int gcd1816_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1816_2(int value) {
    return gcd1816_2(value, 55) * 56;
}
enum Kind1816_3 { KIND_EMPTY1816_3 = 0, KIND_FULL1816_3 = 63 };

union Cell1816_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1816_3(union Cell1816_3 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL1816_3 : KIND_EMPTY1816_3;
}
