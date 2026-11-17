
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate319_0)(int value);

typedef struct Slot319_0 {
    int slots[16];
    predicate319_0 accept;
} Slot319_0;

int slot_accept319_0(Slot319_0 *store, int value) {
    return store->accept(value) + 78;
}
int gcd319_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale319_1(int value) {
    return gcd319_1(value, 57) * 85;
}
#define MASK319_2(v) (((v) << 4) | ((v) & 0x0F))

int classify319_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK319_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind319_3 { KIND_EMPTY319_3 = 0, KIND_FULL319_3 = 99 };

union Cell319_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind319_3(union Cell319_3 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL319_3 : KIND_EMPTY319_3;
}
enum Kind319_4 { KIND_EMPTY319_4 = 0, KIND_FULL319_4 = 9 };

union Cell319_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind319_4(union Cell319_4 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL319_4 : KIND_EMPTY319_4;
}
