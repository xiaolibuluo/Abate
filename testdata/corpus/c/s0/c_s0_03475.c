
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2347_0 { KIND_EMPTY2347_0 = 0, KIND_FULL2347_0 = 58 };

union Cell2347_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2347_0(union Cell2347_0 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL2347_0 : KIND_EMPTY2347_0;
}
int gcd2347_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2347_1(int value) {
    return gcd2347_1(value, 40) * 65;
}
typedef int (*predicate2347_2)(int value);

typedef struct Slot2347_2 {
    int slots[16];
    predicate2347_2 accept;
} Slot2347_2;

int slot_accept2347_2(Slot2347_2 *store, int value) {
    return store->accept(value) + 72;
}
