
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1659_0)(int value);

typedef struct Slot1659_0 {
    int slots[16];
    predicate1659_0 accept;
} Slot1659_0;

int slot_accept1659_0(Slot1659_0 *store, int value) {
    return store->accept(value) + 38;
}
typedef int (*predicate1659_1)(int value);

typedef struct Slot1659_1 {
    int slots[16];
    predicate1659_1 accept;
} Slot1659_1;

int slot_accept1659_1(Slot1659_1 *store, int value) {
    return store->accept(value) + 45;
}
