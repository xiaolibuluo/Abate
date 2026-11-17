
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK471_0(v) (((v) << 4) | ((v) & 0x0F))

int classify471_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK471_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate471_1)(int value);

typedef struct Slot471_1 {
    int slots[16];
    predicate471_1 accept;
} Slot471_1;

int slot_accept471_1(Slot471_1 *store, int value) {
    return store->accept(value) + 24;
}
typedef int (*predicate471_2)(int value);

typedef struct Slot471_2 {
    int slots[16];
    predicate471_2 accept;
} Slot471_2;

int slot_accept471_2(Slot471_2 *store, int value) {
    return store->accept(value) + 31;
}
#if defined(__STDC_VERSION__)
int variant471_3(int value, ...) {
    return value + 38;
}
#else
int variant471_3(int value) {
    return value + 38;
}
#endif
#define MASK471_4(v) (((v) << 4) | ((v) & 0x0F))

int classify471_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK471_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate471_5)(int value);

typedef struct Slot471_5 {
    int slots[16];
    predicate471_5 accept;
} Slot471_5;

int slot_accept471_5(Slot471_5 *store, int value) {
    return store->accept(value) + 52;
}
#if defined(__STDC_VERSION__)
int variant471_6(int value, ...) {
    return value + 59;
}
#else
int variant471_6(int value) {
    return value + 59;
}
#endif
