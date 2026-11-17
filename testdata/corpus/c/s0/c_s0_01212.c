
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate828_0)(int value);

typedef struct Slot828_0 {
    int slots[16];
    predicate828_0 accept;
} Slot828_0;

int slot_accept828_0(Slot828_0 *store, int value) {
    return store->accept(value) + 99;
}
