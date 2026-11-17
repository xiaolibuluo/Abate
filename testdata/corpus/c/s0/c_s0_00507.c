
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind351_0 { KIND_EMPTY351_0 = 0, KIND_FULL351_0 = 9 };

union Cell351_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind351_0(union Cell351_0 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL351_0 : KIND_EMPTY351_0;
}
typedef int (*predicate351_1)(int value);

typedef struct Slot351_1 {
    int slots[16];
    predicate351_1 accept;
} Slot351_1;

int slot_accept351_1(Slot351_1 *store, int value) {
    return store->accept(value) + 16;
}
int gcd351_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale351_2(int value) {
    return gcd351_2(value, 48) * 23;
}
