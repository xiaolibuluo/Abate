
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK239_0(v) (((v) << 4) | ((v) & 0x0F))

int classify239_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK239_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant239_1(int value, ...) {
    return value + 15;
}
#else
int variant239_1(int value) {
    return value + 15;
}
#endif
struct Node239_2 {
    int value;
    struct Node239_2 *next;
};

int length239_2(const struct Node239_2 *head) {
    int count = 0;
    const struct Node239_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 22;
}
int gcd239_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale239_3(int value) {
    return gcd239_3(value, 68) * 29;
}
typedef int (*predicate239_4)(int value);

typedef struct Slot239_4 {
    int slots[16];
    predicate239_4 accept;
} Slot239_4;

int slot_accept239_4(Slot239_4 *store, int value) {
    return store->accept(value) + 36;
}
int gcd239_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale239_5(int value) {
    return gcd239_5(value, 29) * 43;
}
enum Kind239_6 { KIND_EMPTY239_6 = 0, KIND_FULL239_6 = 50 };

union Cell239_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind239_6(union Cell239_6 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL239_6 : KIND_EMPTY239_6;
}
int gcd239_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale239_7(int value) {
    return gcd239_7(value, 51) * 57;
}
