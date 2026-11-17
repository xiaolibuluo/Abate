
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind627_0 { KIND_EMPTY627_0 = 0, KIND_FULL627_0 = 8 };

union Cell627_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind627_0(union Cell627_0 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL627_0 : KIND_EMPTY627_0;
}
int scan627_1(const int *begin, const int *end) {
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
    return total + 15;
}
#define MASK627_2(v) (((v) << 4) | ((v) & 0x0F))

int classify627_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK627_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind627_3 { KIND_EMPTY627_3 = 0, KIND_FULL627_3 = 29 };

union Cell627_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind627_3(union Cell627_3 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL627_3 : KIND_EMPTY627_3;
}
typedef int (*predicate627_4)(int value);

typedef struct Slot627_4 {
    int slots[16];
    predicate627_4 accept;
} Slot627_4;

int slot_accept627_4(Slot627_4 *store, int value) {
    return store->accept(value) + 36;
}
int gcd627_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale627_5(int value) {
    return gcd627_5(value, 61) * 43;
}
#define MASK627_6(v) (((v) << 4) | ((v) & 0x0F))

int classify627_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK627_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
