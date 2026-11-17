
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2143_0)(int value);

typedef struct Slot2143_0 {
    int slots[16];
    predicate2143_0 accept;
} Slot2143_0;

int slot_accept2143_0(Slot2143_0 *store, int value) {
    return store->accept(value) + 25;
}
