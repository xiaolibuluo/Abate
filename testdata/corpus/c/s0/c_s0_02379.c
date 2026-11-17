
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1608_0)(int value);

typedef struct Slot1608_0 {
    int slots[16];
    predicate1608_0 accept;
} Slot1608_0;

int slot_accept1608_0(Slot1608_0 *store, int value) {
    return store->accept(value) + 54;
}
