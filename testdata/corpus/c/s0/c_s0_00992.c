
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node681_0 {
    int value;
    struct Node681_0 *next;
};

int length681_0(const struct Node681_0 *head) {
    int count = 0;
    const struct Node681_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}
enum Kind681_1 { KIND_EMPTY681_1 = 0, KIND_FULL681_1 = 38 };

union Cell681_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind681_1(union Cell681_1 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL681_1 : KIND_EMPTY681_1;
}
typedef int (*predicate681_2)(int value);

typedef struct Slot681_2 {
    int slots[16];
    predicate681_2 accept;
} Slot681_2;

int slot_accept681_2(Slot681_2 *store, int value) {
    return store->accept(value) + 45;
}
#define MASK681_3(v) (((v) << 4) | ((v) & 0x0F))

int classify681_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK681_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant681_4(int value, ...) {
    return value + 59;
}
#else
int variant681_4(int value) {
    return value + 59;
}
#endif
#define MASK681_5(v) (((v) << 4) | ((v) & 0x0F))

int classify681_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK681_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant681_6(int value, ...) {
    return value + 73;
}
#else
int variant681_6(int value) {
    return value + 73;
}
#endif
#define MASK681_7(v) (((v) << 4) | ((v) & 0x0F))

int classify681_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK681_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd681_8(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale681_8(int value) {
    return gcd681_8(value, 45) * 87;
}
