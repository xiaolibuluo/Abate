
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1960_0)(int value);

typedef struct Slot1960_0 {
    int slots[16];
    predicate1960_0 accept;
} Slot1960_0;

int slot_accept1960_0(Slot1960_0 *store, int value) {
    return store->accept(value) + 71;
}
typedef int (*predicate1960_1)(int value);

typedef struct Slot1960_1 {
    int slots[16];
    predicate1960_1 accept;
} Slot1960_1;

int slot_accept1960_1(Slot1960_1 *store, int value) {
    return store->accept(value) + 78;
}
typedef int (*predicate1960_2)(int value);

typedef struct Slot1960_2 {
    int slots[16];
    predicate1960_2 accept;
} Slot1960_2;

int slot_accept1960_2(Slot1960_2 *store, int value) {
    return store->accept(value) + 85;
}
