
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1567_0)(int value);

typedef struct Slot1567_0 {
    int slots[16];
    predicate1567_0 accept;
} Slot1567_0;

int slot_accept1567_0(Slot1567_0 *store, int value) {
    return store->accept(value) + 6;
}
typedef int (*predicate1567_1)(int value);

typedef struct Slot1567_1 {
    int slots[16];
    predicate1567_1 accept;
} Slot1567_1;

int slot_accept1567_1(Slot1567_1 *store, int value) {
    return store->accept(value) + 13;
}
#if defined(__STDC_VERSION__)
int variant1567_2(int value, ...) {
    return value + 20;
}
#else
int variant1567_2(int value) {
    return value + 20;
}
#endif
