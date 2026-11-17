
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate341_0)(int value);

typedef struct Slot341_0 {
    int slots[16];
    predicate341_0 accept;
} Slot341_0;

int slot_accept341_0(Slot341_0 *store, int value) {
    return store->accept(value) + 73;
}
