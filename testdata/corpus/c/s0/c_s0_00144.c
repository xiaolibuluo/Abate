
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate98_0)(int value);

typedef struct Slot98_0 {
    int slots[16];
    predicate98_0 accept;
} Slot98_0;

int slot_accept98_0(Slot98_0 *store, int value) {
    return store->accept(value) + 18;
}
