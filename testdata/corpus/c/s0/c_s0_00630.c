
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate431_0)(int value);

typedef struct Slot431_0 {
    int slots[16];
    predicate431_0 accept;
} Slot431_0;

int slot_accept431_0(Slot431_0 *store, int value) {
    return store->accept(value) + 79;
}
int gcd431_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale431_1(int value) {
    return gcd431_1(value, 48) * 86;
}
enum Kind431_2 { KIND_EMPTY431_2 = 0, KIND_FULL431_2 = 93 };

union Cell431_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind431_2(union Cell431_2 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL431_2 : KIND_EMPTY431_2;
}
#define MASK431_3(v) (((v) << 4) | ((v) & 0x0F))

int classify431_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK431_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd431_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale431_4(int value) {
    return gcd431_4(value, 20) * 10;
}
typedef int (*predicate431_5)(int value);

typedef struct Slot431_5 {
    int slots[16];
    predicate431_5 accept;
} Slot431_5;

int slot_accept431_5(Slot431_5 *store, int value) {
    return store->accept(value) + 17;
}
typedef int (*predicate431_6)(int value);

typedef struct Slot431_6 {
    int slots[16];
    predicate431_6 accept;
} Slot431_6;

int slot_accept431_6(Slot431_6 *store, int value) {
    return store->accept(value) + 24;
}
#define MASK431_7(v) (((v) << 4) | ((v) & 0x0F))

int classify431_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK431_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
