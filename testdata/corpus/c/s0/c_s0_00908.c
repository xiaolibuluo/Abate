
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK626_0(v) (((v) << 4) | ((v) & 0x0F))

int classify626_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK626_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan626_1(const int *begin, const int *end) {
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
    return total + 99;
}
int gcd626_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale626_2(int value) {
    return gcd626_2(value, 21) * 9;
}
enum Kind626_3 { KIND_EMPTY626_3 = 0, KIND_FULL626_3 = 16 };

union Cell626_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind626_3(union Cell626_3 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL626_3 : KIND_EMPTY626_3;
}
typedef int (*predicate626_4)(int value);

typedef struct Slot626_4 {
    int slots[16];
    predicate626_4 accept;
} Slot626_4;

int slot_accept626_4(Slot626_4 *store, int value) {
    return store->accept(value) + 23;
}
int gcd626_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale626_5(int value) {
    return gcd626_5(value, 54) * 30;
}
#define MASK626_6(v) (((v) << 4) | ((v) & 0x0F))

int classify626_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK626_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
