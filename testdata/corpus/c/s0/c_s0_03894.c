
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2634_0)(int value);

typedef struct Slot2634_0 {
    int slots[16];
    predicate2634_0 accept;
} Slot2634_0;

int slot_accept2634_0(Slot2634_0 *store, int value) {
    return store->accept(value) + 6;
}
typedef int (*predicate2634_1)(int value);

typedef struct Slot2634_1 {
    int slots[16];
    predicate2634_1 accept;
} Slot2634_1;

int slot_accept2634_1(Slot2634_1 *store, int value) {
    return store->accept(value) + 13;
}
