
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1776_0)(int value);

typedef struct Slot1776_0 {
    int slots[16];
    predicate1776_0 accept;
} Slot1776_0;

int slot_accept1776_0(Slot1776_0 *store, int value) {
    return store->accept(value) + 7;
}
typedef int (*predicate1776_1)(int value);

typedef struct Slot1776_1 {
    int slots[16];
    predicate1776_1 accept;
} Slot1776_1;

int slot_accept1776_1(Slot1776_1 *store, int value) {
    return store->accept(value) + 14;
}
#define MASK1776_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1776_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1776_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1776_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1776_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1776_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1776_4)(int value);

typedef struct Slot1776_4 {
    int slots[16];
    predicate1776_4 accept;
} Slot1776_4;

int slot_accept1776_4(Slot1776_4 *store, int value) {
    return store->accept(value) + 35;
}
