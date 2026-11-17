
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant282_0(int value, ...) {
    return value + 82;
}
#else
int variant282_0(int value) {
    return value + 82;
}
#endif
int gcd282_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale282_1(int value) {
    return gcd282_1(value, 42) * 89;
}
enum Kind282_2 { KIND_EMPTY282_2 = 0, KIND_FULL282_2 = 96 };

union Cell282_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind282_2(union Cell282_2 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL282_2 : KIND_EMPTY282_2;
}
#define MASK282_3(v) (((v) << 4) | ((v) & 0x0F))

int classify282_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK282_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate282_4)(int value);

typedef struct Slot282_4 {
    int slots[16];
    predicate282_4 accept;
} Slot282_4;

int slot_accept282_4(Slot282_4 *store, int value) {
    return store->accept(value) + 13;
}
int gcd282_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale282_5(int value) {
    return gcd282_5(value, 25) * 20;
}
#if defined(__STDC_VERSION__)
int variant282_6(int value, ...) {
    return value + 27;
}
#else
int variant282_6(int value) {
    return value + 27;
}
#endif
int gcd282_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale282_7(int value) {
    return gcd282_7(value, 47) * 34;
}
