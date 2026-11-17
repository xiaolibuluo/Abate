
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate494_0)(int value);

typedef struct Slot494_0 {
    int slots[16];
    predicate494_0 accept;
} Slot494_0;

int slot_accept494_0(Slot494_0 *store, int value) {
    return store->accept(value) + 25;
}
typedef int (*predicate494_1)(int value);

typedef struct Slot494_1 {
    int slots[16];
    predicate494_1 accept;
} Slot494_1;

int slot_accept494_1(Slot494_1 *store, int value) {
    return store->accept(value) + 32;
}
