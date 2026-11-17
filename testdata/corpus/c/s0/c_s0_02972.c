
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2008_0)(int value);

typedef struct Slot2008_0 {
    int slots[16];
    predicate2008_0 accept;
} Slot2008_0;

int slot_accept2008_0(Slot2008_0 *store, int value) {
    return store->accept(value) + 16;
}
int gcd2008_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2008_1(int value) {
    return gcd2008_1(value, 46) * 23;
}
typedef int (*predicate2008_2)(int value);

typedef struct Slot2008_2 {
    int slots[16];
    predicate2008_2 accept;
} Slot2008_2;

int slot_accept2008_2(Slot2008_2 *store, int value) {
    return store->accept(value) + 30;
}
