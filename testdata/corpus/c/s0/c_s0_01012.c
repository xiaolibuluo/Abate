
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind695_0 { KIND_EMPTY695_0 = 0, KIND_FULL695_0 = 19 };

union Cell695_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind695_0(union Cell695_0 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL695_0 : KIND_EMPTY695_0;
}
int gcd695_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale695_1(int value) {
    return gcd695_1(value, 66) * 26;
}
typedef int (*predicate695_2)(int value);

typedef struct Slot695_2 {
    int slots[16];
    predicate695_2 accept;
} Slot695_2;

int slot_accept695_2(Slot695_2 *store, int value) {
    return store->accept(value) + 33;
}
