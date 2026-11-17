
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1315_0)(int value);

typedef struct Slot1315_0 {
    int slots[16];
    predicate1315_0 accept;
} Slot1315_0;

int slot_accept1315_0(Slot1315_0 *store, int value) {
    return store->accept(value) + 28;
}
