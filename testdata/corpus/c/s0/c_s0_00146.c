
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate100_0)(int value);

typedef struct Slot100_0 {
    int slots[16];
    predicate100_0 accept;
} Slot100_0;

int slot_accept100_0(Slot100_0 *store, int value) {
    return store->accept(value) + 44;
}
