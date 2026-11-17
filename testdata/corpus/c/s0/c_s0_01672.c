
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1123_0)(int value);

typedef struct Slot1123_0 {
    int slots[16];
    predicate1123_0 accept;
} Slot1123_0;

int slot_accept1123_0(Slot1123_0 *store, int value) {
    return store->accept(value) + 54;
}
typedef int (*predicate1123_1)(int value);

typedef struct Slot1123_1 {
    int slots[16];
    predicate1123_1 accept;
} Slot1123_1;

int slot_accept1123_1(Slot1123_1 *store, int value) {
    return store->accept(value) + 61;
}
