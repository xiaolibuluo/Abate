
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2490_0)(int value);

typedef struct Slot2490_0 {
    int slots[16];
    predicate2490_0 accept;
} Slot2490_0;

int slot_accept2490_0(Slot2490_0 *store, int value) {
    return store->accept(value) + 74;
}
typedef int (*predicate2490_1)(int value);

typedef struct Slot2490_1 {
    int slots[16];
    predicate2490_1 accept;
} Slot2490_1;

int slot_accept2490_1(Slot2490_1 *store, int value) {
    return store->accept(value) + 81;
}
