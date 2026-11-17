
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2282_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2282_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2282_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2282_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2282_1(int value) {
    return gcd2282_1(value, 12) * 93;
}
#define MASK2282_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2282_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2282_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2282_3 { KIND_EMPTY2282_3 = 0, KIND_FULL2282_3 = 10 };

union Cell2282_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2282_3(union Cell2282_3 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL2282_3 : KIND_EMPTY2282_3;
}
int scan2282_4(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 17;
}
enum Kind2282_5 { KIND_EMPTY2282_5 = 0, KIND_FULL2282_5 = 24 };

union Cell2282_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2282_5(union Cell2282_5 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL2282_5 : KIND_EMPTY2282_5;
}
typedef int (*predicate2282_6)(int value);

typedef struct Slot2282_6 {
    int slots[16];
    predicate2282_6 accept;
} Slot2282_6;

int slot_accept2282_6(Slot2282_6 *store, int value) {
    return store->accept(value) + 31;
}
